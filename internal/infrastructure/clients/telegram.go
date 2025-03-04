package clients

import (
	"context"
	"fmt"
	"log/slog"
	"path/filepath"

	"github.com/gotd/td/telegram"
	"github.com/gotd/td/telegram/auth"
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

func (t *TelegramClient) History(ctx context.Context, sources []string) error {
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
			t.log.Debug("User first name", "username", firstName)
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("run client: %w", err)
	}

	return nil
}

func (t *TelegramClient) Subscribe(ctx context.Context, sources []string) error {
	// TODO: implement
	return nil
}
