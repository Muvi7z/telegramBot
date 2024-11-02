package messages

import (
	mocks "github.com/Muvi7z/telegramBot.git/internal/mocks/messages"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOnStartCommand_ShouldAnswerWithIntroMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	sender := mocks.NewMockMessageSender(ctrl)
	model := New(sender)

	sender.EXPECT().SendMessage("не знаю эту команду", int64(123))

	err := model.IncomingMessage(Message{
		Text:   "/start",
		UserID: (123),
	})

	assert.NoError(t, err)

}
