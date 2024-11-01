package messages

type MessageSender interface {
	SendMessage(text string, userID int64) error
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
	return nil
}
