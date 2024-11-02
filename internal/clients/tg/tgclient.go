package tg

import (
	"github.com/Muvi7z/telegramBot.git/internal/model/messages"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
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

func (c *Client) SendMessage(userId int64, text string) error {
	_, err := c.client.Send(tgbotapi.NewMessage(userId, text))
	if err != nil {
		return errors.Wrap(err, "client.SendMessage")
	}

	return nil
}

func (c *Client) ListenUpdate(msgModel *messages.Model) error {
	return nil

}
