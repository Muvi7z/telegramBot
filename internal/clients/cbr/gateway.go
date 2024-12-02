package cbr

import (
	"context"
	"encoding/xml"
	"fmt"
	"github.com/Muvi7z/telegramBot.git/internal/domain"
	"github.com/pkg/errors"
	"golang.org/x/text/encoding/charmap"
	"io"
	"net/http"
	"strings"
	"time"
)

type Gateway struct {
	client *http.Client
}

func New() *Gateway {
	return &Gateway{
		client: http.DefaultClient,
	}
}

func (gate *Gateway) FetchRates(ctx context.Context, date time.Time) ([]domain.Rate, error) {
	ctx, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	url := "https://www.cbr-xml-daily.ru/daily_utf8.xml" //fmt.Sprintf("https://www.cbr-xml-daily.ru/daily_eng_utf8.xml?date_req=%s", date.Format("02/01/2006"))

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Content-Type", "application/xml")
	resp, err := gate.client.Do(req)
	if err != nil {
		return nil, err
	}
	fmt.Println(url)

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("failed to get rates on the date %s", date.Format("02/01/2006")))
	}
	defer resp.Body.Close()
	var cbrRates Rates

	d := xml.NewDecoder(resp.Body)

	d.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		switch charset {
		case "windows-1251":
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		default:
			return nil, errors.New("unsupported charset")
		}
	}

	if err := d.Decode(&cbrRates); err != nil {
		return nil, err
	}
	rates := make([]domain.Rate, len(cbrRates.Currencies))

	for _, rate := range cbrRates.Currencies {
		rates = append(rates, domain.Rate{
			Code:    rate.CharCode,
			Nominal: rate.Nominal,
			Course:  strings.Replace(rate.Value, ",", ".", 1),
		})
	}

	return rates, nil
}
