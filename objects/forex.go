package objects

// CryptoCandlePeriod and CryptoSerieTypeLine
const (
	ForexCnadlePeriod1Min  ForexCnadlePeriod = "1min"
	ForexCnadlePeriod5Min  ForexCnadlePeriod = "5min"
	ForexCnadlePeriod15Min ForexCnadlePeriod = "15min"
	ForexCnadlePeriod30Min ForexCnadlePeriod = "30min"
	ForexCnadlePeriod1Hour ForexCnadlePeriod = "1hour"
	ForexCnadlePeriod4Hour ForexCnadlePeriod = "4hour"

	ForexSerieTypeLine ForexSerieType = "line"
)

// ForexCnadlePeriod ...
type ForexCnadlePeriod string

// ForexSerieType ...
type ForexSerieType string

// ForexSymbol ...
type ForexSymbol struct {
	Symbol            string `json:"symbol"`
	Name              string `json:"name"`
	Currency          string `json:"currency"`
	StockExchange     string `json:"stockExchange"`
	ExchangeShortName string `json:"exchangeShortName"`
}

// ForexCandle ...
type ForexCandle struct {
	Date   string  `json:"date"`
	Open   float64 `json:"open"`
	Low    float64 `json:"low"`
	High   float64 `json:"high"`
	Close  float64 `json:"close"`
	Volume int64   `json:"volume"`
}

// ForexBindAsk ...
type ForexBindAsk struct {
	Ticker  string  `json:"ticker"`
	Bid     string  `json:"bid"`
	Ask     string  `json:"ask"`
	Open    string  `json:"open"`
	Low     string  `json:"low"`
	High    string  `json:"high"`
	Changes float64 `json:"changes"`
	Date    string  `json:"date"`
}

// ForexQuote ...
type ForexQuote struct {
	Symbol               string   `json:"symbol"`
	Name                 string   `json:"name"`
	Price                float64  `json:"price"`
	ChangesPercentage    float64  `json:"changesPercentage"`
	Change               float64  `json:"change"`
	DayLow               float64  `json:"dayLow"`
	DayHigh              float64  `json:"dayHigh"`
	YearHigh             float64  `json:"yearHigh"`
	YearLow              float64  `json:"yearLow"`
	MarketCap            *float64 `json:"marketCap"`
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
	SharesOutstanding    *int64   `json:"sharesOutstanding"`
	Timestamp            int64    `json:"timestamp"`
}

// RequestForexCandleList ...
type RequestForexCandleList struct {
	Period ForexCnadlePeriod
	Symbol string
}

// ForexDailyLineList ...
type ForexDailyLineList struct {
	Symbol     string           `json:"symbol"`
	Historical []ForexDailyLine `json:"historical"`
}

// ForexDailyLine ...
type ForexDailyLine struct {
	Date  string  `json:"date"`
	Close float64 `json:"close"`
}

// ForexDailyCandleList ...
type ForexDailyCandleList struct {
	Symbol     string             `json:"symbol"`
	Historical []ForexDailyCandle `json:"historical"`
}

// ForexDailyCandle ...
type ForexDailyCandle struct {
	Date           string  `json:"date"`
	Open           float64 `json:"open"`
	High           float64 `json:"high"`
	Low            float64 `json:"low"`
	Close          float64 `json:"close"`
	AdjClose       float64 `json:"adjClose"`
	Change         float64 `json:"change"`
	ChangePercent  float64 `json:"changePercent"`
	Label          string  `json:"label"`
	ChangeOverTime float64 `json:"changeOverTime"`
}
