package objects

const (
	TechnicalIndicatorResolution1Min  TechnicalIndicatorResolution = "1min"
	TechnicalIndicatorResolution5Min  TechnicalIndicatorResolution = "5min"
	TechnicalIndicatorResolution15Min TechnicalIndicatorResolution = "15min"
	TechnicalIndicatorResolution30Min TechnicalIndicatorResolution = "30min"
	TechnicalIndicatorResolution1Hour TechnicalIndicatorResolution = "1hour"
	TechnicalIndicatorResolution4Hour TechnicalIndicatorResolution = "4hour"
	TechnicalIndicatorResolutionDaily TechnicalIndicatorResolution = "daily"

	TechnicalIndicatorTypeEMA               TechnicalIndicatorType = "ema"
	TechnicalIndicatorTypeSMA               TechnicalIndicatorType = "sma"
	TechnicalIndicatorTypeWMA               TechnicalIndicatorType = "wma"
	TechnicalIndicatorTypeDEMA              TechnicalIndicatorType = "dema"
	TechnicalIndicatorTypeTEMA              TechnicalIndicatorType = "tema"
	TechnicalIndicatorTypeWilliams          TechnicalIndicatorType = "williams"
	TechnicalIndicatorTypeRSI               TechnicalIndicatorType = "rsi"
	TechnicalIndicatorTypeADX               TechnicalIndicatorType = "adx"
	TechnicalIndicatorTypeStandardDeviation TechnicalIndicatorType = "standardDeviation"
)

// TechnicalIndicatorResolution ...
type TechnicalIndicatorResolution string

// TechnicalIndicatorType ...
type TechnicalIndicatorType string

// RequestIndicators ...
type RequestIndicators struct {
	Resolution TechnicalIndicatorResolution
	Indicator  TechnicalIndicatorType
	Symbol     string
	Timeperiod int64
}

// ResponseIndicators ...
type ResponseIndicators struct {
	Date              string   `json:"date"` // 2020-09-16 (daily) || 2020-09-21 14:26:00 (Intraday)
	Open              float64  `json:"open"`
	High              float64  `json:"high"`
	Low               float64  `json:"low"`
	Close             float64  `json:"close"`
	Volume            float64  `json:"volume"`
	SMA               *float64 `json:"ema"`
	EMA               *float64 `json:"sma"`
	WMA               *float64 `json:"wma"`
	DEMA              *float64 `json:"dema"`
	TEAM              *float64 `json:"tema"`
	Williams          *float64 `json:"williams"`
	RSI               *float64 `json:"rsi"`
	ADX               *float64 `json:"adx"`
	StandardDeviation *float64 `json:"standardDeviation"`
}

// TechnicalIndicatorResolution return string
func (t TechnicalIndicatorResolution) String() string {
	return string(t)
}

// TechnicalIndicatorType return string
func (t TechnicalIndicatorType) String() string {
	return string(t)
}
