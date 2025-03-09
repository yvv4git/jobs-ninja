package service

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/yvv4git/jobs-tg-collector/internal/domain"
)

type ClientTelegram interface {
	Authenticate(ctx context.Context) error
	History(ctx context.Context, sources []string) ([]domain.Message, error)
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

func (c *Collector) Authenticate(ctx context.Context) error {
	return c.clientTelegram.Authenticate(ctx)
}

func (c *Collector) FetchHistory(ctx context.Context, sources []string) error {
	// TODO: implement fetching history

	messages, err := c.clientTelegram.History(ctx, sources)
	if err != nil {
		return err
	}

	for i, message := range messages {
		fmt.Printf("Message[%d]: %v\n ", i, message)
	}

	return nil
}

func (c *Collector) Subscribe(ctx context.Context, sources []string) error {
	// TODO: implement subscribing to sources
	return nil
}
