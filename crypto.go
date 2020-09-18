package fmpcloud

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/spacecodewor/fmpcloud-go/objects"
)

// Url const for request
const (
	urlAPICryptoSymbols = "/symbol/available-cryptocurrencies"
	urlAPICryptoQuotes  = "/quotes/crypto"
	urlAPICryptoCandles = "/historical-chart/%s/%s"
	urlAPICryptoDaily   = "/historical-price-full/%s"
)

// Crypto client
type Crypto struct {
	Client *resty.Client
	url    string
	apiKey string
}

// AvalibleSymbols - available symbol list
func (c *Crypto) AvalibleSymbols() (sList []objects.CryptoSymbol, err error) {
	data, err := c.Client.R().
		SetQueryParams(map[string]string{"apikey": c.apiKey}).
		Get(c.url + urlAPICryptoSymbols)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data.Body(), &sList)
	if err != nil {
		return nil, err
	}

	return sList, nil
}

// Quotes - all real-time prices
func (c *Crypto) Quotes() (qList []objects.CryptoQuote, err error) {
	data, err := c.Client.R().
		SetQueryParams(map[string]string{"apikey": c.apiKey}).
		Get(c.url + urlAPICryptoQuotes)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data.Body(), &qList)
	if err != nil {
		return nil, err
	}

	return qList, nil
}

// Candles - Historical candles
func (c *Crypto) Candles(req objects.RequestCryptoCandleList) (cList []objects.CryptoCandle, err error) {
	data, err := c.Client.R().
		SetQueryParams(map[string]string{"apikey": c.apiKey}).
		Get(c.url + fmt.Sprintf(urlAPICryptoCandles, req.Period, req.Symbol))

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data.Body(), &cList)
	if err != nil {
		return nil, err
	}

	return cList, nil
}

// DailyLine - Daily line
func (c *Crypto) DailyLine(symbol string, serieType objects.CryptoSerieType) (cList *objects.CryptoDailyLineList, err error) {
	data, err := c.Client.R().
		SetQueryParams(map[string]string{"apikey": c.apiKey, "serietype": string(serieType)}).
		Get(c.url + fmt.Sprintf(urlAPICryptoDaily, symbol))

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data.Body(), &cList)
	if err != nil {
		return nil, err
	}

	return cList, nil
}

// DailyChangeAndVolume - Daily candle change and volume
func (c *Crypto) DailyChangeAndVolume(symbol string) (cList *objects.CryptoDailyCandleList, err error) {
	data, err := c.Client.R().
		SetQueryParams(map[string]string{"apikey": c.apiKey}).
		Get(c.url + fmt.Sprintf(urlAPICryptoDaily, symbol))

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data.Body(), &cList)
	if err != nil {
		return nil, err
	}

	return cList, nil
}

// DailySpecificPeriod - Daily candle list by specific period
func (c *Crypto) DailySpecificPeriod(symbol string, from time.Time, to time.Time) (cList *objects.CryptoDailyCandleList, err error) {
	data, err := c.Client.R().
		SetQueryParams(map[string]string{
			"apikey": c.apiKey,
			"from":   from.Format("2006-01-02"),
			"to":     to.Format("2006-01-02"),
		}).
		Get(c.url + fmt.Sprintf(urlAPICryptoDaily, symbol))

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data.Body(), &cList)
	if err != nil {
		return nil, err
	}

	return cList, nil
}

// DailyLastNDays - Daily candle list last N days
func (c *Crypto) DailyLastNDays(symbol string, days int) (cList *objects.CryptoDailyCandleList, err error) {
	data, err := c.Client.R().
		SetQueryParams(map[string]string{"apikey": c.apiKey, "timeseries": fmt.Sprint(days)}).
		Get(c.url + fmt.Sprintf(urlAPICryptoDaily, symbol))

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data.Body(), &cList)
	if err != nil {
		return nil, err
	}

	return cList, nil
}
