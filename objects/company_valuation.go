package objects

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
)

// ETFSector ...
type ETFSector string

// CompanyValuationPeriod ...
type CompanyValuationPeriod string

// CompanyValuationHistory ...
type CompanyValuationHistory string

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

// RequestStockNews ...
type RequestStockNews struct {
	SymbolList []string
	Limit      int64
}

// RequestStockScreener ...
type RequestStockScreener struct {
	Exchange           []string
	Sector             *string
	Industry           *string
	Limit              int64
	MarketCapMoreThan  *int64
	MarketCapLowerThan *int64
	VolumeMoreThan     *int64
	VolumeLowerThan    *int64
	BetaMoreThan       *float64
	BetaLowerThan      *float64
	DividendMoreThan   *float64
	DividendLowerThan  *float64
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
}

// SplitCalendar ...
type SplitCalendar struct {
	Date        string `json:"date"` // 2020-09-10
	Label       string `json:"label"`
	Symbol      string `json:"symbol"`
	Numerator   int    `json:"numerator"`
	Denominator int    `json:"denominator"`
}

// DividendCalendar ...
type DividendCalendar struct {
	Date        string  `json:"date"`  // 2020-09-10
	Label       string  `json:"label"` // September 10, 20
	AdjDividend float64 `json:"adjDividend"`
	Symbol      string  `json:"symbol"`
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
	Date                             string  `json:"date"`
	Symbol                           string  `json:"symbol"`
	FillingDate                      string  `json:"fillingDate"`
	AcceptedDate                     string  `json:"acceptedDate"`
	Period                           string  `json:"period"`
	Revenue                          int64   `json:"revenue"`
	CostOfRevenue                    int64   `json:"costOfRevenue"`
	GrossProfit                      int64   `json:"grossProfit"`
	GrossProfitRatio                 float64 `json:"grossProfitRatio"`
	ResearchAndDevelopmentExpenses   int64   `json:"researchAndDevelopmentExpenses"`
	GeneralAndAdministrativeExpenses int64   `json:"generalAndAdministrativeExpenses"`
	SellingAndMarketingExpenses      float64 `json:"sellingAndMarketingExpenses"`
	OtherExpenses                    float64 `json:"otherExpenses"`
	OperatingExpenses                float64 `json:"operatingExpenses"`
	CostAndExpenses                  float64 `json:"costAndExpenses"`
	InterestExpense                  int64   `json:"interestExpense"`
	DepreciationAndAmortization      int64   `json:"depreciationAndAmortization"`
	Ebitda                           int64   `json:"ebitda"`
	Ebitdaratio                      float64 `json:"ebitdaratio"`
	OperatingIncome                  int64   `json:"operatingIncome"`
	OperatingIncomeRatio             float64 `json:"operatingIncomeRatio"`
	TotalOtherIncomeExpensesNet      int64   `json:"totalOtherIncomeExpensesNet"`
	IncomeBeforeTax                  int64   `json:"incomeBeforeTax"`
	IncomeBeforeTaxRatio             float64 `json:"incomeBeforeTaxRatio"`
	IncomeTaxExpense                 int64   `json:"incomeTaxExpense"`
	NetIncome                        int64   `json:"netIncome"`
	NetIncomeRatio                   float64 `json:"netIncomeRatio"`
	Eps                              float64 `json:"eps"`
	Epsdiluted                       float64 `json:"epsdiluted"`
	WeightedAverageShsOut            int64   `json:"weightedAverageShsOut"`
	WeightedAverageShsOutDil         int64   `json:"weightedAverageShsOutDil"`
	Link                             string  `json:"link"`
	FinalLink                        string  `json:"finalLink"`
}

