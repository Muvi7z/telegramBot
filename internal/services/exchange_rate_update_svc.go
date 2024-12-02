package services

import (
	"context"
	"github.com/Muvi7z/telegramBot.git/internal/domain"
	"log"
	"time"
)

type ExchangeRateFetcherGateway interface {
	FetchRates(ctx context.Context, date time.Time) ([]domain.Rate, error)
}

type ConfigGetter interface {
	SupportedCurrencyCodes() []string
}

type ExchangeRateUpdateSvc struct {
	gateway ExchangeRateFetcherGateway
	config  ConfigGetter
}

func NewExchangeRateUpdateSvc(gateway ExchangeRateFetcherGateway, config ConfigGetter) *ExchangeRateUpdateSvc {
	return &ExchangeRateUpdateSvc{
		gateway: gateway,
		config:  config,
	}
}

func (svc *ExchangeRateUpdateSvc) UpdateExchangeRatesOn(ctx context.Context, time time.Time) error {
	rates, err := svc.gateway.FetchRates(ctx, time)
	if err != nil {
		return err
	}

	supportedCodes := svc.config.SupportedCurrencyCodes()

	supportedCodesMap := make(map[string]string, len(supportedCodes))
	for _, currency := range supportedCodes {
		supportedCodesMap[currency] = currency
	}
	for _, rate := range rates {
		if _, ok := supportedCodesMap[rate.Code]; ok {
			//перевод в kopecs
			log.Println(rate.Course, rate.Code)
		}
	}
	return nil
}