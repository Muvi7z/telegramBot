package main

import (
	"context"
	"database/sql"
	"github.com/Muvi7z/telegramBot.git/internal/clients/cbr"
	"github.com/Muvi7z/telegramBot.git/internal/clients/tg"
	"github.com/Muvi7z/telegramBot.git/internal/config"
	"github.com/Muvi7z/telegramBot.git/internal/database"
	"github.com/Muvi7z/telegramBot.git/internal/model/messages"
	"github.com/Muvi7z/telegramBot.git/internal/services"
	"github.com/Muvi7z/telegramBot.git/internal/worker"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log/slog"
	"os"
)

func main() {
	ctx := context.Background()
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}
	db, err := sql.Open("postgress", "host=localhost port=5432 user=postgres password=pass")
	gorm.Open(postgres.Open("host=localhost port=5432 user=postgres password=pass"))
	if err != nil {
		panic(err)
	}

	//err = db.AutoMigrate(&domain.Rate{})
	//if err != nil {
	//	panic(err)
	//}

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	tgClient, err := tg.New(cfg, logger)
	if err != nil {
		panic(err)
	}

	//DATABASE

	rateDB := database.NewRateDB(db)

	//GATEWAY

	cbrGate := cbr.New()

	msgModel := messages.New(tgClient, rateDB)

	//SERVICES

	exchangeRateUpdateSVC := services.NewExchangeRateUpdateSvc(rateDB, cbrGate, cfg)

	//WORKERS

	currencyExchangeRateWorker := worker.New(exchangeRateUpdateSVC)

	currencyExchangeRateWorker.Run(ctx)

	messagesListenerWorker := worker.NewMessageListenerWorker(tgClient, msgModel, logger)

	messagesListenerWorker.Run(ctx)

}
