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
	Time             string  `json:"time"` // timezone
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
	SellingAndMarketingExpenses      int64   `json:"sellingAndMarketingExpenses"`
	OtherExpenses                    int64   `json:"otherExpenses"`
	OperatingExpenses                int64   `json:"operatingExpenses"`
	CostAndExpenses                  int64   `json:"costAndExpenses"`
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
	Date                                    string `json:"date"`
	Symbol                                  string `json:"symbol"`
	FillingDate                             string `json:"fillingDate"`
	AcceptedDate                            string `json:"acceptedDate"`
	Period                                  string `json:"period"`
	CashAndCashEquivalents                  int64  `json:"cashAndCashEquivalents"`
	ShortTermInvestments                    int64  `json:"shortTermInvestments"`
	CashAndShortTermInvestments             int64  `json:"cashAndShortTermInvestments"`
	NetReceivables                          int64  `json:"netReceivables"`
	Inventory                               int64  `json:"inventory"`
	OtherCurrentAssets                      int64  `json:"otherCurrentAssets"`
	TotalCurrentAssets                      int64  `json:"totalCurrentAssets"`
	PropertyPlantEquipmentNet               int64  `json:"propertyPlantEquipmentNet"`
	Goodwill                                int64  `json:"goodwill"`
	IntangibleAssets                        int64  `json:"intangibleAssets"`
	GoodwillAndIntangibleAssets             int64  `json:"goodwillAndIntangibleAssets"`
	LongTermInvestments                     int64  `json:"longTermInvestments"`
	TaxAssets                               int64  `json:"taxAssets"`
	OtherNonCurrentAssets                   int64  `json:"otherNonCurrentAssets"`
	TotalNonCurrentAssets                   int64  `json:"totalNonCurrentAssets"`
	OtherAssets                             int64  `json:"otherAssets"`
	TotalAssets                             int64  `json:"totalAssets"`
	AccountPayables                         int64  `json:"accountPayables"`
	ShortTermDebt                           int64  `json:"shortTermDebt"`
	TaxPayables                             int64  `json:"taxPayables"`
	DeferredRevenue                         int64  `json:"deferredRevenue"`
	OtherCurrentLiabilities                 int64  `json:"otherCurrentLiabilities"`
	TotalCurrentLiabilities                 int64  `json:"totalCurrentLiabilities"`
	LongTermDebt                            int64  `json:"longTermDebt"`
	DeferredRevenueNonCurrent               int64  `json:"deferredRevenueNonCurrent"`
	DeferredTaxLiabilitiesNonCurrent        int64  `json:"deferredTaxLiabilitiesNonCurrent"`
	OtherNonCurrentLiabilities              int64  `json:"otherNonCurrentLiabilities"`
	TotalNonCurrentLiabilities              int64  `json:"totalNonCurrentLiabilities"`
	OtherLiabilities                        int64  `json:"otherLiabilities"`
	TotalLiabilities                        int64  `json:"totalLiabilities"`
	CommonStock                             int64  `json:"commonStock"`
	RetainedEarnings                        int64  `json:"retainedEarnings"`
	AccumulatedOtherComprehensiveIncomeLoss int64  `json:"accumulatedOtherComprehensiveIncomeLoss"`
	OthertotalStockholdersEquity            int64  `json:"othertotalStockholdersEquity"`
	TotalStockholdersEquity                 int64  `json:"totalStockholdersEquity"`
	TotalLiabilitiesAndStockholdersEquity   int64  `json:"totalLiabilitiesAndStockholdersEquity"`
	TotalInvestments                        int64  `json:"totalInvestments"`
	TotalDebt                               int64  `json:"totalDebt"`
	NetDebt                                 int64  `json:"netDebt"`
	Link                                    string `json:"link"`
	FinalLink                               string `json:"finalLink"`
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
	Operatingexpenses                                                                           float64     `json:"operatingexpenses"`
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
	Cashandcashequivalentsatcarryingvalue           float64     `json:"cashandcashequivalentsatcarryingvalue"`
	Retainedearningsaccumulateddeficit              float64     `json:"retainedearningsaccumulateddeficit"`
	Liabilitiesnoncurrent                           interface{} `json:"liabilitiesnoncurrent"`
	Propertyplantandequipmentnet                    float64     `json:"propertyplantandequipmentnet"`
	Commonstocksincludingadditionalpaidincapital    float64     `json:"commonstocksincludingadditionalpaidincapital"`
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
	Date                                                                                                           string  `json:"date"`
	Symbol                                                                                                         string  `json:"symbol"`
	Period                                                                                                         string  `json:"period"`
	Paymentsforrepurchaseofcommonstock                                                                             float64 `json:"paymentsforrepurchaseofcommonstock"`
	Sharebasedcompensation                                                                                         float64 `json:"sharebasedcompensation"`
	Netincomeloss                                                                                                  float64 `json:"netincomeloss"`
	Increasedecreaseinaccountspayable                                                                              float64 `json:"increasedecreaseinaccountspayable"`
	Proceedsfrompaymentsforotherfinancingactivities                                                                float64 `json:"proceedsfrompaymentsforotherfinancingactivities"`
	Paymentsrelatedtotaxwithholdingforsharebasedcompensation                                                       float64 `json:"paymentsrelatedtotaxwithholdingforsharebasedcompensation"`
	Increasedecreaseinotheroperatingliabilities                                                                    float64 `json:"increasedecreaseinotheroperatingliabilities"`
	Othernoncashincomeexpense                                                                                      float64 `json:"othernoncashincomeexpense"`
	Paymentstoacquirebusinessesnetofcashacquired                                                                   float64 `json:"paymentstoacquirebusinessesnetofcashacquired"`
	Deferredincometaxexpensebenefit                                                                                float64 `json:"deferredincometaxexpensebenefit"`
	Cashcashequivalentsrestrictedcashandrestrictedcashequivalents                                                  float64 `json:"cashcashequivalentsrestrictedcashandrestrictedcashequivalents"`
	Cashcashequivalentsrestrictedcashandrestrictedcashequivalentsperiodincreasedecreaseincludingexchangerateeffect float64 `json:"cashcashequivalentsrestrictedcashandrestrictedcashequivalentsperiodincreasedecreaseincludingexchangerateeffect"`
	Netcashprovidedbyusedinoperatingactivities                                                                     float64 `json:"netcashprovidedbyusedinoperatingactivities"`
	Proceedsfromsaleofavailableforsalesecuritiesdebt                                                               float64 `json:"proceedsfromsaleofavailableforsalesecuritiesdebt"`
	Repaymentsoflongtermdebt                                                                                       float64 `json:"repaymentsoflongtermdebt"`
	Incometaxespaidnet                                                                                             float64 `json:"incometaxespaidnet"`
	Proceedsfromissuanceoflongtermdebt                                                                             float64 `json:"proceedsfromissuanceoflongtermdebt"`
	Paymentstoacquireotherinvestments                                                                              float64 `json:"paymentstoacquireotherinvestments"`
	Netcashprovidedbyusedininvestingactivities                                                                     float64 `json:"netcashprovidedbyusedininvestingactivities"`
	Increasedecreaseincontractwithcustomerliability                                                                float64 `json:"increasedecreaseincontractwithcustomerliability"`
	Interestpaidnet                                                                                                float64 `json:"interestpaidnet"`
	Netcashprovidedbyusedinfinancingactivities                                                                     float64 `json:"netcashprovidedbyusedinfinancingactivities"`
	Proceedsfromrepaymentsofcommercialpaper                                                                        float64 `json:"proceedsfromrepaymentsofcommercialpaper"`
	Proceedsfromsaleandmaturityofotherinvestments                                                                  float64 `json:"proceedsfromsaleandmaturityofotherinvestments"`
	Paymentstoacquireavailableforsalesecuritiesdebt                                                                float64 `json:"paymentstoacquireavailableforsalesecuritiesdebt"`
	Paymentstoacquirepropertyplantandequipment                                                                     float64 `json:"paymentstoacquirepropertyplantandequipment"`
	Paymentsforproceedsfromotherinvestingactivities                                                                float64 `json:"paymentsforproceedsfromotherinvestingactivities"`
	Increasedecreaseinotherreceivables                                                                             float64 `json:"increasedecreaseinotherreceivables"`
	Paymentsofdividends                                                                                            float64 `json:"paymentsofdividends"`
	Increasedecreaseininventories                                                                                  float64 `json:"increasedecreaseininventories"`
	Increasedecreaseinaccountsreceivable                                                                           float64 `json:"increasedecreaseinaccountsreceivable"`
	Proceedsfromissuanceofcommonstock                                                                              float64 `json:"proceedsfromissuanceofcommonstock"`
	Depreciationdepletionandamortization                                                                           float64 `json:"depreciationdepletionandamortization"`
	Proceedsfrommaturitiesprepaymentsandcallsofavailableforsalesecurities                                          float64 `json:"proceedsfrommaturitiesprepaymentsandcallsofavailableforsalesecurities"`
	Increasedecreaseinotheroperatingassets                                                                         float64 `json:"increasedecreaseinotheroperatingassets"`
	Proceedsfromothershorttermdebt                                                                                 float64 `json:"proceedsfromothershorttermdebt"`
}

