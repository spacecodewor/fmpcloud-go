package objects

import "time"

// Base const type ETFSector, CompanyValuationPeriod, CompanyValuationHistory ...
const (
	ETFSectorHealthcare                        ETFSector = "Healthcare"
	ETFTelecommunicationsServices              ETFSector = "Telecommunications Services"
	ETFSectorEnergy                            ETFSector = "Energy"
	ETFSectorBasicMaterials                    ETFSector = "Basic Materials"
	ETFSectorConsumerCyclicals                 ETFSector = "Consumer Cyclicals"
	ETFSectorTechnology                        ETFSector = "Technology"
	ETFSectorFinancials                        ETFSector = "Financials"
	ETFSectorUtilities                         ETFSector = "Utilities"
	ETFSectorConsumerNonCyclicals              ETFSector = "Consumer Non-Cyclicals"
	ETFSectorIndustrials                       ETFSector = "Industrials"
	ETFSectorBeverages                         ETFSector = "Beverages"
	ETFSectorSoftwareITServices                ETFSector = "Software IT Services"
	ETFSectorCommunicationsNetworking          ETFSector = "Communications Networking"
	ETFSectorRenewableEnergyEquipmentServices  ETFSector = "Renewable Energy Equipment Services"
	ETFSectorRenewableFuels                    ETFSector = "Renewable Fuels"
	ETFSectorElectricUtilities                 ETFSector = "Electric Utilities"
	ETFSectorMultilineUtilities                ETFSector = "Multiline Utilities"
	ETFSectorIndependentPowerProducers         ETFSector = "Independent Power Producers"
	ETFSectorFoodTobacco                       ETFSector = "Food Tobacco"
	ETFSectorFoodDrugRetailing                 ETFSector = "Food Drug Retailing"
	ETFSectorDiversifiedRetail                 ETFSector = "Diversified Retail"
	ETFSectorPersonalHouseholdProductsServices ETFSector = "Personal Household Products Services"
	ETFSectorOilGasExplorationAndProduction    ETFSector = "Oil Gas Exploration and Production"
	ETFSectorOilGasRefiningAndMarketing        ETFSector = "Oil Gas Refining and Marketing"
	ETFSectorOilGasTransportationServices      ETFSector = "Oil Gas Transportation Services"
	ETFSectorIntegratedOilGas                  ETFSector = "Integrated Oil Gas"
	ETFSectorLeisureProducts                   ETFSector = "Leisure Products"
	ETFSectorMediaPublishing                   ETFSector = "Media Publishing"
	ETFSectorProfessionalCommercialServices    ETFSector = "Professional Commercial Services"
	ETFSectorHotelsEntertainmentServices       ETFSector = "Hotels Entertainment Services"
	ETFSectorResidentialCommercialREITs        ETFSector = "Residential Commercial REITs"
	ETFSectorElectricUtilitiesIPPs             ETFSector = "Electric Utilities IPPs"
	ETFSectorNaturalGasUtilities               ETFSector = "Natural Gas Utilities"
	ETFSectorWaterUtilities                    ETFSector = "Water Utilities"

	CompanyValuationPeriodAnnual  CompanyValuationPeriod = "annual"
	CompanyValuationPeriodQuarter CompanyValuationPeriod = "quarter"

	CompanyValuationHistoryDaily CompanyValuationHistory = "daily"
	CompanyValuationHistoryToday CompanyValuationHistory = "today"

	GradeTypeBuy        GradeType = "Buy"
	GradeTypeOutperform GradeType = "Outperform"
	GradeTypeNeutral    GradeType = "Neutral"
	GradeTypeOverweight GradeType = "Overweight"
	GradeTypeReduce     GradeType = "Reduce"
	GradeTypeHold       GradeType = "Hold"

	StockScreenerSectorConsumerCyclical      StockScreenerSector = "Consumer Cyclical"
	StockScreenerSectorEnergy                StockScreenerSector = "Energy"
	StockScreenerSectorTechnology            StockScreenerSector = "Technology"
	StockScreenerSectorIndustrialsl          StockScreenerSector = "Industrials"
	StockScreenerSectorFinancialServices     StockScreenerSector = "Financial Services"
	StockScreenerSectorBasicMaterials        StockScreenerSector = "Basic Materials"
	StockScreenerSectorCommunicationServices StockScreenerSector = "Communication Services"
	StockScreenerSectorConsumerDefensive     StockScreenerSector = "Consumer Defensive"
	StockScreenerSectorHealthcare            StockScreenerSector = "Healthcare"
	StockScreenerSectorRealEstate            StockScreenerSector = "Real Estate"
	StockScreenerSectorUtilities             StockScreenerSector = "Utilities"
	StockScreenerSectorIndustrialGoods       StockScreenerSector = "Industrial Goods"
	StockScreenerSectorFinancial             StockScreenerSector = "Financial"
	StockScreenerSectorServices              StockScreenerSector = "Services"
	StockScreenerSectorConglomerates         StockScreenerSector = "Conglomerates"

	StockScreenerIndustryAutos                 StockScreenerIndustry = "Autos"
	StockScreenerIndustryBanks                 StockScreenerIndustry = "Banks"
	StockScreenerIndustryBanksDiversified      StockScreenerIndustry = "Banks Diversified"
	StockScreenerIndustrySoftware              StockScreenerIndustry = "Software"
	StockScreenerIndustryBanksRegional         StockScreenerIndustry = "Banks Regional"
	StockScreenerIndustryBeveragesAlcoholic    StockScreenerIndustry = "Beverages Alcoholic"
	StockScreenerIndustryBeveragesBrewers      StockScreenerIndustry = "Beverages Brewers"
	StockScreenerIndustryBeveragesNonAlcoholic StockScreenerIndustry = "Beverages Non-Alcoholic"

	StockScreenerExchangeNYSE       StockScreenerExchange = "nyse"
	StockScreenerExchangeNasdaq     StockScreenerExchange = "nasdaq"
	StockScreenerExchangeAmex       StockScreenerExchange = "amex"
	StockScreenerExchangeEuroNext   StockScreenerExchange = "euronext"
	StockScreenerExchangeTSX        StockScreenerExchange = "tsx"
	StockScreenerExchangeETF        StockScreenerExchange = "etf"
	StockScreenerExchangeMutualFund StockScreenerExchange = "mutual_fund"
)

// ETFSector ...
type ETFSector string

// StockScreenerSector ...
type StockScreenerSector string

// StockScreenerIndustry ...
type StockScreenerIndustry string

// StockScreenerExchange ...
type StockScreenerExchange string

// CompanyValuationPeriod ...
type CompanyValuationPeriod string

// CompanyValuationHistory ...
type CompanyValuationHistory string

// GradeType ...
type GradeType string

// RequestIncomeStatement ...
type RequestIncomeStatement struct {
	Symbol string
	Period CompanyValuationPeriod
	Limit  int64
}

// RequestIncomeStatementGrowth ...
type RequestIncomeStatementGrowth struct {
	Symbol string
	Period CompanyValuationPeriod
	Limit  int64
}

// RequestBalanceSheetStatement ...
type RequestBalanceSheetStatement struct {
	Symbol string
	Period CompanyValuationPeriod
	Limit  int64
}

// RequestBalanceSheetStatementGrowth ...
type RequestBalanceSheetStatementGrowth struct {
	Symbol string
	Period CompanyValuationPeriod
	Limit  int64
}

// RequestCashFlowStatement ...
type RequestCashFlowStatement struct {
	Symbol string
	Period CompanyValuationPeriod
	Limit  int64
}

// RequestCashFlowStatementGrowth ...
type RequestCashFlowStatementGrowth struct {
	Symbol string
	Period CompanyValuationPeriod
	Limit  int64
}

// RequestInternationalFiling ...
type RequestInternationalFiling struct {
	Symbol string
	Period CompanyValuationPeriod
	Limit  int64
}

// RequestIncomeStatementAsReported ...
type RequestIncomeStatementAsReported struct {
	Symbol string
	Period CompanyValuationPeriod
	Limit  int64
}

// RequestBalanceSheetStatementAsReported ...
type RequestBalanceSheetStatementAsReported struct {
	Symbol string
	Period CompanyValuationPeriod
	Limit  int64
}

// RequestCashFlowStatementAsReported ...
type RequestCashFlowStatementAsReported struct {
	Symbol string
	Period CompanyValuationPeriod
	Limit  int64
}

// RequestFullFinancialStatementAsReported ...
type RequestFullFinancialStatementAsReported struct {
	Symbol string
	Period CompanyValuationPeriod
	Limit  int64
}

// RequestFinancialRatios ...
type RequestFinancialRatios struct {
	Symbol string
	Period CompanyValuationPeriod
	Limit  int64
}

// RequestKeyMetrics ...
type RequestKeyMetrics struct {
	Symbol string
	Period CompanyValuationPeriod
	Limit  int64
}

// RequestEnterpriseValue ...
type RequestEnterpriseValue struct {
	Symbol string
	Period CompanyValuationPeriod
	Limit  int64
}

// RequestFinancialStatementsGrowth ...
type RequestFinancialStatementsGrowth struct {
	Symbol string
	Period CompanyValuationPeriod
	Limit  int64
}

// RequestDailyDiscountedCashFlow ...
type RequestDailyDiscountedCashFlow struct {
	Symbol string
	Limit  int64
}

// RequestRating ...
type RequestRating struct {
	Symbol string
	Limit  int64
}

// RequestMarketCapitalization ...
type RequestMarketCapitalization struct {
	Symbol string
	Limit  int64
}

// RequestHistoryDiscountedCashFlow ...
type RequestHistoryDiscountedCashFlow struct {
	Symbol string
	Period CompanyValuationPeriod
	Limit  int64
}

// RequestGrade ...
type RequestGrade struct {
	Symbol string
	Limit  int64
}

// RequestAnalystStockRecommendations ...
type RequestAnalystStockRecommendations struct {
	Symbol string
	Limit  int64
}

// RequestPressReleases ...
type RequestPressReleases struct {
	Symbol string
	Limit  int64
}

// RequestStockNews ...
type RequestStockNews struct {
	SymbolList []string
	Limit      int64
}

// RequestStockScreener ...
type RequestStockScreener struct {
	Exchange           []string // Example: StockScreenerExchangeNYSE
	Country            *string
	Sector             *StockScreenerSector
	Industry           *StockScreenerIndustry
	Limit              int64
	MarketCapMoreThan  *int64
	MarketCapLowerThan *int64
	VolumeMoreThan     *int64
	VolumeLowerThan    *int64
	BetaMoreThan       *float64
	BetaLowerThan      *float64
	DividendMoreThan   *float64
	DividendLowerThan  *float64
	PriceMoreThan      *float64
	PriceLowerThan     *float64
	IsETF              *bool
	IsActivelyTrading  *bool
}

// RequestEarningCallTranscript ...
type RequestEarningCallTranscript struct {
	Symbol  string
	Quarter int64
	Year    int64
}

// RequestAnalystEstimates ...
type RequestAnalystEstimates struct {
	Symbol string
	Period CompanyValuationPeriod
}

// RequestEconomicCalendar ...
type RequestEconomicCalendar struct {
	From *time.Time
	To   *time.Time
}

// RequestSECFilings ...
type RequestSECFilings struct {
	Symbol string
	Type   *string
	Limit  int64
}

// RequestHistoryEconomicCalendar ...
type RequestHistoryEconomicCalendar struct {
	Event   string
	Country string
}

// RssFeed ...
type RssFeed struct {
	Title    string `json:"title"`
	Date     string `json:"date"` // 2020-09-14 17:27:24
	Link     string `json:"link"`
	Cik      string `json:"cik"`
	FormType string `json:"form_type"`
	Ticker   string `json:"ticker"`
}

// ETF ...
type ETF struct {
	Symbol            string  `json:"symbol"`
	Name              string  `json:"name"`
	Price             float64 `json:"price"`
	Exchange          string  `json:"exchange"`
	ExchangeShortName string  `json:"exchangeShortName"`
}

// AvailableTraded ...
type AvailableTraded struct {
	Symbol            string  `json:"symbol"`
	Name              string  `json:"name"`
	Price             float64 `json:"price"`
	Exchange          string  `json:"exchange"`
	ExchangeShortName string  `json:"exchangeShortName"`
	Type              string  `json:"type"`
}

// IPOCalendar ...
type IPOCalendar struct {
	Date       string  `json:"date"` // 2020-09-01
	Company    string  `json:"company"`
	Symbol     string  `json:"symbol"`
	Exchange   string  `json:"exchange"`
	Actions    string  `json:"actions"`
	Shares     int64   `json:"shares"`
	PriceRange string  `json:"priceRange"`
	MarketCap  float64 `json:"marketCap"`
}

// EarningCalendar ...
type EarningCalendar struct {
	Date             string  `json:"date"` // 2020-10-28
	Symbol           string  `json:"symbol"`
	Eps              float64 `json:"eps"`
	EpsEstimated     float64 `json:"epsEstimated"`
	Time             string  `json:"time"` // Indicates whether the earnings is announced before market open(bmo), after market close(amc), or during market hour(dmh).
	Revenue          float64 `json:"revenue"`
	RevenueEstimated float64 `json:"revenueEstimated"`
	UpdatedFromDate  string  `json:"updatedFromDate"`
	FiscalDateEnding string  `json:"fiscalDateEnding"`
}

// EarningCalendarConfirmed ...
type EarningCalendarConfirmed struct {
	Symbol          string `json:"symbol"`
	Exchange        string `json:"exchange"`
	Time            string `json:"time"`
	When            string `json:"when"`
	Date            string `json:"date"`
	PublicationDate string `json:"publicationDate"`
	Title           string `json:"title"`
	URL             string `json:"url"`
}

// EarningCallTranscript ...
type EarningCallTranscript struct {
	Symbol  string `json:"symbol"`
	Quarter int    `json:"quarter"`
	Year    int    `json:"year"`
	Date    string `json:"date"` // 2020-07-31 17:00:00
	Content string `json:"content"`
}

// SplitCalendar ...
type SplitCalendar struct {
	Date        string  `json:"date"` // 2020-09-10
	Label       string  `json:"label"`
	Symbol      string  `json:"symbol"`
	Numerator   float64 `json:"numerator"`
	Denominator float64 `json:"denominator"`
}

// DividendCalendar ...
type DividendCalendar struct {
	Date            string  `json:"date"`  // 2020-09-10
	Label           string  `json:"label"` // September 10, 20
	AdjDividend     float64 `json:"adjDividend"`
	Symbol          string  `json:"symbol"`
	Dividend        float64 `json:"dividend"`
	RecordDate      string  `json:"recordDate"`
	PaymentDate     string  `json:"paymentDate"`
	DeclarationDate string  `json:"declarationDate"`
}

// InstitutionalHolder ...
type InstitutionalHolder struct {
	Holder       string `json:"holder"`
	Shares       int    `json:"shares"`
	DateReported string `json:"dateReported"`
	Change       int    `json:"change"`
}

// MutualFundHolder ...
type MutualFundHolder struct {
	Holder        string  `json:"holder"`
	Shares        int64   `json:"shares"`
	DateReported  string  `json:"dateReported"`
	Change        int64   `json:"change"`
	WeightPercent float64 `json:"weightPercent"`
}

// ETFHolder ...
type ETFHolder struct {
	Asset            string  `json:"asset"`
	SharesNumber     int64   `json:"sharesNumber"`
	WeightPercentage float64 `json:"weightPercentage"`
}

// ETFStockExposure ...
type ETFStockExposure struct {
	ETFSymbol        string  `json:"etfSymbol"`
	AssetExposure    string  `json:"assetExposure"`
	SharesNumber     int     `json:"sharesNumber"`
	WeightPercentage float64 `json:"weightPercentage"`
	MarketValue      float64 `json:"marketValue"`
}

// ETFSectorWeighting ...
type ETFSectorWeighting struct {
	Sector           ETFSector `json:"sector"`
	WeightPercentage string    `json:"weightPercentage"`
}

// ETFCountryWeighting ...
type ETFCountryWeighting struct {
	Country          string `json:"country"`
	WeightPercentage string `json:"weightPercentage"`
}

// IncomeStatement ...
type IncomeStatement struct {
	Date                                    string  `json:"date" csv:"date"`
	Symbol                                  string  `json:"symbol" csv:"symbol"`
	FillingDate                             string  `json:"fillingDate" csv:"fillingDate"`
	Cik                                     string  `json:"cik" csv:"cik"`
	CalendarYear                            string  `json:"calendarYear" csv:"calendarYear"`
	ReportedCurrency                        string  `json:"reportedCurrency" csv:"reportedCurrency"`
	AcceptedDate                            string  `json:"acceptedDate" csv:"acceptedDate"`
	Period                                  string  `json:"period" csv:"period"`
	Revenue                                 float64 `json:"revenue" csv:"revenue"`
	CostOfRevenue                           float64 `json:"costOfRevenue" csv:"costOfRevenue"`
	GrossProfit                             float64 `json:"grossProfit" csv:"grossProfit"`
	GrossProfitRatio                        float64 `json:"grossProfitRatio" csv:"grossProfitRatio"`
	ResearchAndDevelopmentExpenses          float64 `json:"researchAndDevelopmentExpenses" csv:"ResearchAndDevelopmentExpenses"`
	GeneralAndAdministrativeExpenses        float64 `json:"generalAndAdministrativeExpenses" csv:"GeneralAndAdministrativeExpenses"`
	SellingAndMarketingExpenses             float64 `json:"sellingAndMarketingExpenses" csv:"SellingAndMarketingExpenses"`
	SellingGeneralAndAdministrativeExpenses float64 `json:"sellingGeneralAndAdministrativeExpenses" csv:"sellingGeneralAndAdministrativeExpenses"`
	OtherExpenses                           float64 `json:"otherExpenses" csv:"otherExpenses"`
	OperatingExpenses                       float64 `json:"operatingExpenses" csv:"operatingExpenses"`
	CostAndExpenses                         float64 `json:"costAndExpenses" csv:"costAndExpenses"`
	InterestExpense                         float64 `json:"interestExpense" csv:"interestExpense"`
	DepreciationAndAmortization             float64 `json:"depreciationAndAmortization" csv:"depreciationAndAmortization"`
	Ebitda                                  float64 `json:"ebitda" csv:"EBITDA"`
	Ebitdaratio                             float64 `json:"ebitdaratio" csv:"EBITDARatio"`
	OperatingIncome                         float64 `json:"operatingIncome" csv:"operatingIncome"`
	OperatingIncomeRatio                    float64 `json:"operatingIncomeRatio" csv:"operatingIncomeRatio"`
	TotalOtherIncomeExpensesNet             float64 `json:"totalOtherIncomeExpensesNet" csv:"totalOtherIncomeExpensesNet"`
	IncomeBeforeTax                         float64 `json:"incomeBeforeTax" csv:"incomeBeforeTax"`
	IncomeBeforeTaxRatio                    float64 `json:"incomeBeforeTaxRatio" csv:"incomeBeforeTaxRatio"`
	IncomeTaxExpense                        float64 `json:"incomeTaxExpense" csv:"incomeTaxExpense"`
	NetIncome                               float64 `json:"netIncome" csv:"netIncome"`
	NetIncomeRatio                          float64 `json:"netIncomeRatio" csv:"netIncomeRatio"`
	Eps                                     float64 `json:"eps" csv:"EPS"`
	Epsdiluted                              float64 `json:"epsdiluted" csv:"EPSDiluted"`
	WeightedAverageShsOut                   float64 `json:"weightedAverageShsOut" csv:"weightedAverageShsOut"`
	WeightedAverageShsOutDil                float64 `json:"weightedAverageShsOutDil" csv:"weightedAverageShsOutDil"`
	Link                                    string  `json:"link" csv:"link"`
	FinalLink                               string  `json:"finalLink" csv:"finalLink"`
}

// IncomeStatementGrowth ...
type IncomeStatementGrowth struct {
	Date                                   string  `json:"date"`
	Symbol                                 string  `json:"symbol"`
	Period                                 string  `json:"period"`
	GrowthRevenue                          float64 `json:"growthRevenue"`
	ReportedCurrency                       string  `json:"reportedCurrency"`
	GrowthCostOfRevenue                    float64 `json:"growthCostOfRevenue"`
	GrowthGrossProfit                      float64 `json:"growthGrossProfit"`
	GrowthGrossProfitRatio                 float64 `json:"growthGrossProfitRatio"`
	GrowthResearchAndDevelopmentExpenses   float64 `json:"growthResearchAndDevelopmentExpenses"`
	GrowthGeneralAndAdministrativeExpenses float64 `json:"growthGeneralAndAdministrativeExpenses"`
	GrowthSellingAndMarketingExpenses      float64 `json:"growthSellingAndMarketingExpenses"`
	GrowthOtherExpenses                    float64 `json:"growthOtherExpenses"`
	GrowthOperatingExpenses                float64 `json:"growthOperatingExpenses"`
	GrowthCostAndExpenses                  float64 `json:"growthCostAndExpenses"`
	GrowthInterestExpense                  float64 `json:"growthInterestExpense"`
	GrowthDepreciationAndAmortization      float64 `json:"growthDepreciationAndAmortization"`
	GrowthEBITDA                           float64 `json:"growthEBITDA"`
	GrowthEBITDARatio                      float64 `json:"growthEBITDARatio"`
	GrowthOperatingIncome                  float64 `json:"growthOperatingIncome"`
	GrowthOperatingIncomeRatio             float64 `json:"growthOperatingIncomeRatio"`
	GrowthTotalOtherIncomeExpensesNet      float64 `json:"growthTotalOtherIncomeExpensesNet"`
	GrowthIncomeBeforeTax                  float64 `json:"growthIncomeBeforeTax"`
	GrowthIncomeBeforeTaxRatio             float64 `json:"growthIncomeBeforeTaxRatio"`
	GrowthIncomeTaxExpense                 float64 `json:"growthIncomeTaxExpense"`
	GrowthNetIncome                        float64 `json:"growthNetIncome"`
	GrowthNetIncomeRatio                   float64 `json:"growthNetIncomeRatio"`
	GrowthEPS                              float64 `json:"growthEPS"`
	GrowthEPSDiluted                       float64 `json:"growthEPSDiluted"`
	GrowthWeightedAverageShsOut            float64 `json:"growthWeightedAverageShsOut"`
	GrowthWeightedAverageShsOutDil         float64 `json:"growthWeightedAverageShsOutDil"`
}

