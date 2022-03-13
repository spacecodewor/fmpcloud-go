package objects

// EconomicsMarketRisk ...
type EconomicsMarketRisk struct {
	Country                string  `json:"country"`
	Continent              string  `json:"continent"`
	TotalEquityRiskPremium float64 `json:"totalEquityRiskPremium"`
	CountryRiskPremium     float64 `json:"countryRiskPremium"`
}

// EconomicsTreasuryRates ...
type EconomicsTreasuryRates struct {
	Date   string  `json:"date"`
	Month1 float64 `json:"month1"`
	Month2 float64 `json:"month2"`
	Month3 float64 `json:"month3"`
	Month6 float64 `json:"month6"`
	Year1  float64 `json:"year1"`
	Year2  float64 `json:"year2"`
	Year3  float64 `json:"year3"`
	Year5  float64 `json:"year5"`
	Year7  float64 `json:"year7"`
	Year10 float64 `json:"year10"`
	Year20 float64 `json:"year20"`
	Year30 float64 `json:"year30"`
}

// EconomicsIndicator ...
type EconomicsIndicator struct {
	Date  string  `json:"date"`
	Value float64 `json:"value"`
}
