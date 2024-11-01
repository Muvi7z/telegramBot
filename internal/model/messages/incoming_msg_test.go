package messages

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOnStartCommand_ShouldAnswerWithIntroMessage(t *testing.T) {
	model := New(nil)

	err := model.IncomingMessage(Message{
		Text:   "/start",
		UserID: 1,
	})

	assert.NoError(t, err)

}
