package messages

import (
	"context"
	"strings"
	"time"
)

// /add цена; описание; дата
func (s *Model) addExpense(ctx context.Context, msg Message) []string {
	date := time.Now()

	parts := strings.Split(strings.TrimPrefix(msg.Text, "/add"), ";")

	if len(parts) != 2 {

	}

	return parts
}
