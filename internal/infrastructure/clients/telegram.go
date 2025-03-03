package clients

import (
	"context"
	"log/slog"
)

type TelegramClient struct {
	log *slog.Logger
}

func NewTelegramClient(log *slog.Logger) *TelegramClient {
	return &TelegramClient{log: log}
}

func (t *TelegramClient) SessionStart(ctx context.Context) error {
	// TODO: implement
	return nil
}

func (t *TelegramClient) History(ctx context.Context, source string) error {
	// TODO: implement
	return nil
}

func (t *TelegramClient) Subscribe(ctx context.Context, sources []string) error {
	// TODO: implement
	return nil
}
