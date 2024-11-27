package worker

import (
	"context"
	"log"
	"time"
)

type CurrencyExchangeRateWorker struct {
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

				}
			}
		}
	}()
}
