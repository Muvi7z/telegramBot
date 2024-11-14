package messages

import (
	"context"
	"github.com/Muvi7z/telegramBot.git/helpers"
	"github.com/Muvi7z/telegramBot.git/internal/domain"
	"github.com/pkg/errors"
	"log"
	"strings"
	"time"
)

const (
	dateFormat = "02.01.2006"
)

const FormatAddedExpenseMessage = "Расход добавлен: %v %s %s %s"

// /add цена; описание; дата
func (s *Model) addExpense(ctx context.Context, msg Message) (string, error) {
	date := time.Now()
	title := ""

	parts := strings.Split(strings.TrimPrefix(msg.Text, "/add"), ";")

	if len(parts) <= 1 {
		return "", errors.New("invalid command")
	}

	if len(parts) >= 2 {
		title = strings.TrimSpace(parts[1])
	}

	kopecks, err := helpers.ConvertStringAmountToKopecks(parts[0])
	if err != nil {
		return "", helpers.ErrInvalidAmount
	}

	if len(parts) >= 3 {
		date, err = time.ParseInLocation(dateFormat, parts[2], time.UTC)
		if err != nil {
			log.Printf("[%d]: %s", msg.UserID, err.Error())
			return "", errors.New("invalid date")
		}
	}

	//Добавление в бд

	e := domain.Expense{
		Title:  title,
		Date:   date,
		Amount: kopecks,
	}

	_ = e

	return "", nil
}
