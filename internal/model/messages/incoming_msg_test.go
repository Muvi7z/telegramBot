package messages

import (
	"context"
	mocks_msg "github.com/Muvi7z/telegramBot.git/internal/mocks/messages"
	"github.com/gojuno/minimock/v3"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func parseDate(date string) time.Time {
	v, _ := time.ParseInLocation(dateFormat, "01.10.2022", time.UTC)
	return v
}

func TestOnStartCommand_ShouldAnswerWithIntroMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := minimock.NewController(t)
	defer m.Finish()

	userID := int64(123)

	testCases := []struct {
		name    string
		command string
		amount  string
		kopecks int64
		title   string
		date    interface{}
		answer  string
	}{
		{
			name:    "normal",
			amount:  "100.0",
			kopecks: 10000,
			title:   "расход",
			date:    parseDate("01.10.2022"),
			command: "/add 100.0; расход; 01.10.2022",
			answer:  "Расход добавлен: 100.00 RUB расход 01.10.2022",
		},
		{
			name:    "without title",
			amount:  "100.0",
			kopecks: 10000,
			title:   "",
			date:    parseDate("01.10.2022"),
			command: "/add 100.0; ;01.10.2022",
			answer:  "Расход добавлен: 100.00 RUB  01.10.2022",
		},
		{
			name:    "without date",
			amount:  "100.0",
			kopecks: 10000,
			title:   "расход",
			date:    gomock.Any(),
			command: "/add 100.0; расход;",
			answer:  "Расход добавлен: 100.00 RUB расход " + time.Now().UTC().Format(dateFormat),
		},
		{
			name:    "only amount with semicolon",
			amount:  "100.0",
			kopecks: 10000,
			title:   "",
			date:    gomock.Any(),
			command: "/add 100.0;",
			answer:  "Расход добавлен: 100.00 RUB  " + time.Now().UTC().Format(dateFormat),
		},
		{
			name:    "only amount without semicolon",
			amount:  "100.0",
			kopecks: 10000,
			title:   "",
			date:    gomock.Any(),
			command: "/add 100.0",
			answer:  "Расход добавлен: 100.00 RUB  " + time.Now().UTC().Format(dateFormat),
		},
		{
			name:    "without amount",
			amount:  "",
			kopecks: 0,
			title:   "расход",
			date:    gomock.Any(),
			command: "/add ; расход; 01.10.2022",
			answer:  InvalidAmountMessage,
		},
		{
			name:    "invalid amount",
			amount:  "100.0.0",
			kopecks: 0,
			title:   "расход",
			date:    gomock.Any(),
			command: "/add 100.0.0; расход; 01.10.2022",
			answer:  InvalidAmountMessage,
		},
		{
			name:    "invalid date",
			amount:  "100.0",
			kopecks: 10000,
			title:   "расход",
			date:    gomock.Any(),
			command: "/add 100.0; расход; 01.24.2022",
			answer:  InvalidDateMessage,
		},
		{
			name:    "invalid date format",
			amount:  "100.0",
			kopecks: 10000,
			title:   "расход",
			date:    gomock.Any(),
			command: "/add 100.0; расход; 01.10.2022.0",
			answer:  InvalidDateMessage,
		},
		{
			name:    "empty command",
			command: "/add",
			date:    gomock.Any(),
			answer:  InvalidCommandMessage,
		},
	}

	baseCurrency := "RUB"
	sender := mocks_msg.NewMockMessageSender(ctrl)
	config := mocks_msg.NewMockConfigGetter(ctrl)
	expenseDB := mocks_msg.NewMockExpenseDB(ctrl)
	userDB := mocks_msg.NewMockUserDB(ctrl)
	rateDB := NewRateStorageMock(m)
	updater := mocks_msg.NewMockCurrencyExchangeRateUpdater(ctrl)
	model := New(sender, rateDB)

	config.EXPECT().GetBaseCurrency().Return(baseCurrency).AnyTimes()
	userDB.EXPECT().GetDefaultCurrency(gomock.Any(), userID).Return(baseCurrency, nil).AnyTimes()
	userDB.EXPECT().UserExist(gomock.Any(), gomock.Any()).Return(true).AnyTimes()
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			expenseDB.EXPECT().AddExpense(gomock.Any(), userID, tc.kopecks, tc.title, tc.date).Return(nil).AnyTimes()
			sender.EXPECT().SendMessage(tc.answer, userID).Return(nil)

			err := model.IncomingMessage(context.TODO(), Message{
				Text:   tc.command,
				UserID: userID,
			})

			assert.NoError(t, err)
		})
	}
}
