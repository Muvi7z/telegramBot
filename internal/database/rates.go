package database

import (
	"context"
	"github.com/Muvi7z/telegramBot.git/internal/domain"
	"gorm.io/gorm"
	"time"
)

type RatesDB struct {
	db *gorm.DB
}

func NewRateDB(db *gorm.DB) *RatesDB {
	return &RatesDB{
		db: db,
	}
}

func (db *RatesDB) AddRate(ctx context.Context, time time.Time, rate domain.Rate) error {
	db.db.Create(&rate)
	return nil
}

func (db *RatesDB) GetRate(ctx context.Context, code string, time time.Time) (domain.Rate, error) {
	var resp domain.Rate
	db.db.Model(domain.Rate{Code: code, Ts: time}).Last(&resp)
	return resp, nil
}