// BalanceSheetStatement ...
type BalanceSheetStatement struct {
	Date                                    string  `json:"date" csv:"date"`
	Symbol                                  string  `json:"symbol" csv:"symbol"`
	ReportedCurrency                        string  `json:"reportedCurrency" csv:"reportedCurrency"`
	FillingDate                             string  `json:"fillingDate" csv:"fillingDate"`
	AcceptedDate                            string  `json:"acceptedDate" csv:"acceptedDate"`
	Period                                  string  `json:"period" csv:"period"`
	CashAndCashEquivalents                  float64 `json:"cashAndCashEquivalents" csv:"cashAndCashEquivalents"`
	ShortTermInvestments                    float64 `json:"shortTermInvestments" csv:"shortTermInvestments"`
	CashAndShortTermInvestments             float64 `json:"cashAndShortTermInvestments" csv:"cashAndShortTermInvestments"`
	NetReceivables                          float64 `json:"netReceivables" csv:"netReceivables"`
	Inventory                               float64 `json:"inventory" csv:"inventory"`
	OtherCurrentAssets                      float64 `json:"otherCurrentAssets" csv:"otherCurrentAssets"`
	TotalCurrentAssets                      float64 `json:"totalCurrentAssets" csv:"totalCurrentAssets"`
	PropertyPlantEquipmentNet               float64 `json:"propertyPlantEquipmentNet" csv:"propertyPlantEquipmentNet"`
	Goodwill                                float64 `json:"goodwill" csv:"goodwill"`
	IntangibleAssets                        float64 `json:"intangibleAssets" csv:"intangibleAssets"`
	GoodwillAndIntangibleAssets             float64 `json:"goodwillAndIntangibleAssets" csv:"goodwillAndIntangibleAssets"`
	LongTermInvestments                     float64 `json:"longTermInvestments" csv:"longTermInvestments"`
	TaxAssets                               float64 `json:"taxAssets" csv:"taxAssets"`
	OtherNonCurrentAssets                   float64 `json:"otherNonCurrentAssets" csv:"otherNonCurrentAssets"`
	TotalNonCurrentAssets                   float64 `json:"totalNonCurrentAssets" csv:"totalNonCurrentAssets"`
	OtherAssets                             float64 `json:"otherAssets" csv:"otherAssets"`
	TotalAssets                             float64 `json:"totalAssets" csv:"totalAssets"`
	AccountPayables                         float64 `json:"accountPayables" csv:"accountPayables"`
	ShortTermDebt                           float64 `json:"shortTermDebt" csv:"shortTermDebt"`
	TaxPayables                             float64 `json:"taxPayables" csv:"taxPayables"`
	DeferredRevenue                         float64 `json:"deferredRevenue" csv:"deferredRevenue"`
	OtherCurrentLiabilities                 float64 `json:"otherCurrentLiabilities" csv:"otherCurrentLiabilities"`
	TotalCurrentLiabilities                 float64 `json:"totalCurrentLiabilities" csv:"totalCurrentLiabilities"`
	LongTermDebt                            float64 `json:"longTermDebt" csv:"longTermDebt"`
	DeferredRevenueNonCurrent               float64 `json:"deferredRevenueNonCurrent" csv:"deferredRevenueNonCurrent"`
	DeferredTaxLiabilitiesNonCurrent        float64 `json:"deferredTaxLiabilitiesNonCurrent" csv:"deferredTaxLiabilitiesNonCurrent"`
	OtherNonCurrentLiabilities              float64 `json:"otherNonCurrentLiabilities" csv:"otherNonCurrentLiabilities"`
	TotalNonCurrentLiabilities              float64 `json:"totalNonCurrentLiabilities" csv:"totalNonCurrentLiabilities"`
	OtherLiabilities                        float64 `json:"otherLiabilities" csv:"otherLiabilities"`
	TotalLiabilities                        float64 `json:"totalLiabilities" csv:"totalLiabilities"`
	CommonStock                             float64 `json:"commonStock" csv:"commonStock"`
	RetainedEarnings                        float64 `json:"retainedEarnings" csv:"retainedEarnings"`
	AccumulatedOtherComprehensiveIncomeLoss float64 `json:"accumulatedOtherComprehensiveIncomeLoss" csv:"accumulatedOtherComprehensiveIncomeLoss"`
	OthertotalStockholdersEquity            float64 `json:"othertotalStockholdersEquity" csv:"othertotalStockholdersEquity"`
	TotalStockholdersEquity                 float64 `json:"totalStockholdersEquity" csv:"totalStockholdersEquity"`
	TotalLiabilitiesAndStockholdersEquity   float64 `json:"totalLiabilitiesAndStockholdersEquity" csv:"totalLiabilitiesAndStockholdersEquity"`
	TotalInvestments                        float64 `json:"totalInvestments" csv:"totalInvestments"`
	TotalDebt                               float64 `json:"totalDebt" csv:"totalDebt"`
	NetDebt                                 float64 `json:"netDebt" csv:"netDebt"`
	Link                                    string  `json:"link" csv:"link"`
	FinalLink                               string  `json:"finalLink" csv:"finalLink"`
}

// BalanceSheetStatementGrowth ...
type BalanceSheetStatementGrowth struct {
	Date                                          string  `json:"date"`
	Symbol                                        string  `json:"symbol"`
	Period                                        string  `json:"period"`
	ReportedCurrency                              string  `json:"reportedCurrency"`
	GrowthCashAndCashEquivalents                  float64 `json:"growthCashAndCashEquivalents"`
	GrowthShortTermInvestments                    float64 `json:"growthShortTermInvestments"`
	GrowthCashAndShortTermInvestments             float64 `json:"growthCashAndShortTermInvestments"`
	GrowthNetReceivables                          float64 `json:"growthNetReceivables"`
	GrowthInventory                               float64 `json:"growthInventory"`
	GrowthOtherCurrentAssets                      float64 `json:"growthOtherCurrentAssets"`
	GrowthTotalCurrentAssets                      float64 `json:"growthTotalCurrentAssets"`
	GrowthPropertyPlantEquipmentNet               float64 `json:"growthPropertyPlantEquipmentNet"`
	GrowthGoodwill                                float64 `json:"growthGoodwill"`
	GrowthIntangibleAssets                        float64 `json:"growthIntangibleAssets"`
	GrowthGoodwillAndIntangibleAssets             float64 `json:"growthGoodwillAndIntangibleAssets"`
	GrowthLongTermInvestments                     float64 `json:"growthLongTermInvestments"`
	GrowthTaxAssets                               float64 `json:"growthTaxAssets"`
	GrowthOtherNonCurrentAssets                   float64 `json:"growthOtherNonCurrentAssets"`
	GrowthTotalNonCurrentAssets                   float64 `json:"growthTotalNonCurrentAssets"`
	GrowthOtherAssets                             float64 `json:"growthOtherAssets"`
	GrowthTotalAssets                             float64 `json:"growthTotalAssets"`
	GrowthAccountPayables                         float64 `json:"growthAccountPayables"`
	GrowthShortTermDebt                           float64 `json:"growthShortTermDebt"`
	GrowthTaxPayables                             float64 `json:"growthTaxPayables"`
	GrowthDeferredRevenue                         float64 `json:"growthDeferredRevenue"`
	GrowthOtherCurrentLiabilities                 float64 `json:"growthOtherCurrentLiabilities"`
	GrowthTotalCurrentLiabilities                 float64 `json:"growthTotalCurrentLiabilities"`
	GrowthLongTermDebt                            float64 `json:"growthLongTermDebt"`
	GrowthDeferredRevenueNonCurrent               float64 `json:"growthDeferredRevenueNonCurrent"`
	GrowthDeferrredTaxLiabilitiesNonCurrent       float64 `json:"growthDeferrredTaxLiabilitiesNonCurrent"`
	GrowthOtherNonCurrentLiabilities              float64 `json:"growthOtherNonCurrentLiabilities"`
	GrowthTotalNonCurrentLiabilities              float64 `json:"growthTotalNonCurrentLiabilities"`
	GrowthOtherLiabilities                        float64 `json:"growthOtherLiabilities"`
	GrowthTotalLiabilities                        float64 `json:"growthTotalLiabilities"`
	GrowthCommonStock                             float64 `json:"growthCommonStock"`
	GrowthRetainedEarnings                        float64 `json:"growthRetainedEarnings"`
	GrowthAccumulatedOtherComprehensiveIncomeLoss float64 `json:"growthAccumulatedOtherComprehensiveIncomeLoss"`
	GrowthOthertotalStockholdersEquity            float64 `json:"growthOthertotalStockholdersEquity"`
	GrowthTotalStockholdersEquity                 float64 `json:"growthTotalStockholdersEquity"`
	GrowthTotalLiabilitiesAndStockholdersEquity   float64 `json:"growthTotalLiabilitiesAndStockholdersEquity"`
	GrowthTotalInvestments                        float64 `json:"growthTotalInvestments"`
	GrowthTotalDebt                               float64 `json:"growthTotalDebt"`
	GrowthNetDebt                                 float64 `json:"growthNetDebt"`
}

// CashFlowStatement ...
type CashFlowStatement struct {
	Date                                     string  `json:"date" csv:"date"`
	Symbol                                   string  `json:"symbol" csv:"symbol"`
	FillingDate                              string  `json:"fillingDate" csv:"fillingDate"`
	AcceptedDate                             string  `json:"acceptedDate" csv:"acceptedDate"`
	Period                                   string  `json:"period" csv:"period"`
	ReportedCurrency                         string  `json:"reportedCurrency" csv:"reportedCurrency"`
	NetIncome                                float64 `json:"netIncome" csv:"netIncome"`
	DepreciationAndAmortization              float64 `json:"depreciationAndAmortization" csv:"depreciationAndAmortization"`
	DeferredIncomeTax                        float64 `json:"deferredIncomeTax" csv:"deferredIncomeTax"`
	StockBasedCompensation                   float64 `json:"stockBasedCompensation" csv:"stockBasedCompensation"`
	ChangeInWorkingCapital                   float64 `json:"changeInWorkingCapital" csv:"changeInWorkingCapital"`
	AccountsReceivables                      float64 `json:"accountsReceivables" csv:"accountsReceivables"`
	Inventory                                float64 `json:"inventory" csv:"inventory"`
	AccountsPayables                         float64 `json:"accountsPayables" csv:"accountsPayables"`
	OtherWorkingCapital                      float64 `json:"otherWorkingCapital" csv:"otherWorkingCapital"`
	OtherNonCashItems                        float64 `json:"otherNonCashItems" csv:"otherNonCashItems"`
	NetCashProvidedByOperatingActivities     float64 `json:"netCashProvidedByOperatingActivities" csv:"netCashProvidedByOperatingActivities"`
	InvestmentsInPropertyPlantAndEquipment   float64 `json:"investmentsInPropertyPlantAndEquipment" csv:"investmentsInPropertyPlantAndEquipment"`
	AcquisitionsNet                          float64 `json:"acquisitionsNet" csv:"acquisitionsNet"`
	PurchasesOfInvestments                   float64 `json:"purchasesOfInvestments" csv:"purchasesOfInvestments"`
	SalesMaturitiesOfInvestments             float64 `json:"salesMaturitiesOfInvestments" csv:"salesMaturitiesOfInvestments"`
	OtherInvestingActivites                  float64 `json:"otherInvestingActivites" csv:"otherInvestingActivites"`
	NetCashUsedForInvestingActivites         float64 `json:"netCashUsedForInvestingActivites" csv:"netCashUsedForInvestingActivites"`
	DebtRepayment                            float64 `json:"debtRepayment" csv:"debtRepayment"`
	CommonStockIssued                        float64 `json:"commonStockIssued" csv:"commonStockIssued"`
	CommonStockRepurchased                   float64 `json:"commonStockRepurchased" csv:"commonStockRepurchased"`
	DividendsPaid                            float64 `json:"dividendsPaid" csv:"dividendsPaid"`
	OtherFinancingActivites                  float64 `json:"otherFinancingActivites" csv:"otherFinancingActivites"`
	NetCashUsedProvidedByFinancingActivities float64 `json:"netCashUsedProvidedByFinancingActivities" csv:"netCashUsedProvidedByFinancingActivities"`
	EffectOfForexChangesOnCash               float64 `json:"effectOfForexChangesOnCash" csv:"effectOfForexChangesOnCash"`
	NetChangeInCash                          float64 `json:"netChangeInCash" csv:"netChangeInCash"`
	CashAtEndOfPeriod                        float64 `json:"cashAtEndOfPeriod" csv:"cashAtEndOfPeriod"`
	CashAtBeginningOfPeriod                  float64 `json:"cashAtBeginningOfPeriod" csv:"cashAtBeginningOfPeriod"`
	OperatingCashFlow                        float64 `json:"operatingCashFlow" csv:"operatingCashFlow"`
	CapitalExpenditure                       float64 `json:"capitalExpenditure" csv:"capitalExpenditure"`
	FreeCashFlow                             float64 `json:"freeCashFlow" csv:"freeCashFlow"`
	Link                                     string  `json:"link" csv:"link"`
	FinalLink                                string  `json:"finalLink" csv:"finalLink"`
}

// CashFlowStatementGrowth ...
type CashFlowStatementGrowth struct {
	Date                                           string  `json:"date"`
	Symbol                                         string  `json:"symbol"`
	Period                                         string  `json:"period"`
	ReportedCurrency                               string  `json:"reportedCurrency"`
	GrowthNetIncome                                float64 `json:"growthNetIncome"`
	GrowthDepreciationAndAmortization              float64 `json:"growthDepreciationAndAmortization"`
	GrowthDeferredIncomeTax                        float64 `json:"growthDeferredIncomeTax"`
	GrowthStockBasedCompensation                   float64 `json:"growthStockBasedCompensation"`
	GrowthChangeInWorkingCapital                   float64 `json:"growthChangeInWorkingCapital"`
	GrowthAccountsReceivables                      float64 `json:"growthAccountsReceivables"`
	GrowthInventory                                float64 `json:"growthInventory"`
	GrowthAccountsPayables                         float64 `json:"growthAccountsPayables"`
	GrowthOtherWorkingCapital                      float64 `json:"growthOtherWorkingCapital"`
	GrowthOtherNonCashItems                        float64 `json:"growthOtherNonCashItems"`
	GrowthNetCashProvidedByOperatingActivites      float64 `json:"growthNetCashProvidedByOperatingActivites"`
	GrowthInvestmentsInPropertyPlantAndEquipment   float64 `json:"growthInvestmentsInPropertyPlantAndEquipment"`
	GrowthAcquisitionsNet                          float64 `json:"growthAcquisitionsNet"`
	GrowthPurchasesOfInvestments                   float64 `json:"growthPurchasesOfInvestments"`
	GrowthSalesMaturitiesOfInvestments             float64 `json:"growthSalesMaturitiesOfInvestments"`
	GrowthOtherInvestingActivites                  float64 `json:"growthOtherInvestingActivites"`
	GrowthNetCashUsedForInvestingActivites         float64 `json:"growthNetCashUsedForInvestingActivites"`
	GrowthDebtRepayment                            float64 `json:"growthDebtRepayment"`
	GrowthCommonStockIssued                        float64 `json:"growthCommonStockIssued"`
	GrowthCommonStockRepurchased                   float64 `json:"growthCommonStockRepurchased"`
	GrowthDividendsPaid                            float64 `json:"growthDividendsPaid"`
	GrowthOtherFinancingActivites                  float64 `json:"growthOtherFinancingActivites"`
	GrowthNetCashUsedProvidedByFinancingActivities float64 `json:"growthNetCashUsedProvidedByFinancingActivities"`
	GrowthEffectOfForexChangesOnCash               float64 `json:"growthEffectOfForexChangesOnCash"`
	GrowthNetChangeInCash                          float64 `json:"growthNetChangeInCash"`
	GrowthCashAtEndOfPeriod                        float64 `json:"growthCashAtEndOfPeriod"`
	GrowthCashAtBeginningOfPeriod                  float64 `json:"growthCashAtBeginningOfPeriod"`
	GrowthOperatingCashFlow                        float64 `json:"growthOperatingCashFlow"`
	GrowthCapitalExpenditure                       float64 `json:"growthCapitalExpenditure"`
	GrowthFreeCashFlow                             float64 `json:"growthFreeCashFlow"`
}

// IncomeStatementAsReported ...
type IncomeStatementAsReported struct {
	Date                                                                                        string      `json:"date"`
	Symbol                                                                                      string      `json:"symbol"`
	Period                                                                                      string      `json:"period"`
	Costofgoodsandservicessold                                                                  interface{} `json:"costofgoodsandservicessold"`
	Netincomeloss                                                                               float64     `json:"netincomeloss"`
	Researchanddevelopmentexpense                                                               float64     `json:"researchanddevelopmentexpense"`
	Grossprofit                                                                                 interface{} `json:"grossprofit"`
	Othercomprehensiveincomelossreclassificationadjustmentfromaociforsaleofsecuritiesnetoftax   float64     `json:"othercomprehensiveincomelossreclassificationadjustmentfromaociforsaleofsecuritiesnetoftax"`
	Othercomprehensiveincomelossavailableforsalesecuritiesadjustmentnetoftax                    float64     `json:"othercomprehensiveincomelossavailableforsalesecuritiesadjustmentnetoftax"`
	Othercomprehensiveincomelossderivativesqualifyingashedgesnetoftax                           float64     `json:"othercomprehensiveincomelossderivativesqualifyingashedgesnetoftax"`
	Othercomprehensiveincomelossforeigncurrencytransactionandtranslationadjustmentnetoftax      float64     `json:"othercomprehensiveincomelossforeigncurrencytransactionandtranslationadjustmentnetoftax"`
	Othercomprehensiveincomelossderivativeinstrumentgainlossbeforereclassificationaftertax      float64     `json:"othercomprehensiveincomelossderivativeinstrumentgainlossbeforereclassificationaftertax"`
	Weightedaveragenumberofdilutedsharesoutstanding                                             float64     `json:"weightedaveragenumberofdilutedsharesoutstanding"`
	Weightedaveragenumberofsharesoutstandingbasic                                               float64     `json:"weightedaveragenumberofsharesoutstandingbasic"`
	Othercomprehensiveincomeunrealizedgainlossonderivativesarisingduringperiodnetoftax          float64     `json:"othercomprehensiveincomeunrealizedgainlossonderivativesarisingduringperiodnetoftax"`
	Operatingincomeloss                                                                         float64     `json:"operatingincomeloss"`
	Othercomprehensiveincomelossreclassificationadjustmentfromaocionderivativesnetoftax         float64     `json:"othercomprehensiveincomelossreclassificationadjustmentfromaocionderivativesnetoftax"`
	Incomelossfromcontinuingoperationsbeforeincometaxesextraordinaryitemsnoncontrollinginterest float64     `json:"incomelossfromcontinuingoperationsbeforeincometaxesextraordinaryitemsnoncontrollinginterest"`
	Earningspersharebasic                                                                       float64     `json:"earningspersharebasic"`
	Incometaxexpensebenefit                                                                     float64     `json:"incometaxexpensebenefit"`
	Revenuefromcontractwithcustomerexcludingassessedtax                                         interface{} `json:"revenuefromcontractwithcustomerexcludingassessedtax"`
	Nonoperatingincomeexpense                                                                   float64     `json:"nonoperatingincomeexpense"`
	Operatingexpenses                                                                           interface{} `json:"operatingexpenses"`
	Earningspersharediluted                                                                     float64     `json:"earningspersharediluted"`
	Othercomprehensiveincomeunrealizedholdinggainlossonsecuritiesarisingduringperiodnetoftax    float64     `json:"othercomprehensiveincomeunrealizedholdinggainlossonsecuritiesarisingduringperiodnetoftax"`
	Sellinggeneralandadministrativeexpense                                                      float64     `json:"sellinggeneralandadministrativeexpense"`
	Othercomprehensiveincomelossnetoftaxportionattributabletoparent                             float64     `json:"othercomprehensiveincomelossnetoftaxportionattributabletoparent"`
	Comprehensiveincomenetoftax                                                                 float64     `json:"comprehensiveincomenetoftax"`
	Othercomprehensiveincomelossderivativeinstrumentgainlossafterreclassificationandtax         float64     `json:"othercomprehensiveincomelossderivativeinstrumentgainlossafterreclassificationandtax"`
	Othercomprehensiveincomelosscashflowhedgegainlossreclassificationaftertax                   float64     `json:"othercomprehensiveincomelosscashflowhedgegainlossreclassificationaftertax"`
}

// BalanceSheetStatementAsReported ...
type BalanceSheetStatementAsReported struct {
	Date                                            string      `json:"date"`
	Symbol                                          string      `json:"symbol"`
	Period                                          string      `json:"period"`
	Liabilitiesandstockholdersequity                interface{} `json:"liabilitiesandstockholdersequity"`
	Liabilities                                     interface{} `json:"liabilities"`
	Liabilitiescurrent                              interface{} `json:"liabilitiescurrent"`
	Commonstocksharesauthorized                     float64     `json:"commonstocksharesauthorized"`
	Cashandcashequivalentsatcarryingvalue           interface{} `json:"cashandcashequivalentsatcarryingvalue"`
	Retainedearningsaccumulateddeficit              interface{} `json:"retainedearningsaccumulateddeficit"`
	Liabilitiesnoncurrent                           interface{} `json:"liabilitiesnoncurrent"`
	Propertyplantandequipmentnet                    float64     `json:"propertyplantandequipmentnet"`
	Commonstocksincludingadditionalpaidincapital    interface{} `json:"commonstocksincludingadditionalpaidincapital"`
	Commercialpaper                                 float64     `json:"commercialpaper"`
	Longtermdebtcurrent                             float64     `json:"longtermdebtcurrent"`
	Commonstocksharesoutstanding                    float64     `json:"commonstocksharesoutstanding"`
	Otherliabilitiesnoncurrent                      float64     `json:"otherliabilitiesnoncurrent"`
	Marketablesecuritiescurrent                     float64     `json:"marketablesecuritiescurrent"`
	Otherliabilitiescurrent                         float64     `json:"otherliabilitiescurrent"`
	Assetscurrent                                   interface{} `json:"assetscurrent"`
	Longtermdebtnoncurrent                          float64     `json:"longtermdebtnoncurrent"`
	Contractwithcustomerliabilitycurrent            float64     `json:"contractwithcustomerliabilitycurrent"`
	Nontradereceivablescurrent                      float64     `json:"nontradereceivablescurrent"`
	Commonstocksharesissued                         float64     `json:"commonstocksharesissued"`
	Stockholdersequity                              interface{} `json:"stockholdersequity"`
	Accountsreceivablenetcurrent                    float64     `json:"accountsreceivablenetcurrent"`
	Accountspayablecurrent                          float64     `json:"accountspayablecurrent"`
	Assets                                          interface{} `json:"assets"`
	Assetsnoncurrent                                interface{} `json:"assetsnoncurrent"`
	Otherassetscurrent                              float64     `json:"otherassetscurrent"`
	Otherassetsnoncurrent                           float64     `json:"otherassetsnoncurrent"`
	Inventorynet                                    float64     `json:"inventorynet"`
	Marketablesecuritiesnoncurrent                  interface{} `json:"marketablesecuritiesnoncurrent"`
	Accumulatedothercomprehensiveincomelossnetoftax float64     `json:"accumulatedothercomprehensiveincomelossnetoftax"`
	Othershorttermborrowings                        float64     `json:"othershorttermborrowings"`
}

