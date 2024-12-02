package worker

import (
	"context"
	"github.com/Muvi7z/telegramBot.git/internal/model/messages"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"log/slog"
)

type MessageFetcher interface {
	Start() tgbotapi.UpdatesChannel
	Stop()
	Request(callback tgbotapi.Chattable) (*tgbotapi.APIResponse, error)
}

type MessageProcessor interface {
	IncomingMessage(msg messages.Message) error
}

type MessageListenerWorker struct {
	messageFetcher MessageFetcher
	processor      MessageProcessor
	logger         *slog.Logger
}

func NewMessageListenerWorker(fetcher MessageFetcher, processor MessageProcessor, logger *slog.Logger) *MessageListenerWorker {
	return &MessageListenerWorker{
		messageFetcher: fetcher,
		processor:      processor,
		logger:         logger,
	}
}

func (m *MessageListenerWorker) Run(ctx context.Context) {

	log.Println("listening for messages")

	for update := range m.messageFetcher.Start() {
		select {
		case <-ctx.Done():
			m.messageFetcher.Stop()
			log.Println("stopped listening for messages")
			return
		default:
			if err := m.processing(ctx, update); err != nil {
				m.logger.Error("error processing update:", err)
			}
		}
	}
}

func (m *MessageListenerWorker) processing(ctx context.Context, update tgbotapi.Update) error {
	if update.Message != nil {
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		err := m.processor.IncomingMessage(messages.Message{
			Text:   update.Message.Text,
			UserID: update.Message.From.ID,
		})
		if err != nil {
			return err
		}
	} else if update.CallbackQuery != nil {
		callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
		if _, err := m.messageFetcher.Request(callback); err != nil {
			return err
		}
		err := m.processor.IncomingMessage(messages.Message{
			Text:   update.CallbackQuery.Data,
			UserID: update.CallbackQuery.From.ID,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
