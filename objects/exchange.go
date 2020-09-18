package objects

// Exchange ...
type Exchange struct {
	StockExchangeName       string             `json:"stockExchangeName"`
	StockMarketHours        ExchangeHours      `json:"stockMarketHours"`
	StockMarketHolidays     []ExchangeHolidays `json:"stockMarketHolidays"`
	IsTheStockMarketOpen    bool               `json:"isTheStockMarketOpen"`
	IsTheEuronextMarketOpen bool               `json:"isTheEuronextMarketOpen"`
	IsTheForexMarketOpen    bool               `json:"isTheForexMarketOpen"`
	IsTheCryptoMarketOpen   bool               `json:"isTheCryptoMarketOpen"`
}

// ExchangeHours ...
type ExchangeHours struct {
	OpeningHour string `json:"openingHour"`
	ClosingHour string `json:"closingHour"`
}

// ExchangeHolidays ...
type ExchangeHolidays struct {
	Year            int    `json:"year"`
	NewYearsDay     string `json:"New Years Day"`
	GoodFriday      string `json:"Good Friday"`
	MemorialDay     string `json:"Memorial Day"`
	IndependenceDay string `json:"Independence Day"`
	LaborDay        string `json:"Labor Day"`
	ThanksgivingDay string `json:"Thanksgiving Day"`
	Christmas       string `json:"Christmas"`
}