// CashFlowStatementAsReported ...
type CashFlowStatementAsReported struct {
	Date                                                                                                           string      `json:"date"`
	Symbol                                                                                                         string      `json:"symbol"`
	Period                                                                                                         string      `json:"period"`
	Paymentsforrepurchaseofcommonstock                                                                             float64     `json:"paymentsforrepurchaseofcommonstock"`
	Sharebasedcompensation                                                                                         float64     `json:"sharebasedcompensation"`
	Netincomeloss                                                                                                  float64     `json:"netincomeloss"`
	Increasedecreaseinaccountspayable                                                                              float64     `json:"increasedecreaseinaccountspayable"`
	Proceedsfrompaymentsforotherfinancingactivities                                                                float64     `json:"proceedsfrompaymentsforotherfinancingactivities"`
	Paymentsrelatedtotaxwithholdingforsharebasedcompensation                                                       float64     `json:"paymentsrelatedtotaxwithholdingforsharebasedcompensation"`
	Increasedecreaseinotheroperatingliabilities                                                                    float64     `json:"increasedecreaseinotheroperatingliabilities"`
	Othernoncashincomeexpense                                                                                      float64     `json:"othernoncashincomeexpense"`
	Paymentstoacquirebusinessesnetofcashacquired                                                                   float64     `json:"paymentstoacquirebusinessesnetofcashacquired"`
	Deferredincometaxexpensebenefit                                                                                float64     `json:"deferredincometaxexpensebenefit"`
	Cashcashequivalentsrestrictedcashandrestrictedcashequivalents                                                  interface{} `json:"cashcashequivalentsrestrictedcashandrestrictedcashequivalents"`
	Cashcashequivalentsrestrictedcashandrestrictedcashequivalentsperiodincreasedecreaseincludingexchangerateeffect float64     `json:"cashcashequivalentsrestrictedcashandrestrictedcashequivalentsperiodincreasedecreaseincludingexchangerateeffect"`
	Netcashprovidedbyusedinoperatingactivities                                                                     float64     `json:"netcashprovidedbyusedinoperatingactivities"`
	Proceedsfromsaleofavailableforsalesecuritiesdebt                                                               float64     `json:"proceedsfromsaleofavailableforsalesecuritiesdebt"`
	Repaymentsoflongtermdebt                                                                                       float64     `json:"repaymentsoflongtermdebt"`
	Incometaxespaidnet                                                                                             float64     `json:"incometaxespaidnet"`
	Proceedsfromissuanceoflongtermdebt                                                                             float64     `json:"proceedsfromissuanceoflongtermdebt"`
	Paymentstoacquireotherinvestments                                                                              float64     `json:"paymentstoacquireotherinvestments"`
	Netcashprovidedbyusedininvestingactivities                                                                     interface{} `json:"netcashprovidedbyusedininvestingactivities"`
	Increasedecreaseincontractwithcustomerliability                                                                float64     `json:"increasedecreaseincontractwithcustomerliability"`
	Interestpaidnet                                                                                                float64     `json:"interestpaidnet"`
	Netcashprovidedbyusedinfinancingactivities                                                                     interface{} `json:"netcashprovidedbyusedinfinancingactivities"`
	Proceedsfromrepaymentsofcommercialpaper                                                                        float64     `json:"proceedsfromrepaymentsofcommercialpaper"`
	Proceedsfromsaleandmaturityofotherinvestments                                                                  float64     `json:"proceedsfromsaleandmaturityofotherinvestments"`
	Paymentstoacquireavailableforsalesecuritiesdebt                                                                interface{} `json:"paymentstoacquireavailableforsalesecuritiesdebt"`
	Paymentstoacquirepropertyplantandequipment                                                                     float64     `json:"paymentstoacquirepropertyplantandequipment"`
	Paymentsforproceedsfromotherinvestingactivities                                                                float64     `json:"paymentsforproceedsfromotherinvestingactivities"`
	Increasedecreaseinotherreceivables                                                                             float64     `json:"increasedecreaseinotherreceivables"`
	Paymentsofdividends                                                                                            float64     `json:"paymentsofdividends"`
	Increasedecreaseininventories                                                                                  float64     `json:"increasedecreaseininventories"`
	Increasedecreaseinaccountsreceivable                                                                           float64     `json:"increasedecreaseinaccountsreceivable"`
	Proceedsfromissuanceofcommonstock                                                                              float64     `json:"proceedsfromissuanceofcommonstock"`
	Depreciationdepletionandamortization                                                                           float64     `json:"depreciationdepletionandamortization"`
	Proceedsfrommaturitiesprepaymentsandcallsofavailableforsalesecurities                                          float64     `json:"proceedsfrommaturitiesprepaymentsandcallsofavailableforsalesecurities"`
	Increasedecreaseinotheroperatingassets                                                                         float64     `json:"increasedecreaseinotheroperatingassets"`
	Proceedsfromothershorttermdebt                                                                                 float64     `json:"proceedsfromothershorttermdebt"`
}

