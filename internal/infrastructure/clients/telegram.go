package clients

import (
	"context"
	"fmt"
	"log/slog"
	"path/filepath"
	"time"

	"github.com/gotd/td/telegram"
	"github.com/gotd/td/telegram/auth"
	"github.com/gotd/td/tg"
	"github.com/yvv4git/jobs-tg-collector/internal/config"
)

type TelegramClient struct {
	log      *slog.Logger
	cfg      config.ClientTelegram
	client   *telegram.Client
	authFlow auth.Flow
}

func NewTelegramClient(log *slog.Logger, cfg config.ClientTelegram) *TelegramClient {
	sessionStorage := &telegram.FileSessionStorage{
		Path: filepath.Join(".", cfg.SessionFile),
	}

	client := telegram.NewClient(cfg.APIID, cfg.APIHash, telegram.Options{SessionStorage: sessionStorage})

	flow := auth.NewFlow(
		termAuth{phone: cfg.Phone},
		auth.SendCodeOptions{},
	)

	return &TelegramClient{
		log:      log,
		cfg:      cfg,
		client:   client,
		authFlow: flow,
	}
}

func (t *TelegramClient) Authenticate(ctx context.Context) error {
	err := t.client.Run(ctx, func(ctx context.Context) error {
		if err := t.client.Auth().IfNecessary(ctx, t.authFlow); err != nil {
			return err
		}

		user, err := t.client.Self(ctx)
		if err != nil {
			return fmt.Errorf("get self: %w", err)
		}

		firstName, ok := user.GetFirstName()
		if ok {
			t.log.Debug("get first name", "username", firstName)
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("run client: %w", err)
	}

	return nil
}

func (t *TelegramClient) History(ctx context.Context, sources []string) error {
	err := t.client.Run(ctx, func(ctx context.Context) error {
		if err := t.client.Auth().IfNecessary(ctx, t.authFlow); err != nil {
			return err
		}

		for _, source := range sources {
			fmt.Printf("\nHistory[%s]: \n", source)

			err := t.historyByNames(ctx, source, 10)
			if err != nil {
				t.log.Error("get history by name", "error", err)
			}
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("run client: %w", err)
	}

	return nil
}

func (t *TelegramClient) historyByNames(ctx context.Context, name string, msgCount int) error {
	resolveRequest := &tg.ContactsResolveUsernameRequest{
		Username: name,
	}

	api := t.client.API()

	group, err := api.ContactsResolveUsername(ctx, resolveRequest)
	if err != nil {
		return fmt.Errorf("resolve username: %w", err)
	}

	peer := group.GetPeer()
	peerChannel, ok := peer.(*tg.PeerChannel)
	if !ok {
		return fmt.Errorf("peer is not a channel or group")
	}

	accessHash := int64(0)
	for _, chat := range group.Chats {
		if channel, ok := chat.(*tg.Channel); ok {
			if channel.ID == peerChannel.ChannelID {
				accessHash = channel.AccessHash
				break
			}
		}
	}

	if accessHash == 0 {
		return fmt.Errorf("get access hash for the channel")
	}

	inputPeer := &tg.InputPeerChannel{
		ChannelID:  peerChannel.ChannelID,
		AccessHash: accessHash,
	}

	history, err := api.MessagesGetHistory(ctx, &tg.MessagesGetHistoryRequest{
		Peer:       inputPeer,
		OffsetID:   0,
		OffsetDate: 0,
		AddOffset:  0,
		Limit:      msgCount,
		MaxID:      0,
		MinID:      0,
		Hash:       0,
	})
	if err != nil {
		return fmt.Errorf("failed to get history: %w", err)
	}

	// TODO: add processing of messages
	switch result := history.(type) {
	case *tg.MessagesMessages: // Обычные сообщения
		for _, msg := range result.Messages {
			switch m := msg.(type) {
			case *tg.Message:
				fmt.Printf("Message[%v]: %s\n", time.Unix(int64(m.Date), 0), m.Message)
			case *tg.MessageService:
				fmt.Printf("Service message: %v\n", m)
			default:
				fmt.Printf("Unknown message type: %T\n", m)
			}
		}
	case *tg.MessagesChannelMessages: // Сообщения из канала
		idx := 1
		for _, msg := range result.Messages {
			switch m := msg.(type) {
			case *tg.Message:
				fmt.Printf("[%d] Message channel[%v]: %s\n", idx, time.Unix(int64(m.Date), 0), m.Message)
			case *tg.MessageService:
				fmt.Printf("[%d] Service message channel: %v\n", idx, m)
			default:
				fmt.Printf("Unknown message type: %T\n", m)
			}
			idx++
		}
	default:
		return fmt.Errorf("unsupported message type: %T", result)
	}

	return nil
}

func (t *TelegramClient) Subscribe(ctx context.Context, sources []string) error {
	// TODO: implement
	return nil
}
