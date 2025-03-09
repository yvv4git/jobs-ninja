package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/alecthomas/kingpin"
	"github.com/yvv4git/jobs-tg-collector/internal/config"
	"github.com/yvv4git/jobs-tg-collector/internal/infrastructure/clients"
	"github.com/yvv4git/jobs-tg-collector/internal/infrastructure/logger"
	"github.com/yvv4git/jobs-tg-collector/internal/service"
	"github.com/yvv4git/jobs-tg-collector/internal/utils"
)

type appMode string

const (
	appModeAuth      appMode = "auth"
	appModeHistory   appMode = "history"
	appModeSubscribe appMode = "subscribe"
)

func main() {
	configPath := kingpin.Flag("config", "Path to config file").Short('c').Default("config.toml").String()
	mode := kingpin.Arg("mode", "Mode of operation: history or subscribe").Required().Enum(string(appModeHistory), string(appModeSubscribe), string(appModeAuth))

	kingpin.Parse()

	logDefault := logger.SetupDefaultLogger()

	fmt.Println(os.Getwd())

	cfg := config.NewConfig(logDefault)
	if err := cfg.Load(utils.Deref(configPath)); err != nil {
		logDefault.Error("can't load config", "err", err)
		return
	}

	fmt.Println("CFG: ", cfg)

	log := logger.SetupLoggerWithLevel(logger.ParseLogLevel(cfg.Level))

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	clientTelegram := clients.NewTelegramClient(log, cfg.Collector.ClientTelegram)

	svcCollector := service.NewCollector(log, clientTelegram)

	switch appMode(utils.Deref(mode)) {
	case appModeAuth:
		log.Info("authenticating")

		if err := svcCollector.Authenticate(ctx); err != nil {
			log.Error("can't authenticate", "err", err)
			return
		}

	case appModeHistory:
		log.Info("fetching history")

		if err := svcCollector.FetchHistory(ctx, cfg.Collector.ClientTelegram.HistoryList); err != nil {
			log.Error("can't fetch history", "err", err)
			return
		}

	case appModeSubscribe:
		log.Info("subscribing")

		if err := clientTelegram.Subscribe(ctx, cfg.Collector.ClientTelegram.SubscribeList); err != nil {
			log.Error("can't subscribe", "err", err)
			return
		}
	}
}
