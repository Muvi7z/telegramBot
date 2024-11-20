package tg

import (
	"github.com/Muvi7z/telegramBot.git/internal/model/messages"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
	"log"
)

type Client struct {
	client *tgbotapi.BotAPI
}

type TokenGetter interface {
	Token() string
}

func New(tokenGetter TokenGetter) (*Client, error) {
	client, err := tgbotapi.NewBotAPI(tokenGetter.Token())

	if err != nil {
		return nil, errors.Wrap(err, "NewBotAPI")
	}

	return &Client{
		client: client,
	}, nil
}

func (c *Client) SendMessage(userId int64, text string, buttons ...map[string]string) error {
	msg := tgbotapi.NewMessage(userId, text)

	if len(buttons) != 0 {
		var rows [][]tgbotapi.InlineKeyboardButton
		for _, button := range buttons {
			var row []tgbotapi.InlineKeyboardButton
			for text, data := range button {
				row = append(row, tgbotapi.NewInlineKeyboardButtonData(text, data))
			}
			rows = append(rows, row)
		}

		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(rows...)
	}

	_, err := c.client.Send(msg)
	if err != nil {
		return errors.Wrap(err, "client.SendMessage")
	}

	return nil
}

func (c *Client) ListenUpdate(msgModel *messages.Model) error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := c.client.GetUpdatesChan(u)

	log.Println("listening for messages")

	for update := range updates {
		if update.Message == nil {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			err := msgModel.IncomingMessage(messages.Message{
				Text:   update.Message.Text,
				UserID: update.Message.From.ID,
			})
			if err != nil {
				log.Println("error processing message:", err)
			}

		}
	}
	return nil

}