// FullFinancialStatementAsReported ...
type FullFinancialStatementAsReported struct {
	Date                                                                                                                                        string      `json:"date"`
	Symbol                                                                                                                                      string      `json:"symbol"`
	Period                                                                                                                                      string      `json:"period"`
	Numberofsignificantvendors                                                                                                                  float64     `json:"numberofsignificantvendors"`
	Entitycurrentreportingstatus                                                                                                                string      `json:"entitycurrentreportingstatus"`
	Machineryequipmentandinternalusesoftwaremember                                                                                              float64     `json:"machineryequipmentandinternalusesoftwaremember"`
	Otherliabilitiesmember                                                                                                                      float64     `json:"otherliabilitiesmember"`
	Entityemerginggrowthcompany                                                                                                                 string      `json:"entityemerginggrowthcompany"`
	Incometaxreconciliationtaxcreditsresearch                                                                                                   float64     `json:"incometaxreconciliationtaxcreditsresearch"`
	Sharebasedcompensation                                                                                                                      float64     `json:"sharebasedcompensation"`
	Unrecordedunconditionalpurchaseobligationbalanceonsecondanniversary                                                                         float64     `json:"unrecordedunconditionalpurchaseobligationbalanceonsecondanniversary"`
	Nonoperatingincomeexpense                                                                                                                   float64     `json:"nonoperatingincomeexpense"`
	Incometaxespaidnet                                                                                                                          float64     `json:"incometaxespaidnet"`
	Federalincometaxexpensebenefitcontinuingoperations                                                                                          float64     `json:"federalincometaxexpensebenefitcontinuingoperations"`
	Japansegmentmember                                                                                                                          float64     `json:"japansegmentmember"`
	Currentforeigntaxexpensebenefit                                                                                                             float64     `json:"currentforeigntaxexpensebenefit"`
	Debtsecuritiesavailableforsalecontinuousunrealizedlossposition12Monthsorlonger                                                              float64     `json:"debtsecuritiesavailableforsalecontinuousunrealizedlossposition12monthsorlonger"`
	Proceedsfromsaleofavailableforsalesecuritiesdebt                                                                                            float64     `json:"proceedsfromsaleofavailableforsalesecuritiesdebt"`
	Commonstocksharesoutstanding                                                                                                                float64     `json:"commonstocksharesoutstanding"`
	Unrecognizedtaxbenefitsthatwouldimpacteffectivetaxrate                                                                                      float64     `json:"unrecognizedtaxbenefitsthatwouldimpacteffectivetaxrate"`
	Decreaseinunrecognizedtaxbenefitsisreasonablypossible                                                                                       float64     `json:"decreaseinunrecognizedtaxbenefitsisreasonablypossible"`
	Deferredtaxassetsnet                                                                                                                        float64     `json:"deferredtaxassetsnet"`
	Deferredtaxassetsgoodwillandintangibleassets                                                                                                float64     `json:"deferredtaxassetsgoodwillandintangibleassets"`
	Interestratecontractmember                                                                                                                  float64     `json:"interestratecontractmember"`
	Othercomprehensiveincomelossavailableforsalesecuritiesadjustmentnetoftax                                                                    float64     `json:"othercomprehensiveincomelossavailableforsalesecuritiesadjustmentnetoftax"`
	Entitytaxidentificationnumber                                                                                                               string      `json:"entitytaxidentificationnumber"`
	Incometaxreconciliationforeignincometaxratedifferential                                                                                     float64     `json:"incometaxreconciliationforeignincometaxratedifferential"`
	Paymentsrelatedtotaxwithholdingforsharebasedcompensation                                                                                    float64     `json:"paymentsrelatedtotaxwithholdingforsharebasedcompensation"`
	Increasedecreaseinotheroperatingliabilities                                                                                                 float64     `json:"increasedecreaseinotheroperatingliabilities"`
	Deferredtaxassetsliabilitiesnet                                                                                                             float64     `json:"deferredtaxassetsliabilitiesnet"`
	Debtinstrumentmaturityyearrangestart                                                                                                        float64     `json:"debtinstrumentmaturityyearrangestart"`
	Propertyplantandequipmentnet                                                                                                                float64     `json:"propertyplantandequipmentnet"`
	Commonstocksincludingadditionalpaidincapital                                                                                                float64     `json:"commonstocksincludingadditionalpaidincapital"`
	Weightedaveragenumberofdilutedsharesoutstanding                                                                                             float64     `json:"weightedaveragenumberofdilutedsharesoutstanding"`
	Commercialpaper                                                                                                                             float64     `json:"commercialpaper"`
	Commonstockmember                                                                                                                           interface{} `json:"commonstockmember"`
	Employeeservicesharebasedcompensationnonvestedawardstotalcompensationcostnotyetrecognized                                                   float64     `json:"employeeservicesharebasedcompensationnonvestedawardstotalcompensationcostnotyetrecognized"`
	Stockrepurchasedandretiredduringperiodshares                                                                                                float64     `json:"stockrepurchasedandretiredduringperiodshares"`
	Greaterchinasegmentmember                                                                                                                   float64     `json:"greaterchinasegmentmember"`
	Costofgoodsandservicessold                                                                                                                  interface{} `json:"costofgoodsandservicessold"`
	Longtermdebtcurrent                                                                                                                         float64     `json:"longtermdebtcurrent"`
	Accumulatedtranslationadjustmentmember                                                                                                      float64     `json:"accumulatedtranslationadjustmentmember"`
	Cashcashequivalentsrestrictedcashandrestrictedcashequivalentsperiodincreasedecreaseincludingexchangerateeffect                              float64     `json:"cashcashequivalentsrestrictedcashandrestrictedcashequivalentsperiodincreasedecreaseincludingexchangerateeffect"`
	Debtinstrumentunamortizeddiscountpremiumanddebtissuancecostsnet                                                                             float64     `json:"debtinstrumentunamortizeddiscountpremiumanddebtissuancecostsnet"`
	Repaymentsoflongtermdebt                                                                                                                    float64     `json:"repaymentsoflongtermdebt"`
	Earningspersharediluted                                                                                                                     float64     `json:"earningspersharediluted"`
	Netincomeloss                                                                                                                               float64     `json:"netincomeloss"`
	Proceedsfromissuanceoflongtermdebt                                                                                                          float64     `json:"proceedsfromissuanceoflongtermdebt"`
	Unrecordedunconditionalpurchaseobligationbalanceonfifthanniversary                                                                          float64     `json:"unrecordedunconditionalpurchaseobligationbalanceonfifthanniversary"`
	Incomelossfromcontinuingoperationsbeforeincometaxesextraordinaryitemsnoncontrollinginterest                                                 float64     `json:"incomelossfromcontinuingoperationsbeforeincometaxesextraordinaryitemsnoncontrollinginterest"`
	Incomelossfromcontinuingoperationsbeforeincometaxesforeign                                                                                  float64     `json:"incomelossfromcontinuingoperationsbeforeincometaxesforeign"`
	Crosscurrencyinterestratecontractmember                                                                                                     string      `json:"crosscurrencyinterestratecontractmember"`
	Localphonenumber                                                                                                                            string      `json:"localphonenumber"`
	Collateralalreadyreceivedaggregatefairvalue                                                                                                 float64     `json:"collateralalreadyreceivedaggregatefairvalue"`
	Interestpaidnet                                                                                                                             float64     `json:"interestpaidnet"`
	Gainlossfromcomponentsexcludedfromassessmentoffairvaluehedgeeffectivenessnet                                                                float64     `json:"gainlossfromcomponentsexcludedfromassessmentoffairvaluehedgeeffectivenessnet"`
	Accumulateddepreciationdepletionandamortizationpropertyplantandequipment                                                                    float64     `json:"accumulateddepreciationdepletionandamortizationpropertyplantandequipment"`
	Othercomprehensiveincomeunrealizedgainlossonderivativesarisingduringperiodbeforetax                                                         float64     `json:"othercomprehensiveincomeunrealizedgainlossonderivativesarisingduringperiodbeforetax"`
	Proceedsfromrepaymentsofcommercialpaper                                                                                                     float64     `json:"proceedsfromrepaymentsofcommercialpaper"`
	Proceedsfromsaleandmaturityofotherinvestments                                                                                               float64     `json:"proceedsfromsaleandmaturityofotherinvestments"`
	Longtermdebtnoncurrent                                                                                                                      float64     `json:"longtermdebtnoncurrent"`
	Maximumlengthoftimeforeigncurrencycashflowhedge                                                                                             string      `json:"maximumlengthoftimeforeigncurrencycashflowhedge"`
	Stockissuedduringperiodvaluenewissues                                                                                                       float64     `json:"stockissuedduringperiodvaluenewissues"`
	Netcashprovidedbyusedinoperatingactivities                                                                                                  float64     `json:"netcashprovidedbyusedinoperatingactivities"`
	Operatingleasesfutureminimumpaymentsdue                                                                                                     float64     `json:"operatingleasesfutureminimumpaymentsdue"`
	Proceedsfromrepaymentsofshorttermdebtmaturinginthreemonthsorless                                                                            float64     `json:"proceedsfromrepaymentsofshorttermdebtmaturinginthreemonthsorless"`
	Stockrepurchasedandretiredduringperiodvalue                                                                                                 float64     `json:"stockrepurchasedandretiredduringperiodvalue"`
	Proceedsfromrepaymentsofshorttermdebtmaturinginmorethanthreemonths                                                                          float64     `json:"proceedsfromrepaymentsofshorttermdebtmaturinginmorethanthreemonths"`
	Restrictedinvestments                                                                                                                       float64     `json:"restrictedinvestments"`
	Leaseholdimprovementsmember                                                                                                                 interface{} `json:"leaseholdimprovementsmember"`
	Standardproductwarrantyaccrualpayments                                                                                                      float64     `json:"standardproductwarrantyaccrualpayments"`
	Accountsreceivablenetcurrent                                                                                                                float64     `json:"accountsreceivablenetcurrent"`
	Performanceobligationsinarrangements                                                                                                        float64     `json:"performanceobligationsinarrangements"`
	Sharespaidfortaxwithholdingforsharebasedcompensation                                                                                        float64     `json:"sharespaidfortaxwithholdingforsharebasedcompensation"`
	Currentstateandlocaltaxexpensebenefit                                                                                                       float64     `json:"currentstateandlocaltaxexpensebenefit"`
	Repaymentsofshorttermdebtmaturinginmorethanthreemonths                                                                                      float64     `json:"repaymentsofshorttermdebtmaturinginmorethanthreemonths"`
	Derivativefairvalueofderivativenet                                                                                                          float64     `json:"derivativefairvalueofderivativenet"`
	Upfrontpaymentunderacceleratedsharerepurchasearrangement                                                                                    float64     `json:"upfrontpaymentunderacceleratedsharerepurchasearrangement"`
	Americassegmentmember                                                                                                                       interface{} `json:"americassegmentmember"`
	Deferredfederalincometaxexpensebenefit                                                                                                      float64     `json:"deferredfederalincometaxexpensebenefit"`
	Operatingleasesfutureminimumpaymentsduethereafter                                                                                           float64     `json:"operatingleasesfutureminimumpaymentsduethereafter"`
	Longtermmarketablesecuritiesmaturitiesterm                                                                                                  string      `json:"longtermmarketablesecuritiesmaturitiesterm"`
	Cashandcashequivalentsatcarryingvalue                                                                                                       float64     `json:"cashandcashequivalentsatcarryingvalue"`
	Documentannualreport                                                                                                                        string      `json:"documentannualreport"`
	Otherassetscurrent                                                                                                                          float64     `json:"otherassetscurrent"`
	Unfavorableinvestigationoutcomeeustateaidrulesmember                                                                                        float64     `json:"unfavorableinvestigationoutcomeeustateaidrulesmember"`
	Othernonoperatingincomeexpense                                                                                                              float64     `json:"othernonoperatingincomeexpense"`
	Maximumlengthoftimehedgedininterestratecashflowhedge1                                                                                       string      `json:"maximumlengthoftimehedgedininterestratecashflowhedge1"`
	Deferredtaxassetsunrealizedlossesonavailableforsalesecuritiesgross                                                                          float64     `json:"deferredtaxassetsunrealizedlossesonavailableforsalesecuritiesgross"`
	Definedcontributionplanmaximumannualcontributionsperemployeeamount                                                                          float64     `json:"definedcontributionplanmaximumannualcontributionsperemployeeamount"`
	Debtsecuritiesavailableforsaleunrealizedlossposition                                                                                        float64     `json:"debtsecuritiesavailableforsaleunrealizedlossposition"`
	Acceleratedsharerepurchasearrangementfebruary2019Member                                                                                     float64     `json:"acceleratedsharerepurchasearrangementfebruary2019member"`
	Increasedecreaseinotheroperatingassets                                                                                                      float64     `json:"increasedecreaseinotheroperatingassets"`
	Accumulatednetunrealizedinvestmentgainlossmember                                                                                            float64     `json:"accumulatednetunrealizedinvestmentgainlossmember"`
	Restofasiapacificsegmentmember                                                                                                              float64     `json:"restofasiapacificsegmentmember"`
	Unrecordedunconditionalpurchaseobligationbalanceonfourthanniversary                                                                         float64     `json:"unrecordedunconditionalpurchaseobligationbalanceonfourthanniversary"`
	Contractwithcustomerliability                                                                                                               float64     `json:"contractwithcustomerliability"`
	Tradeaccountsreceivablemember                                                                                                               float64     `json:"tradeaccountsreceivablemember"`
	Securityexchangename                                                                                                                        string      `json:"securityexchangename"`
	Proceedsfromissuanceofcommonstock                                                                                                           float64     `json:"proceedsfromissuanceofcommonstock"`
	Interestcostsincurred                                                                                                                       float64     `json:"interestcostsincurred"`
	Adjustmentstoadditionalpaidincapitaltaxeffectfromsharebasedcompensationincludingtransferpricing                                             float64     `json:"adjustmentstoadditionalpaidincapitaltaxeffectfromsharebasedcompensationincludingtransferpricing"`
	Revenueremainingperformanceobligationpercentage                                                                                             float64     `json:"revenueremainingperformanceobligationpercentage"`
	Liabilitiescurrent                                                                                                                          interface{} `json:"liabilitiescurrent"`
	Sharebasedcompensationarrangementbysharebasedpaymentawardofferingterm                                                                       string      `json:"sharebasedcompensationarrangementbysharebasedpaymentawardofferingterm"`
	Increasedecreaseinaccountspayable                                                                                                           float64     `json:"increasedecreaseinaccountspayable"`
	Grossprofit                                                                                                                                 interface{} `json:"grossprofit"`
	Reclassificationfromaocicurrentperiodbeforetaxattributabletoparent                                                                          float64     `json:"reclassificationfromaocicurrentperiodbeforetaxattributabletoparent"`
	Sharebasedcompensationarrangementbysharebasedpaymentawardequityinstrumentsotherthanoptionsvestedinperiod                                    float64     `json:"sharebasedcompensationarrangementbysharebasedpaymentawardequityinstrumentsotherthanoptionsvestedinperiod"`
	Retainedearningsaccumulateddeficit                                                                                                          interface{} `json:"retainedearningsaccumulateddeficit"`
	Proceedsfrompaymentsforotherfinancingactivities                                                                                             float64     `json:"proceedsfrompaymentsforotherfinancingactivities"`
	Definedcontributionplanemployermatchingcontributionpercentofmatch                                                                           float64     `json:"definedcontributionplanemployermatchingcontributionpercentofmatch"`
	Othercomprehensiveincomelossreclassificationadjustmentfromaocionderivativesnetoftax                                                         float64     `json:"othercomprehensiveincomelossreclassificationadjustmentfromaocionderivativesnetoftax"`
	Losscontingencyestimateofpossibleloss                                                                                                       float64     `json:"losscontingencyestimateofpossibleloss"`
	Currentfiscalyearenddate                                                                                                                    string      `json:"currentfiscalyearenddate"`
	Employeestockpurchaseplanmaximumannualpurchasesperemployeeamount                                                                            float64     `json:"employeestockpurchaseplanmaximumannualpurchasesperemployeeamount"`
	Longtermdebtfairvalue                                                                                                                       interface{} `json:"longtermdebtfairvalue"`
	Hedgeaccountingadjustmentsrelatedtolongtermdebt                                                                                             float64     `json:"hedgeaccountingadjustmentsrelatedtolongtermdebt"`
	Paymentstoacquirebusinessesnetofcashacquired                                                                                                float64     `json:"paymentstoacquirebusinessesnetofcashacquired"`
	Investmentincomeinterestanddividend                                                                                                         float64     `json:"investmentincomeinterestanddividend"`
	Othercomprehensiveincomeunrealizedgainlossonderivativesarisingduringperiodnetoftax                                                          float64     `json:"othercomprehensiveincomeunrealizedgainlossonderivativesarisingduringperiodnetoftax"`
	Changeinunrealizedgainlossonhedgediteminfairvaluehedge1                                                                                     float64     `json:"changeinunrealizedgainlossonhedgediteminfairvaluehedge1"`
	Accruedincometaxesnoncurrent                                                                                                                float64     `json:"accruedincometaxesnoncurrent"`
	Operatingexpenses                                                                                                                           interface{} `json:"operatingexpenses"`
	Cashmember                                                                                                                                  float64     `json:"cashmember"`
	Operatingincomeloss                                                                                                                         float64     `json:"operatingincomeloss"`
	Deferredtaxassetstaxdeferredexpensecompensationandbenefitssharebasedcompensationcost                                                        float64     `json:"deferredtaxassetstaxdeferredexpensecompensationandbenefitssharebasedcompensationcost"`
	A20132018Debtissuancesmember                                                                                                                float64     `json:"a20132018debtissuancesmember"`
	Noncashactivitiesinvolvingpropertyplantandequipmentnetincreasedecreasetoaccountspayableandothercurrentliabilities                           float64     `json:"noncashactivitiesinvolvingpropertyplantandequipmentnetincreasedecreasetoaccountspayableandothercurrentliabilities"`
	Sharebasedcompensationarrangementbysharebasedpaymentawardequityinstrumentsotherthanoptionsgrantsinperiod                                    float64     `json:"sharebasedcompensationarrangementbysharebasedpaymentawardequityinstrumentsotherthanoptionsgrantsinperiod"`
	Security12Btitle                                                                                                                            string      `json:"security12btitle"`
	Derivativefairvalueofderivativeasset                                                                                                        float64     `json:"derivativefairvalueofderivativeasset"`
	Otherassetsmember                                                                                                                           float64     `json:"otherassetsmember"`
	Paymentsforrepurchaseofcommonstock                                                                                                          float64     `json:"paymentsforrepurchaseofcommonstock"`
	Paymentstoacquireotherinvestments                                                                                                           float64     `json:"paymentstoacquireotherinvestments"`
	Netcashprovidedbyusedininvestingactivities                                                                                                  float64     `json:"netcashprovidedbyusedininvestingactivities"`
	Employeestockmember                                                                                                                         float64     `json:"employeestockmember"`
	Unrecordedunconditionalpurchaseobligationbalanceonfirstanniversary                                                                          float64     `json:"unrecordedunconditionalpurchaseobligationbalanceonfirstanniversary"`
	Assetscurrent                                                                                                                               interface{} `json:"assetscurrent"`
	Sharebasedcompensationarrangementbysharebasedpaymentawardnumberofsharesavailableforgrant                                                    float64     `json:"sharebasedcompensationarrangementbysharebasedpaymentawardnumberofsharesavailableforgrant"`
	Othernoncashincomeexpense                                                                                                                   float64     `json:"othernoncashincomeexpense"`
	Netcashprovidedbyusedinfinancingactivities                                                                                                  float64     `json:"netcashprovidedbyusedinfinancingactivities"`
	Paymentstoacquireavailableforsalesecuritiesdebt                                                                                             float64     `json:"paymentstoacquireavailableforsalesecuritiesdebt"`
	Derivativeinstrumentsgainlossreclassifiedfromaccumulatedociintoincomeeffectiveportionnet                                                    float64     `json:"derivativeinstrumentsgainlossreclassifiedfromaccumulatedociintoincomeeffectiveportionnet"`
	Restrictedcashandcashequivalentsatcarryingvalue                                                                                             float64     `json:"restrictedcashandcashequivalentsatcarryingvalue"`
	Restrictedcashandcashequivalentsnoncurrent                                                                                                  float64     `json:"restrictedcashandcashequivalentsnoncurrent"`
	Debtsecuritiesavailableforsalecontinuousunrealizedlosspositionlessthan12Monthsaccumulatedloss                                               float64     `json:"debtsecuritiesavailableforsalecontinuousunrealizedlosspositionlessthan12monthsaccumulatedloss"`
	Interestexpense                                                                                                                             float64     `json:"interestexpense"`
	Entityinteractivedatacurrent                                                                                                                string      `json:"entityinteractivedatacurrent"`
	Debtsecuritiesavailableforsaleunrealizedlosspositionaccumulatedloss                                                                         float64     `json:"debtsecuritiesavailableforsaleunrealizedlosspositionaccumulatedloss"`
	Debtinstrumentcarryingamount                                                                                                                interface{} `json:"debtinstrumentcarryingamount"`
	Contractwithcustomerliabilitycurrent                                                                                                        float64     `json:"contractwithcustomerliabilitycurrent"`
	Weightedaveragenumberofsharesoutstandingbasic                                                                                               float64     `json:"weightedaveragenumberofsharesoutstandingbasic"`
	Documentfiscalyearfocus                                                                                                                     float64     `json:"documentfiscalyearfocus"`
	Sharebasedcompensationarrangementbysharebasedpaymentawardmaximumemployeesubscriptionrate                                                    float64     `json:"sharebasedcompensationarrangementbysharebasedpaymentawardmaximumemployeesubscriptionrate"`
	Documentperiodenddate                                                                                                                       string      `json:"documentperiodenddate"`
	Sharebasedcompensationarrangementbysharebasedpaymentawardequityinstrumentsotherthanoptionsnonvestednumber                                   float64     `json:"sharebasedcompensationarrangementbysharebasedpaymentawardequityinstrumentsotherthanoptionsnonvestednumber"`
	Unrecognizedtaxbenefitsincometaxpenaltiesandinterestexpense                                                                                 float64     `json:"unrecognizedtaxbenefitsincometaxpenaltiesandinterestexpense"`
	Liabilities                                                                                                                                 interface{} `json:"liabilities"`
	Sharebasedcompensationarrangementbysharebasedpaymentawardequityinstrumentsotherthanoptionsgrantsinperiodweightedaveragegrantdatefairvalue   float64     `json:"sharebasedcompensationarrangementbysharebasedpaymentawardequityinstrumentsotherthanoptionsgrantsinperiodweightedaveragegrantdatefairvalue"`
	Debtinstrumentinterestratestatedpercentage                                                                                                  float64     `json:"debtinstrumentinterestratestatedpercentage"`
	Losscontingencyestimateofpossiblelossreductioninperiod                                                                                      float64     `json:"losscontingencyestimateofpossiblelossreductioninperiod"`
	Operatingleasesfutureminimumpaymentsdueinfouryears                                                                                          float64     `json:"operatingleasesfutureminimumpaymentsdueinfouryears"`
	Commonstockparorstatedvaluepershare                                                                                                         float64     `json:"commonstockparorstatedvaluepershare"`
	Availableforsaledebtsecuritiesaccumulatedgrossunrealizedgainbeforetax                                                                       float64     `json:"availableforsaledebtsecuritiesaccumulatedgrossunrealizedgainbeforetax"`
	Sharebasedcompensationarrangementbysharebasedpaymentawardpurchasepriceofcommonstockpercent                                                  float64     `json:"sharebasedcompensationarrangementbysharebasedpaymentawardpurchasepriceofcommonstockpercent"`
	Incometaxreconciliationstateandlocalincometaxes                                                                                             float64     `json:"incometaxreconciliationstateandlocalincometaxes"`
	Fairvalueinputslevel1Member                                                                                                                 float64     `json:"fairvalueinputslevel1member"`
	Lesseeoperatingleasetermofcontract                                                                                                          string      `json:"lesseeoperatingleasetermofcontract"`
	Debtinstrumentinterestrateeffectivepercentage                                                                                               float64     `json:"debtinstrumentinterestrateeffectivepercentage"`
	Contractwithcustomerliabilityrevenuerecognized                                                                                              float64     `json:"contractwithcustomerliabilityrevenuerecognized"`
	Unrecognizedtaxbenefitsreductionsresultingfromlapseofapplicablestatuteoflimitations                                                         float64     `json:"unrecognizedtaxbenefitsreductionsresultingfromlapseofapplicablestatuteoflimitations"`
	Entityfilercategory                                                                                                                         string      `json:"entityfilercategory"`
	Longtermdebtmaturitiesrepaymentsofprincipalinyeartwo                                                                                        float64     `json:"longtermdebtmaturitiesrepaymentsofprincipalinyeartwo"`
	Accountspayablecurrent                                                                                                                      float64     `json:"accountspayablecurrent"`
	Equitysecuritieswithoutreadilydeterminablefairvalueamount                                                                                   float64     `json:"equitysecuritieswithoutreadilydeterminablefairvalueamount"`
	Sharebasedcompensationarrangementbysharebasedpaymentawardequityinstrumentsotherthanoptionsnumberofsharesofcommonsharesawardeduponsettlement float64     `json:"sharebasedcompensationarrangementbysharebasedpaymentawardequityinstrumentsotherthanoptionsnumberofsharesofcommonsharesawardeduponsettlement"`
	Unrecordedunconditionalpurchaseobligationdueafterfiveyears                                                                                  float64     `json:"unrecordedunconditionalpurchaseobligationdueafterfiveyears"`
	Weightedaveragenumberdilutedsharesoutstandingadjustment                                                                                     float64     `json:"weightedaveragenumberdilutedsharesoutstandingadjustment"`
	Accumulatedothercomprehensiveincomemember                                                                                                   float64     `json:"accumulatedothercomprehensiveincomemember"`
	Fairvaluehedgingmember                                                                                                                      float64     `json:"fairvaluehedgingmember"`
	Derivativeassetsreductionformasternettingarrangements                                                                                       float64     `json:"derivativeassetsreductionformasternettingarrangements"`
	Researchanddevelopmentexpense                                                                                                               float64     `json:"researchanddevelopmentexpense"`
	Allocatedsharebasedcompensationexpense                                                                                                      float64     `json:"allocatedsharebasedcompensationexpense"`
	Availableforsalesecuritiesdebtsecurities                                                                                                    interface{} `json:"availableforsalesecuritiesdebtsecurities"`
	Liabilitiesandstockholdersequity                                                                                                            interface{} `json:"liabilitiesandstockholdersequity"`
	Availableforsaledebtsecuritiesamortizedcostbasis                                                                                            interface{} `json:"availableforsaledebtsecuritiesamortizedcostbasis"`
	Nontradereceivablemember                                                                                                                    float64     `json:"nontradereceivablemember"`
	Incometaxreconciliationincometaxexpensebenefitatfederalstatutoryincometaxrate                                                               float64     `json:"incometaxreconciliationincometaxexpensebenefitatfederalstatutoryincometaxrate"`
	Entityaddresscityortown                                                                                                                     string      `json:"entityaddresscityortown"`
	Derivativenotionalamount                                                                                                                    interface{} `json:"derivativenotionalamount"`
	Othercomprehensiveincomelossreclassificationadjustmentfromaociforsaleofsecuritiesnetoftax                                                   float64     `json:"othercomprehensiveincomelossreclassificationadjustmentfromaociforsaleofsecuritiesnetoftax"`
	Foreigncurrencydebtmember                                                                                                                   float64     `json:"foreigncurrencydebtmember"`
	Othercomprehensiveincomelosstaxportionattributabletoparent1                                                                                 float64     `json:"othercomprehensiveincomelosstaxportionattributabletoparent1"`
	Employeeservicesharebasedcompensationnonvestedawardstotalcompensationcostnotyetrecognizedperiodforrecognition1                              string      `json:"employeeservicesharebasedcompensationnonvestedawardstotalcompensationcostnotyetrecognizedperiodforrecognition1"`
	Entityregistrantname                                                                                                                        string      `json:"entityregistrantname"`
	Cashcashequivalentsrestrictedcashandrestrictedcashequivalents                                                                               float64     `json:"cashcashequivalentsrestrictedcashandrestrictedcashequivalents"`
	Entityaddressaddressline1                                                                                                                   string      `json:"entityaddressaddressline1"`
	Reclassificationoutofaccumulatedothercomprehensiveincomemember                                                                              float64     `json:"reclassificationoutofaccumulatedothercomprehensiveincomemember"`
	Stockissuedduringperiodsharessharebasedpaymentarrangementnetofshareswithheldfortaxes                                                        float64     `json:"stockissuedduringperiodsharessharebasedpaymentarrangementnetofshareswithheldfortaxes"`
	Landandbuildingmember                                                                                                                       float64     `json:"landandbuildingmember"`
	Deferredtaxliabilitiesminimumtaxonforeignearnings                                                                                           float64     `json:"deferredtaxliabilitiesminimumtaxonforeignearnings"`
	Europesegmentmember                                                                                                                         float64     `json:"europesegmentmember"`
	Effectiveincometaxratereconciliationsharebasedcompensationexcesstaxbenefitamount                                                            float64     `json:"effectiveincometaxratereconciliationsharebasedcompensationexcesstaxbenefitamount"`
	Numberofcustomerswithsignificantaccountsreceivablebalance                                                                                   float64     `json:"numberofcustomerswithsignificantaccountsreceivablebalance"`
	Marketablesecuritiesnoncurrent                                                                                                              interface{} `json:"marketablesecuritiesnoncurrent"`
	Sharebasedcompensationarrangementbysharebasedpaymentawardawardvestingperiod1                                                                string      `json:"sharebasedcompensationarrangementbysharebasedpaymentawardawardvestingperiod1"`
	Derivativecounterpartycreditriskexposure                                                                                                    float64     `json:"derivativecounterpartycreditriskexposure"`
	Changeinunrealizedgainlossonfairvaluehedginginstruments1                                                                                    float64     `json:"changeinunrealizedgainlossonfairvaluehedginginstruments1"`
	Operatingleasesrentexpensenet                                                                                                               float64     `json:"operatingleasesrentexpensenet"`
	Othercurrentliabilitiesmember                                                                                                               float64     `json:"othercurrentliabilitiesmember"`
	Amendmentflag                                                                                                                               string      `json:"amendmentflag"`
	Sharebasedcompensationarrangementbysharebasedpaymentawardequityinstrumentsotherthanoptionsnonvestedweightedaveragegrantdatefairvalue        float64     `json:"sharebasedcompensationarrangementbysharebasedpaymentawardequityinstrumentsotherthanoptionsnonvestedweightedaveragegrantdatefairvalue"`
	Foreignexchangecontractmember                                                                                                               interface{} `json:"foreignexchangecontractmember"`
	Unrecognizedtaxbenefits                                                                                                                     float64     `json:"unrecognizedtaxbenefits"`
	Increasedecreaseincontractwithcustomerliability                                                                                             float64     `json:"increasedecreaseincontractwithcustomerliability"`
	Sharebasedcompensationarrangementbysharebasedpaymentawardequityinstrumentsotherthanoptionsforfeituresweightedaveragegrantdatefairvalue      float64     `json:"sharebasedcompensationarrangementbysharebasedpaymentawardequityinstrumentsotherthanoptionsforfeituresweightedaveragegrantdatefairvalue"`
	Deferredincometaxexpensebenefit                                                                                                             float64     `json:"deferredincometaxexpensebenefit"`
	Entityaddressstateorprovince                                                                                                                string      `json:"entityaddressstateorprovince"`
	Propertyplantandequipmentusefullife                                                                                                         string      `json:"propertyplantandequipmentusefullife"`
	Unrecognizedtaxbenefitsincreasesresultingfrompriorperiodtaxpositions                                                                        float64     `json:"unrecognizedtaxbenefitsincreasesresultingfrompriorperiodtaxpositions"`
	Documenttype                                                                                                                                string      `json:"documenttype"`
	Paymentsforproceedsfromotherinvestingactivities                                                                                             float64     `json:"paymentsforproceedsfromotherinvestingactivities"`
	Entitysmallbusiness                                                                                                                         string      `json:"entitysmallbusiness"`
	Revenuefromcontractwithcustomerexcludingassessedtax                                                                                         string      `json:"revenuefromcontractwithcustomerexcludingassessedtax"`
	Currentfederaltaxexpensebenefit                                                                                                             float64     `json:"currentfederaltaxexpensebenefit"`
	Sharebasedcompensationarrangementbysharebasedpaymentawardequityinstrumentsotherthanoptionsaggregateintrinsicvaluenonvested                  float64     `json:"sharebasedcompensationarrangementbysharebasedpaymentawardequityinstrumentsotherthanoptionsaggregateintrinsicvaluenonvested"`
	Amountutilizedundersharerepurchaseprogram                                                                                                   interface{} `json:"amountutilizedundersharerepurchaseprogram"`
	Debtinstrumentmaturityyearrangeend                                                                                                          float64     `json:"debtinstrumentmaturityyearrangeend"`
	Commonstocksharesauthorized                                                                                                                 float64     `json:"commonstocksharesauthorized"`
	Deferredtaxassetsdeferredincome                                                                                                             float64     `json:"deferredtaxassetsdeferredincome"`
	Paymentsofdividends                                                                                                                         float64     `json:"paymentsofdividends"`
	Debtsecuritiesavailableforsalecontinuousunrealizedlossposition12Monthsorlongeraccumulatedloss                                               float64     `json:"debtsecuritiesavailableforsalecontinuousunrealizedlossposition12monthsorlongeraccumulatedloss"`
	Entityincorporationstatecountrycode                                                                                                         string      `json:"entityincorporationstatecountrycode"`
	Effectiveincometaxratecontinuingoperations                                                                                                  float64     `json:"effectiveincometaxratecontinuingoperations"`
	Factorbywhicheachrsugrantedreducesandeachrsucanceledorsharewithheldfortaxesincreasessharesavailableforgrant                                 float64     `json:"factorbywhicheachrsugrantedreducesandeachrsucanceledorsharewithheldfortaxesincreasessharesavailableforgrant"`
	Increasedecreaseininventories                                                                                                               float64     `json:"increasedecreaseininventories"`
	Ocibeforereclassificationsbeforetaxattributabletoparent                                                                                     float64     `json:"ocibeforereclassificationsbeforetaxattributabletoparent"`
	Dividends                                                                                                                                   float64     `json:"dividends"`
	Entitywellknownseasonedissuer                                                                                                               string      `json:"entitywellknownseasonedissuer"`
	Increasedecreaseinaccountsreceivable                                                                                                        float64     `json:"increasedecreaseinaccountsreceivable"`
	Entityvoluntaryfilers                                                                                                                       string      `json:"entityvoluntaryfilers"`
	Otherliabilitiescurrent                                                                                                                     float64     `json:"otherliabilitiescurrent"`
	Otheraccruedliabilitiesnoncurrent                                                                                                           float64     `json:"otheraccruedliabilitiesnoncurrent"`
	Operatingleasesfutureminimumpaymentsduecurrent                                                                                              float64     `json:"operatingleasesfutureminimumpaymentsduecurrent"`
	Proceedsfrommaturitiesprepaymentsandcallsofavailableforsalesecurities                                                                       float64     `json:"proceedsfrommaturitiesprepaymentsandcallsofavailableforsalesecurities"`
	Revenues                                                                                                                                    interface{} `json:"revenues"`
	Retainedearningsmember                                                                                                                      float64     `json:"retainedearningsmember"`
	Depreciation                                                                                                                                float64     `json:"depreciation"`
	Longtermdebtmaturitiesrepaymentsofprincipalafteryearfive                                                                                    float64     `json:"longtermdebtmaturitiesrepaymentsofprincipalafteryearfive"`
	Cityareacode                                                                                                                                float64     `json:"cityareacode"`
	Employeeservicesharebasedcompensationtaxbenefitfromcompensationexpense                                                                      float64     `json:"employeeservicesharebasedcompensationtaxbenefitfromcompensationexpense"`
	Accumulatednetgainlossfromdesignatedorqualifyingcashflowhedgesmember                                                                        float64     `json:"accumulatednetgainlossfromdesignatedorqualifyingcashflowhedgesmember"`
	Sharebasedcompensationarrangementbysharebasedpaymentawardequityinstrumentsotherthanoptionsforfeitedinperiod                                 float64     `json:"sharebasedcompensationarrangementbysharebasedpaymentawardequityinstrumentsotherthanoptionsforfeitedinperiod"`
	Entityaddresspostalzipcode                                                                                                                  interface{} `json:"entityaddresspostalzipcode"`
	Comprehensiveincomenetoftax                                                                                                                 float64     `json:"comprehensiveincomenetoftax"`
	Adjustmentstoadditionalpaidincapitalsharebasedcompensationrequisiteserviceperiodrecognitionvalue                                            float64     `json:"adjustmentstoadditionalpaidincapitalsharebasedcompensationrequisiteserviceperiodrecognitionvalue"`
	Fairvalueinputslevel2Member                                                                                                                 interface{} `json:"fairvalueinputslevel2member"`
	Deferredincometaxliabilities                                                                                                                float64     `json:"deferredincometaxliabilities"`
	Othercomprehensiveincomelossforeigncurrencytransactionandtranslationadjustmentnetoftax                                                      float64     `json:"othercomprehensiveincomelossforeigncurrencytransactionandtranslationadjustmentnetoftax"`
	Concentrationriskpercentage1                                                                                                                float64     `json:"concentrationriskpercentage1"`
	Stockholdersequity                                                                                                                          interface{} `json:"stockholdersequity"`
	Effectiveincometaxratereconciliationtaxcutsandjobsactof2017Amount                                                                           float64     `json:"effectiveincometaxratereconciliationtaxcutsandjobsactof2017amount"`
	Cashflowhedgingmember                                                                                                                       float64     `json:"cashflowhedgingmember"`
	Propertyplantandequipmentgross                                                                                                              float64     `json:"propertyplantandequipmentgross"`
	Tradingsymbol                                                                                                                               string      `json:"tradingsymbol"`
	Foreignincometaxexpensebenefitcontinuingoperations                                                                                          float64     `json:"foreignincometaxexpensebenefitcontinuingoperations"`
	Operatingleasesfutureminimumpaymentsdueintwoyears                                                                                           float64     `json:"operatingleasesfutureminimumpaymentsdueintwoyears"`
	Sharebasedcompensationarrangementbysharebasedpaymentawardequityinstrumentsotherthanoptionsvestedinperiodweightedaveragegrantdatefairvalue   float64     `json:"sharebasedcompensationarrangementbysharebasedpaymentawardequityinstrumentsotherthanoptionsvestedinperiodweightedaveragegrantdatefairvalue"`
	Employeestockplan2014Planmember                                                                                                             float64     `json:"employeestockplan2014planmember"`
	A2019Debtissuancemember                                                                                                                     float64     `json:"a2019debtissuancemember"`
	Documenttransitionreport                                                                                                                    string      `json:"documenttransitionreport"`
	Othercomprehensiveincomelossderivativesqualifyingashedgesnetoftax                                                                           float64     `json:"othercomprehensiveincomelossderivativesqualifyingashedgesnetoftax"`
	Liabilitiesnoncurrent                                                                                                                       interface{} `json:"liabilitiesnoncurrent"`
	Shorttermdebtweightedaverageinterestrate                                                                                                    float64     `json:"shorttermdebtweightedaverageinterestrate"`
	Incometaxreconciliationotheradjustments                                                                                                     float64     `json:"incometaxreconciliationotheradjustments"`
	Depreciationdepletionandamortization                                                                                                        float64     `json:"depreciationdepletionandamortization"`
	Deferredforeignincometaxexpensebenefit                                                                                                      float64     `json:"deferredforeignincometaxexpensebenefit"`
	Unrecordedunconditionalpurchaseobligationbalanceonthirdanniversary                                                                          float64     `json:"unrecordedunconditionalpurchaseobligationbalanceonthirdanniversary"`
	Operatingleasesfutureminimumpaymentsdueinthreeyears                                                                                         float64     `json:"operatingleasesfutureminimumpaymentsdueinthreeyears"`
	Definedcontributionplanemployermatchingcontributionpercent                                                                                  float64     `json:"definedcontributionplanemployermatchingcontributionpercent"`
	Debtsecuritiesavailableforsalecontinuousunrealizedlosspositionlessthan12Months                                                              float64     `json:"debtsecuritiesavailableforsalecontinuousunrealizedlosspositionlessthan12months"`
	Deferredtaxassetstaxdeferredexpensereservesandaccruals                                                                                      float64     `json:"deferredtaxassetstaxdeferredexpensereservesandaccruals"`
	Unrecognizedtaxbenefitsdecreasesresultingfrompriorperiodtaxpositions                                                                        float64     `json:"unrecognizedtaxbenefitsdecreasesresultingfrompriorperiodtaxpositions"`
	Stateandlocalincometaxexpensebenefitcontinuingoperations                                                                                    float64     `json:"stateandlocalincometaxexpensebenefitcontinuingoperations"`
	Deferredstateandlocalincometaxexpensebenefit                                                                                                float64     `json:"deferredstateandlocalincometaxexpensebenefit"`
	Otherliabilitiesnoncurrent                                                                                                                  float64     `json:"otherliabilitiesnoncurrent"`
	Commonstockincludingadditionalpaidincapitalmember                                                                                           float64     `json:"commonstockincludingadditionalpaidincapitalmember"`
	Deferredtaxassetsother                                                                                                                      float64     `json:"deferredtaxassetsother"`
	Standardproductwarrantyaccrual                                                                                                              float64     `json:"standardproductwarrantyaccrual"`
	Deferredtaxliabilitiesundistributedforeignearnings                                                                                          float64     `json:"deferredtaxliabilitiesundistributedforeignearnings"`
	Marketablesecuritiescurrent                                                                                                                 float64     `json:"marketablesecuritiescurrent"`
	Derivativefairvalueofderivativeliability                                                                                                    float64     `json:"derivativefairvalueofderivativeliability"`
	Unrecognizedtaxbenefitsdecreasesresultingfromsettlementswithtaxingauthorities                                                               float64     `json:"unrecognizedtaxbenefitsdecreasesresultingfromsettlementswithtaxingauthorities"`
	Longtermdebtmaturitiesrepaymentsofprincipalinyearfour                                                                                       float64     `json:"longtermdebtmaturitiesrepaymentsofprincipalinyearfour"`
	Othergeneralandadministrativeexpense                                                                                                        float64     `json:"othergeneralandadministrativeexpense"`
	Longtermdebtmaturitiesrepaymentsofprincipalinyearfive                                                                                       float64     `json:"longtermdebtmaturitiesrepaymentsofprincipalinyearfive"`
	Commonstockdividendspersharedeclared                                                                                                        float64     `json:"commonstockdividendspersharedeclared"`
	Directorplanmember                                                                                                                          float64     `json:"directorplanmember"`
	Adjustmentsrelatedtotaxwithholdingforsharebasedcompensation                                                                                 float64     `json:"adjustmentsrelatedtotaxwithholdingforsharebasedcompensation"`
	Unrecordedunconditionalpurchaseobligationbalancesheetamount                                                                                 float64     `json:"unrecordedunconditionalpurchaseobligationbalancesheetamount"`
	Incometaxexpensebenefit                                                                                                                     float64     `json:"incometaxexpensebenefit"`
	Othercomprehensiveincomelossnetoftaxportionattributabletoparent                                                                             float64     `json:"othercomprehensiveincomelossnetoftaxportionattributabletoparent"`
	Restrictedstockunitsrsumember                                                                                                               interface{} `json:"restrictedstockunitsrsumember"`
	Derivativeliabilitiesreductionformasternettingarrangements                                                                                  float64     `json:"derivativeliabilitiesreductionformasternettingarrangements"`
	Nontradereceivablescurrent                                                                                                                  float64     `json:"nontradereceivablescurrent"`
	Sharebasedcompensationarrangementbysharebasedpaymentawardequityinstrumentsotherthanoptionsvestedinperiodtotalfairvalue                      float64     `json:"sharebasedcompensationarrangementbysharebasedpaymentawardequityinstrumentsotherthanoptionsvestedinperiodtotalfairvalue"`
	Othercomprehensiveincomeunrealizedholdinggainlossonsecuritiesarisingduringperiodnetoftax                                                    float64     `json:"othercomprehensiveincomeunrealizedholdinggainlossonsecuritiesarisingduringperiodnetoftax"`
	Inventorynet                                                                                                                                float64     `json:"inventorynet"`
	Increasedecreaseinotherreceivables                                                                                                          float64     `json:"increasedecreaseinotherreceivables"`
	Entitycentralindexkey                                                                                                                       float64     `json:"entitycentralindexkey"`
	Paymentstoacquirepropertyplantandequipment                                                                                                  float64     `json:"paymentstoacquirepropertyplantandequipment"`
	Debtinstrumentterm                                                                                                                          string      `json:"debtinstrumentterm"`
	Revenueremainingperformanceobligationexpectedtimingofsatisfactionperiod1                                                                    string      `json:"revenueremainingperformanceobligationexpectedtimingofsatisfactionperiod1"`
	Availableforsaledebtsecuritiesaccumulatedgrossunrealizedlossbeforetax                                                                       float64     `json:"availableforsaledebtsecuritiesaccumulatedgrossunrealizedlossbeforetax"`
	Unrecognizedtaxbenefitsincometaxpenaltiesandinterestaccrued                                                                                 float64     `json:"unrecognizedtaxbenefitsincometaxpenaltiesandinterestaccrued"`
	Entityfilenumber                                                                                                                            string      `json:"entityfilenumber"`
	Assets                                                                                                                                      interface{} `json:"assets"`
	Seniornotes                                                                                                                                 float64     `json:"seniornotes"`
	Assetsnoncurrent                                                                                                                            interface{} `json:"assetsnoncurrent"`
	Earningspersharebasic                                                                                                                       float64     `json:"earningspersharebasic"`
	Documentfiscalperiodfocus                                                                                                                   string      `json:"documentfiscalperiodfocus"`
	Standardproductwarrantyaccrualwarrantiesissued                                                                                              float64     `json:"standardproductwarrantyaccrualwarrantiesissued"`
	Proceedsfromshorttermdebtmaturinginmorethanthreemonths                                                                                      float64     `json:"proceedsfromshorttermdebtmaturinginmorethanthreemonths"`
	Entityshellcompany                                                                                                                          string      `json:"entityshellcompany"`
	Sellinggeneralandadministrativeexpense                                                                                                      float64     `json:"sellinggeneralandadministrativeexpense"`
	Deferredtaxliabilitiesother                                                                                                                 float64     `json:"deferredtaxliabilitiesother"`
	Operatingleasesfutureminimumpaymentsdueinfiveyears                                                                                          float64     `json:"operatingleasesfutureminimumpaymentsdueinfiveyears"`
	Netinvestmenthedgingmember                                                                                                                  float64     `json:"netinvestmenthedgingmember"`
	Longtermdebtmaturitiesrepaymentsofprincipalinyearthree                                                                                      float64     `json:"longtermdebtmaturitiesrepaymentsofprincipalinyearthree"`
	Effectiveincometaxratereconciliationatfederalstatutoryincometaxrate                                                                         float64     `json:"effectiveincometaxratereconciliationatfederalstatutoryincometaxrate"`
	Noncurrentassets                                                                                                                            float64     `json:"noncurrentassets"`
	Otherassetsnoncurrent                                                                                                                       float64     `json:"otherassetsnoncurrent"`
	Commonstocksharesissued                                                                                                                     float64     `json:"commonstocksharesissued"`
	Antidilutivesecuritiesexcludedfromcomputationofearningspershareamount                                                                       float64     `json:"antidilutivesecuritiesexcludedfromcomputationofearningspershareamount"`
	Unrecognizedtaxbenefitsincreasesresultingfromcurrentperiodtaxpositions                                                                      float64     `json:"unrecognizedtaxbenefitsincreasesresultingfromcurrentperiodtaxpositions"`
	Longtermdebtmaturitiesrepaymentsofprincipalinnexttwelvemonths                                                                               float64     `json:"longtermdebtmaturitiesrepaymentsofprincipalinnexttwelvemonths"`
	Accumulatedothercomprehensiveincomelossnetoftax                                                                                             float64     `json:"accumulatedothercomprehensiveincomelossnetoftax"`
	Lesseeoperatingandfinanceleaseliabilitypaymentsdueafteryearfive                                                                             float64     `json:"lesseeoperatingandfinanceleaseliabilitypaymentsdueafteryearfive"`
	Operatingandfinanceleaseweightedaverageremainingleaseterm                                                                                   *string     `json:"operatingandfinanceleaseweightedaverageremainingleaseterm"`
	Acceleratedsharerepurchasearrangementnovember2019Member                                                                                     float64     `json:"acceleratedsharerepurchasearrangementnovember2019member"`
	Stockrepurchaseprogramauthorizedamount1                                                                                                     interface{} `json:"stockrepurchaseprogramauthorizedamount1"`
	Lesseeoperatingandfinanceleaseliabilitypaymentsdueyearfive                                                                                  float64     `json:"lesseeoperatingandfinanceleaseliabilitypaymentsdueyearfive"`
	Othercomprehensiveincomelossnetinvestmenthedgegainlossbeforereclassificationandtax                                                          float64     `json:"othercomprehensiveincomelossnetinvestmenthedgegainlossbeforereclassificationandtax"`
	Lesseeoperatingleaseleasenotyetcommencedpaymentsdue                                                                                         float64     `json:"lesseeoperatingleaseleasenotyetcommencedpaymentsdue"`
	Debtinstrumentfaceamount                                                                                                                    float64     `json:"debtinstrumentfaceamount"`
	Documentquarterlyreport                                                                                                                     *string     `json:"documentquarterlyreport"`
	Lesseeoperatingandfinanceleaseliabilitypaymentsdueyearfour                                                                                  float64     `json:"lesseeoperatingandfinanceleaseliabilitypaymentsdueyearfour"`
	Othercomprehensiveincomelosscashflowhedgegainlossbeforereclassificationandtax                                                               float64     `json:"othercomprehensiveincomelosscashflowhedgegainlossbeforereclassificationandtax"`
	Lesseeoperatingandfinanceleaseliabilityundiscountedexcessamount                                                                             float64     `json:"lesseeoperatingandfinanceleaseliabilityundiscountedexcessamount"`
	Operatingleaseliabilitycurrent                                                                                                              float64     `json:"operatingleaseliabilitycurrent"`
	Operatingandfinanceleaseweightedaveragediscountratepercent                                                                                  float64     `json:"operatingandfinanceleaseweightedaveragediscountratepercent"`
	Lesseeoperatingleaseliabilitypaymentsdueafteryearfive                                                                                       float64     `json:"lesseeoperatingleaseliabilitypaymentsdueafteryearfive"`
	Lesseeoperatingleaseliabilitypaymentsremainderoffiscalyear                                                                                  float64     `json:"lesseeoperatingleaseliabilitypaymentsremainderoffiscalyear"`
	Rightofuseassetsobtainedinexchangeforoperatingandfinanceleaseliabilities                                                                    float64     `json:"rightofuseassetsobtainedinexchangeforoperatingandfinanceleaseliabilities"`
	Othernoncurrentliabilitiesmember                                                                                                            float64     `json:"othernoncurrentliabilitiesmember"`
	Lesseeoperatingleaseliabilitypaymentsdueyeartwo                                                                                             float64     `json:"lesseeoperatingleaseliabilitypaymentsdueyeartwo"`
	Lesseeoperatingandfinanceleaseliabilitypaymentsremainderoffiscalyear                                                                        float64     `json:"lesseeoperatingandfinanceleaseliabilitypaymentsremainderoffiscalyear"`
	Operatingleasepayments                                                                                                                      float64     `json:"operatingleasepayments"`
	Operatingleaseliabilitynoncurrent                                                                                                           float64     `json:"operatingleaseliabilitynoncurrent"`
	Secureddebtrepurchaseagreements                                                                                                             float64     `json:"secureddebtrepurchaseagreements"`
	Variableleasecost                                                                                                                           float64     `json:"variableleasecost"`
	Othercomprehensiveincomelossderivativeexcludedcomponentincreasedecreasebeforeadjustmentsandtax                                              float64     `json:"othercomprehensiveincomelossderivativeexcludedcomponentincreasedecreasebeforeadjustmentsandtax"`
	Financeleaseliabilitycurrent                                                                                                                float64     `json:"financeleaseliabilitycurrent"`
	Financeleaseliabilitynoncurrent                                                                                                             float64     `json:"financeleaseliabilitynoncurrent"`
	Currenttermdebtandnoncurrenttermdebtmember                                                                                                  float64     `json:"currenttermdebtandnoncurrenttermdebtmember"`
	Hedgedliabilityfairvaluehedgecumulativeincreasedecrease                                                                                     float64     `json:"hedgedliabilityfairvaluehedgecumulativeincreasedecrease"`
	Othernoncurrentassetsmember                                                                                                                 float64     `json:"othernoncurrentassetsmember"`
	Hedgedassetfairvaluehedgecumulativeincreasedecrease                                                                                         float64     `json:"hedgedassetfairvaluehedgecumulativeincreasedecrease"`
	Nonoperatingincomeexpensemember                                                                                                             float64     `json:"nonoperatingincomeexpensemember"`
	Acceleratedsharerepurchasearrangementmay2020Member                                                                                          float64     `json:"acceleratedsharerepurchasearrangementmay2020member"`
	Securitiessoldunderagreementstorepurchasefairvalueofcollateral                                                                              float64     `json:"securitiessoldunderagreementstorepurchasefairvalueofcollateral"`
	Lesseeoperatingleaseliabilitypaymentsdueyearthree                                                                                           float64     `json:"lesseeoperatingleaseliabilitypaymentsdueyearthree"`
	Lesseeoperatingleaseliabilitypaymentsdueyearfive                                                                                            float64     `json:"lesseeoperatingleaseliabilitypaymentsdueyearfive"`
	Commercialpapermember                                                                                                                       float64     `json:"commercialpapermember"`
	Lesseeoperatingleaseliabilityundiscountedexcessamount                                                                                       float64     `json:"lesseeoperatingleaseliabilityundiscountedexcessamount"`
	Accumulatedgainlossnetderivativeinstrumentparentmember                                                                                      float64     `json:"accumulatedgainlossnetderivativeinstrumentparentmember"`
	Sharebasedcompensationarrangementbysharebasedpaymentawardequityinstrumentsotherthanoptionsnumberofsharesofcommonstockawardeduponsettlement  float64     `json:"sharebasedcompensationarrangementbysharebasedpaymentawardequityinstrumentsotherthanoptionsnumberofsharesofcommonstockawardeduponsettlement"`
	Lesseeoperatingleaseliabilitypaymentsdueyearfour                                                                                            float64     `json:"lesseeoperatingleaseliabilitypaymentsdueyearfour"`
	Lesseeoperatingandfinanceleasetermofcontract                                                                                                *string     `json:"lesseeoperatingandfinanceleasetermofcontract"`
	Lesseeoperatingandfinanceleaseliabilitypaymentsdueyearthree                                                                                 float64     `json:"lesseeoperatingandfinanceleaseliabilitypaymentsdueyearthree"`
	Othercomprehensiveincomelosscashflowhedgegainlossreclassificationbeforetax                                                                  float64     `json:"othercomprehensiveincomelosscashflowhedgegainlossreclassificationbeforetax"`
	Lesseeoperatingleaseleasenotyetcommencedtermofcontract1                                                                                     *string     `json:"lesseeoperatingleaseleasenotyetcommencedtermofcontract1"`
	Operatingandfinanceleaserightofuseasset                                                                                                     float64     `json:"operatingandfinanceleaserightofuseasset"`
	Operatingleasecost                                                                                                                          float64     `json:"operatingleasecost"`
	Currentmarketablesecuritiesandnoncurrentmarketablesecuritiesmember                                                                          float64     `json:"currentmarketablesecuritiesandnoncurrentmarketablesecuritiesmember"`
	Lesseeoperatingandfinanceleaseliabilitypaymentsdueyeartwo                                                                                   float64     `json:"lesseeoperatingandfinanceleaseliabilitypaymentsdueyeartwo"`
	Securitiessoldunderagreementstorepurchasemember                                                                                             *string     `json:"securitiessoldunderagreementstorepurchasemember"`
	A20132019Debtissuancesmember                                                                                                                float64     `json:"a20132019debtissuancesmember"`
	Operatingleaserightofuseasset                                                                                                               float64     `json:"operatingleaserightofuseasset"`
	Firstquarter2020Debtissuancemember                                                                                                          float64     `json:"firstquarter2020debtissuancemember"`
	Thirdquarter2020Debtissuancemember                                                                                                          float64     `json:"thirdquarter2020debtissuancemember"`
	Propertyplantandequipmentmember                                                                                                             float64     `json:"propertyplantandequipmentmember"`
}

