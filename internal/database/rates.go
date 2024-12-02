package database

import (
	"context"
	"github.com/Muvi7z/telegramBot.git/internal/domain"
	"time"
)

type RatesDB struct {
	store map[int64][]domain.Rate
}

func NewRateDB() (*RatesDB, error) {
	return &RatesDB{store: make(map[int64][]domain.Rate)}, nil
}

func (db *RatesDB) AddRate(ctx context.Context, time time.Time, rate domain.Rate) error {
	return nil
}

func (db *RatesDB) GetRates() ([]domain.Rate, error) {

}
