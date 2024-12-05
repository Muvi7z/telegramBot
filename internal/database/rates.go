package database

import (
	"context"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/Muvi7z/telegramBot.git/internal/domain"
	"time"
)

type RatesDB struct {
	db *sql.DB
}

func NewRateDB(db *sql.DB) *RatesDB {
	return &RatesDB{
		db: db,
	}
}

func (s *RatesDB) AddRate(ctx context.Context, date time.Time, rate domain.Rate) error {
	builder := sq.Insert("rates").Columns(
		"created_at",
		"code",
		"nominal",
		"kopecks",
		"course",
		"ts",
	).Values(
		time.Now(),
		rate.Code,
		rate.Nominal,
		rate.Kopecks,
		rate.Course,
		rate.Ts,
	)
	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	_, err = s.db.ExecContext(ctx, query, args...)

	return nil
}

func (s *RatesDB) GetRate(ctx context.Context, code string, date time.Time) (*domain.Rate, error) {
	builder := sq.Select(
		"id",
		"code",
		"nominal",
		"kopecks",
		"course",
		"ts",
		"created_at",
		"updated_at",
		"deleted_at",
	).From("rates").Where(sq.Eq{"code": code})

	if !date.IsZero() {
		builder = builder.Where(sq.Eq{"ts": date})
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	var rate domain.Rate
	err = s.db.QueryRowContext(ctx, query, args...).Scan(&rate)
	return &rate, nil
}
