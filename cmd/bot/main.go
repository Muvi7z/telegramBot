package main

import (
	"context"
	"github.com/Muvi7z/telegramBot.git/internal/clients/cbr"
	"github.com/Muvi7z/telegramBot.git/internal/clients/tg"
	"github.com/Muvi7z/telegramBot.git/internal/config"
	"github.com/Muvi7z/telegramBot.git/internal/model/messages"
	"github.com/Muvi7z/telegramBot.git/internal/services"
	"github.com/Muvi7z/telegramBot.git/internal/worker"
	"log/slog"
	"os"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	tgClient, err := tg.New(cfg, logger)
	if err != nil {
		panic(err)
	}

	//GATEWAY

	cbrGate := cbr.New()

	msgModel := messages.New(tgClient)

	//SERVICES

	exchangeRateUpdateSVC := services.NewExchangeRateUpdateSvc(cbrGate, cfg)

	//WORKERS

	currencyExchangeRateWorker := worker.New(exchangeRateUpdateSVC)

	currencyExchangeRateWorker.Run(ctx)

	messagesListenerWorker := worker.NewMessageListenerWorker(tgClient, msgModel, logger)

	messagesListenerWorker.Run(ctx)

}