// FinancialRatios ...
type FinancialRatios struct {
	Symbol                             string  `json:"symbol" csv:"symbol"`
	Date                               string  `json:"date" csv:"date"`
	Period                             string  `json:"period" csv:"period"`
	CurrentRatio                       float64 `json:"currentRatio" csv:"currentRatio"`
	QuickRatio                         float64 `json:"quickRatio" csv:"quickRatio"`
	CashRatio                          float64 `json:"cashRatio" csv:"cashRatio"`
	DaysOfSalesOutstanding             float64 `json:"daysOfSalesOutstanding" csv:"daysOfSalesOutstanding"`
	DaysOfInventoryOutstanding         float64 `json:"daysOfInventoryOutstanding" csv:"daysOfInventoryOutstanding"`
	OperatingCycle                     float64 `json:"operatingCycle" csv:"operatingCycle"`
	DaysOfPayablesOutstanding          float64 `json:"daysOfPayablesOutstanding" csv:"daysOfPayablesOutstanding"`
	CashConversionCycle                float64 `json:"cashConversionCycle" csv:"cashConversionCycle"`
	GrossProfitMargin                  float64 `json:"grossProfitMargin" csv:"grossProfitMargin"`
	OperatingProfitMargin              float64 `json:"operatingProfitMargin" csv:"operatingProfitMargin"`
	PretaxProfitMargin                 float64 `json:"pretaxProfitMargin" csv:"pretaxProfitMargin"`
	NetProfitMargin                    float64 `json:"netProfitMargin" csv:"netProfitMargin"`
	EffectiveTaxRate                   float64 `json:"effectiveTaxRate" csv:"effectiveTaxRate"`
	ReturnOnAssets                     float64 `json:"returnOnAssets" csv:"returnOnAssets"`
	ReturnOnEquity                     float64 `json:"returnOnEquity" csv:"returnOnEquity"`
	ReturnOnCapitalEmployed            float64 `json:"returnOnCapitalEmployed" csv:"returnOnCapitalEmployed"`
	NetIncomePerEBT                    float64 `json:"netIncomePerEBT" csv:"netIncomePerEBT"`
	EbtPerEbit                         float64 `json:"ebtPerEbit" csv:"ebtPerEbit"`
	EbitPerRevenue                     float64 `json:"ebitPerRevenue" csv:"ebitPerRevenue"`
	DebtRatio                          float64 `json:"debtRatio" csv:"debtRatio"`
	DebtEquityRatio                    float64 `json:"debtEquityRatio" csv:"debtEquityRatio"`
	LongTermDebtToCapitalization       float64 `json:"longTermDebtToCapitalization" csv:"longTermDebtToCapitalization"`
	TotalDebtToCapitalization          float64 `json:"totalDebtToCapitalization" csv:"totalDebtToCapitalization"`
	InterestCoverage                   float64 `json:"interestCoverage" csv:"interestCoverage"`
	CashFlowToDebtRatio                float64 `json:"cashFlowToDebtRatio" csv:"cashFlowToDebtRatio"`
	CompanyEquityMultiplier            float64 `json:"companyEquityMultiplier" csv:"companyEquityMultiplier"`
	ReceivablesTurnover                float64 `json:"receivablesTurnover" csv:"receivablesTurnover"`
	PayablesTurnover                   float64 `json:"payablesTurnover" csv:"payablesTurnover"`
	InventoryTurnover                  float64 `json:"inventoryTurnover" csv:"inventoryTurnover"`
	FixedAssetTurnover                 float64 `json:"fixedAssetTurnover" csv:"fixedAssetTurnover"`
	AssetTurnover                      float64 `json:"assetTurnover" csv:"assetTurnover"`
	OperatingCashFlowPerShare          float64 `json:"operatingCashFlowPerShare" csv:"operatingCashFlowPerShare"`
	FreeCashFlowPerShare               float64 `json:"freeCashFlowPerShare" csv:"freeCashFlowPerShare"`
	CashPerShare                       float64 `json:"cashPerShare" csv:"cashPerShare"`
	PayoutRatio                        float64 `json:"payoutRatio" csv:"payoutRatio"`
	OperatingCashFlowSalesRatio        float64 `json:"operatingCashFlowSalesRatio" csv:"operatingCashFlowSalesRatio"`
	FreeCashFlowOperatingCashFlowRatio float64 `json:"freeCashFlowOperatingCashFlowRatio" csv:"freeCashFlowOperatingCashFlowRatio"`
	CashFlowCoverageRatios             float64 `json:"cashFlowCoverageRatios" csv:"cashFlowCoverageRatios"`
	ShortTermCoverageRatios            float64 `json:"shortTermCoverageRatios" csv:"shortTermCoverageRatios"`
	CapitalExpenditureCoverageRatio    float64 `json:"capitalExpenditureCoverageRatio" csv:"capitalExpenditureCoverageRatio"`
	DividendPaidAndCapexCoverageRatio  float64 `json:"dividendPaidAndCapexCoverageRatio" csv:"dividendPaidAndCapexCoverageRatio"`
	DividendPayoutRatio                float64 `json:"dividendPayoutRatio" csv:"dividendPayoutRatio"`
	PriceBookValueRatio                float64 `json:"priceBookValueRatio" csv:"priceBookValueRatio"`
	PriceToBookRatio                   float64 `json:"priceToBookRatio" csv:"priceToBookRatio"`
	PriceToSalesRatio                  float64 `json:"priceToSalesRatio" csv:"priceToSalesRatio"`
	PriceEarningsRatio                 float64 `json:"priceEarningsRatio" csv:"priceEarningsRatio"`
	PriceToFreeCashFlowsRatio          float64 `json:"priceToFreeCashFlowsRatio" csv:"priceToFreeCashFlowsRatio"`
	PriceToOperatingCashFlowsRatio     float64 `json:"priceToOperatingCashFlowsRatio" csv:"priceToOperatingCashFlowsRatio"`
	PriceCashFlowRatio                 float64 `json:"priceCashFlowRatio" csv:"priceCashFlowRatio"`
	PriceEarningsToGrowthRatio         float64 `json:"priceEarningsToGrowthRatio" csv:"priceEarningsToGrowthRatio"`
	PriceSalesRatio                    float64 `json:"priceSalesRatio" csv:"priceSalesRatio"`
	DividendYield                      float64 `json:"dividendYield" csv:"dividendYield"`
	EnterpriseValueMultiple            float64 `json:"enterpriseValueMultiple" csv:"enterpriseValueMultiple"`
	PriceFairValue                     float64 `json:"priceFairValue" csv:"priceFairValue"`
}

