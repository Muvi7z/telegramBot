package domain

import (
	"gorm.io/gorm"
	"time"
)

type Rate struct {
	gorm.Model
	Code    string
	Nominal int64
	Course  string
	Kopecks int64
	Ts      time.Time
}
