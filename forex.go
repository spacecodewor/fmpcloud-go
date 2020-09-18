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
	urlAPIForexListAndQuotes = "/fx"
	urlAPIForexSymbols       = "/symbol/available-forex-currency-pairs"
	urlAPIForexQuotes        = "/quotes/forex"
	urlAPIForexCandles       = "/historical-chart/%s/%s"
	urlAPIForexDaily         = "/historical-price-full/%s"
)

// Forex client
type Forex struct {
	Client *resty.Client
	url    string
	apiKey string
}

// AvalibleSymbols - available symbol list
func (f *Forex) AvalibleSymbols() (sList []objects.ForexSymbol, err error) {
	data, err := f.Client.R().
		SetQueryParams(map[string]string{"apikey": f.apiKey}).
		Get(f.url + urlAPIForexSymbols)

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
func (f *Forex) Quotes() (qList []objects.ForexQuote, err error) {
	data, err := f.Client.R().
		SetQueryParams(map[string]string{"apikey": f.apiKey}).
		Get(f.url + urlAPIForexQuotes)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data.Body(), &qList)
	if err != nil {
		return nil, err
	}

	return qList, nil
}

// ListSymbolsAndQuotes - Forex List And Price (Get last bid/ask data)
func (f *Forex) ListSymbolsAndQuotes() (bList []objects.ForexBindAsk, err error) {
	data, err := f.Client.R().
		SetQueryParams(map[string]string{"apikey": f.apiKey}).
		Get(f.url + urlAPIForexListAndQuotes)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data.Body(), &bList)
	if err != nil {
		return nil, err
	}

	return bList, nil
}

// Candles - historical candles
func (f *Forex) Candles(req objects.RequestForexCandleList) (cList []objects.ForexCandle, err error) {
	data, err := f.Client.R().
		SetQueryParams(map[string]string{"apikey": f.apiKey}).
		Get(f.url + fmt.Sprintf(urlAPIForexCandles, req.Period, req.Symbol))

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data.Body(), &cList)
	if err != nil {
		return nil, err
	}

	return cList, nil
}

// DailyLine - faily line
func (f *Forex) DailyLine(symbol string, serieType objects.ForexSerieType) (cList *objects.ForexDailyLineList, err error) {
	data, err := f.Client.R().
		SetQueryParams(map[string]string{"apikey": f.apiKey, "serietype": string(serieType)}).
		Get(f.url + fmt.Sprintf(urlAPIForexDaily, symbol))

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data.Body(), &cList)
	if err != nil {
		return nil, err
	}

	return cList, nil
}

// DailyChangeAndVolume - daily candle change and volume
func (f *Forex) DailyChangeAndVolume(symbol string) (cList *objects.ForexDailyCandleList, err error) {
	data, err := f.Client.R().
		SetQueryParams(map[string]string{"apikey": f.apiKey}).
		Get(f.url + fmt.Sprintf(urlAPIForexDaily, symbol))

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data.Body(), &cList)
	if err != nil {
		return nil, err
	}

	return cList, nil
}

// DailySpecificPeriod - daily candle list by specific period
func (f *Forex) DailySpecificPeriod(symbol string, from time.Time, to time.Time) (cList *objects.ForexDailyCandleList, err error) {
	data, err := f.Client.R().
		SetQueryParams(map[string]string{
			"apikey": f.apiKey,
			"from":   from.Format("2006-01-02"),
			"to":     to.Format("2006-01-02"),
		}).
		Get(f.url + fmt.Sprintf(urlAPIForexDaily, symbol))

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data.Body(), &cList)
	if err != nil {
		return nil, err
	}

	return cList, nil
}

// DailyLastNDays - daily candle list last N days
func (f *Forex) DailyLastNDays(symbol string, days int) (cList *objects.ForexDailyCandleList, err error) {
	data, err := f.Client.R().
		SetQueryParams(map[string]string{"apikey": f.apiKey, "timeseries": fmt.Sprint(days)}).
		Get(f.url + fmt.Sprintf(urlAPIForexDaily, symbol))

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data.Body(), &cList)
	if err != nil {
		return nil, err
	}

	return cList, nil
}