// IncomeStatementGrowth ...
type IncomeStatementGrowth struct {
	Date                                   string  `json:"date"`
	Symbol                                 string  `json:"symbol"`
	Period                                 string  `json:"period"`
	GrowthRevenue                          float64 `json:"growthRevenue"`
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
	Date                                    string  `json:"date"`
	Symbol                                  string  `json:"symbol"`
	FillingDate                             string  `json:"fillingDate"`
	AcceptedDate                            string  `json:"acceptedDate"`
	Period                                  string  `json:"period"`
	CashAndCashEquivalents                  int64   `json:"cashAndCashEquivalents"`
	ShortTermInvestments                    float64 `json:"shortTermInvestments"`
	CashAndShortTermInvestments             int64   `json:"cashAndShortTermInvestments"`
	NetReceivables                          int64   `json:"netReceivables"`
	Inventory                               int64   `json:"inventory"`
	OtherCurrentAssets                      int64   `json:"otherCurrentAssets"`
	TotalCurrentAssets                      int64   `json:"totalCurrentAssets"`
	PropertyPlantEquipmentNet               int64   `json:"propertyPlantEquipmentNet"`
	Goodwill                                int64   `json:"goodwill"`
	IntangibleAssets                        int64   `json:"intangibleAssets"`
	GoodwillAndIntangibleAssets             int64   `json:"goodwillAndIntangibleAssets"`
	LongTermInvestments                     float64 `json:"longTermInvestments"`
	TaxAssets                               int64   `json:"taxAssets"`
	OtherNonCurrentAssets                   int64   `json:"otherNonCurrentAssets"`
	TotalNonCurrentAssets                   int64   `json:"totalNonCurrentAssets"`
	OtherAssets                             int64   `json:"otherAssets"`
	TotalAssets                             int64   `json:"totalAssets"`
	AccountPayables                         int64   `json:"accountPayables"`
	ShortTermDebt                           int64   `json:"shortTermDebt"`
	TaxPayables                             int64   `json:"taxPayables"`
	DeferredRevenue                         int64   `json:"deferredRevenue"`
	OtherCurrentLiabilities                 int64   `json:"otherCurrentLiabilities"`
	TotalCurrentLiabilities                 int64   `json:"totalCurrentLiabilities"`
	LongTermDebt                            int64   `json:"longTermDebt"`
	DeferredRevenueNonCurrent               int64   `json:"deferredRevenueNonCurrent"`
	DeferredTaxLiabilitiesNonCurrent        int64   `json:"deferredTaxLiabilitiesNonCurrent"`
	OtherNonCurrentLiabilities              int64   `json:"otherNonCurrentLiabilities"`
	TotalNonCurrentLiabilities              int64   `json:"totalNonCurrentLiabilities"`
	OtherLiabilities                        int64   `json:"otherLiabilities"`
	TotalLiabilities                        int64   `json:"totalLiabilities"`
	CommonStock                             int64   `json:"commonStock"`
	RetainedEarnings                        int64   `json:"retainedEarnings"`
	AccumulatedOtherComprehensiveIncomeLoss int64   `json:"accumulatedOtherComprehensiveIncomeLoss"`
	OthertotalStockholdersEquity            int64   `json:"othertotalStockholdersEquity"`
	TotalStockholdersEquity                 int64   `json:"totalStockholdersEquity"`
	TotalLiabilitiesAndStockholdersEquity   int64   `json:"totalLiabilitiesAndStockholdersEquity"`
	TotalInvestments                        int64   `json:"totalInvestments"`
	TotalDebt                               int64   `json:"totalDebt"`
	NetDebt                                 int64   `json:"netDebt"`
	Link                                    string  `json:"link"`
	FinalLink                               string  `json:"finalLink"`
}

