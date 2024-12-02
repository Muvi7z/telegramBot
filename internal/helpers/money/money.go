package money

import (
	"github.com/pkg/errors"
	"regexp"
	"strconv"
)

var ErrInvalidAmount = errors.New("invalid amount")

// 1,000,500.10 -> 1000500.10 || 1 000 500.100 -> 1000500.100
var regexpNoNumber = regexp.MustCompile(`[^\d\.]`)

// ConvertStringAmountToKopecks - convert amount to kopecks
func ConvertStringAmountToKopecks(amount string) (int64, error) {
	v, err := strconv.ParseFloat(regexpNoNumber.ReplaceAllString(amount, ""), 64)
	if err != nil {
		return 0, ErrInvalidAmount
	}

	return ConvertFloat64AmountToKopecks(v)
}

// ConvertFloat64AmountToKopecks - convert amount to kopecks
func ConvertFloat64AmountToKopecks(amount float64) (int64, error) {
	return int64(amount * 100), nil
}
