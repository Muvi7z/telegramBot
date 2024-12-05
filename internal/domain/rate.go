package domain

import (
	"time"
)

type Rate struct {
	ID        int
	Code      string
	Nominal   int64
	Course    string
	Kopecks   int64
	Ts        time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
