package messages

import (
	"context"
	"fmt"
	mocks "github.com/Muvi7z/telegramBot.git/internal/mocks/messages"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestAddExpense(t *testing.T) {
	ctrl := gomock.NewController(t)
	sender := mocks.NewMockMessageSender(ctrl)
	model := New(sender)
	l := model.addExpense(context.Background(), Message{
		Text:   "/add цена; описание; дата",
		UserID: 0,
	})

	fmt.Println(l[0])

}
