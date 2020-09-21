package fmpcloud

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/spacecodewor/fmpcloud-go/objects"
)

// Url const for request
const (
	urlAPITechnicalIndicatorSymbol = "/technical_indicator/%s/%s"
)

// TechnicalIndicator client
type TechnicalIndicator struct {
	Client *resty.Client
	url    string
	apiKey string
}

// Daily Indicators. Types: SMA - EMA - WMA - DEMA - TEMA - williams - RSI - ADX - standardDeviation
func (t *TechnicalIndicator) Indicators(req objects.RequestIndicators) (iList []objects.ResponseIndicators, err error) {
	data, err := t.Client.R().
		SetQueryParams(map[string]string{
			"apikey": t.apiKey,
			"type":   string(req.Indicator),
			"period": fmt.Sprint(req.Timeperiod),
		}).
		Get(t.url + fmt.Sprintf(urlAPITechnicalIndicatorSymbol, string(req.Resolution), req.Symbol))

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data.Body(), &iList)
	if err != nil {
		return nil, err
	}

	return iList, nil
}
