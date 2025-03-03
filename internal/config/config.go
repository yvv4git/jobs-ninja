package config

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/yvv4git/jobs-tg-collector/internal/infrastructure/logger"
)

type (
	Config struct {
		log       *slog.Logger
		Level     logger.LogLevel `toml:"level"`
		Collector Collector       `toml:"collector"`
	}

	Collector struct {
		ClientTelegram ClientTelegram `toml:"client_telegram"`
	}

	ClientTelegram struct {
		Phone       string `toml:"phone"`
		APIID       string `toml:"api_id"`
		APIHash     string `toml:"api_hash"`
		SessionFile string `toml:"session_file"`
	}
)

func NewConfig(log *slog.Logger) *Config {
	return &Config{
		log: log,
	}
}

func (c *Config) Load(filename string) error {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return fmt.Errorf("config file not found: %s", filename)
	}

	if _, err := toml.DecodeFile(filename, c); err != nil {
		return fmt.Errorf("failed to parse config file: %w", err)
	}

	return nil
}
