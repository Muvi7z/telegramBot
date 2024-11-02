package messages

//Дз добавил возмодность добавления нового расхода, должно быть сумма, группа, дата,

type MessageSender interface {
	SendMessage(userID int64, text string) error
}

func New(tgClient MessageSender) *Model {
	return &Model{
		tgClient: tgClient,
	}
}

type Model struct {
	tgClient MessageSender
}

type Message struct {
	Text   string
	UserID int64
}

func (s *Model) IncomingMessage(msg Message) error {
	if msg.Text == "/start" {
		return s.tgClient.SendMessage(123, "hello")
	}
	return s.tgClient.SendMessage(123, "не знаю эту команду")
}
