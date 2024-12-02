package tg

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
	"log/slog"
)

type Client struct {
	client *tgbotapi.BotAPI
	logger *slog.Logger
}

type TokenGetter interface {
	Token() string
}

func New(tokenGetter TokenGetter, logger *slog.Logger) (*Client, error) {
	client, err := tgbotapi.NewBotAPI(tokenGetter.Token())

	if err != nil {
		return nil, errors.Wrap(err, "NewBotAPI")
	}

	return &Client{
		client: client,
		logger: logger,
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

func (c *Client) Start() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return c.client.GetUpdatesChan(u)
}

func (c *Client) Stop() {
	c.client.StopReceivingUpdates()
}

func (c *Client) Request(callback tgbotapi.Chattable) (*tgbotapi.APIResponse, error) {
	request, err := c.client.Request(callback)
	if err != nil {
		return nil, err
	}
	return request, nil
}
