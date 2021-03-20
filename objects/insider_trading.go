package objects

// RequestInsiderTrading ...
type RequestInsiderTrading struct {
	Symbol       string
	ReportingCik string
	CompanyCik   string
	Limit        int64
}

// InsiderTrading ...
type InsiderTrading struct {
	Symbol                  string   `json:"symbol"`
	TransactionDate         string   `json:"transactionDate"`
	ReportingCik            string   `json:"reportingCik"`
	TransactionType         string   `json:"transactionType"`
	SecuritiesOwned         float64  `json:"securitiesOwned"`
	CompanyCik              string   `json:"companyCik"`
	ReportingName           string   `json:"reportingName"`
	AcquistionOrDisposition string   `json:"acquistionOrDisposition"`
	FormType                string   `json:"formType"`
	TypeOfOwner             string   `json:"typeOfOwner"`
	SecuritiesTransacted    float64  `json:"securitiesTransacted"`
	SecurityName            string   `json:"securityName"`
	Link                    string   `json:"link"`
	Price                   *float64 `json:"price"`
}

// InsiderTradingRSSFeed ...
type InsiderTradingRSSFeed struct {
	Title        string `json:"title"`
	FillingDate  string `json:"fillingDate"`
	Symbol       string `json:"symbol"`
	Link         string `json:"link"`
	IssuerCik    string `json:"issuerCik"`
	ReportingCik string `json:"reportingCik"`
}

// InsiderTradingMapperCikCompany ...
type InsiderTradingMapperCikCompany struct {
	Symbol     string `json:"symbol"`
	CompanyCik string `json:"companyCik"`
}

// InsiderTradingMapperCikName ...
type InsiderTradingMapperCikName struct {
	ReportingCik  string `json:"reportingCik"`
	ReportingName string `json:"reportingName"`
}