// BalanceSheetStatementGrowth ...
type BalanceSheetStatementGrowth struct {
	Date                                          string  `json:"date"`
	Symbol                                        string  `json:"symbol"`
	Period                                        string  `json:"period"`
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
	Date                                     string `json:"date"`
	Symbol                                   string `json:"symbol"`
	FillingDate                              string `json:"fillingDate"`
	AcceptedDate                             string `json:"acceptedDate"`
	Period                                   string `json:"period"`
	NetIncome                                int64  `json:"netIncome"`
	DepreciationAndAmortization              int64  `json:"depreciationAndAmortization"`
	DeferredIncomeTax                        int    `json:"deferredIncomeTax"`
	StockBasedCompensation                   int64  `json:"stockBasedCompensation"`
	ChangeInWorkingCapital                   int64  `json:"changeInWorkingCapital"`
	AccountsReceivables                      int    `json:"accountsReceivables"`
	Inventory                                int    `json:"inventory"`
	AccountsPayables                         int    `json:"accountsPayables"`
	OtherWorkingCapital                      int64  `json:"otherWorkingCapital"`
	OtherNonCashItems                        int64  `json:"otherNonCashItems"`
	NetCashProvidedByOperatingActivities     int64  `json:"netCashProvidedByOperatingActivities"`
	InvestmentsInPropertyPlantAndEquipment   int64  `json:"investmentsInPropertyPlantAndEquipment"`
	AcquisitionsNet                          int    `json:"acquisitionsNet"`
	PurchasesOfInvestments                   int64  `json:"purchasesOfInvestments"`
	SalesMaturitiesOfInvestments             int64  `json:"salesMaturitiesOfInvestments"`
	OtherInvestingActivites                  int64  `json:"otherInvestingActivites"`
	NetCashUsedForInvestingActivites         int64  `json:"netCashUsedForInvestingActivites"`
	DebtRepayment                            int64  `json:"debtRepayment"`
	CommonStockIssued                        int    `json:"commonStockIssued"`
	CommonStockRepurchased                   int64  `json:"commonStockRepurchased"`
	DividendsPaid                            int64  `json:"dividendsPaid"`
	OtherFinancingActivites                  int    `json:"otherFinancingActivites"`
	NetCashUsedProvidedByFinancingActivities int64  `json:"netCashUsedProvidedByFinancingActivities"`
	EffectOfForexChangesOnCash               int    `json:"effectOfForexChangesOnCash"`
	NetChangeInCash                          int64  `json:"netChangeInCash"`
	CashAtEndOfPeriod                        int64  `json:"cashAtEndOfPeriod"`
	CashAtBeginningOfPeriod                  int64  `json:"cashAtBeginningOfPeriod"`
	OperatingCashFlow                        int64  `json:"operatingCashFlow"`
	CapitalExpenditure                       int64  `json:"capitalExpenditure"`
	FreeCashFlow                             int64  `json:"freeCashFlow"`
	Link                                     string `json:"link"`
	FinalLink                                string `json:"finalLink"`
}

// CashFlowStatementGrowth ...
type CashFlowStatementGrowth struct {
	Date                                           string  `json:"date"`
	Symbol                                         string  `json:"symbol"`
	Period                                         string  `json:"period"`
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
	Symbol                             string  `json:"symbol"`
	Date                               string  `json:"date"`
	CurrentRatio                       float64 `json:"currentRatio"`
	QuickRatio                         float64 `json:"quickRatio"`
	CashRatio                          float64 `json:"cashRatio"`
	DaysOfSalesOutstanding             float64 `json:"daysOfSalesOutstanding"`
	DaysOfInventoryOutstanding         float64 `json:"daysOfInventoryOutstanding"`
	OperatingCycle                     float64 `json:"operatingCycle"`
	DaysOfPayablesOutstanding          float64 `json:"daysOfPayablesOutstanding"`
	CashConversionCycle                float64 `json:"cashConversionCycle"`
	GrossProfitMargin                  float64 `json:"grossProfitMargin"`
	OperatingProfitMargin              float64 `json:"operatingProfitMargin"`
	PretaxProfitMargin                 float64 `json:"pretaxProfitMargin"`
	NetProfitMargin                    float64 `json:"netProfitMargin"`
	EffectiveTaxRate                   float64 `json:"effectiveTaxRate"`
	ReturnOnAssets                     float64 `json:"returnOnAssets"`
	ReturnOnEquity                     float64 `json:"returnOnEquity"`
	ReturnOnCapitalEmployed            float64 `json:"returnOnCapitalEmployed"`
	NetIncomePerEBT                    float64 `json:"netIncomePerEBT"`
	EbtPerEbit                         int64   `json:"ebtPerEbit"`
	EbitPerRevenue                     float64 `json:"ebitPerRevenue"`
	DebtRatio                          float64 `json:"debtRatio"`
	DebtEquityRatio                    float64 `json:"debtEquityRatio"`
	LongTermDebtToCapitalization       float64 `json:"longTermDebtToCapitalization"`
	TotalDebtToCapitalization          float64 `json:"totalDebtToCapitalization"`
	InterestCoverage                   float64 `json:"interestCoverage"`
	CashFlowToDebtRatio                float64 `json:"cashFlowToDebtRatio"`
	CompanyEquityMultiplier            float64 `json:"companyEquityMultiplier"`
	ReceivablesTurnover                float64 `json:"receivablesTurnover"`
	PayablesTurnover                   float64 `json:"payablesTurnover"`
	InventoryTurnover                  float64 `json:"inventoryTurnover"`
	FixedAssetTurnover                 float64 `json:"fixedAssetTurnover"`
	AssetTurnover                      float64 `json:"assetTurnover"`
	OperatingCashFlowPerShare          float64 `json:"operatingCashFlowPerShare"`
	FreeCashFlowPerShare               float64 `json:"freeCashFlowPerShare"`
	CashPerShare                       float64 `json:"cashPerShare"`
	PayoutRatio                        float64 `json:"payoutRatio"`
	OperatingCashFlowSalesRatio        float64 `json:"operatingCashFlowSalesRatio"`
	FreeCashFlowOperatingCashFlowRatio float64 `json:"freeCashFlowOperatingCashFlowRatio"`
	CashFlowCoverageRatios             float64 `json:"cashFlowCoverageRatios"`
	ShortTermCoverageRatios            float64 `json:"shortTermCoverageRatios"`
	CapitalExpenditureCoverageRatio    float64 `json:"capitalExpenditureCoverageRatio"`
	DividendPaidAndCapexCoverageRatio  float64 `json:"dividendPaidAndCapexCoverageRatio"`
	DividendPayoutRatio                float64 `json:"dividendPayoutRatio"`
	PriceBookValueRatio                float64 `json:"priceBookValueRatio"`
	PriceToBookRatio                   float64 `json:"priceToBookRatio"`
	PriceToSalesRatio                  float64 `json:"priceToSalesRatio"`
	PriceEarningsRatio                 float64 `json:"priceEarningsRatio"`
	PriceToFreeCashFlowsRatio          float64 `json:"priceToFreeCashFlowsRatio"`
	PriceToOperatingCashFlowsRatio     float64 `json:"priceToOperatingCashFlowsRatio"`
	PriceCashFlowRatio                 float64 `json:"priceCashFlowRatio"`
	PriceEarningsToGrowthRatio         float64 `json:"priceEarningsToGrowthRatio"`
	PriceSalesRatio                    float64 `json:"priceSalesRatio"`
	DividendYield                      float64 `json:"dividendYield"`
	EnterpriseValueMultiple            float64 `json:"enterpriseValueMultiple"`
	PriceFairValue                     float64 `json:"priceFairValue"`
}

