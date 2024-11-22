package main

import (
	"context"
	"fmt"
	"github.com/Muvi7z/telegramBot.git/internal/clients/cbr"
	"github.com/Muvi7z/telegramBot.git/internal/clients/tg"
	"github.com/Muvi7z/telegramBot.git/internal/config"
	"github.com/Muvi7z/telegramBot.git/internal/model/messages"
	"log/slog"
	"os"
	"time"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	tgClient, err := tg.New(cfg, logger)
	if err != nil {
		panic(err)
	}

	cbrGate := cbr.New()

	rs, err := cbrGate.FetchRates(context.Background(), time.Now())
	if err != nil {
		panic(err)
	}

	fmt.Println(rs)

	msgModel := messages.New(tgClient)

	tgClient.ListenUpdate(msgModel)
}