// FinancialRatiosTTM ...
type FinancialRatiosTTM struct {
	Symbol                                string  `csv:"symbol"`
	DividendYielTTM                       float64 `json:"dividendYielTTM" csv:"dividendYielTTM"`
	DividendYielPercentageTTM             float64 `json:"dividendYielPercentageTTM" csv:"dividendYielPercentageTTM"`
	PeRatioTTM                            float64 `json:"peRatioTTM" csv:"peRatioTTM"`
	PegRatioTTM                           float64 `json:"pegRatioTTM" csv:"pegRatioTTM"`
	PayoutRatioTTM                        float64 `json:"payoutRatioTTM" csv:"payoutRatioTTM"`
	CurrentRatioTTM                       float64 `json:"currentRatioTTM" csv:"currentRatioTTM"`
	QuickRatioTTM                         float64 `json:"quickRatioTTM" csv:"quickRatioTTM"`
	CashRatioTTM                          float64 `json:"cashRatioTTM" csv:"cashRatioTTM"`
	DaysOfSalesOutstandingTTM             float64 `json:"daysOfSalesOutstandingTTM" csv:"daysOfSalesOutstandingTTM"`
	DaysOfInventoryOutstandingTTM         float64 `json:"daysOfInventoryOutstandingTTM" csv:"daysOfInventoryOutstandingTTM"`
	OperatingCycleTTM                     float64 `json:"operatingCycleTTM" csv:"operatingCycleTTM"`
	DaysOfPayablesOutstandingTTM          float64 `json:"daysOfPayablesOutstandingTTM" csv:"daysOfPayablesOutstandingTTM"`
	CashConversionCycleTTM                float64 `json:"cashConversionCycleTTM" csv:"cashConversionCycleTTM"`
	GrossProfitMarginTTM                  float64 `json:"grossProfitMarginTTM" csv:"grossProfitMarginTTM"`
	OperatingProfitMarginTTM              float64 `json:"operatingProfitMarginTTM" csv:"operatingProfitMarginTTM"`
	PretaxProfitMarginTTM                 float64 `json:"pretaxProfitMarginTTM" csv:"pretaxProfitMarginTTM"`
	NetProfitMarginTTM                    float64 `json:"netProfitMarginTTM" csv:"netProfitMarginTTM"`
	EffectiveTaxRateTTM                   float64 `json:"effectiveTaxRateTTM" csv:"effectiveTaxRateTTM"`
	ReturnOnAssetsTTM                     float64 `json:"returnOnAssetsTTM" csv:"returnOnAssetsTTM"`
	ReturnOnEquityTTM                     float64 `json:"returnOnEquityTTM" csv:"returnOnEquityTTM"`
	ReturnOnCapitalEmployedTTM            float64 `json:"returnOnCapitalEmployedTTM" csv:"returnOnCapitalEmployedTTM"`
	NetIncomePerEBTTTM                    float64 `json:"netIncomePerEBTTTM" csv:"netIncomePerEBTTTM"`
	EbtPerEbitTTM                         float64 `json:"ebtPerEbitTTM" csv:"ebtPerEbitTTM"`
	EbitPerRevenueTTM                     float64 `json:"ebitPerRevenueTTM" csv:"ebitPerRevenueTTM"`
	DebtRatioTTM                          float64 `json:"debtRatioTTM" csv:"debtRatioTTM"`
	DebtEquityRatioTTM                    float64 `json:"debtEquityRatioTTM" csv:"debtEquityRatioTTM"`
	LongTermDebtToCapitalizationTTM       float64 `json:"longTermDebtToCapitalizationTTM" csv:"longTermDebtToCapitalizationTTM"`
	TotalDebtToCapitalizationTTM          float64 `json:"totalDebtToCapitalizationTTM" csv:"totalDebtToCapitalizationTTM"`
	InterestCoverageTTM                   float64 `json:"interestCoverageTTM" csv:"interestCoverageTTM"`
	CashFlowToDebtRatioTTM                float64 `json:"cashFlowToDebtRatioTTM" csv:"cashFlowToDebtRatioTTM"`
	CompanyEquityMultiplierTTM            float64 `json:"companyEquityMultiplierTTM" csv:"companyEquityMultiplierTTM"`
	ReceivablesTurnoverTTM                float64 `json:"receivablesTurnoverTTM" csv:"receivablesTurnoverTTM"`
	PayablesTurnoverTTM                   float64 `json:"payablesTurnoverTTM" csv:"payablesTurnoverTTM"`
	InventoryTurnoverTTM                  float64 `json:"inventoryTurnoverTTM" csv:"inventoryTurnoverTTM"`
	FixedAssetTurnoverTTM                 float64 `json:"fixedAssetTurnoverTTM" csv:"fixedAssetTurnoverTTM"`
	AssetTurnoverTTM                      float64 `json:"assetTurnoverTTM" csv:"assetTurnoverTTM"`
	OperatingCashFlowPerShareTTM          float64 `json:"operatingCashFlowPerShareTTM" csv:"operatingCashFlowPerShareTTM"`
	FreeCashFlowPerShareTTM               float64 `json:"freeCashFlowPerShareTTM" csv:"freeCashFlowPerShareTTM"`
	CashPerShareTTM                       float64 `json:"cashPerShareTTM" csv:"cashPerShareTTM"`
	OperatingCashFlowSalesRatioTTM        float64 `json:"operatingCashFlowSalesRatioTTM" csv:"operatingCashFlowSalesRatioTTM"`
	FreeCashFlowOperatingCashFlowRatioTTM float64 `json:"freeCashFlowOperatingCashFlowRatioTTM" csv:"freeCashFlowOperatingCashFlowRatioTTM"`
	CashFlowCoverageRatiosTTM             float64 `json:"cashFlowCoverageRatiosTTM" csv:"cashFlowCoverageRatiosTTM"`
	ShortTermCoverageRatiosTTM            float64 `json:"shortTermCoverageRatiosTTM" csv:"shortTermCoverageRatiosTTM"`
	CapitalExpenditureCoverageRatioTTM    float64 `json:"capitalExpenditureCoverageRatioTTM" csv:"capitalExpenditureCoverageRatioTTM"`
	DividendPaidAndCapexCoverageRatioTTM  float64 `json:"dividendPaidAndCapexCoverageRatioTTM" csv:"dividendPaidAndCapexCoverageRatioTTM"`
	PriceBookValueRatioTTM                float64 `json:"priceBookValueRatioTTM" csv:"priceBookValueRatioTTM"`
	PriceToBookRatioTTM                   float64 `json:"priceToBookRatioTTM" csv:"priceToBookRatioTTM"`
	PriceToSalesRatioTTM                  float64 `json:"priceToSalesRatioTTM" csv:"priceToSalesRatioTTM"`
	PriceEarningsRatioTTM                 float64 `json:"priceEarningsRatioTTM" csv:"priceEarningsRatioTTM"`
	PriceToFreeCashFlowsRatioTTM          float64 `json:"priceToFreeCashFlowsRatioTTM" csv:"priceToFreeCashFlowsRatioTTM"`
	PriceToOperatingCashFlowsRatioTTM     float64 `json:"priceToOperatingCashFlowsRatioTTM" csv:"priceToOperatingCashFlowsRatioTTM"`
	PriceCashFlowRatioTTM                 float64 `json:"priceCashFlowRatioTTM" csv:"priceCashFlowRatioTTM"`
	PriceEarningsToGrowthRatioTTM         float64 `json:"priceEarningsToGrowthRatioTTM" csv:"priceEarningsToGrowthRatioTTM"`
	PriceSalesRatioTTM                    float64 `json:"priceSalesRatioTTM" csv:"priceSalesRatioTTM"`
	DividendYieldTTM                      float64 `json:"dividendYieldTTM" csv:"dividendYieldTTM"`
	EnterpriseValueMultipleTTM            float64 `json:"enterpriseValueMultipleTTM" csv:"enterpriseValueMultipleTTM"`
	PriceFairValueTTM                     float64 `json:"priceFairValueTTM" csv:"priceFairValueTTM"`
	DividendPerShareTTM                   float64 `json:"dividendPerShareTTM" csv:"dividendPerShareTTM"`
}

// KeyMetricsTTM ...
type KeyMetricsTTM struct {
	RevenuePerShareTTM                        float64 `json:"revenuePerShareTTM"`
	NetIncomePerShareTTM                      float64 `json:"netIncomePerShareTTM"`
	OperatingCashFlowPerShareTTM              float64 `json:"operatingCashFlowPerShareTTM"`
	FreeCashFlowPerShareTTM                   float64 `json:"freeCashFlowPerShareTTM"`
	CashPerShareTTM                           float64 `json:"cashPerShareTTM"`
	BookValuePerShareTTM                      float64 `json:"bookValuePerShareTTM"`
	TangibleBookValuePerShareTTM              float64 `json:"tangibleBookValuePerShareTTM"`
	ShareholdersEquityPerShareTTM             float64 `json:"shareholdersEquityPerShareTTM"`
	InterestDebtPerShareTTM                   float64 `json:"interestDebtPerShareTTM"`
	MarketCapTTM                              float64 `json:"marketCapTTM"`
	EnterpriseValueTTM                        float64 `json:"enterpriseValueTTM"`
	PeRatioTTM                                float64 `json:"peRatioTTM"`
	PriceToSalesRatioTTM                      float64 `json:"priceToSalesRatioTTM"`
	PocfratioTTM                              float64 `json:"pocfratioTTM"`
	PfcfRatioTTM                              float64 `json:"pfcfRatioTTM"`
	PbRatioTTM                                float64 `json:"pbRatioTTM"`
	PtbRatioTTM                               float64 `json:"ptbRatioTTM"`
	EvToSalesTTM                              float64 `json:"evToSalesTTM"`
	EnterpriseValueOverEBITDATTM              float64 `json:"enterpriseValueOverEBITDATTM"`
	EvToOperatingCashFlowTTM                  float64 `json:"evToOperatingCashFlowTTM"`
	EvToFreeCashFlowTTM                       float64 `json:"evToFreeCashFlowTTM"`
	EarningsYieldTTM                          float64 `json:"earningsYieldTTM"`
	FreeCashFlowYieldTTM                      float64 `json:"freeCashFlowYieldTTM"`
	DebtToEquityTTM                           float64 `json:"debtToEquityTTM"`
	DividendPerShareTTM                       float64 `json:"dividendPerShareTTM"`
	DebtToAssetsTTM                           float64 `json:"debtToAssetsTTM"`
	NetDebtToEBITDATTM                        float64 `json:"netDebtToEBITDATTM"`
	CurrentRatioTTM                           float64 `json:"currentRatioTTM"`
	InterestCoverageTTM                       float64 `json:"interestCoverageTTM"`
	IncomeQualityTTM                          float64 `json:"incomeQualityTTM"`
	DividendYieldTTM                          float64 `json:"dividendYieldTTM"`
	DividendYieldPercentageTTM                float64 `json:"dividendYieldPercentageTTM"`
	PayoutRatioTTM                            float64 `json:"payoutRatioTTM"`
	SalesGeneralAndAdministrativeToRevenueTTM float64 `json:"salesGeneralAndAdministrativeToRevenueTTM"`
	ResearchAndDevelopementToRevenueTTM       float64 `json:"researchAndDevelopementToRevenueTTM"`
	IntangiblesToTotalAssetsTTM               float64 `json:"intangiblesToTotalAssetsTTM"`
	CapexToOperatingCashFlowTTM               float64 `json:"capexToOperatingCashFlowTTM"`
	CapexToRevenueTTM                         float64 `json:"capexToRevenueTTM"`
	CapexToDepreciationTTM                    float64 `json:"capexToDepreciationTTM"`
	StockBasedCompensationToRevenueTTM        float64 `json:"stockBasedCompensationToRevenueTTM"`
	GrahamNumberTTM                           float64 `json:"grahamNumberTTM"`
	RoicTTM                                   float64 `json:"roicTTM"`
	ReturnOnTangibleAssetsTTM                 float64 `json:"returnOnTangibleAssetsTTM"`
	GrahamNetNetTTM                           float64 `json:"grahamNetNetTTM"`
	WorkingCapitalTTM                         float64 `json:"workingCapitalTTM"`
	TangibleAssetValueTTM                     float64 `json:"tangibleAssetValueTTM"`
	NetCurrentAssetValueTTM                   float64 `json:"netCurrentAssetValueTTM"`
	InvestedCapitalTTM                        float64 `json:"investedCapitalTTM"`
	AverageReceivablesTTM                     float64 `json:"averageReceivablesTTM"`
	AveragePayablesTTM                        float64 `json:"averagePayablesTTM"`
	AverageInventoryTTM                       float64 `json:"averageInventoryTTM"`
	DaysSalesOutstandingTTM                   float64 `json:"daysSalesOutstandingTTM"`
	DaysPayablesOutstandingTTM                float64 `json:"daysPayablesOutstandingTTM"`
	DaysOfInventoryOnHandTTM                  float64 `json:"daysOfInventoryOnHandTTM"`
	ReceivablesTurnoverTTM                    float64 `json:"receivablesTurnoverTTM"`
	PayablesTurnoverTTM                       float64 `json:"payablesTurnoverTTM"`
	InventoryTurnoverTTM                      float64 `json:"inventoryTurnoverTTM"`
	RoeTTM                                    float64 `json:"roeTTM"`
	CapexPerShareTTM                          float64 `json:"capexPerShareTTM"`
}

// KeyMetrics ...
type KeyMetrics struct {
	Symbol                                 string  `json:"symbol" csv:"symbol"`
	Date                                   string  `json:"date" csv:"date"`
	Period                                 string  `json:"period" csv:"period"`
	RevenuePerShare                        float64 `json:"revenuePerShare" csv:"revenuePerShare"`
	NetIncomePerShare                      float64 `json:"netIncomePerShare" csv:"netIncomePerShare"`
	OperatingCashFlowPerShare              float64 `json:"operatingCashFlowPerShare" csv:"operatingCashFlowPerShare"`
	FreeCashFlowPerShare                   float64 `json:"freeCashFlowPerShare" csv:"freeCashFlowPerShare"`
	CashPerShare                           float64 `json:"cashPerShare" csv:"cashPerShare"`
	BookValuePerShare                      float64 `json:"bookValuePerShare" csv:"bookValuePerShare"`
	TangibleBookValuePerShare              float64 `json:"tangibleBookValuePerShare" csv:"tangibleBookValuePerShare"`
	ShareholdersEquityPerShare             float64 `json:"shareholdersEquityPerShare" csv:"shareholdersEquityPerShare"`
	InterestDebtPerShare                   float64 `json:"interestDebtPerShare" csv:"interestDebtPerShare"`
	MarketCap                              float64 `json:"marketCap" csv:"marketCap"`
	EnterpriseValue                        float64 `json:"enterpriseValue" csv:"enterpriseValue"`
	PeRatio                                float64 `json:"peRatio" csv:"peRatio"`
	PriceToSalesRatio                      float64 `json:"priceToSalesRatio" csv:"priceToSalesRatio"`
	Pocfratio                              float64 `json:"pocfratio" csv:"pocfratio"`
	PfcfRatio                              float64 `json:"pfcfRatio" csv:"pfcfRatio"`
	PbRatio                                float64 `json:"pbRatio" csv:"pbRatio"`
	PtbRatio                               float64 `json:"ptbRatio" csv:"ptbRatio"`
	EvToSales                              float64 `json:"evToSales" csv:"evToSales"`
	EnterpriseValueOverEBITDA              float64 `json:"enterpriseValueOverEBITDA" csv:"enterpriseValueOverEBITDA"`
	EvToOperatingCashFlow                  float64 `json:"evToOperatingCashFlow" csv:"evToOperatingCashFlow"`
	EvToFreeCashFlow                       float64 `json:"evToFreeCashFlow" csv:"evToFreeCashFlow"`
	EarningsYield                          float64 `json:"earningsYield" csv:"earningsYield"`
	FreeCashFlowYield                      float64 `json:"freeCashFlowYield" csv:"freeCashFlowYield"`
	DebtToEquity                           float64 `json:"debtToEquity" csv:"debtToEquity"`
	DebtToAssets                           float64 `json:"debtToAssets" csv:"debtToAssets"`
	NetDebtToEBITDA                        float64 `json:"netDebtToEBITDA" csv:"netDebtToEBITDA"`
	CurrentRatio                           float64 `json:"currentRatio" csv:"currentRatio"`
	InterestCoverage                       float64 `json:"interestCoverage" csv:"interestCoverage"`
	IncomeQuality                          float64 `json:"incomeQuality" csv:"incomeQuality"`
	DividendYield                          float64 `json:"dividendYield" csv:"dividendYield"`
	PayoutRatio                            float64 `json:"payoutRatio" csv:"payoutRatio"`
	SalesGeneralAndAdministrativeToRevenue float64 `json:"salesGeneralAndAdministrativeToRevenue" csv:"salesGeneralAndAdministrativeToRevenue"`
	ResearchAndDdevelopementToRevenue      float64 `json:"researchAndDdevelopementToRevenue" csv:"researchAndDdevelopementToRevenue"`
	IntangiblesToTotalAssets               float64 `json:"intangiblesToTotalAssets" csv:"intangiblesToTotalAssets"`
	CapexToOperatingCashFlow               float64 `json:"capexToOperatingCashFlow" csv:"capexToOperatingCashFlow"`
	CapexToRevenue                         float64 `json:"capexToRevenue" csv:"capexToRevenue"`
	CapexToDepreciation                    float64 `json:"capexToDepreciation" csv:"capexToDepreciation"`
	StockBasedCompensationToRevenue        float64 `json:"stockBasedCompensationToRevenue" csv:"stockBasedCompensationToRevenue"`
	GrahamNumber                           float64 `json:"grahamNumber" csv:"grahamNumber"`
	Roic                                   float64 `json:"roic" csv:"roic"`
	ReturnOnTangibleAssets                 float64 `json:"returnOnTangibleAssets" csv:"returnOnTangibleAssets"`
	GrahamNetNet                           float64 `json:"grahamNetNet" csv:"grahamNetNet"`
	WorkingCapital                         float64 `json:"workingCapital" csv:"workingCapital"`
	TangibleAssetValue                     float64 `json:"tangibleAssetValue" csv:"tangibleAssetValue"`
	NetCurrentAssetValue                   float64 `json:"netCurrentAssetValue" csv:"netCurrentAssetValue"`
	InvestedCapital                        float64 `json:"investedCapital" csv:"investedCapital"`
	AverageReceivables                     float64 `json:"averageReceivables" csv:"averageReceivables"`
	AveragePayables                        float64 `json:"averagePayables" csv:"averagePayables"`
	AverageInventory                       float64 `json:"averageInventory" csv:"averageInventory"`
	DaysSalesOutstanding                   float64 `json:"daysSalesOutstanding" csv:"daysSalesOutstanding"`
	DaysPayablesOutstanding                float64 `json:"daysPayablesOutstanding" csv:"daysPayablesOutstanding"`
	DaysOfInventoryOnHand                  float64 `json:"daysOfInventoryOnHand" csv:"daysOfInventoryOnHand"`
	ReceivablesTurnover                    float64 `json:"receivablesTurnover" csv:"receivablesTurnover"`
	PayablesTurnover                       float64 `json:"payablesTurnover" csv:"payablesTurnover"`
	InventoryTurnover                      float64 `json:"inventoryTurnover" csv:"inventoryTurnover"`
	Roe                                    float64 `json:"roe" csv:"roe"`
	CapexPerShare                          float64 `json:"capexPerShare" csv:"capexPerShare"`
}

