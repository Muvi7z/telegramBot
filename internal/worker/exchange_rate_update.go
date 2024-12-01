package worker

import (
	"context"
	"log"
	"time"
)

type CurrencyExchangeRateUpdate interface {
	UpdateExchangeRatesOn(ctx context.Context, time time.Time) error
}

type CurrencyExchangeRateWorker struct {
	updater CurrencyExchangeRateUpdate
}

func New(updater CurrencyExchangeRateUpdate) *CurrencyExchangeRateWorker {
	return &CurrencyExchangeRateWorker{updater: updater}
}

func (worker *CurrencyExchangeRateWorker) Run(ctx context.Context) {
	ticker := time.NewTicker(5 * time.Second)

	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Println("stopped receiving exchange rates")
				return
			case <-ticker.C:
				select {
				case <-ctx.Done():
					log.Println("stopped receiving exchange rates")
					return
				default:
					err := worker.updater.UpdateExchangeRatesOn(ctx, time.Now())
					if err != nil {
						log.Println(err)
					}
				}
			}
		}
	}()
}
