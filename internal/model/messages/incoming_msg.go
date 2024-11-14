package messages

import "strings"

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

type MessageSender interface {
	SendMessage(userID int64, text string) error
}

type Model struct {
	tgClient MessageSender
}

func New(tgClient MessageSender) *Model {
	return &Model{
		tgClient: tgClient,
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
		return s.tgClient.SendMessage(msg.UserID, introMessage)
	}

	if strings.HasPrefix(msg.Text, "/set_currency") {
		return s.tgClient.SendMessage(msg.UserID, introMessage)
	}

	return s.tgClient.SendMessage(123, "не знаю эту команду")
}
