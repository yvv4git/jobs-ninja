package clients

import (
	"context"
	"log/slog"

	"github.com/yvv4git/jobs-tg-collector/internal/config"
)

type TelegramClient struct {
	log *slog.Logger
	cfg config.ClientTelegram
}

func NewTelegramClient(log *slog.Logger, cfg config.ClientTelegram) *TelegramClient {
	return &TelegramClient{
		log: log,
		cfg: cfg,
	}
}

func (t *TelegramClient) SessionStart(ctx context.Context) error {
	// TODO: implement
	return nil
}

func (t *TelegramClient) History(ctx context.Context, sources []string) error {
	// TODO: implement
	return nil
}

func (t *TelegramClient) Subscribe(ctx context.Context, sources []string) error {
	// TODO: implement
	return nil
}