// EnterpriseValue ...
type EnterpriseValue struct {
	Symbol                      string  `json:"symbol"`
	Date                        string  `json:"date"` // 2018-09-29
	StockPrice                  float64 `json:"stockPrice"`
	NumberOfShares              int64   `json:"numberOfShares"`
	MarketCapitalization        float64 `json:"marketCapitalization"`
	MinusCashAndCashEquivalents int64   `json:"minusCashAndCashEquivalents"`
	AddTotalDebt                int64   `json:"addTotalDebt"`
	EnterpriseValue             float64 `json:"enterpriseValue"`
}

// FinancialStatementsGrowth ...
type FinancialStatementsGrowth struct {
	Symbol                                 string  `json:"symbol"`
	Date                                   string  `json:"date"` // 2019-09-28
	Period                                 string  `json:"period"`
	RevenueGrowth                          float64 `json:"revenueGrowth"`
	GrossProfitGrowth                      float64 `json:"grossProfitGrowth"`
	Ebitgrowth                             float64 `json:"ebitgrowth"`
	OperatingIncomeGrowth                  float64 `json:"operatingIncomeGrowth"`
	NetIncomeGrowth                        float64 `json:"netIncomeGrowth"`
	Epsgrowth                              float64 `json:"epsgrowth"`
	EpsdilutedGrowth                       float64 `json:"epsdilutedGrowth"`
	WeightedAverageSharesGrowth            float64 `json:"weightedAverageSharesGrowth"`
	WeightedAverageSharesDilutedGrowth     float64 `json:"weightedAverageSharesDilutedGrowth"`
	DividendsperShareGrowth                float64 `json:"dividendsperShareGrowth"`
	OperatingCashFlowGrowth                float64 `json:"operatingCashFlowGrowth"`
	FreeCashFlowGrowth                     float64 `json:"freeCashFlowGrowth"`
	TenYRevenueGrowthPerShare              float64 `json:"tenYRevenueGrowthPerShare"`
	FiveYRevenueGrowthPerShare             float64 `json:"fiveYRevenueGrowthPerShare"`
	ThreeYRevenueGrowthPerShare            float64 `json:"threeYRevenueGrowthPerShare"`
	TenYOperatingCFGrowthPerShare          float64 `json:"tenYOperatingCFGrowthPerShare"`
	FiveYOperatingCFGrowthPerShare         float64 `json:"fiveYOperatingCFGrowthPerShare"`
	ThreeYOperatingCFGrowthPerShare        float64 `json:"threeYOperatingCFGrowthPerShare"`
	TenYNetIncomeGrowthPerShare            float64 `json:"tenYNetIncomeGrowthPerShare"`
	FiveYNetIncomeGrowthPerShare           float64 `json:"fiveYNetIncomeGrowthPerShare"`
	ThreeYNetIncomeGrowthPerShare          float64 `json:"threeYNetIncomeGrowthPerShare"`
	TenYShareholdersEquityGrowthPerShare   float64 `json:"tenYShareholdersEquityGrowthPerShare"`
	FiveYShareholdersEquityGrowthPerShare  float64 `json:"fiveYShareholdersEquityGrowthPerShare"`
	ThreeYShareholdersEquityGrowthPerShare float64 `json:"threeYShareholdersEquityGrowthPerShare"`
	TenYDividendperShareGrowthPerShare     float64 `json:"tenYDividendperShareGrowthPerShare"`
	FiveYDividendperShareGrowthPerShare    float64 `json:"fiveYDividendperShareGrowthPerShare"`
	ThreeYDividendperShareGrowthPerShare   float64 `json:"threeYDividendperShareGrowthPerShare"`
	ReceivablesGrowth                      float64 `json:"receivablesGrowth"`
	InventoryGrowth                        float64 `json:"inventoryGrowth"`
	AssetGrowth                            float64 `json:"assetGrowth"`
	BookValueperShareGrowth                float64 `json:"bookValueperShareGrowth"`
	DebtGrowth                             float64 `json:"debtGrowth"`
	RdexpenseGrowth                        float64 `json:"rdexpenseGrowth"`
	SgaexpensesGrowth                      float64 `json:"sgaexpensesGrowth"`
}

// DiscountedCashFlow ...
type DiscountedCashFlow struct {
	Symbol     string  `json:"symbol"`
	Date       string  `json:"date"`
	Dcf        float64 `json:"dcf"`
	StockPrice float64 `json:"Stock Price"`
}

// DailyDiscountedCashFlow ...
type DailyDiscountedCashFlow struct {
	Symbol string  `json:"symbol" csv:"symbol"`
	Date   string  `json:"date" csv:"date"`
	Dcf    float64 `json:"dcf" csv:"DCF"`
}

// HistoryDiscountedCashFlow ...
type HistoryDiscountedCashFlow struct {
	Symbol string  `json:"symbol"`
	Date   string  `json:"date"`
	Price  float64 `json:"price"`
	Dcf    float64 `json:"dcf"`
}

// Rating ...
type Rating struct {
	Symbol                         string `json:"symbol" csv:"symbol"`
	Date                           string `json:"date" csv:"date"` // 2020-09-14
	Rating                         string `json:"rating" csv:"rating"`
	RatingScore                    int64  `json:"ratingScore" csv:"ratingScore"`
	RatingRecommendation           string `json:"ratingRecommendation" csv:"ratingRecommendation"`
	RatingDetailsDCFScore          int64  `json:"ratingDetailsDCFScore" csv:"ratingDetailsDCFScore"`
	RatingDetailsDCFRecommendation string `json:"ratingDetailsDCFRecommendation" csv:"ratingDetailsDCFRecommendation"`
	RatingDetailsROEScore          int64  `json:"ratingDetailsROEScore" csv:"ratingDetailsROEScore"`
	RatingDetailsROERecommendation string `json:"ratingDetailsROERecommendation" csv:"ratingDetailsROERecommendation"`
	RatingDetailsROAScore          int64  `json:"ratingDetailsROAScore" csv:"ratingDetailsROAScore"`
	RatingDetailsROARecommendation string `json:"ratingDetailsROARecommendation" csv:"ratingDetailsROARecommendation"`
	RatingDetailsDEScore           int64  `json:"ratingDetailsDEScore" csv:"ratingDetailsDEScore"`
	RatingDetailsDERecommendation  string `json:"ratingDetailsDERecommendation" csv:"ratingDetailsDERecommendation"`
	RatingDetailsPEScore           int64  `json:"ratingDetailsPEScore" csv:"ratingDetailsPEScore"`
	RatingDetailsPERecommendation  string `json:"ratingDetailsPERecommendation" csv:"ratingDetailsPERecommendation"`
	RatingDetailsPBScore           int64  `json:"ratingDetailsPBScore" csv:"ratingDetailsPBScore"`
	RatingDetailsPBRecommendation  string `json:"ratingDetailsPBRecommendation" csv:"ratingDetailsPBRecommendation"`
}

// MarketCapitalization ...
type MarketCapitalization struct {
	Symbol    string  `json:"symbol"`
	Date      string  `json:"date"`
	MarketCap float64 `json:"marketCap"`
}

// DelstedCompany ...
type DelstedCompany struct {
	Symbol       string `json:"symbol"`
	CompanyName  string `json:"companyName"`
	Exchange     string `json:"exchange"`
	IpoDate      string `json:"ipoDate"`
	DelistedDate string `json:"delistedDate"`
}

// StockNews ...
type StockNews struct {
	Symbol        string `json:"symbol"`
	PublishedDate string `json:"publishedDate"` // 2020-09-15 14:51:03
	Title         string `json:"title"`
	Image         string `json:"image"`
	Site          string `json:"site"`
	Text          string `json:"text"`
	URL           string `json:"url"`
}

// StockScreener ...
type StockScreener struct {
	Symbol             string  `json:"symbol"`
	CompanyName        string  `json:"companyName"`
	MarketCap          int64   `json:"marketCap"`
	Sector             string  `json:"sector"`
	Industry           string  `json:"industry"`
	Beta               float64 `json:"beta"`
	Price              float64 `json:"price"`
	LastAnnualDividend float64 `json:"lastAnnualDividend"`
	Volume             float64 `json:"volume"`
	Exchange           string  `json:"exchange"`
	ExchangeShortName  string  `json:"exchangeShortName"`
}

// AnalystEstimates ...
type AnalystEstimates struct {
	Symbol                        string  `json:"symbol"`
	Date                          string  `json:"date"` // 2020-09-15
	EstimatedRevenueLow           float64 `json:"estimatedRevenueLow"`
	EstimatedRevenueHigh          float64 `json:"estimatedRevenueHigh"`
	EstimatedRevenueAvg           float64 `json:"estimatedRevenueAvg"`
	EstimatedEbitdaLow            float64 `json:"estimatedEbitdaLow"`
	EstimatedEbitdaHigh           float64 `json:"estimatedEbitdaHigh"`
	EstimatedEbitdaAvg            float64 `json:"estimatedEbitdaAvg"`
	EstimatedEbitLow              float64 `json:"estimatedEbitLow"`
	EstimatedEbitHigh             float64 `json:"estimatedEbitHigh"`
	EstimatedEbitAvg              float64 `json:"estimatedEbitAvg"`
	EstimatedNetIncomeLow         float64 `json:"estimatedNetIncomeLow"`
	EstimatedNetIncomeHigh        float64 `json:"estimatedNetIncomeHigh"`
	EstimatedNetIncomeAvg         float64 `json:"estimatedNetIncomeAvg"`
	EstimatedSgaExpenseLow        float64 `json:"estimatedSgaExpenseLow"`
	EstimatedSgaExpenseHigh       float64 `json:"estimatedSgaExpenseHigh"`
	EstimatedSgaExpenseAvg        float64 `json:"estimatedSgaExpenseAvg"`
	EstimatedEpsAvg               float64 `json:"estimatedEpsAvg"`
	EstimatedEpsHigh              float64 `json:"estimatedEpsHigh"`
	EstimatedEpsLow               float64 `json:"estimatedEpsLow"`
	NumberAnalystEstimatedRevenue int64   `json:"numberAnalystEstimatedRevenue"`
	NumberAnalystsEstimatedEps    int64   `json:"numberAnalystsEstimatedEps"`
}

// Grade ...
type Grade struct {
	Symbol         string    `json:"symbol"`
	Date           string    `json:"date"` // 2020-09-22
	GradingCompany string    `json:"gradingCompany"`
	PreviousGrade  GradeType `json:"previousGrade"`
	NewGrade       GradeType `json:"newGrade"`
}

// AnalystStockRecommendations ...
type AnalystStockRecommendations struct {
	Symbol                   string `json:"symbol"`
	Date                     string `json:"date"` // 2020-08-01
	AnalystRatingsbuy        int64  `json:"analystRatingsbuy"`
	AnalystRatingsHold       int64  `json:"analystRatingsHold"`
	AnalystRatingsSell       int64  `json:"analystRatingsSell"`
	AnalystRatingsStrongSell int64  `json:"analystRatingsStrongSell"`
	AnalystRatingsStrongBuy  int64  `json:"analystRatingsStrongBuy"`
}

// PressReleases ...
type PressReleases struct {
	Symbol string `json:"symbol"`
	Date   string `json:"date"` // 2020-09-15 16:30:00
	Title  string `json:"title"`
	Text   string `json:"text"`
}

// EconomicCalendarEventList ...
type EconomicCalendarEventList struct {
	Event   string `json:"event"`
	Country string `json:"country"`
}

// EconomicCalendar ...
type EconomicCalendar struct {
	Event            string   `json:"event"`
	Date             string   `json:"date"` // 2020-09-16 11:00:00
	Country          string   `json:"country"`
	Impact           *string  `json:"impact"`
	Actual           *float64 `json:"actual"`
	Previous         *float64 `json:"previous"`
	Change           *float64 `json:"change"`
	ChangePercentage *float64 `json:"changePercentage"`
	Estimate         *float64 `json:"estimate"`
}

// EarningSurprise ...
type EarningSurprise struct {
	Date                string  `json:"date" csv:"date"` // 2020-09-08
	Symbol              string  `json:"symbol" csv:"symbol"`
	ActualEarningResult float64 `json:"actualEarningResult" csv:"actualEarningResult"`
	EstimatedEarning    float64 `json:"estimatedEarning" csv:"estimatedEarning"`
}

// HistoryEconomicCalendar ...
type HistoryEconomicCalendar struct {
	Event            string  `json:"event"`
	Date             string  `json:"date"` // 2020-09-16
	Country          string  `json:"country"`
	Actual           float64 `json:"actual"`
	Previous         float64 `json:"previous"`
	Change           float64 `json:"change"`
	ChangePercentage float64 `json:"changePercentage"`
}

// SECFiling ...
type SECFiling struct {
	Symbol       string `json:"symbol"`
	FillingDate  string `json:"fillingDate"`
	AcceptedDate string `json:"acceptedDate"`
	Cik          string `json:"cik"`
	Type         string `json:"type"`
	Link         string `json:"link"`
	FinalLink    string `json:"finalLink"`
}

