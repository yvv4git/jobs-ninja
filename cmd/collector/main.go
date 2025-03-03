package main

import (
	"context"
	"fmt"

	"github.com/yvv4git/jobs-tg-collector/internal/config"
	"github.com/yvv4git/jobs-tg-collector/internal/infrastructure/logger"
	"github.com/yvv4git/jobs-tg-collector/internal/service"
)

const (
	configPath = "config.toml"
)

func main() {
	log := logger.SetupDefaultLogger()

	cfg := config.NewConfig(log)
	err := cfg.Load(configPath)
	if err != nil {
		log.Error("can't load config", "err", err)
		return
	}

	fmt.Println("CFG: ", cfg)

	svcCollector := service.NewCollector(log, nil)

	svcCollector.FetchHistory(context.Background(), []string{"@jobs_tg_channel"})
}
