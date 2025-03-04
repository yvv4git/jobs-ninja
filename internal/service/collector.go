package service

import (
	"context"
	"log/slog"
)

type ClientTelegram interface {
	History(ctx context.Context, sources []string) error
	Subscribe(ctx context.Context, sources []string) error
}

type Collector struct {
	log            *slog.Logger
	clientTelegram ClientTelegram
}

func NewCollector(log *slog.Logger, clientTelegram ClientTelegram) *Collector {
	return &Collector{
		log:            log,
		clientTelegram: clientTelegram,
	}
}

func (c *Collector) FetchHistory(ctx context.Context, sources []string) error {
	// TODO: implement fetching history

	err := c.clientTelegram.History(ctx, sources)
	if err != nil {
		return err
	}

	return nil
}

func (c *Collector) Subscribe(ctx context.Context, sources []string) error {
	// TODO: implement subscribing to sources
	return nil
}
