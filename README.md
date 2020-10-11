# Go API client for Financial Modeling Prep (fmpcloud.io)

![GitHub release (latest by date)](https://img.shields.io/github/v/release/spacecodewor/fmpcloud-go) ![GitHub](https://img.shields.io/github/license/spacecodewor/fmpcloud-go) ![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/spacecodewor/fmpcloud-go) [![Go Report Card](https://goreportcard.com/badge/github.com/spacecodewor/fmpcloud-go)](https://goreportcard.com/report/github.com/spacecodewor/fmpcloud-go)

## Overview
- API documentation [FmpCloud](https://fmpcloud.io/documentation) or [Financial Modeling Prep](https://financialmodelingprep.com/developer/docs/)
- API version: 3
- Package version: 1.0.12

## Support Methods 
* Company Valuation
* Calendars
* Institutional Fund
* International Filings
* Stock Time Series
* Market Indexes
* Commodities
* ETF
* Mutual Funds
* EuroNext
* TSX
* Stock Market
* Cryptocurrencies
* Forex (FX)
* Technical Indicators
* Stock Analysis
* Economic Calendar
* [Formulas](https://financialmodelingprep.com/developer/docs/formula/)
* [Status](https://financialmodelingprep.com/developer/docs/status/)

## Installation
```sh
go get -u github.com/spacecodewor/fmpcloud-go
```

Example:
```go
APIClient, err := NewAPIClient(Config{APIKey: "YOU_KEY"})
if err != nil {
    log.Println("Error init api client: " + err.Error())
}

// Get rating by symbol
rating, err := APIClient.CompanyValuation.Rating("AAPL")
if err != nil {
    log.Println("Error get rating: " + err.Error())
}

// Get Company profile by symbol
profile, err := APIClient.Stock.CompanyProfile("AAPL")
if err != nil {
    log.Println("Error get company profile: " + err.Error())
}

// Get real-time single quote
quote, err := APIClient.Stock.Quote("AAPL")
if err != nil {
    log.Println("Error get quote: " + err.Error())
}

// Crypto avalible symbol list
cList, err := APIClient.Crypto.AvalibleSymbols()
if err != nil {
    log.Println("Error get crypto symbols list: " + err.Error())
}

// Crypto quotes list
cQuotes, err := APIClient.Crypto.Quotes()
if err != nil {
    log.Println("Error get crypto quotes: " + err.Error())
}

// Forex avalible symbol list
fList, err := APIClient.Forex.AvalibleSymbols()
if err != nil {
    log.Println("Error get forex symbols list: " + err.Error())
}

// Forex quotes list
fQuotes, err := APIClient.Forex.Quotes()
if err != nil {
    log.Println("Error get forex quotes: " + err.Error())
}
```

Example custom client:
```go
// Init new instance for zap logger and config custom
cfg := zap.NewProductionConfig()
cfg.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder

logger, err := cfg.Build()
if err != nil {
    return nil, errors.Wrap(err, "Logger Error: init")
}

return logger, nil

// Init your custome API client
APIClient, err := NewAPIClient(Config{
    APIKey:  "YOU_KEY",      // Set Your API Key from site, default: demo
    Debug:   true,           // Set flag for debug request and response, default: false
    Timeout: 60,             // Set timeout for http client, default: 25
    APIUrl:  APIFmpcloudURL, // Set custom url (APIFmpcloudURL || APIFinancialModelingPrepURL), default: APIFinancialModelingPrepURL
    Logger:  logger,         // Set your (zap) logger, default: init new
    Version: "v4",           // Set custom API version, default: v3
})

if err != nil {
    log.Println("Error init api client: " + err.Error())
}
```

## FAQ
Historical candles support (count) (Daily from 1980):
* 1 minute 400 periods
* 5 minutes 400 periods
* 15 min 200 periods
* 30 minutes 200 periods
* 1 hour 200 periods
* 4 hours 200 periods

## License
[Apache](https://github.com/spacecodewor/fmpcloud-go/blob/master/LICENSE)