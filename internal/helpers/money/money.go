package money

import (
	"github.com/pkg/errors"
	"regexp"
)

var ErrInvalidAmount = errors.New("invalid amount")

// 1,000,500.10 -> 1000500.10 || 1 000 500.100 -> 1000500.100
var regexpNoNumber = regexp.MustCompile(`[^\d\.]`)
