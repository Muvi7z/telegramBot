package main

import (
	"github.com/Muvi7z/telegramBot.git/internal/clients/tg"
	"github.com/Muvi7z/telegramBot.git/internal/config"
	"github.com/Muvi7z/telegramBot.git/internal/model/messages"
)

func main() {
	config, err := config.New()
	if err != nil {
		panic(err)
	}

	tgClient, err := tg.New(config)
	if err != nil {
		panic(err)
	}

	msgModel := messages.New(tgClient)

	tgClient.ListenUpdate(msgModel)
}