// FinancialRatiosTTM ...
type FinancialRatiosTTM struct {
	DividendYielTTM float64 `json:"dividendYielTTM"`
	PeRatioTTM      float64 `json:"peRatioTTM"`
	PegRatioTTM     float64 `json:"pegRatioTTM"`
	PayoutRatioTTM  float64 `json:"payoutRatioTTM"`
}

// KeyMetrics ...
type KeyMetrics struct {
	Symbol                                 string   `json:"symbol"`
	Date                                   string   `json:"date"`
	RevenuePerShare                        float64  `json:"revenuePerShare"`
	NetIncomePerShare                      float64  `json:"netIncomePerShare"`
	OperatingCashFlowPerShare              float64  `json:"operatingCashFlowPerShare"`
	FreeCashFlowPerShare                   float64  `json:"freeCashFlowPerShare"`
	CashPerShare                           float64  `json:"cashPerShare"`
	BookValuePerShare                      float64  `json:"bookValuePerShare"`
	TangibleBookValuePerShare              float64  `json:"tangibleBookValuePerShare"`
	ShareholdersEquityPerShare             float64  `json:"shareholdersEquityPerShare"`
	InterestDebtPerShare                   float64  `json:"interestDebtPerShare"`
	MarketCap                              float64  `json:"marketCap"`
	EnterpriseValue                        float64  `json:"enterpriseValue"`
	PeRatio                                float64  `json:"peRatio"`
	PriceToSalesRatio                      float64  `json:"priceToSalesRatio"`
	Pocfratio                              float64  `json:"pocfratio"`
	PfcfRatio                              float64  `json:"pfcfRatio"`
	PbRatio                                float64  `json:"pbRatio"`
	PtbRatio                               float64  `json:"ptbRatio"`
	EvToSales                              float64  `json:"evToSales"`
	EnterpriseValueOverEBITDA              float64  `json:"enterpriseValueOverEBITDA"`
	EvToOperatingCashFlow                  float64  `json:"evToOperatingCashFlow"`
	EvToFreeCashFlow                       float64  `json:"evToFreeCashFlow"`
	EarningsYield                          float64  `json:"earningsYield"`
	FreeCashFlowYield                      float64  `json:"freeCashFlowYield"`
	DebtToEquity                           float64  `json:"debtToEquity"`
	DebtToAssets                           float64  `json:"debtToAssets"`
	NetDebtToEBITDA                        float64  `json:"netDebtToEBITDA"`
	CurrentRatio                           float64  `json:"currentRatio"`
	InterestCoverage                       float64  `json:"interestCoverage"`
	IncomeQuality                          float64  `json:"incomeQuality"`
	DividendYield                          *float64 `json:"dividendYield"`
	PayoutRatio                            float64  `json:"payoutRatio"`
	SalesGeneralAndAdministrativeToRevenue float64  `json:"salesGeneralAndAdministrativeToRevenue"`
	ResearchAndDdevelopementToRevenue      float64  `json:"researchAndDdevelopementToRevenue"`
	IntangiblesToTotalAssets               float64  `json:"intangiblesToTotalAssets"`
	CapexToOperatingCashFlow               float64  `json:"capexToOperatingCashFlow"`
	CapexToRevenue                         float64  `json:"capexToRevenue"`
	CapexToDepreciation                    float64  `json:"capexToDepreciation"`
	StockBasedCompensationToRevenue        float64  `json:"stockBasedCompensationToRevenue"`
	GrahamNumber                           float64  `json:"grahamNumber"`
	Roic                                   float64  `json:"roic"`
	ReturnOnTangibleAssets                 float64  `json:"returnOnTangibleAssets"`
	GrahamNetNet                           float64  `json:"grahamNetNet"`
	WorkingCapital                         float64  `json:"workingCapital"`
	TangibleAssetValue                     float64  `json:"tangibleAssetValue"`
	NetCurrentAssetValue                   float64  `json:"netCurrentAssetValue"`
	InvestedCapital                        *float64 `json:"investedCapital"`
	AverageReceivables                     float64  `json:"averageReceivables"`
	AveragePayables                        float64  `json:"averagePayables"`
	AverageInventory                       float64  `json:"averageInventory"`
	DaysSalesOutstanding                   float64  `json:"daysSalesOutstanding"`
	DaysPayablesOutstanding                float64  `json:"daysPayablesOutstanding"`
	DaysOfInventoryOnHand                  float64  `json:"daysOfInventoryOnHand"`
	ReceivablesTurnover                    float64  `json:"receivablesTurnover"`
	PayablesTurnover                       float64  `json:"payablesTurnover"`
	InventoryTurnover                      float64  `json:"inventoryTurnover"`
	Roe                                    float64  `json:"roe"`
	CapexPerShare                          float64  `json:"capexPerShare"`
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
	Symbol string  `json:"symbol"`
	Date   string  `json:"date"`
	Dcf    float64 `json:"dcf"`
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
	Symbol                         string `json:"symbol"`
	Date                           string `json:"date"` // 2020-09-14
	Rating                         string `json:"rating"`
	RatingScore                    int64  `json:"ratingScore"`
	RatingRecommendation           string `json:"ratingRecommendation"`
	RatingDetailsDCFScore          int64  `json:"ratingDetailsDCFScore"`
	RatingDetailsDCFRecommendation string `json:"ratingDetailsDCFRecommendation"`
	RatingDetailsROEScore          int64  `json:"ratingDetailsROEScore"`
	RatingDetailsROERecommendation string `json:"ratingDetailsROERecommendation"`
	RatingDetailsROAScore          int64  `json:"ratingDetailsROAScore"`
	RatingDetailsROARecommendation string `json:"ratingDetailsROARecommendation"`
	RatingDetailsDEScore           int64  `json:"ratingDetailsDEScore"`
	RatingDetailsDERecommendation  string `json:"ratingDetailsDERecommendation"`
	RatingDetailsPEScore           int64  `json:"ratingDetailsPEScore"`
	RatingDetailsPERecommendation  string `json:"ratingDetailsPERecommendation"`
	RatingDetailsPBScore           int64  `json:"ratingDetailsPBScore"`
	RatingDetailsPBRecommendation  string `json:"ratingDetailsPBRecommendation"`
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
	Volume             int64   `json:"volume"`
	Exchange           string  `json:"exchange"`
	ExchangeShortName  string  `json:"exchangeShortName"`
}
