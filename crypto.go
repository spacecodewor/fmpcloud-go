package fmpcloud

import (
	"fmt"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/spacecodewor/fmpcloud-go/objects"
)

// Url const for request
const (
	urlAPICryptoSymbols = "/v3/symbol/available-cryptocurrencies"
	urlAPICryptoQuotes  = "/v3/quotes/crypto"
	urlAPICryptoCandles = "/v3/historical-chart/%s/%s"
	urlAPICryptoDaily   = "/v3/historical-price-full/%s"
)

// Crypto client
type Crypto struct {
	Client *HTTPClient
}

// AvalibleSymbols - available symbol list
func (c *Crypto) AvalibleSymbols() (sList []objects.CryptoSymbol, err error) {
	data, err := c.Client.Get(urlAPICryptoSymbols, nil)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &sList)
	if err != nil {
		return nil, err
	}

	return sList, nil
}

// Quotes - all real-time prices
func (c *Crypto) Quotes() (qList []objects.CryptoQuote, err error) {
	data, err := c.Client.Get(urlAPICryptoQuotes, nil)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &qList)
	if err != nil {
		return nil, err
	}

	return qList, nil
}

// Candles - Historical candles
func (c *Crypto) Candles(req objects.RequestCryptoCandleList) (cList []objects.CryptoCandle, err error) {
	reqParam := make(map[string]string)
	if req.From != nil {
		reqParam["from"] = req.From.Format("2006-01-02")
	}

	if req.To != nil {
		reqParam["to"] = req.To.Format("2006-01-02")
	}

	data, err := c.Client.Get(fmt.Sprintf(urlAPICryptoCandles, req.Period, req.Symbol), reqParam)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &cList)
	if err != nil {
		return nil, err
	}

	return cList, nil
}

// DailyLine - Daily line
func (c *Crypto) DailyLine(symbol string, serieType objects.CryptoSerieType) (cList *objects.CryptoDailyLineList, err error) {
	data, err := c.Client.Get(
		fmt.Sprintf(urlAPICryptoDaily, symbol),
		map[string]string{
			"serietype": string(serieType),
		})
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &cList)
	if err != nil {
		return nil, err
	}

	return cList, nil
}

// DailyChangeAndVolume - Daily candle change and volume
func (c *Crypto) DailyChangeAndVolume(symbol string) (cList *objects.CryptoDailyCandleList, err error) {
	data, err := c.Client.Get(fmt.Sprintf(urlAPICryptoDaily, symbol), nil)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &cList)
	if err != nil {
		return nil, err
	}

	return cList, nil
}

// DailySpecificPeriod - Daily candle list by specific period
func (c *Crypto) DailySpecificPeriod(symbol string, from time.Time, to time.Time) (cList *objects.CryptoDailyCandleList, err error) {
	data, err := c.Client.Get(
		fmt.Sprintf(urlAPICryptoDaily, symbol),
		map[string]string{
			"from": from.Format("2006-01-02"),
			"to":   to.Format("2006-01-02"),
		})
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &cList)
	if err != nil {
		return nil, err
	}

	return cList, nil
}

// DailyLastNDays - Daily candle list last N days
func (c *Crypto) DailyLastNDays(symbol string, days int) (cList *objects.CryptoDailyCandleList, err error) {
	data, err := c.Client.Get(
		fmt.Sprintf(urlAPICryptoDaily, symbol),
		map[string]string{
			"timeseries": fmt.Sprint(days),
		})
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &cList)
	if err != nil {
		return nil, err
	}

	return cList, nil
}
