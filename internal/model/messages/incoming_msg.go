package messages

import (
	"context"
	"github.com/Muvi7z/telegramBot.git/internal/domain"
	"strings"
	"time"
)

//Дз добавил возмодность добавления нового расхода, должно быть сумма, группа, дата,

const introMessage = "Привет! Я умею учитывать твои траты.\n\n" + helpMessage

const helpMessage = `Для работы с ботом тебе могут потребоваться следующие команды:
Чтобы изменить выбранную валюту необходимо выполнить команду /change_currency

Чтобы добавить новый расход, отправь мне сообщение в формате:
/add цена; описание; дата (дд.мм.гггг, опционально) - если не указать дату, то расход будет добавлен на сегодняшний день

Чтобы посмотреть расходы отправь:
/list -  за всё время.
/list_day - за день.
/list_week - за неделю.
/list_month - за месяц.
/list_year - за год.`

const (
	InvalidCommandMessage       = "Неверный формат команды, исправьте и повторите команду"
	InvalidAmountMessage        = "Неверный формат суммы, исправьте и повторите команду"
	InvalidDateMessage          = "Неверный формат даты, исправьте и повторите команду"
	FailedWriteMessage          = "Не удалось записать расход, повторите попытку позже"
	FailedMessage               = "Я временно не работаю, повторите попытку позже"
	FailedChangeCurrencyMessage = "Не удалось изменить текущую валюту, повторите попытку позже"
)

type MessageSender interface {
	SendMessage(userID int64, text string, buttons ...map[string]string) error
}

type RateStorage interface {
	GetRate(ctx context.Context, Code string, time time.Time) (*domain.Rate, error)
}

type Model struct {
	tgClient    MessageSender
	rateStorage RateStorage
}

func New(tgClient MessageSender, rateStorage RateStorage) *Model {
	return &Model{
		tgClient:    tgClient,
		rateStorage: rateStorage,
	}
}

type Message struct {
	Text   string
	UserID int64
}

func (s *Model) IncomingMessage(msg Message) error {
	if msg.Text == "/start" {
		return s.tgClient.SendMessage(msg.UserID, introMessage)
	}

	if msg.Text == "/help" {
		return s.tgClient.SendMessage(msg.UserID, helpMessage)
	}

	if strings.HasPrefix(msg.Text, "/add") {
		expense, err := s.addExpense(context.Background(), msg)
		if err != nil {
			return err
		}
		return s.tgClient.SendMessage(msg.UserID, expense)
	}

	if strings.HasPrefix(msg.Text, "/set_currency") {
		return s.tgClient.SendMessage(msg.UserID, introMessage)
	}

	if strings.HasPrefix(msg.Text, "/change_currency") {
		answer, buttons := s.changeDefaultCurrency()
		return s.tgClient.SendMessage(msg.UserID, answer, buttons...)
	}

	return s.tgClient.SendMessage(123, "не знаю эту команду")
}