type CompanyOutlook struct {
	Profile struct {
		Symbol            string  `json:"symbol"`
		Range             string  `json:"range"`
		CompanyName       string  `json:"companyName"`
		Currency          string  `json:"currency"`
		Cik               string  `json:"cik"`
		Isin              string  `json:"isin"`
		Cusip             string  `json:"cusip"`
		Exchange          string  `json:"exchange"`
		ExchangeShortName string  `json:"exchangeShortName"`
		Industry          string  `json:"industry"`
		Website           string  `json:"website"`
		Description       string  `json:"description"`
		Ceo               string  `json:"ceo"`
		Sector            string  `json:"sector"`
		Country           string  `json:"country"`
		FullTimeEmployees string  `json:"fullTimeEmployees"`
		Phone             string  `json:"phone"`
		Address           string  `json:"address"`
		City              string  `json:"city"`
		State             string  `json:"state"`
		Zip               string  `json:"zip"`
		Image             string  `json:"image"`
		IpoDate           string  `json:"ipoDate"`
		Price             float64 `json:"price"`
		Beta              float64 `json:"beta"`
		VolAvg            float64 `json:"volAvg"`
		MktCap            float64 `json:"mktCap"`
		LastDiv           float64 `json:"lastDiv"`
		Changes           float64 `json:"changes"`
		DcfDiff           float64 `json:"dcfDiff"`
		Dcf               float64 `json:"dcf"`
		DefaultImage      bool    `json:"defaultImage"`
		IsEtf             bool    `json:"isEtf"`
		IsActivelyTrading bool    `json:"isActivelyTrading"`
	} `json:"profile"`
	InsideTrades []struct {
		Symbol                  string  `json:"symbol"`
		TransactionDate         string  `json:"transactionDate"`
		ReportingCik            string  `json:"reportingCik"`
		TransactionType         string  `json:"transactionType"`
		CompanyCik              string  `json:"companyCik"`
		ReportingName           string  `json:"reportingName"`
		TypeOfOwner             string  `json:"typeOfOwner"`
		AcquistionOrDisposition string  `json:"acquistionOrDisposition"`
		FormType                string  `json:"formType"`
		SecurityName            string  `json:"securityName"`
		Link                    string  `json:"link"`
		SecuritiesOwned         float64 `json:"securitiesOwned"`
		SecuritiesTransacted    float64 `json:"securitiesTransacted"`
		Price                   float64 `json:"price"`
	} `json:"insideTrades"`
	KeyExecutives []struct {
		Title       string  `json:"title"`
		Name        string  `json:"name"`
		Gender      string  `json:"gender"`
		CurrencyPay string  `json:"currencyPay"`
		Pay         float64 `json:"pay"`
	} `json:"keyExecutives"`
	SplitHistory []struct {
		Date        string  `json:"date"`
		Label       string  `json:"label"`
		Numerator   float64 `json:"numerator"`
		Denominator float64 `json:"denominator"`
	} `json:"splitHistory"`
	StockDividend []struct {
		Date            string  `json:"date"`
		Label           string  `json:"label"`
		RecordDate      string  `json:"recordDate"`
		PaymentDate     string  `json:"paymentDate"`
		DeclarationDate string  `json:"declarationDate"`
		AdjDividend     float64 `json:"adjDividend"`
		Dividend        float64 `json:"dividend"`
	} `json:"stockDividend"`
	StockNews []struct {
		Symbol        string `json:"symbol"`
		PublishedDate string `json:"publishedDate"`
		Title         string `json:"title"`
		Image         string `json:"image"`
		Site          string `json:"site"`
		Text          string `json:"text"`
		URL           string `json:"url"`
	} `json:"stockNews"`
	Metrics struct {
		DividendYielTTM float64 `json:"dividendYielTTM"`
		Volume          float64 `json:"volume"`
		YearHigh        float64 `json:"yearHigh"`
		YearLow         float64 `json:"yearLow"`
	} `json:"metrics"`
	Ratios []struct {
		DividendYielTTM                       float64 `json:"dividendYielTTM"`
		DividendYielPercentageTTM             float64 `json:"dividendYielPercentageTTM"`
		PeRatioTTM                            float64 `json:"peRatioTTM"`
		PegRatioTTM                           float64 `json:"pegRatioTTM"`
		PayoutRatioTTM                        float64 `json:"payoutRatioTTM"`
		CurrentRatioTTM                       float64 `json:"currentRatioTTM"`
		QuickRatioTTM                         float64 `json:"quickRatioTTM"`
		CashRatioTTM                          float64 `json:"cashRatioTTM"`
		DaysOfSalesOutstandingTTM             float64 `json:"daysOfSalesOutstandingTTM"`
		DaysOfInventoryOutstandingTTM         float64 `json:"daysOfInventoryOutstandingTTM"`
		OperatingCycleTTM                     float64 `json:"operatingCycleTTM"`
		DaysOfPayablesOutstandingTTM          float64 `json:"daysOfPayablesOutstandingTTM"`
		CashConversionCycleTTM                float64 `json:"cashConversionCycleTTM"`
		GrossProfitMarginTTM                  float64 `json:"grossProfitMarginTTM"`
		OperatingProfitMarginTTM              float64 `json:"operatingProfitMarginTTM"`
		PretaxProfitMarginTTM                 float64 `json:"pretaxProfitMarginTTM"`
		NetProfitMarginTTM                    float64 `json:"netProfitMarginTTM"`
		EffectiveTaxRateTTM                   float64 `json:"effectiveTaxRateTTM"`
		ReturnOnAssetsTTM                     float64 `json:"returnOnAssetsTTM"`
		ReturnOnEquityTTM                     float64 `json:"returnOnEquityTTM"`
		ReturnOnCapitalEmployedTTM            float64 `json:"returnOnCapitalEmployedTTM"`
		NetIncomePerEBTTTM                    float64 `json:"netIncomePerEBTTTM"`
		EbtPerEbitTTM                         float64 `json:"ebtPerEbitTTM"`
		EbitPerRevenueTTM                     float64 `json:"ebitPerRevenueTTM"`
		DebtRatioTTM                          float64 `json:"debtRatioTTM"`
		DebtEquityRatioTTM                    float64 `json:"debtEquityRatioTTM"`
		LongTermDebtToCapitalizationTTM       float64 `json:"longTermDebtToCapitalizationTTM"`
		TotalDebtToCapitalizationTTM          float64 `json:"totalDebtToCapitalizationTTM"`
		InterestCoverageTTM                   float64 `json:"interestCoverageTTM"`
		CashFlowToDebtRatioTTM                float64 `json:"cashFlowToDebtRatioTTM"`
		CompanyEquityMultiplierTTM            float64 `json:"companyEquityMultiplierTTM"`
		ReceivablesTurnoverTTM                float64 `json:"receivablesTurnoverTTM"`
		PayablesTurnoverTTM                   float64 `json:"payablesTurnoverTTM"`
		InventoryTurnoverTTM                  float64 `json:"inventoryTurnoverTTM"`
		FixedAssetTurnoverTTM                 float64 `json:"fixedAssetTurnoverTTM"`
		AssetTurnoverTTM                      float64 `json:"assetTurnoverTTM"`
		OperatingCashFlowPerShareTTM          float64 `json:"operatingCashFlowPerShareTTM"`
		FreeCashFlowPerShareTTM               float64 `json:"freeCashFlowPerShareTTM"`
		CashPerShareTTM                       float64 `json:"cashPerShareTTM"`
		OperatingCashFlowSalesRatioTTM        float64 `json:"operatingCashFlowSalesRatioTTM"`
		FreeCashFlowOperatingCashFlowRatioTTM float64 `json:"freeCashFlowOperatingCashFlowRatioTTM"`
		CashFlowCoverageRatiosTTM             float64 `json:"cashFlowCoverageRatiosTTM"`
		ShortTermCoverageRatiosTTM            float64 `json:"shortTermCoverageRatiosTTM"`
		CapitalExpenditureCoverageRatioTTM    float64 `json:"capitalExpenditureCoverageRatioTTM"`
		DividendPaidAndCapexCoverageRatioTTM  float64 `json:"dividendPaidAndCapexCoverageRatioTTM"`
		PriceBookValueRatioTTM                float64 `json:"priceBookValueRatioTTM"`
		PriceToBookRatioTTM                   float64 `json:"priceToBookRatioTTM"`
		PriceToSalesRatioTTM                  float64 `json:"priceToSalesRatioTTM"`
		PriceEarningsRatioTTM                 float64 `json:"priceEarningsRatioTTM"`
		PriceToFreeCashFlowsRatioTTM          float64 `json:"priceToFreeCashFlowsRatioTTM"`
		PriceToOperatingCashFlowsRatioTTM     float64 `json:"priceToOperatingCashFlowsRatioTTM"`
		PriceCashFlowRatioTTM                 float64 `json:"priceCashFlowRatioTTM"`
		PriceEarningsToGrowthRatioTTM         float64 `json:"priceEarningsToGrowthRatioTTM"`
		PriceSalesRatioTTM                    float64 `json:"priceSalesRatioTTM"`
		DividendYieldTTM                      float64 `json:"dividendYieldTTM"`
		EnterpriseValueMultipleTTM            float64 `json:"enterpriseValueMultipleTTM"`
		PriceFairValueTTM                     float64 `json:"priceFairValueTTM"`
	} `json:"ratios"`
	Rating []struct {
		Symbol                         string `json:"symbol"`
		Date                           string `json:"date"`
		Rating                         string `json:"rating"`
		Ratingscore                    int    `json:"ratingScore"`
		Ratingrecommendation           string `json:"ratingRecommendation"`
		Ratingdetailsdcfscore          int    `json:"ratingDetailsDCFScore"`
		Ratingdetailsdcfrecommendation string `json:"ratingDetailsDCFRecommendation"`
		Ratingdetailsroescore          int    `json:"ratingDetailsROEScore"`
		Ratingdetailsroerecommendation string `json:"ratingDetailsROERecommendation"`
		Ratingdetailsroascore          int    `json:"ratingDetailsROAScore"`
		Ratingdetailsroarecommendation string `json:"ratingDetailsROARecommendation"`
		Ratingdetailsdescore           int    `json:"ratingDetailsDEScore"`
		Ratingdetailsderecommendation  string `json:"ratingDetailsDERecommendation"`
		Ratingdetailspescore           int    `json:"ratingDetailsPEScore"`
		Ratingdetailsperecommendation  string `json:"ratingDetailsPERecommendation"`
		Ratingdetailspbscore           int    `json:"ratingDetailsPBScore"`
		Ratingdetailspbrecommendation  string `json:"ratingDetailsPBRecommendation"`
	} `json:"rating"`
	FinancialsAnnual struct {
		Income []struct {
			Date                             string  `json:"date"`
			Symbol                           string  `json:"symbol"`
			ReportedCurrency                 string  `json:"reportedCurrency"`
			FillingDate                      string  `json:"fillingDate"`
			AcceptedDate                     string  `json:"acceptedDate"`
			Period                           string  `json:"period"`
			Link                             string  `json:"link"`
			FinalLink                        string  `json:"finalLink"`
			Revenue                          float64 `json:"revenue"`
			GrosspPofit                      float64 `json:"grossProfit"`
			CostOfRevenue                    float64 `json:"costOfRevenue"`
			GrosspPofitRatio                 float64 `json:"grossProfitRatio"`
			ResearchAndDevelopmentExpenses   float64 `json:"researchAndDevelopmentExpenses"`
			GeneralAndAdministrativeExpenses float64 `json:"generalAndAdministrativeExpenses"`
			SellingAndMarketingExpenses      float64 `json:"sellingAndMarketingExpenses"`
			OtherExpenses                    float64 `json:"otherExpenses"`
			OperatingExpenses                float64 `json:"operatingExpenses"`
			CostAndExpenses                  float64 `json:"costAndExpenses"`
			InterestExpense                  float64 `json:"interestExpense"`
			DepreciationAndAmortization      float64 `json:"depreciationAndAmortization"`
			Ebitda                           float64 `json:"ebitda"`
			EbitdaRatio                      float64 `json:"ebitdaratio"`
			OperatingIncome                  float64 `json:"operatingIncome"`
			OperatingIncomeRatio             float64 `json:"operatingIncomeRatio"`
			TotalOtherIncomeExpensesNet      float64 `json:"totalOtherIncomeExpensesNet"`
			IncomeBeforeTax                  float64 `json:"incomeBeforeTax"`
			IncomeBeforeTaxRatio             float64 `json:"incomeBeforeTaxRatio"`
			IncomeTaxExpense                 float64 `json:"incomeTaxExpense"`
			NetIncome                        float64 `json:"netIncome"`
			NetIncomeRatio                   float64 `json:"netIncomeRatio"`
			Eps                              float64 `json:"eps"`
			Epsdiluted                       float64 `json:"epsdiluted"`
			WeightedAverageShsOut            float64 `json:"weightedAverageShsOut"`
			WeightedAverageShsOutDil         float64 `json:"weightedAverageShsOutDil"`
		} `json:"income"`
		Balance []struct {
			Date                                    string  `json:"date"`
			Symbol                                  string  `json:"symbol"`
			ReportedCurrency                        string  `json:"reportedCurrency"`
			FillingDate                             string  `json:"fillingDate"`
			AcceptedDate                            string  `json:"acceptedDate"`
			Period                                  string  `json:"period"`
			Link                                    string  `json:"link"`
			FinalLink                               string  `json:"finalLink"`
			CashAndCashEquivalents                  float64 `json:"cashAndCashEquivalents"`
			ShortTermInvestments                    float64 `json:"shortTermInvestments"`
			CashAndShortTermInvestments             float64 `json:"cashAndShortTermInvestments"`
			NetReceivables                          float64 `json:"netReceivables"`
			Inventory                               float64 `json:"inventory"`
			OtherCurrentAssets                      float64 `json:"otherCurrentAssets"`
			TotalCurrentAssets                      float64 `json:"totalCurrentAssets"`
			PropertyPlantEquipmentNet               float64 `json:"propertyPlantEquipmentNet"`
			Goodwill                                float64 `json:"goodwill"`
			IntangibleAssets                        float64 `json:"intangibleAssets"`
			GoodwillAndIntangibleAssets             float64 `json:"goodwillAndIntangibleAssets"`
			LongTermInvestments                     float64 `json:"longTermInvestments"`
			TaxAssets                               float64 `json:"taxAssets"`
			OtherNonCurrentAssets                   float64 `json:"otherNonCurrentAssets"`
			TotalNonCurrentAssets                   float64 `json:"totalNonCurrentAssets"`
			OtherAssets                             float64 `json:"otherAssets"`
			TotalAssets                             float64 `json:"totalAssets"`
			AccountPayables                         float64 `json:"accountPayables"`
			ShortTermDebt                           float64 `json:"shortTermDebt"`
			TaxPayables                             float64 `json:"taxPayables"`
			DeferredRevenue                         float64 `json:"deferredRevenue"`
			OtherCurrentLiabilities                 float64 `json:"otherCurrentLiabilities"`
			TotalCurrentLiabilities                 float64 `json:"totalCurrentLiabilities"`
			LongTermDebt                            float64 `json:"longTermDebt"`
			DeferredRevenueNonCurrent               float64 `json:"deferredRevenueNonCurrent"`
			DeferredTaxLiabilitiesNonCurrent        float64 `json:"deferredTaxLiabilitiesNonCurrent"`
			OtherNonCurrentLiabilities              float64 `json:"otherNonCurrentLiabilities"`
			TotalNonCurrentLiabilities              float64 `json:"totalNonCurrentLiabilities"`
			OtherLiabilities                        float64 `json:"otherLiabilities"`
			TotalLiabilities                        float64 `json:"totalLiabilities"`
			CommonStock                             float64 `json:"commonStock"`
			RetainedEarnings                        float64 `json:"retainedEarnings"`
			AccumulatedOtherComprehensiveIncomeLoss float64 `json:"accumulatedOtherComprehensiveIncomeLoss"`
			OthertotalStockholdersEquity            float64 `json:"othertotalStockholdersEquity"`
			TotalStockholdersEquity                 float64 `json:"totalStockholdersEquity"`
			TotalLiabilitiesAndStockholdersEquity   float64 `json:"totalLiabilitiesAndStockholdersEquity"`
			TotalInvestments                        float64 `json:"totalInvestments"`
			TotalDebt                               float64 `json:"totalDebt"`
			NetDebt                                 float64 `json:"netDebt"`
		} `json:"balance"`
		Cash []struct {
			Date                                     string  `json:"date"`
			Symbol                                   string  `json:"symbol"`
			ReportedCurrency                         string  `json:"reportedCurrency"`
			FillingDate                              string  `json:"fillingDate"`
			AcceptedDate                             string  `json:"acceptedDate"`
			Period                                   string  `json:"period"`
			NetIncome                                float64 `json:"netIncome"`
			DepreciationAndAmortization              float64 `json:"depreciationAndAmortization"`
			DeferredIncomeTax                        float64 `json:"deferredIncomeTax"`
			StockBasedCompensation                   float64 `json:"stockBasedCompensation"`
			ChangeInWorkingCapital                   float64 `json:"changeInWorkingCapital"`
			AccountsReceivables                      float64 `json:"accountsReceivables"`
			Inventory                                float64 `json:"inventory"`
			AccountsPayables                         float64 `json:"accountsPayables"`
			OtherWorkingCapital                      float64 `json:"otherWorkingCapital"`
			OtherNonCashItems                        float64 `json:"otherNonCashItems"`
			NetCashProvidedByOperatingActivities     float64 `json:"netCashProvidedByOperatingActivities"`
			InvestmentsInPropertyPlantAndEquipment   float64 `json:"investmentsInPropertyPlantAndEquipment"`
			AcquisitionsNet                          float64 `json:"acquisitionsNet"`
			PurchasesOfInvestments                   float64 `json:"purchasesOfInvestments"`
			SalesMaturitiesOfInvestments             float64 `json:"salesMaturitiesOfInvestments"`
			OtherInvestingActivites                  float64 `json:"otherInvestingActivites"`
			NetCashUsedForInvestingActivites         float64 `json:"netCashUsedForInvestingActivites"`
			DebtRepayment                            float64 `json:"debtRepayment"`
			CommonStockIssued                        float64 `json:"commonStockIssued"`
			CommonStockRepurchased                   float64 `json:"commonStockRepurchased"`
			DividendsPaid                            float64 `json:"dividendsPaid"`
			OtherFinancingActivites                  float64 `json:"otherFinancingActivites"`
			NetCashUsedProvidedByFinancingActivities float64 `json:"netCashUsedProvidedByFinancingActivities"`
			EffectOfForexChangesOnCash               float64 `json:"effectOfForexChangesOnCash"`
			NetChangeInCash                          float64 `json:"netChangeInCash"`
			CashAtEndOfPeriod                        float64 `json:"cashAtEndOfPeriod"`
			CashAtBeginningOfPeriod                  float64 `json:"cashAtBeginningOfPeriod"`
			OperatingCashFlow                        float64 `json:"operatingCashFlow"`
			CapitalExpenditure                       float64 `json:"capitalExpenditure"`
			FreeCashFlow                             float64 `json:"freeCashFlow"`
			Link                                     string  `json:"link"`
			FinalLink                                string  `json:"finalLink"`
		} `json:"cash"`
	} `json:"financialsAnnual"`
	FinancialsQuarter struct {
		Income []struct {
			Date                             string  `json:"date"`
			Symbol                           string  `json:"symbol"`
			ReportedCurrency                 string  `json:"reportedCurrency"`
			FillingDate                      string  `json:"fillingDate"`
			AcceptedDate                     string  `json:"acceptedDate"`
			Period                           string  `json:"period"`
			Link                             string  `json:"link"`
			DinalLink                        string  `json:"finalLink"`
			Revenue                          float64 `json:"revenue"`
			CostOfRevenue                    float64 `json:"costOfRevenue"`
			GrossProfit                      float64 `json:"grossProfit"`
			GrossProfitRatio                 float64 `json:"grossProfitRatio"`
			ResearchAndDevelopmentExpenses   float64 `json:"researchAndDevelopmentExpenses"`
			GeneralAndAdministrativeExpenses float64 `json:"generalAndAdministrativeExpenses"`
			SellingAndMarketingExpenses      float64 `json:"sellingAndMarketingExpenses"`
			OtherExpenses                    float64 `json:"otherExpenses"`
			OperatingExpenses                float64 `json:"operatingExpenses"`
			CostAndExpenses                  float64 `json:"costAndExpenses"`
			InterestExpense                  float64 `json:"interestExpense"`
			DepreciationAndAmortization      float64 `json:"depreciationAndAmortization"`
			Ebitda                           float64 `json:"ebitda"`
			EbitdaRatio                      float64 `json:"ebitdaratio"`
			OperatingIncome                  float64 `json:"operatingIncome"`
			OperatingIncomeRatio             float64 `json:"operatingIncomeRatio"`
			TotalOtherIncomeExpensesNet      float64 `json:"totalOtherIncomeExpensesNet"`
			IncomeBeforeTax                  float64 `json:"incomeBeforeTax"`
			IncomeBeforeTaxRatio             float64 `json:"incomeBeforeTaxRatio"`
			IncomeTaxExpense                 float64 `json:"incomeTaxExpense"`
			NetIncome                        float64 `json:"netIncome"`
			NetIncomeRatio                   float64 `json:"netIncomeRatio"`
			Eps                              float64 `json:"eps"`
			Epsdiluted                       float64 `json:"epsdiluted"`
			WeightedAverageShsOut            float64 `json:"weightedAverageShsOut"`
			WeightedAverageShsOutDil         float64 `json:"weightedAverageShsOutDil"`
		} `json:"income"`
		Balance []struct {
			Date                                    string  `json:"date"`
			Symbol                                  string  `json:"symbol"`
			ReportedCurrency                        string  `json:"reportedCurrency"`
			FillingDate                             string  `json:"fillingDate"`
			AcceptedDate                            string  `json:"acceptedDate"`
			Period                                  string  `json:"period"`
			Link                                    string  `json:"link"`
			FinalLink                               string  `json:"finalLink"`
			CashAndCashEquivalents                  float64 `json:"cashAndCashEquivalents"`
			ShortTermInvestments                    float64 `json:"shortTermInvestments"`
			CashAndShortTermInvestments             float64 `json:"cashAndShortTermInvestments"`
			NetReceivables                          float64 `json:"netReceivables"`
			Inventory                               float64 `json:"inventory"`
			OtherCurrentAssets                      float64 `json:"otherCurrentAssets"`
			Totalcurrentassets                      float64 `json:"totalCurrentAssets"`
			TotalCurrentAssets                      float64 `json:"propertyPlantEquipmentNet"`
			Goodwill                                float64 `json:"goodwill"`
			IntangibleAssets                        float64 `json:"intangibleAssets"`
			GoodwillAndIntangibleAssets             float64 `json:"goodwillAndIntangibleAssets"`
			LongTermInvestments                     float64 `json:"longTermInvestments"`
			TaxAssets                               float64 `json:"taxAssets"`
			OtherNonCurrentAssets                   float64 `json:"otherNonCurrentAssets"`
			TotalNonCurrentAssets                   float64 `json:"totalNonCurrentAssets"`
			OtherAssets                             float64 `json:"otherAssets"`
			TotalAssets                             float64 `json:"totalAssets"`
			AccountPayables                         float64 `json:"accountPayables"`
			ShortTermDebt                           float64 `json:"shortTermDebt"`
			TaxPayables                             float64 `json:"taxPayables"`
			DeferredRevenue                         float64 `json:"deferredRevenue"`
			OtherCurrentLiabilities                 float64 `json:"otherCurrentLiabilities"`
			TotalCurrentLiabilities                 float64 `json:"totalCurrentLiabilities"`
			LongTermDebt                            float64 `json:"longTermDebt"`
			DeferredRevenueNonCurrent               float64 `json:"deferredRevenueNonCurrent"`
			DeferredTaxLiabilitiesNonCurrent        float64 `json:"deferredTaxLiabilitiesNonCurrent"`
			OtherNonCurrentLiabilities              float64 `json:"otherNonCurrentLiabilities"`
			TotalNonCurrentLiabilities              float64 `json:"totalNonCurrentLiabilities"`
			OtherLiabilities                        float64 `json:"otherLiabilities"`
			TotalLiabilities                        float64 `json:"totalLiabilities"`
			CommonStock                             float64 `json:"commonStock"`
			RetainedEarnings                        float64 `json:"retainedEarnings"`
			AccumulatedOtherComprehensiveIncomeLoss float64 `json:"accumulatedOtherComprehensiveIncomeLoss"`
			OtherTotalStockholdersEquity            float64 `json:"othertotalStockholdersEquity"`
			TotalStockholdersEquity                 float64 `json:"totalStockholdersEquity"`
			TotalLiabilitiesAndStockholdersEquity   float64 `json:"totalLiabilitiesAndStockholdersEquity"`
			TotalInvestments                        float64 `json:"totalInvestments"`
			TotalDebt                               float64 `json:"totalDebt"`
			NetDebt                                 float64 `json:"netDebt"`
		} `json:"balance"`
		Cash []struct {
			Date                                     string  `json:"date"`
			Symbol                                   string  `json:"symbol"`
			ReportedCurrency                         string  `json:"reportedCurrency"`
			FillingDate                              string  `json:"fillingDate"`
			AcceptedDate                             string  `json:"acceptedDate"`
			Period                                   string  `json:"period"`
			Link                                     string  `json:"link"`
			FinalLink                                string  `json:"finalLink"`
			NetIncome                                float64 `json:"netIncome"`
			DepreciationAndAmortization              float64 `json:"depreciationAndAmortization"`
			DeferredIncomeTax                        float64 `json:"deferredIncomeTax"`
			StockBasedCompensation                   float64 `json:"stockBasedCompensation"`
			ChangeInWorkingCapital                   float64 `json:"changeInWorkingCapital"`
			AccountsReceivables                      float64 `json:"accountsReceivables"`
			Inventory                                float64 `json:"inventory"`
			AccountsPayables                         float64 `json:"accountsPayables"`
			OtherWorkingCapital                      float64 `json:"otherWorkingCapital"`
			OtherNonCashItems                        float64 `json:"otherNonCashItems"`
			NetCashProvidedByOperatingActivities     float64 `json:"netCashProvidedByOperatingActivities"`
			Investmentsinpropertyplantandequipment   float64 `json:"investmentsInPropertyPlantAndEquipment"`
			AcquisitionsNet                          float64 `json:"acquisitionsNet"`
			PurchasesOfInvestments                   float64 `json:"purchasesOfInvestments"`
			SalesMaturitiesOfInvestments             float64 `json:"salesMaturitiesOfInvestments"`
			OtherInvestingActivites                  float64 `json:"otherInvestingActivites"`
			NetCashUsedForInvestingActivites         float64 `json:"netCashUsedForInvestingActivites"`
			DebtrPayment                             float64 `json:"debtRepayment"`
			CommonStockIssued                        float64 `json:"commonStockIssued"`
			CommonStockRepurchased                   float64 `json:"commonStockRepurchased"`
			DividendsPaid                            float64 `json:"dividendsPaid"`
			OtherFinancingActivites                  float64 `json:"otherFinancingActivites"`
			NetCashUsedProvidedByFinancingActivities float64 `json:"netCashUsedProvidedByFinancingActivities"`
			EffectOfForexChangesOnCash               float64 `json:"effectOfForexChangesOnCash"`
			Netchangeincash                          float64 `json:"netChangeInCash"`
			CashAtEndOfPeriod                        float64 `json:"cashAtEndOfPeriod"`
			CashAtBeginningOfPeriod                  float64 `json:"cashAtBeginningOfPeriod"`
			OperatingCashFlow                        float64 `json:"operatingCashFlow"`
			CapitalExpenditure                       float64 `json:"capitalExpenditure"`
			FreeCashFlow                             float64 `json:"freeCashFlow"`
		} `json:"cash"`
	} `json:"financialsQuarter"`
}

// EmployeeCount ...
type EmployeeCount struct {
	Symbol         string `json:"symbol"`
	Cik            string `json:"cik"`
	AcceptanceTime string `json:"acceptanceTime"`
	PeriodOfReport string `json:"periodOfReport"`
	CompanyName    string `json:"companyName"`
	FormType       string `json:"formType"`
	FilingDate     string `json:"filingDate"`
	Source         string `json:"source"`
	EmployeeCount  int    `json:"employeeCount"`
}

// IPOCalendarConfirmed ...
type IPOCalendarConfirmed struct {
	Symbol            string `json:"symbol"`
	Cik               string `json:"cik"`
	Form              string `json:"form"`
	FilingDate        string `json:"filingDate"`
	AcceptedDate      string `json:"acceptedDate"`
	EffectivenessDate string `json:"effectivenessDate"`
	URL               string `json:"url"`
}

// IPOCalendarProspectus ...
type IPOCalendarProspectus struct {
	Symbol                          string  `json:"symbol"`
	Cik                             string  `json:"cik"`
	Form                            string  `json:"form"`
	FilingDate                      string  `json:"filingDate"`
	AcceptedDate                    string  `json:"acceptedDate"`
	IpoDate                         string  `json:"ipoDate"`
	URL                             string  `json:"url"`
	PricePublicPerShare             float64 `json:"pricePublicPerShare"`
	PricePublicTotal                float64 `json:"pricePublicTotal"`
	DiscountsAndCommissionsPerShare float64 `json:"discountsAndCommissionsPerShare"`
	DiscountsAndCommissionsTotal    float64 `json:"discountsAndCommissionsTotal"`
	ProceedsBeforeExpensesPerShare  float64 `json:"proceedsBeforeExpensesPerShare"`
	ProceedsBeforeExpensesTotal     float64 `json:"proceedsBeforeExpensesTotal"`
}

// SocialSentiment ...
type SocialSentiment struct {
	Date                  string  `json:"date"`
	Symbol                string  `json:"symbol"`
	StocktwitsPosts       int     `json:"stocktwitsPosts"`
	TwitterPosts          int     `json:"twitterPosts"`
	StocktwitsComments    int     `json:"stocktwitsComments"`
	TwitterComments       int     `json:"twitterComments"`
	StocktwitsLikes       int     `json:"stocktwitsLikes"`
	TwitterLikes          int     `json:"twitterLikes"`
	StocktwitsImpressions int     `json:"stocktwitsImpressions"`
	TwitterImpressions    int     `json:"twitterImpressions"`
	StocktwitsSentiment   float64 `json:"stocktwitsSentiment"`
	TwitterSentiment      float64 `json:"twitterSentiment"`
}

// SocialSentimentChange ...
type SocialSentimentChange struct {
	Symbol          string  `json:"symbol"`
	Name            string  `json:"name"`
	Rank            int     `json:"rank"`
	Sentiment       float64 `json:"sentiment"`
	SentimentChange float64 `json:"sentimentChange"`
}

// Score ...
type Score struct {
	Symbol           string  `json:"symbol" csv:"symbol"`
	AltmanZScore     float64 `json:"altmanZScore" csv:"altmanZScore"`
	PiotroskiScore   float64 `json:"piotroskiScore" csv:"piotroskiScore"`
	WorkingCapital   string  `json:"workingCapital" csv:"workingCapital"`
	TotalAssets      string  `json:"totalAssets" csv:"totalAssets"`
	RetainedEarnings string  `json:"retainedEarnings" csv:"retainedEarnings"`
	Ebit             string  `json:"ebit" csv:"ebit"`
	MarketCap        string  `json:"marketCap" csv:"marketCap"`
	TotalLiabilities string  `json:"totalLiabilities" csv:"totalLiabilities"`
	Revenue          string  `json:"revenue" csv:"revenue"`
}

// SharesFloat ...
type SharesFloat struct {
	Symbol            string  `json:"symbol"`
	Date              string  `json:"date"`
	FreeFloat         float64 `json:"freeFloat"`
	FloatShares       int64   `json:"floatShares"`
	OutstandingShares int64   `json:"outstandingShares"`
	Source            string  `json:"source"`
}