// FullFinancialStatementAsReported ...
type FullFinancialStatementAsReported struct {
	Date                                                                                                                                        string  `json:"date"`
	Symbol                                                                                                                                      string  `json:"symbol"`
	Period                                                                                                                                      string  `json:"period"`
	Numberofsignificantvendors                                                                                                                  int64   `json:"numberofsignificantvendors"`
	Entitycurrentreportingstatus                                                                                                                string  `json:"entitycurrentreportingstatus"`
	Machineryequipmentandinternalusesoftwaremember                                                                                              int64   `json:"machineryequipmentandinternalusesoftwaremember"`
	Otherliabilitiesmember                                                                                                                      int64   `json:"otherliabilitiesmember"`
	Entityemerginggrowthcompany                                                                                                                 string  `json:"entityemerginggrowthcompany"`
	Incometaxreconciliationtaxcreditsresearch                                                                                                   int64   `json:"incometaxreconciliationtaxcreditsresearch"`
	Sharebasedcompensation                                                                                                                      int64   `json:"sharebasedcompensation"`
	Unrecordedunconditionalpurchaseobligationbalanceonsecondanniversary                                                                         int64   `json:"unrecordedunconditionalpurchaseobligationbalanceonsecondanniversary"`
	Nonoperatingincomeexpense                                                                                                                   int64   `json:"nonoperatingincomeexpense"`
	Incometaxespaidnet                                                                                                                          int64   `json:"incometaxespaidnet"`
	Federalincometaxexpensebenefitcontinuingoperations                                                                                          int64   `json:"federalincometaxexpensebenefitcontinuingoperations"`
	Japansegmentmember                                                                                                                          int64   `json:"japansegmentmember"`
	Currentforeigntaxexpensebenefit                                                                                                             int64   `json:"currentforeigntaxexpensebenefit"`
	Debtsecuritiesavailableforsalecontinuousunrealizedlossposition12Monthsorlonger                                                              int64   `json:"debtsecuritiesavailableforsalecontinuousunrealizedlossposition12monthsorlonger"`
	Proceedsfromsaleofavailableforsalesecuritiesdebt                                                                                            int64   `json:"proceedsfromsaleofavailableforsalesecuritiesdebt"`
	Commonstocksharesoutstanding                                                                                                                int64   `json:"commonstocksharesoutstanding"`
	Unrecognizedtaxbenefitsthatwouldimpacteffectivetaxrate                                                                                      int64   `json:"unrecognizedtaxbenefitsthatwouldimpacteffectivetaxrate"`
	Decreaseinunrecognizedtaxbenefitsisreasonablypossible                                                                                       int64   `json:"decreaseinunrecognizedtaxbenefitsisreasonablypossible"`
	Deferredtaxassetsnet                                                                                                                        int64   `json:"deferredtaxassetsnet"`
	Deferredtaxassetsgoodwillandintangibleassets                                                                                                int64   `json:"deferredtaxassetsgoodwillandintangibleassets"`
	Interestratecontractmember                                                                                                                  int64   `json:"interestratecontractmember"`
	Othercomprehensiveincomelossavailableforsalesecuritiesadjustmentnetoftax                                                                    int64   `json:"othercomprehensiveincomelossavailableforsalesecuritiesadjustmentnetoftax"`
	Entitytaxidentificationnumber                                                                                                               string  `json:"entitytaxidentificationnumber"`
	Incometaxreconciliationforeignincometaxratedifferential                                                                                     int64   `json:"incometaxreconciliationforeignincometaxratedifferential"`
	Paymentsrelatedtotaxwithholdingforsharebasedcompensation                                                                                    int64   `json:"paymentsrelatedtotaxwithholdingforsharebasedcompensation"`
	Increasedecreaseinotheroperatingliabilities                                                                                                 int64   `json:"increasedecreaseinotheroperatingliabilities"`
	Deferredtaxassetsliabilitiesnet                                                                                                             int64   `json:"deferredtaxassetsliabilitiesnet"`
	Debtinstrumentmaturityyearrangestart                                                                                                        int64   `json:"debtinstrumentmaturityyearrangestart"`
	Propertyplantandequipmentnet                                                                                                                int64   `json:"propertyplantandequipmentnet"`
	Commonstocksincludingadditionalpaidincapital                                                                                                int64   `json:"commonstocksincludingadditionalpaidincapital"`
	Weightedaveragenumberofdilutedsharesoutstanding                                                                                             int64   `json:"weightedaveragenumberofdilutedsharesoutstanding"`
	Commercialpaper                                                                                                                             int64   `json:"commercialpaper"`
	Commonstockmember                                                                                                                           string  `json:"commonstockmember"`
	Employeeservicesharebasedcompensationnonvestedawardstotalcompensationcostnotyetrecognized                                                   int64   `json:"employeeservicesharebasedcompensationnonvestedawardstotalcompensationcostnotyetrecognized"`
	Stockrepurchasedandretiredduringperiodshares                                                                                                int64   `json:"stockrepurchasedandretiredduringperiodshares"`
	Greaterchinasegmentmember                                                                                                                   int64   `json:"greaterchinasegmentmember"`
	Costofgoodsandservicessold                                                                                                                  string  `json:"costofgoodsandservicessold"`
	Longtermdebtcurrent                                                                                                                         int64   `json:"longtermdebtcurrent"`
	Accumulatedtranslationadjustmentmember                                                                                                      int64   `json:"accumulatedtranslationadjustmentmember"`
	Cashcashequivalentsrestrictedcashandrestrictedcashequivalentsperiodincreasedecreaseincludingexchangerateeffect                              int64   `json:"cashcashequivalentsrestrictedcashandrestrictedcashequivalentsperiodincreasedecreaseincludingexchangerateeffect"`
	Debtinstrumentunamortizeddiscountpremiumanddebtissuancecostsnet                                                                             int64   `json:"debtinstrumentunamortizeddiscountpremiumanddebtissuancecostsnet"`
	Repaymentsoflongtermdebt                                                                                                                    int64   `json:"repaymentsoflongtermdebt"`
	Earningspersharediluted                                                                                                                     float64 `json:"earningspersharediluted"`
	Netincomeloss                                                                                                                               int64   `json:"netincomeloss"`
	Proceedsfromissuanceoflongtermdebt                                                                                                          int64   `json:"proceedsfromissuanceoflongtermdebt"`
	Unrecordedunconditionalpurchaseobligationbalanceonfifthanniversary                                                                          int64   `json:"unrecordedunconditionalpurchaseobligationbalanceonfifthanniversary"`
	Incomelossfromcontinuingoperationsbeforeincometaxesextraordinaryitemsnoncontrollinginterest                                                 int64   `json:"incomelossfromcontinuingoperationsbeforeincometaxesextraordinaryitemsnoncontrollinginterest"`
	Incomelossfromcontinuingoperationsbeforeincometaxesforeign                                                                                  int64   `json:"incomelossfromcontinuingoperationsbeforeincometaxesforeign"`
	Crosscurrencyinterestratecontractmember                                                                                                     string  `json:"crosscurrencyinterestratecontractmember"`
	Localphonenumber                                                                                                                            string  `json:"localphonenumber"`
	Collateralalreadyreceivedaggregatefairvalue                                                                                                 int64   `json:"collateralalreadyreceivedaggregatefairvalue"`
	Interestpaidnet                                                                                                                             int64   `json:"interestpaidnet"`
	Gainlossfromcomponentsexcludedfromassessmentoffairvaluehedgeeffectivenessnet                                                                int64   `json:"gainlossfromcomponentsexcludedfromassessmentoffairvaluehedgeeffectivenessnet"`
	Accumulateddepreciationdepletionandamortizationpropertyplantandequipment                                                                    int64   `json:"accumulateddepreciationdepletionandamortizationpropertyplantandequipment"`
	Othercomprehensiveincomeunrealizedgainlossonderivativesarisingduringperiodbeforetax                                                         int64   `json:"othercomprehensiveincomeunrealizedgainlossonderivativesarisingduringperiodbeforetax"`
	Proceedsfromrepaymentsofcommercialpaper                                                                                                     int64   `json:"proceedsfromrepaymentsofcommercialpaper"`
	Proceedsfromsaleandmaturityofotherinvestments                                                                                               int64   `json:"proceedsfromsaleandmaturityofotherinvestments"`
	Longtermdebtnoncurrent                                                                                                                      int64   `json:"longtermdebtnoncurrent"`
	Maximumlengthoftimeforeigncurrencycashflowhedge                                                                                             string  `json:"maximumlengthoftimeforeigncurrencycashflowhedge"`
	Stockissuedduringperiodvaluenewissues                                                                                                       int64   `json:"stockissuedduringperiodvaluenewissues"`
	Netcashprovidedbyusedinoperatingactivities                                                                                                  int64   `json:"netcashprovidedbyusedinoperatingactivities"`
	Operatingleasesfutureminimumpaymentsdue                                                                                                     int64   `json:"operatingleasesfutureminimumpaymentsdue"`
	Proceedsfromrepaymentsofshorttermdebtmaturinginthreemonthsorless                                                                            int64   `json:"proceedsfromrepaymentsofshorttermdebtmaturinginthreemonthsorless"`
	Stockrepurchasedandretiredduringperiodvalue                                                                                                 int64   `json:"stockrepurchasedandretiredduringperiodvalue"`
	Proceedsfromrepaymentsofshorttermdebtmaturinginmorethanthreemonths                                                                          int64   `json:"proceedsfromrepaymentsofshorttermdebtmaturinginmorethanthreemonths"`
	Restrictedinvestments                                                                                                                       int64   `json:"restrictedinvestments"`
	Leaseholdimprovementsmember                                                                                                                 int64   `json:"leaseholdimprovementsmember"`
	Standardproductwarrantyaccrualpayments                                                                                                      int64   `json:"standardproductwarrantyaccrualpayments"`
	Accountsreceivablenetcurrent                                                                                                                int64   `json:"accountsreceivablenetcurrent"`
	Performanceobligationsinarrangements                                                                                                        int64   `json:"performanceobligationsinarrangements"`
	Sharespaidfortaxwithholdingforsharebasedcompensation                                                                                        int64   `json:"sharespaidfortaxwithholdingforsharebasedcompensation"`
	Currentstateandlocaltaxexpensebenefit                                                                                                       int64   `json:"currentstateandlocaltaxexpensebenefit"`
	Repaymentsofshorttermdebtmaturinginmorethanthreemonths                                                                                      int64   `json:"repaymentsofshorttermdebtmaturinginmorethanthreemonths"`
	Derivativefairvalueofderivativenet                                                                                                          int64   `json:"derivativefairvalueofderivativenet"`
	Upfrontpaymentunderacceleratedsharerepurchasearrangement                                                                                    int64   `json:"upfrontpaymentunderacceleratedsharerepurchasearrangement"`
	Americassegmentmember                                                                                                                       int64   `json:"americassegmentmember"`
	Deferredfederalincometaxexpensebenefit                                                                                                      int64   `json:"deferredfederalincometaxexpensebenefit"`
	Operatingleasesfutureminimumpaymentsduethereafter                                                                                           int64   `json:"operatingleasesfutureminimumpaymentsduethereafter"`
	Longtermmarketablesecuritiesmaturitiesterm                                                                                                  string  `json:"longtermmarketablesecuritiesmaturitiesterm"`
	Cashandcashequivalentsatcarryingvalue                                                                                                       int64   `json:"cashandcashequivalentsatcarryingvalue"`
	Documentannualreport                                                                                                                        string  `json:"documentannualreport"`
	Otherassetscurrent                                                                                                                          int64   `json:"otherassetscurrent"`
	Unfavorableinvestigationoutcomeeustateaidrulesmember                                                                                        int64   `json:"unfavorableinvestigationoutcomeeustateaidrulesmember"`
	Othernonoperatingincomeexpense                                                                                                              int64   `json:"othernonoperatingincomeexpense"`
	Maximumlengthoftimehedgedininterestratecashflowhedge1                                                                                       string  `json:"maximumlengthoftimehedgedininterestratecashflowhedge1"`
	Deferredtaxassetsunrealizedlossesonavailableforsalesecuritiesgross                                                                          int64   `json:"deferredtaxassetsunrealizedlossesonavailableforsalesecuritiesgross"`
	Definedcontributionplanmaximumannualcontributionsperemployeeamount                                                                          int64   `json:"definedcontributionplanmaximumannualcontributionsperemployeeamount"`
	Debtsecuritiesavailableforsaleunrealizedlossposition                                                                                        int64   `json:"debtsecuritiesavailableforsaleunrealizedlossposition"`
	Acceleratedsharerepurchasearrangementfebruary2019Member                                                                                     int64   `json:"acceleratedsharerepurchasearrangementfebruary2019member"`
	Increasedecreaseinotheroperatingassets                                                                                                      int64   `json:"increasedecreaseinotheroperatingassets"`
	Accumulatednetunrealizedinvestmentgainlossmember                                                                                            int64   `json:"accumulatednetunrealizedinvestmentgainlossmember"`
	Restofasiapacificsegmentmember                                                                                                              int64   `json:"restofasiapacificsegmentmember"`
	Unrecordedunconditionalpurchaseobligationbalanceonfourthanniversary                                                                         int64   `json:"unrecordedunconditionalpurchaseobligationbalanceonfourthanniversary"`
	Contractwithcustomerliability                                                                                                               int64   `json:"contractwithcustomerliability"`
	Tradeaccountsreceivablemember                                                                                                               int64   `json:"tradeaccountsreceivablemember"`
	Securityexchangename                                                                                                                        string  `json:"securityexchangename"`
	Proceedsfromissuanceofcommonstock                                                                                                           int64   `json:"proceedsfromissuanceofcommonstock"`
	Interestcostsincurred                                                                                                                       int64   `json:"interestcostsincurred"`
	Adjustmentstoadditionalpaidincapitaltaxeffectfromsharebasedcompensationincludingtransferpricing                                             int64   `json:"adjustmentstoadditionalpaidincapitaltaxeffectfromsharebasedcompensationincludingtransferpricing"`
	Revenueremainingperformanceobligationpercentage                                                                                             float64 `json:"revenueremainingperformanceobligationpercentage"`
	Liabilitiescurrent                                                                                                                          string  `json:"liabilitiescurrent"`
	Sharebasedcompensationarrangementbysharebasedpaymentawardofferingterm                                                                       string  `json:"sharebasedcompensationarrangementbysharebasedpaymentawardofferingterm"`
	Increasedecreaseinaccountspayable                                                                                                           int64   `json:"increasedecreaseinaccountspayable"`
	Grossprofit                                                                                                                                 int64   `json:"grossprofit"`
	Reclassificationfromaocicurrentperiodbeforetaxattributabletoparent                                                                          int64   `json:"reclassificationfromaocicurrentperiodbeforetaxattributabletoparent"`
	Sharebasedcompensationarrangementbysharebasedpaymentawardequityinstrumentsotherthanoptionsvestedinperiod                                    int64   `json:"sharebasedcompensationarrangementbysharebasedpaymentawardequityinstrumentsotherthanoptionsvestedinperiod"`
	Retainedearningsaccumulateddeficit                                                                                                          int64   `json:"retainedearningsaccumulateddeficit"`
	Proceedsfrompaymentsforotherfinancingactivities                                                                                             int64   `json:"proceedsfrompaymentsforotherfinancingactivities"`
	Definedcontributionplanemployermatchingcontributionpercentofmatch                                                                           float64 `json:"definedcontributionplanemployermatchingcontributionpercentofmatch"`
	Othercomprehensiveincomelossreclassificationadjustmentfromaocionderivativesnetoftax                                                         int64   `json:"othercomprehensiveincomelossreclassificationadjustmentfromaocionderivativesnetoftax"`
	Losscontingencyestimateofpossibleloss                                                                                                       int64   `json:"losscontingencyestimateofpossibleloss"`
	Currentfiscalyearenddate                                                                                                                    string  `json:"currentfiscalyearenddate"`
	Employeestockpurchaseplanmaximumannualpurchasesperemployeeamount                                                                            int64   `json:"employeestockpurchaseplanmaximumannualpurchasesperemployeeamount"`
	Longtermdebtfairvalue                                                                                                                       string  `json:"longtermdebtfairvalue"`
	Hedgeaccountingadjustmentsrelatedtolongtermdebt                                                                                             int64   `json:"hedgeaccountingadjustmentsrelatedtolongtermdebt"`
	Paymentstoacquirebusinessesnetofcashacquired                                                                                                int64   `json:"paymentstoacquirebusinessesnetofcashacquired"`
	Investmentincomeinterestanddividend                                                                                                         int64   `json:"investmentincomeinterestanddividend"`
	Othercomprehensiveincomeunrealizedgainlossonderivativesarisingduringperiodnetoftax                                                          int64   `json:"othercomprehensiveincomeunrealizedgainlossonderivativesarisingduringperiodnetoftax"`
	Changeinunrealizedgainlossonhedgediteminfairvaluehedge1                                                                                     int64   `json:"changeinunrealizedgainlossonhedgediteminfairvaluehedge1"`
	Accruedincometaxesnoncurrent                                                                                                                int64   `json:"accruedincometaxesnoncurrent"`
	Operatingexpenses                                                                                                                           int64   `json:"operatingexpenses"`
	Cashmember                                                                                                                                  int64   `json:"cashmember"`
	Operatingincomeloss                                                                                                                         int64   `json:"operatingincomeloss"`
	Deferredtaxassetstaxdeferredexpensecompensationandbenefitssharebasedcompensationcost                                                        int64   `json:"deferredtaxassetstaxdeferredexpensecompensationandbenefitssharebasedcompensationcost"`
	A20132018Debtissuancesmember                                                                                                                int64   `json:"a20132018debtissuancesmember"`
	Noncashactivitiesinvolvingpropertyplantandequipmentnetincreasedecreasetoaccountspayableandothercurrentliabilities                           int64   `json:"noncashactivitiesinvolvingpropertyplantandequipmentnetincreasedecreasetoaccountspayableandothercurrentliabilities"`
	Sharebasedcompensationarrangementbysharebasedpaymentawardequityinstrumentsotherthanoptionsgrantsinperiod                                    int64   `json:"sharebasedcompensationarrangementbysharebasedpaymentawardequityinstrumentsotherthanoptionsgrantsinperiod"`
	Security12Btitle                                                                                                                            string  `json:"security12btitle"`
	Derivativefairvalueofderivativeasset                                                                                                        int64   `json:"derivativefairvalueofderivativeasset"`
	Otherassetsmember                                                                                                                           int64   `json:"otherassetsmember"`
	Paymentsforrepurchaseofcommonstock                                                                                                          int64   `json:"paymentsforrepurchaseofcommonstock"`
	Paymentstoacquireotherinvestments                                                                                                           int64   `json:"paymentstoacquireotherinvestments"`
	Netcashprovidedbyusedininvestingactivities                                                                                                  int64   `json:"netcashprovidedbyusedininvestingactivities"`
	Employeestockmember                                                                                                                         int64   `json:"employeestockmember"`
	Unrecordedunconditionalpurchaseobligationbalanceonfirstanniversary                                                                          int64   `json:"unrecordedunconditionalpurchaseobligationbalanceonfirstanniversary"`
	Assetscurrent                                                                                                                               string  `json:"assetscurrent"`
	Sharebasedcompensationarrangementbysharebasedpaymentawardnumberofsharesavailableforgrant                                                    int64   `json:"sharebasedcompensationarrangementbysharebasedpaymentawardnumberofsharesavailableforgrant"`
	Othernoncashincomeexpense                                                                                                                   int64   `json:"othernoncashincomeexpense"`
	Netcashprovidedbyusedinfinancingactivities                                                                                                  int64   `json:"netcashprovidedbyusedinfinancingactivities"`
	Paymentstoacquireavailableforsalesecuritiesdebt                                                                                             int64   `json:"paymentstoacquireavailableforsalesecuritiesdebt"`
	Derivativeinstrumentsgainlossreclassifiedfromaccumulatedociintoincomeeffectiveportionnet                                                    int64   `json:"derivativeinstrumentsgainlossreclassifiedfromaccumulatedociintoincomeeffectiveportionnet"`
	Restrictedcashandcashequivalentsatcarryingvalue                                                                                             int64   `json:"restrictedcashandcashequivalentsatcarryingvalue"`
	Restrictedcashandcashequivalentsnoncurrent                                                                                                  int64   `json:"restrictedcashandcashequivalentsnoncurrent"`
	Debtsecuritiesavailableforsalecontinuousunrealizedlosspositionlessthan12Monthsaccumulatedloss                                               int64   `json:"debtsecuritiesavailableforsalecontinuousunrealizedlosspositionlessthan12monthsaccumulatedloss"`
	Interestexpense                                                                                                                             int64   `json:"interestexpense"`
	Entityinteractivedatacurrent                                                                                                                string  `json:"entityinteractivedatacurrent"`
	Debtsecuritiesavailableforsaleunrealizedlosspositionaccumulatedloss                                                                         int64   `json:"debtsecuritiesavailableforsaleunrealizedlosspositionaccumulatedloss"`
	Debtinstrumentcarryingamount                                                                                                                string  `json:"debtinstrumentcarryingamount"`
	Contractwithcustomerliabilitycurrent                                                                                                        int64   `json:"contractwithcustomerliabilitycurrent"`
	Weightedaveragenumberofsharesoutstandingbasic                                                                                               int64   `json:"weightedaveragenumberofsharesoutstandingbasic"`
	Documentfiscalyearfocus                                                                                                                     int64   `json:"documentfiscalyearfocus"`
	Sharebasedcompensationarrangementbysharebasedpaymentawardmaximumemployeesubscriptionrate                                                    float64 `json:"sharebasedcompensationarrangementbysharebasedpaymentawardmaximumemployeesubscriptionrate"`
	Documentperiodenddate                                                                                                                       string  `json:"documentperiodenddate"`
	Sharebasedcompensationarrangementbysharebasedpaymentawardequityinstrumentsotherthanoptionsnonvestednumber                                   int64   `json:"sharebasedcompensationarrangementbysharebasedpaymentawardequityinstrumentsotherthanoptionsnonvestednumber"`
	Unrecognizedtaxbenefitsincometaxpenaltiesandinterestexpense                                                                                 int64   `json:"unrecognizedtaxbenefitsincometaxpenaltiesandinterestexpense"`
	Liabilities                                                                                                                                 string  `json:"liabilities"`
	Sharebasedcompensationarrangementbysharebasedpaymentawardequityinstrumentsotherthanoptionsgrantsinperiodweightedaveragegrantdatefairvalue   float64 `json:"sharebasedcompensationarrangementbysharebasedpaymentawardequityinstrumentsotherthanoptionsgrantsinperiodweightedaveragegrantdatefairvalue"`
	Debtinstrumentinterestratestatedpercentage                                                                                                  float64 `json:"debtinstrumentinterestratestatedpercentage"`
	Losscontingencyestimateofpossiblelossreductioninperiod                                                                                      int64   `json:"losscontingencyestimateofpossiblelossreductioninperiod"`
	Operatingleasesfutureminimumpaymentsdueinfouryears                                                                                          int64   `json:"operatingleasesfutureminimumpaymentsdueinfouryears"`
	Commonstockparorstatedvaluepershare                                                                                                         float64 `json:"commonstockparorstatedvaluepershare"`
	Availableforsaledebtsecuritiesaccumulatedgrossunrealizedgainbeforetax                                                                       int64   `json:"availableforsaledebtsecuritiesaccumulatedgrossunrealizedgainbeforetax"`
	Sharebasedcompensationarrangementbysharebasedpaymentawardpurchasepriceofcommonstockpercent                                                  float64 `json:"sharebasedcompensationarrangementbysharebasedpaymentawardpurchasepriceofcommonstockpercent"`
	Incometaxreconciliationstateandlocalincometaxes                                                                                             int64   `json:"incometaxreconciliationstateandlocalincometaxes"`
	Fairvalueinputslevel1Member                                                                                                                 int64   `json:"fairvalueinputslevel1member"`
	Lesseeoperatingleasetermofcontract                                                                                                          string  `json:"lesseeoperatingleasetermofcontract"`
	Debtinstrumentinterestrateeffectivepercentage                                                                                               float64 `json:"debtinstrumentinterestrateeffectivepercentage"`
	Contractwithcustomerliabilityrevenuerecognized                                                                                              int64   `json:"contractwithcustomerliabilityrevenuerecognized"`
	Unrecognizedtaxbenefitsreductionsresultingfromlapseofapplicablestatuteoflimitations                                                         int64   `json:"unrecognizedtaxbenefitsreductionsresultingfromlapseofapplicablestatuteoflimitations"`
	Entityfilercategory                                                                                                                         string  `json:"entityfilercategory"`
	Longtermdebtmaturitiesrepaymentsofprincipalinyeartwo                                                                                        int64   `json:"longtermdebtmaturitiesrepaymentsofprincipalinyeartwo"`
	Accountspayablecurrent                                                                                                                      int64   `json:"accountspayablecurrent"`
	Equitysecuritieswithoutreadilydeterminablefairvalueamount                                                                                   int64   `json:"equitysecuritieswithoutreadilydeterminablefairvalueamount"`
	Sharebasedcompensationarrangementbysharebasedpaymentawardequityinstrumentsotherthanoptionsnumberofsharesofcommonsharesawardeduponsettlement int64   `json:"sharebasedcompensationarrangementbysharebasedpaymentawardequityinstrumentsotherthanoptionsnumberofsharesofcommonsharesawardeduponsettlement"`
	Unrecordedunconditionalpurchaseobligationdueafterfiveyears                                                                                  int64   `json:"unrecordedunconditionalpurchaseobligationdueafterfiveyears"`
	Weightedaveragenumberdilutedsharesoutstandingadjustment                                                                                     int64   `json:"weightedaveragenumberdilutedsharesoutstandingadjustment"`
	Accumulatedothercomprehensiveincomemember                                                                                                   int64   `json:"accumulatedothercomprehensiveincomemember"`
	Fairvaluehedgingmember                                                                                                                      int64   `json:"fairvaluehedgingmember"`
	Derivativeassetsreductionformasternettingarrangements                                                                                       int64   `json:"derivativeassetsreductionformasternettingarrangements"`
	Researchanddevelopmentexpense                                                                                                               int64   `json:"researchanddevelopmentexpense"`
	Allocatedsharebasedcompensationexpense                                                                                                      int64   `json:"allocatedsharebasedcompensationexpense"`
	Availableforsalesecuritiesdebtsecurities                                                                                                    string  `json:"availableforsalesecuritiesdebtsecurities"`
	Liabilitiesandstockholdersequity                                                                                                            string  `json:"liabilitiesandstockholdersequity"`
	Availableforsaledebtsecuritiesamortizedcostbasis                                                                                            string  `json:"availableforsaledebtsecuritiesamortizedcostbasis"`
	Nontradereceivablemember                                                                                                                    int64   `json:"nontradereceivablemember"`
	Incometaxreconciliationincometaxexpensebenefitatfederalstatutoryincometaxrate                                                               int64   `json:"incometaxreconciliationincometaxexpensebenefitatfederalstatutoryincometaxrate"`
	Entityaddresscityortown                                                                                                                     string  `json:"entityaddresscityortown"`
	Derivativenotionalamount                                                                                                                    int64   `json:"derivativenotionalamount"`
	Othercomprehensiveincomelossreclassificationadjustmentfromaociforsaleofsecuritiesnetoftax                                                   int64   `json:"othercomprehensiveincomelossreclassificationadjustmentfromaociforsaleofsecuritiesnetoftax"`
	Foreigncurrencydebtmember                                                                                                                   int64   `json:"foreigncurrencydebtmember"`
	Othercomprehensiveincomelosstaxportionattributabletoparent1                                                                                 int64   `json:"othercomprehensiveincomelosstaxportionattributabletoparent1"`
	Employeeservicesharebasedcompensationnonvestedawardstotalcompensationcostnotyetrecognizedperiodforrecognition1                              string  `json:"employeeservicesharebasedcompensationnonvestedawardstotalcompensationcostnotyetrecognizedperiodforrecognition1"`
	Entityregistrantname                                                                                                                        string  `json:"entityregistrantname"`
	Cashcashequivalentsrestrictedcashandrestrictedcashequivalents                                                                               int64   `json:"cashcashequivalentsrestrictedcashandrestrictedcashequivalents"`
	Entityaddressaddressline1                                                                                                                   string  `json:"entityaddressaddressline1"`
	Reclassificationoutofaccumulatedothercomprehensiveincomemember                                                                              int64   `json:"reclassificationoutofaccumulatedothercomprehensiveincomemember"`
	Stockissuedduringperiodsharessharebasedpaymentarrangementnetofshareswithheldfortaxes                                                        int64   `json:"stockissuedduringperiodsharessharebasedpaymentarrangementnetofshareswithheldfortaxes"`
	Landandbuildingmember                                                                                                                       int64   `json:"landandbuildingmember"`
	Deferredtaxliabilitiesminimumtaxonforeignearnings                                                                                           int64   `json:"deferredtaxliabilitiesminimumtaxonforeignearnings"`
	Europesegmentmember                                                                                                                         int64   `json:"europesegmentmember"`
	Effectiveincometaxratereconciliationsharebasedcompensationexcesstaxbenefitamount                                                            int64   `json:"effectiveincometaxratereconciliationsharebasedcompensationexcesstaxbenefitamount"`
	Numberofcustomerswithsignificantaccountsreceivablebalance                                                                                   int64   `json:"numberofcustomerswithsignificantaccountsreceivablebalance"`
	Marketablesecuritiesnoncurrent                                                                                                              string  `json:"marketablesecuritiesnoncurrent"`
	Sharebasedcompensationarrangementbysharebasedpaymentawardawardvestingperiod1                                                                string  `json:"sharebasedcompensationarrangementbysharebasedpaymentawardawardvestingperiod1"`
	Derivativecounterpartycreditriskexposure                                                                                                    int64   `json:"derivativecounterpartycreditriskexposure"`
	Changeinunrealizedgainlossonfairvaluehedginginstruments1                                                                                    int64   `json:"changeinunrealizedgainlossonfairvaluehedginginstruments1"`
	Operatingleasesrentexpensenet                                                                                                               int64   `json:"operatingleasesrentexpensenet"`
	Othercurrentliabilitiesmember                                                                                                               int64   `json:"othercurrentliabilitiesmember"`
	Amendmentflag                                                                                                                               string  `json:"amendmentflag"`
	Sharebasedcompensationarrangementbysharebasedpaymentawardequityinstrumentsotherthanoptionsnonvestedweightedaveragegrantdatefairvalue        float64 `json:"sharebasedcompensationarrangementbysharebasedpaymentawardequityinstrumentsotherthanoptionsnonvestedweightedaveragegrantdatefairvalue"`
	Foreignexchangecontractmember                                                                                                               string  `json:"foreignexchangecontractmember"`
	Unrecognizedtaxbenefits                                                                                                                     int64   `json:"unrecognizedtaxbenefits"`
	Increasedecreaseincontractwithcustomerliability                                                                                             int64   `json:"increasedecreaseincontractwithcustomerliability"`
	Sharebasedcompensationarrangementbysharebasedpaymentawardequityinstrumentsotherthanoptionsforfeituresweightedaveragegrantdatefairvalue      float64 `json:"sharebasedcompensationarrangementbysharebasedpaymentawardequityinstrumentsotherthanoptionsforfeituresweightedaveragegrantdatefairvalue"`
	Deferredincometaxexpensebenefit                                                                                                             int64   `json:"deferredincometaxexpensebenefit"`
	Entityaddressstateorprovince                                                                                                                string  `json:"entityaddressstateorprovince"`
	Propertyplantandequipmentusefullife                                                                                                         string  `json:"propertyplantandequipmentusefullife"`
	Unrecognizedtaxbenefitsincreasesresultingfrompriorperiodtaxpositions                                                                        int64   `json:"unrecognizedtaxbenefitsincreasesresultingfrompriorperiodtaxpositions"`
	Documenttype                                                                                                                                string  `json:"documenttype"`
	Paymentsforproceedsfromotherinvestingactivities                                                                                             int64   `json:"paymentsforproceedsfromotherinvestingactivities"`
	Entitysmallbusiness                                                                                                                         string  `json:"entitysmallbusiness"`
	Revenuefromcontractwithcustomerexcludingassessedtax                                                                                         string  `json:"revenuefromcontractwithcustomerexcludingassessedtax"`
	Currentfederaltaxexpensebenefit                                                                                                             int64   `json:"currentfederaltaxexpensebenefit"`
	Sharebasedcompensationarrangementbysharebasedpaymentawardequityinstrumentsotherthanoptionsaggregateintrinsicvaluenonvested                  int64   `json:"sharebasedcompensationarrangementbysharebasedpaymentawardequityinstrumentsotherthanoptionsaggregateintrinsicvaluenonvested"`
	Amountutilizedundersharerepurchaseprogram                                                                                                   int64   `json:"amountutilizedundersharerepurchaseprogram"`
	Debtinstrumentmaturityyearrangeend                                                                                                          int64   `json:"debtinstrumentmaturityyearrangeend"`
	Commonstocksharesauthorized                                                                                                                 int64   `json:"commonstocksharesauthorized"`
	Deferredtaxassetsdeferredincome                                                                                                             int64   `json:"deferredtaxassetsdeferredincome"`
	Paymentsofdividends                                                                                                                         int64   `json:"paymentsofdividends"`
	Debtsecuritiesavailableforsalecontinuousunrealizedlossposition12Monthsorlongeraccumulatedloss                                               int64   `json:"debtsecuritiesavailableforsalecontinuousunrealizedlossposition12monthsorlongeraccumulatedloss"`
	Entityincorporationstatecountrycode                                                                                                         string  `json:"entityincorporationstatecountrycode"`
	Effectiveincometaxratecontinuingoperations                                                                                                  float64 `json:"effectiveincometaxratecontinuingoperations"`
	Factorbywhicheachrsugrantedreducesandeachrsucanceledorsharewithheldfortaxesincreasessharesavailableforgrant                                 int64   `json:"factorbywhicheachrsugrantedreducesandeachrsucanceledorsharewithheldfortaxesincreasessharesavailableforgrant"`
	Increasedecreaseininventories                                                                                                               int64   `json:"increasedecreaseininventories"`
	Ocibeforereclassificationsbeforetaxattributabletoparent                                                                                     int64   `json:"ocibeforereclassificationsbeforetaxattributabletoparent"`
	Dividends                                                                                                                                   int64   `json:"dividends"`
	Entitywellknownseasonedissuer                                                                                                               string  `json:"entitywellknownseasonedissuer"`
	Increasedecreaseinaccountsreceivable                                                                                                        int64   `json:"increasedecreaseinaccountsreceivable"`
	Entityvoluntaryfilers                                                                                                                       string  `json:"entityvoluntaryfilers"`
	Otherliabilitiescurrent                                                                                                                     int64   `json:"otherliabilitiescurrent"`
	Otheraccruedliabilitiesnoncurrent                                                                                                           int64   `json:"otheraccruedliabilitiesnoncurrent"`
	Operatingleasesfutureminimumpaymentsduecurrent                                                                                              int64   `json:"operatingleasesfutureminimumpaymentsduecurrent"`
	Proceedsfrommaturitiesprepaymentsandcallsofavailableforsalesecurities                                                                       int64   `json:"proceedsfrommaturitiesprepaymentsandcallsofavailableforsalesecurities"`
	Revenues                                                                                                                                    int64   `json:"revenues"`
	Retainedearningsmember                                                                                                                      int64   `json:"retainedearningsmember"`
	Depreciation                                                                                                                                int64   `json:"depreciation"`
	Longtermdebtmaturitiesrepaymentsofprincipalafteryearfive                                                                                    int64   `json:"longtermdebtmaturitiesrepaymentsofprincipalafteryearfive"`
	Cityareacode                                                                                                                                int64   `json:"cityareacode"`
	Employeeservicesharebasedcompensationtaxbenefitfromcompensationexpense                                                                      int64   `json:"employeeservicesharebasedcompensationtaxbenefitfromcompensationexpense"`
	Accumulatednetgainlossfromdesignatedorqualifyingcashflowhedgesmember                                                                        int64   `json:"accumulatednetgainlossfromdesignatedorqualifyingcashflowhedgesmember"`
	Sharebasedcompensationarrangementbysharebasedpaymentawardequityinstrumentsotherthanoptionsforfeitedinperiod                                 int64   `json:"sharebasedcompensationarrangementbysharebasedpaymentawardequityinstrumentsotherthanoptionsforfeitedinperiod"`
	Entityaddresspostalzipcode                                                                                                                  int64   `json:"entityaddresspostalzipcode"`
	Comprehensiveincomenetoftax                                                                                                                 int64   `json:"comprehensiveincomenetoftax"`
	Adjustmentstoadditionalpaidincapitalsharebasedcompensationrequisiteserviceperiodrecognitionvalue                                            int64   `json:"adjustmentstoadditionalpaidincapitalsharebasedcompensationrequisiteserviceperiodrecognitionvalue"`
	Fairvalueinputslevel2Member                                                                                                                 string  `json:"fairvalueinputslevel2member"`
	Deferredincometaxliabilities                                                                                                                int64   `json:"deferredincometaxliabilities"`
	Othercomprehensiveincomelossforeigncurrencytransactionandtranslationadjustmentnetoftax                                                      int64   `json:"othercomprehensiveincomelossforeigncurrencytransactionandtranslationadjustmentnetoftax"`
	Concentrationriskpercentage1                                                                                                                float64 `json:"concentrationriskpercentage1"`
	Stockholdersequity                                                                                                                          int64   `json:"stockholdersequity"`
	Effectiveincometaxratereconciliationtaxcutsandjobsactof2017Amount                                                                           int64   `json:"effectiveincometaxratereconciliationtaxcutsandjobsactof2017amount"`
	Cashflowhedgingmember                                                                                                                       int64   `json:"cashflowhedgingmember"`
	Propertyplantandequipmentgross                                                                                                              int64   `json:"propertyplantandequipmentgross"`
	Tradingsymbol                                                                                                                               string  `json:"tradingsymbol"`
	Foreignincometaxexpensebenefitcontinuingoperations                                                                                          int64   `json:"foreignincometaxexpensebenefitcontinuingoperations"`
	Operatingleasesfutureminimumpaymentsdueintwoyears                                                                                           int64   `json:"operatingleasesfutureminimumpaymentsdueintwoyears"`
	Sharebasedcompensationarrangementbysharebasedpaymentawardequityinstrumentsotherthanoptionsvestedinperiodweightedaveragegrantdatefairvalue   float64 `json:"sharebasedcompensationarrangementbysharebasedpaymentawardequityinstrumentsotherthanoptionsvestedinperiodweightedaveragegrantdatefairvalue"`
	Employeestockplan2014Planmember                                                                                                             int64   `json:"employeestockplan2014planmember"`
	A2019Debtissuancemember                                                                                                                     int64   `json:"a2019debtissuancemember"`
	Documenttransitionreport                                                                                                                    string  `json:"documenttransitionreport"`
	Othercomprehensiveincomelossderivativesqualifyingashedgesnetoftax                                                                           int64   `json:"othercomprehensiveincomelossderivativesqualifyingashedgesnetoftax"`
	Liabilitiesnoncurrent                                                                                                                       string  `json:"liabilitiesnoncurrent"`
	Shorttermdebtweightedaverageinterestrate                                                                                                    float64 `json:"shorttermdebtweightedaverageinterestrate"`
	Incometaxreconciliationotheradjustments                                                                                                     int64   `json:"incometaxreconciliationotheradjustments"`
	Depreciationdepletionandamortization                                                                                                        int64   `json:"depreciationdepletionandamortization"`
	Deferredforeignincometaxexpensebenefit                                                                                                      int64   `json:"deferredforeignincometaxexpensebenefit"`
	Unrecordedunconditionalpurchaseobligationbalanceonthirdanniversary                                                                          int64   `json:"unrecordedunconditionalpurchaseobligationbalanceonthirdanniversary"`
	Operatingleasesfutureminimumpaymentsdueinthreeyears                                                                                         int64   `json:"operatingleasesfutureminimumpaymentsdueinthreeyears"`
	Definedcontributionplanemployermatchingcontributionpercent                                                                                  float64 `json:"definedcontributionplanemployermatchingcontributionpercent"`
	Debtsecuritiesavailableforsalecontinuousunrealizedlosspositionlessthan12Months                                                              int64   `json:"debtsecuritiesavailableforsalecontinuousunrealizedlosspositionlessthan12months"`
	Deferredtaxassetstaxdeferredexpensereservesandaccruals                                                                                      int64   `json:"deferredtaxassetstaxdeferredexpensereservesandaccruals"`
	Unrecognizedtaxbenefitsdecreasesresultingfrompriorperiodtaxpositions                                                                        int64   `json:"unrecognizedtaxbenefitsdecreasesresultingfrompriorperiodtaxpositions"`
	Stateandlocalincometaxexpensebenefitcontinuingoperations                                                                                    int64   `json:"stateandlocalincometaxexpensebenefitcontinuingoperations"`
	Deferredstateandlocalincometaxexpensebenefit                                                                                                int64   `json:"deferredstateandlocalincometaxexpensebenefit"`
	Otherliabilitiesnoncurrent                                                                                                                  int64   `json:"otherliabilitiesnoncurrent"`
	Commonstockincludingadditionalpaidincapitalmember                                                                                           int64   `json:"commonstockincludingadditionalpaidincapitalmember"`
	Deferredtaxassetsother                                                                                                                      int64   `json:"deferredtaxassetsother"`
	Standardproductwarrantyaccrual                                                                                                              int64   `json:"standardproductwarrantyaccrual"`
	Deferredtaxliabilitiesundistributedforeignearnings                                                                                          int64   `json:"deferredtaxliabilitiesundistributedforeignearnings"`
	Marketablesecuritiescurrent                                                                                                                 int64   `json:"marketablesecuritiescurrent"`
	Derivativefairvalueofderivativeliability                                                                                                    int64   `json:"derivativefairvalueofderivativeliability"`
	Unrecognizedtaxbenefitsdecreasesresultingfromsettlementswithtaxingauthorities                                                               int64   `json:"unrecognizedtaxbenefitsdecreasesresultingfromsettlementswithtaxingauthorities"`
	Longtermdebtmaturitiesrepaymentsofprincipalinyearfour                                                                                       int64   `json:"longtermdebtmaturitiesrepaymentsofprincipalinyearfour"`
	Othergeneralandadministrativeexpense                                                                                                        int64   `json:"othergeneralandadministrativeexpense"`
	Longtermdebtmaturitiesrepaymentsofprincipalinyearfive                                                                                       int64   `json:"longtermdebtmaturitiesrepaymentsofprincipalinyearfive"`
	Commonstockdividendspersharedeclared                                                                                                        int64   `json:"commonstockdividendspersharedeclared"`
	Directorplanmember                                                                                                                          int64   `json:"directorplanmember"`
	Adjustmentsrelatedtotaxwithholdingforsharebasedcompensation                                                                                 int64   `json:"adjustmentsrelatedtotaxwithholdingforsharebasedcompensation"`
	Unrecordedunconditionalpurchaseobligationbalancesheetamount                                                                                 int64   `json:"unrecordedunconditionalpurchaseobligationbalancesheetamount"`
	Incometaxexpensebenefit                                                                                                                     int64   `json:"incometaxexpensebenefit"`
	Othercomprehensiveincomelossnetoftaxportionattributabletoparent                                                                             int64   `json:"othercomprehensiveincomelossnetoftaxportionattributabletoparent"`
	Restrictedstockunitsrsumember                                                                                                               int64   `json:"restrictedstockunitsrsumember"`
	Derivativeliabilitiesreductionformasternettingarrangements                                                                                  int64   `json:"derivativeliabilitiesreductionformasternettingarrangements"`
	Nontradereceivablescurrent                                                                                                                  int64   `json:"nontradereceivablescurrent"`
	Sharebasedcompensationarrangementbysharebasedpaymentawardequityinstrumentsotherthanoptionsvestedinperiodtotalfairvalue                      int64   `json:"sharebasedcompensationarrangementbysharebasedpaymentawardequityinstrumentsotherthanoptionsvestedinperiodtotalfairvalue"`
	Othercomprehensiveincomeunrealizedholdinggainlossonsecuritiesarisingduringperiodnetoftax                                                    int64   `json:"othercomprehensiveincomeunrealizedholdinggainlossonsecuritiesarisingduringperiodnetoftax"`
	Inventorynet                                                                                                                                int64   `json:"inventorynet"`
	Increasedecreaseinotherreceivables                                                                                                          int64   `json:"increasedecreaseinotherreceivables"`
	Entitycentralindexkey                                                                                                                       int64   `json:"entitycentralindexkey"`
	Paymentstoacquirepropertyplantandequipment                                                                                                  int64   `json:"paymentstoacquirepropertyplantandequipment"`
	Debtinstrumentterm                                                                                                                          string  `json:"debtinstrumentterm"`
	Revenueremainingperformanceobligationexpectedtimingofsatisfactionperiod1                                                                    string  `json:"revenueremainingperformanceobligationexpectedtimingofsatisfactionperiod1"`
	Availableforsaledebtsecuritiesaccumulatedgrossunrealizedlossbeforetax                                                                       int64   `json:"availableforsaledebtsecuritiesaccumulatedgrossunrealizedlossbeforetax"`
	Unrecognizedtaxbenefitsincometaxpenaltiesandinterestaccrued                                                                                 int64   `json:"unrecognizedtaxbenefitsincometaxpenaltiesandinterestaccrued"`
	Entityfilenumber                                                                                                                            string  `json:"entityfilenumber"`
	Assets                                                                                                                                      string  `json:"assets"`
	Seniornotes                                                                                                                                 int64   `json:"seniornotes"`
	Assetsnoncurrent                                                                                                                            string  `json:"assetsnoncurrent"`
	Earningspersharebasic                                                                                                                       float64 `json:"earningspersharebasic"`
	Documentfiscalperiodfocus                                                                                                                   string  `json:"documentfiscalperiodfocus"`
	Standardproductwarrantyaccrualwarrantiesissued                                                                                              int64   `json:"standardproductwarrantyaccrualwarrantiesissued"`
	Proceedsfromshorttermdebtmaturinginmorethanthreemonths                                                                                      int64   `json:"proceedsfromshorttermdebtmaturinginmorethanthreemonths"`
	Entityshellcompany                                                                                                                          string  `json:"entityshellcompany"`
	Sellinggeneralandadministrativeexpense                                                                                                      int64   `json:"sellinggeneralandadministrativeexpense"`
	Deferredtaxliabilitiesother                                                                                                                 int64   `json:"deferredtaxliabilitiesother"`
	Operatingleasesfutureminimumpaymentsdueinfiveyears                                                                                          int64   `json:"operatingleasesfutureminimumpaymentsdueinfiveyears"`
	Netinvestmenthedgingmember                                                                                                                  int64   `json:"netinvestmenthedgingmember"`
	Longtermdebtmaturitiesrepaymentsofprincipalinyearthree                                                                                      int64   `json:"longtermdebtmaturitiesrepaymentsofprincipalinyearthree"`
	Effectiveincometaxratereconciliationatfederalstatutoryincometaxrate                                                                         float64 `json:"effectiveincometaxratereconciliationatfederalstatutoryincometaxrate"`
	Noncurrentassets                                                                                                                            int64   `json:"noncurrentassets"`
	Otherassetsnoncurrent                                                                                                                       int64   `json:"otherassetsnoncurrent"`
	Commonstocksharesissued                                                                                                                     int64   `json:"commonstocksharesissued"`
	Antidilutivesecuritiesexcludedfromcomputationofearningspershareamount                                                                       int64   `json:"antidilutivesecuritiesexcludedfromcomputationofearningspershareamount"`
	Unrecognizedtaxbenefitsincreasesresultingfromcurrentperiodtaxpositions                                                                      int64   `json:"unrecognizedtaxbenefitsincreasesresultingfromcurrentperiodtaxpositions"`
	Longtermdebtmaturitiesrepaymentsofprincipalinnexttwelvemonths                                                                               int64   `json:"longtermdebtmaturitiesrepaymentsofprincipalinnexttwelvemonths"`
	Accumulatedothercomprehensiveincomelossnetoftax                                                                                             int64   `json:"accumulatedothercomprehensiveincomelossnetoftax"`
	Lesseeoperatingandfinanceleaseliabilitypaymentsdueafteryearfive                                                                             int64   `json:"lesseeoperatingandfinanceleaseliabilitypaymentsdueafteryearfive"`
	Operatingandfinanceleaseweightedaverageremainingleaseterm                                                                                   string  `json:"operatingandfinanceleaseweightedaverageremainingleaseterm"`
	Acceleratedsharerepurchasearrangementnovember2019Member                                                                                     int64   `json:"acceleratedsharerepurchasearrangementnovember2019member"`
	Stockrepurchaseprogramauthorizedamount1                                                                                                     string  `json:"stockrepurchaseprogramauthorizedamount1"`
	Lesseeoperatingandfinanceleaseliabilitypaymentsdueyearfive                                                                                  int64   `json:"lesseeoperatingandfinanceleaseliabilitypaymentsdueyearfive"`
	Othercomprehensiveincomelossnetinvestmenthedgegainlossbeforereclassificationandtax                                                          int64   `json:"othercomprehensiveincomelossnetinvestmenthedgegainlossbeforereclassificationandtax"`
	Lesseeoperatingleaseleasenotyetcommencedpaymentsdue                                                                                         int64   `json:"lesseeoperatingleaseleasenotyetcommencedpaymentsdue"`
	Debtinstrumentfaceamount                                                                                                                    int64   `json:"debtinstrumentfaceamount"`
	Documentquarterlyreport                                                                                                                     string  `json:"documentquarterlyreport"`
	Lesseeoperatingandfinanceleaseliabilitypaymentsdueyearfour                                                                                  int64   `json:"lesseeoperatingandfinanceleaseliabilitypaymentsdueyearfour"`
	Othercomprehensiveincomelosscashflowhedgegainlossbeforereclassificationandtax                                                               int64   `json:"othercomprehensiveincomelosscashflowhedgegainlossbeforereclassificationandtax"`
	Lesseeoperatingandfinanceleaseliabilityundiscountedexcessamount                                                                             int64   `json:"lesseeoperatingandfinanceleaseliabilityundiscountedexcessamount"`
	Operatingleaseliabilitycurrent                                                                                                              int64   `json:"operatingleaseliabilitycurrent"`
	Operatingandfinanceleaseweightedaveragediscountratepercent                                                                                  float64 `json:"operatingandfinanceleaseweightedaveragediscountratepercent"`
	Lesseeoperatingleaseliabilitypaymentsdueafteryearfive                                                                                       int64   `json:"lesseeoperatingleaseliabilitypaymentsdueafteryearfive"`
	Lesseeoperatingleaseliabilitypaymentsremainderoffiscalyear                                                                                  int64   `json:"lesseeoperatingleaseliabilitypaymentsremainderoffiscalyear"`
	Rightofuseassetsobtainedinexchangeforoperatingandfinanceleaseliabilities                                                                    int64   `json:"rightofuseassetsobtainedinexchangeforoperatingandfinanceleaseliabilities"`
	Othernoncurrentliabilitiesmember                                                                                                            int64   `json:"othernoncurrentliabilitiesmember"`
	Lesseeoperatingleaseliabilitypaymentsdueyeartwo                                                                                             int64   `json:"lesseeoperatingleaseliabilitypaymentsdueyeartwo"`
	Lesseeoperatingandfinanceleaseliabilitypaymentsremainderoffiscalyear                                                                        int64   `json:"lesseeoperatingandfinanceleaseliabilitypaymentsremainderoffiscalyear"`
	Operatingleasepayments                                                                                                                      int64   `json:"operatingleasepayments"`
	Operatingleaseliabilitynoncurrent                                                                                                           int64   `json:"operatingleaseliabilitynoncurrent"`
	Secureddebtrepurchaseagreements                                                                                                             int64   `json:"secureddebtrepurchaseagreements"`
	Variableleasecost                                                                                                                           int64   `json:"variableleasecost"`
	Othercomprehensiveincomelossderivativeexcludedcomponentincreasedecreasebeforeadjustmentsandtax                                              int64   `json:"othercomprehensiveincomelossderivativeexcludedcomponentincreasedecreasebeforeadjustmentsandtax"`
	Financeleaseliabilitycurrent                                                                                                                int64   `json:"financeleaseliabilitycurrent"`
	Financeleaseliabilitynoncurrent                                                                                                             int64   `json:"financeleaseliabilitynoncurrent"`
	Currenttermdebtandnoncurrenttermdebtmember                                                                                                  int64   `json:"currenttermdebtandnoncurrenttermdebtmember"`
	Hedgedliabilityfairvaluehedgecumulativeincreasedecrease                                                                                     int64   `json:"hedgedliabilityfairvaluehedgecumulativeincreasedecrease"`
	Othernoncurrentassetsmember                                                                                                                 int64   `json:"othernoncurrentassetsmember"`
	Hedgedassetfairvaluehedgecumulativeincreasedecrease                                                                                         int64   `json:"hedgedassetfairvaluehedgecumulativeincreasedecrease"`
	Nonoperatingincomeexpensemember                                                                                                             int64   `json:"nonoperatingincomeexpensemember"`
	Acceleratedsharerepurchasearrangementmay2020Member                                                                                          int64   `json:"acceleratedsharerepurchasearrangementmay2020member"`
	Securitiessoldunderagreementstorepurchasefairvalueofcollateral                                                                              int64   `json:"securitiessoldunderagreementstorepurchasefairvalueofcollateral"`
	Lesseeoperatingleaseliabilitypaymentsdueyearthree                                                                                           int64   `json:"lesseeoperatingleaseliabilitypaymentsdueyearthree"`
	Lesseeoperatingleaseliabilitypaymentsdueyearfive                                                                                            int64   `json:"lesseeoperatingleaseliabilitypaymentsdueyearfive"`
	Commercialpapermember                                                                                                                       string  `json:"commercialpapermember"`
	Lesseeoperatingleaseliabilityundiscountedexcessamount                                                                                       int64   `json:"lesseeoperatingleaseliabilityundiscountedexcessamount"`
	Accumulatedgainlossnetderivativeinstrumentparentmember                                                                                      int64   `json:"accumulatedgainlossnetderivativeinstrumentparentmember"`
	Sharebasedcompensationarrangementbysharebasedpaymentawardequityinstrumentsotherthanoptionsnumberofsharesofcommonstockawardeduponsettlement  int64   `json:"sharebasedcompensationarrangementbysharebasedpaymentawardequityinstrumentsotherthanoptionsnumberofsharesofcommonstockawardeduponsettlement"`
	Lesseeoperatingleaseliabilitypaymentsdueyearfour                                                                                            int64   `json:"lesseeoperatingleaseliabilitypaymentsdueyearfour"`
	Lesseeoperatingandfinanceleasetermofcontract                                                                                                string  `json:"lesseeoperatingandfinanceleasetermofcontract"`
	Lesseeoperatingandfinanceleaseliabilitypaymentsdueyearthree                                                                                 int64   `json:"lesseeoperatingandfinanceleaseliabilitypaymentsdueyearthree"`
	Othercomprehensiveincomelosscashflowhedgegainlossreclassificationbeforetax                                                                  int64   `json:"othercomprehensiveincomelosscashflowhedgegainlossreclassificationbeforetax"`
	Lesseeoperatingleaseleasenotyetcommencedtermofcontract1                                                                                     string  `json:"lesseeoperatingleaseleasenotyetcommencedtermofcontract1"`
	Operatingandfinanceleaserightofuseasset                                                                                                     int64   `json:"operatingandfinanceleaserightofuseasset"`
	Operatingleasecost                                                                                                                          int64   `json:"operatingleasecost"`
	Currentmarketablesecuritiesandnoncurrentmarketablesecuritiesmember                                                                          int64   `json:"currentmarketablesecuritiesandnoncurrentmarketablesecuritiesmember"`
	Lesseeoperatingandfinanceleaseliabilitypaymentsdueyeartwo                                                                                   int64   `json:"lesseeoperatingandfinanceleaseliabilitypaymentsdueyeartwo"`
	Securitiessoldunderagreementstorepurchasemember                                                                                             string  `json:"securitiessoldunderagreementstorepurchasemember"`
	A20132019Debtissuancesmember                                                                                                                int64   `json:"a20132019debtissuancesmember"`
	Operatingleaserightofuseasset                                                                                                               int64   `json:"operatingleaserightofuseasset"`
	Firstquarter2020Debtissuancemember                                                                                                          int64   `json:"firstquarter2020debtissuancemember"`
	Thirdquarter2020Debtissuancemember                                                                                                          int64   `json:"thirdquarter2020debtissuancemember"`
	Propertyplantandequipmentmember                                                                                                             int64   `json:"propertyplantandequipmentmember"`
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
	TenYDividendperShareGrowthPerShare     int64   `json:"tenYDividendperShareGrowthPerShare"`
	FiveYDividendperShareGrowthPerShare    float64 `json:"fiveYDividendperShareGrowthPerShare"`
	ThreeYDividendperShareGrowthPerShare   float64 `json:"threeYDividendperShareGrowthPerShare"`
	ReceivablesGrowth                      float64 `json:"receivablesGrowth"`
	InventoryGrowth                        float64 `json:"inventoryGrowth"`
	AssetGrowth                            float64 `json:"assetGrowth"`
	BookValueperShareGrowth                int64   `json:"bookValueperShareGrowth"`
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
