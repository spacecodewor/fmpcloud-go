package objects

import "time"

// CryptoCandlePeriod and CryptoSerieTypeLine
const (
	CryptoCandlePeriod1Min  CryptoCandlePeriod = "1min"
	CryptoCandlePeriod5Min  CryptoCandlePeriod = "5min"
	CryptoCandlePeriod15Min CryptoCandlePeriod = "15min"
	CryptoCandlePeriod30Min CryptoCandlePeriod = "30min"
	CryptoCandlePeriod1Hour CryptoCandlePeriod = "1hour"
	CryptoCnadlePeriod4Hour CryptoCandlePeriod = "4hour"

	CryptoSerieTypeLine CryptoSerieType = "line"
)

// CryptoCandlePeriod ...
type CryptoCandlePeriod string

// CryptoSerieType ...
type CryptoSerieType string

// CryptoSymbol ...
type CryptoSymbol struct {
	Symbol            string `json:"symbol"`
	Name              string `json:"name"`
	Currency          string `json:"currency"`
	StockExchange     string `json:"stockExchange"`
	ExchangeShortName string `json:"exchangeShortName"`
}

// CryptoCandle ...
type CryptoCandle struct {
	Date   string  `json:"date"` // 2020-09-14 07:27:00
	Open   float64 `json:"open"`
	Low    float64 `json:"low"`
	High   float64 `json:"high"`
	Close  float64 `json:"close"`
	Volume int64   `json:"volume"`
}

// CryptoQuote ...
type CryptoQuote struct {
	Symbol               string   `json:"symbol"`
	Name                 string   `json:"name"`
	Price                float64  `json:"price"`
	ChangesPercentage    float64  `json:"changesPercentage"`
	Change               float64  `json:"change"`
	DayLow               float64  `json:"dayLow"`
	DayHigh              float64  `json:"dayHigh"`
	YearHigh             float64  `json:"yearHigh"`
	YearLow              float64  `json:"yearLow"`
	MarketCap            float64  `json:"marketCap"`
	PriceAvg50           float64  `json:"priceAvg50"`
	PriceAvg200          float64  `json:"priceAvg200"`
	Volume               int64    `json:"volume"`
	AvgVolume            int64    `json:"avgVolume"`
	Exchange             string   `json:"exchange"`
	Open                 float64  `json:"open"`
	PreviousClose        float64  `json:"previousClose"`
	Eps                  *float64 `json:"eps"`
	Pe                   *float64 `json:"pe"`
	EarningsAnnouncement *float64 `json:"earningsAnnouncement"`
	SharesOutstanding    int64    `json:"sharesOutstanding"`
	Timestamp            int64    `json:"timestamp"`
}

// RequestCryptoCandleList ...
type RequestCryptoCandleList struct {
	Period CryptoCandlePeriod
	Symbol string
	From   *time.Time
	To     *time.Time
}

// CryptoDailyLineList ...
type CryptoDailyLineList struct {
	Symbol     string            `json:"symbol"`
	Historical []CryptoDailyLine `json:"historical"`
}

// CryptoDailyLine ...
type CryptoDailyLine struct {
	Date  string  `json:"date"`
	Close float64 `json:"close"`
}

// CryptoDailyCandleList ...
type CryptoDailyCandleList struct {
	Symbol     string              `json:"symbol"`
	Historical []CryptoDailyCandle `json:"historical"`
}

// CryptoDailyCandle ...
type CryptoDailyCandle struct {
	Date             string  `json:"date"` // 2019-03-11
	Open             float64 `json:"open"`
	High             float64 `json:"high"`
	Low              float64 `json:"low"`
	Close            float64 `json:"close"`
	AdjClose         float64 `json:"adjClose"`
	Volume           float64 `json:"volume"`
	UnadjustedVolume float64 `json:"unadjustedVolume"`
	Change           float64 `json:"change"`
	ChangePercent    float64 `json:"changePercent"`
	Vwap             float64 `json:"vwap"`
	Label            string  `json:"label"` // March 12, 19
	ChangeOverTime   float64 `json:"changeOverTime"`
}
