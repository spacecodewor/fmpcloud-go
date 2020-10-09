package fmpcloud

import (
	"testing"

	"github.com/spacecodewor/fmpcloud-go/objects"
)

// Varibles for test check cases
var (
	testCaseSymbolList       = []string{"AAPL", "GM", "NVDA", "TSLA", "ADBE", "JPM", "BAC"}
	testCaseETFList          = []string{"QQQ", "SPY"}
	testCaseAPIConfig        = Config{Debug: true, Timeout: 60}
	testCaseLimit      int64 = 5
)

func TestRssFeed(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.CompanyValuation.RssFeed()
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestEarningCalendar(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.CompanyValuation.EarningCalendar(nil, nil)
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestEarningSurpriseList(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	for _, symbol := range testCaseSymbolList {
		_, err = APIClient.CompanyValuation.EarningSurpriseList(symbol)
		if err != nil {
			t.Fatal(err.Error())
		}
	}
}

func TestEarningCallTranscript(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	for _, symbol := range testCaseSymbolList {
		_, err = APIClient.CompanyValuation.EarningCallTranscript(objects.RequestEarningCallTranscript{
			Symbol:  symbol,
			Year:    2019,
			Quarter: 2,
		})
		if err != nil {
			t.Fatal(err.Error())
		}
	}
}

func TestHistoryEarningCalendar(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	for _, symbol := range testCaseSymbolList {
		_, err = APIClient.CompanyValuation.HistoryEarningCalendar(symbol)
		if err != nil {
			t.Fatal(err.Error())
		}
	}
}

func TestIPOCalendar(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.CompanyValuation.IPOCalendar(nil, nil)
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestSplitCalendar(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.CompanyValuation.SplitCalendar(nil, nil)
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestInstitutionalHolders(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	for _, symbol := range testCaseSymbolList {
		_, err = APIClient.CompanyValuation.InstitutionalHolders(symbol)
		if err != nil {
			t.Fatal(err.Error())
		}
	}
}

func TestMutualFundHolders(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	for _, symbol := range testCaseSymbolList {
		_, err = APIClient.CompanyValuation.MutualFundHolders(symbol)
		if err != nil {
			t.Fatal(err.Error())
		}
	}
}

func TestETFHolders(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	for _, etf := range testCaseETFList {
		_, err = APIClient.CompanyValuation.ETFHolders(etf)
		if err != nil {
			t.Fatal(err.Error())
		}
	}
}

func TestETFSectorWeightings(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	for _, etf := range testCaseETFList {
		_, err = APIClient.CompanyValuation.ETFSectorWeightings(etf)
		if err != nil {
			t.Fatal(err.Error())
		}
	}
}

func TestETFCountryWeightings(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	for _, etf := range testCaseETFList {
		_, err = APIClient.CompanyValuation.ETFCountryWeightings(etf)
		if err != nil {
			t.Fatal(err.Error())
		}
	}
}

func TestIncomeStatement(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	for _, symbol := range testCaseSymbolList {
		_, err := APIClient.CompanyValuation.IncomeStatement(objects.RequestIncomeStatement{
			Symbol: symbol,
			Period: objects.CompanyValuationPeriodAnnual,
			Limit:  testCaseLimit,
		})
		if err != nil {
			t.Fatal(err.Error())
		}
	}
}

func TestIncomeStatementGrowth(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	for _, symbol := range testCaseSymbolList {
		_, err = APIClient.CompanyValuation.IncomeStatementGrowth(objects.RequestIncomeStatementGrowth{
			Symbol: symbol,
			Period: objects.CompanyValuationPeriodAnnual,
			Limit:  testCaseLimit,
		})
		if err != nil {
			t.Fatal(err.Error())
		}
	}
}

func TestBalanceSheetStatement(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	for _, symbol := range testCaseSymbolList {
		_, err = APIClient.CompanyValuation.BalanceSheetStatement(objects.RequestBalanceSheetStatement{
			Symbol: symbol,
			Period: objects.CompanyValuationPeriodAnnual,
			Limit:  testCaseLimit,
		})
		if err != nil {
			t.Fatal(err.Error())
		}
	}
}

func TestBalanceSheetStatementGrowth(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	for _, symbol := range testCaseSymbolList {
		_, err = APIClient.CompanyValuation.BalanceSheetStatementGrowth(objects.RequestBalanceSheetStatementGrowth{
			Symbol: symbol,
			Period: objects.CompanyValuationPeriodAnnual,
			Limit:  testCaseLimit,
		})
		if err != nil {
			t.Fatal(err.Error())
		}
	}
}

func TestCashFlowStatement(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	for _, symbol := range testCaseSymbolList {
		_, err = APIClient.CompanyValuation.CashFlowStatement(objects.RequestCashFlowStatement{
			Symbol: symbol,
			Period: objects.CompanyValuationPeriodAnnual,
			Limit:  testCaseLimit,
		})
		if err != nil {
			t.Fatal(err.Error())
		}
	}
}

func TestCashFlowStatementGrowth(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	for _, symbol := range testCaseSymbolList {
		_, err = APIClient.CompanyValuation.CashFlowStatementGrowth(objects.RequestCashFlowStatementGrowth{
			Symbol: symbol,
			Period: objects.CompanyValuationPeriodAnnual,
			Limit:  testCaseLimit,
		})
		if err != nil {
			t.Fatal(err.Error())
		}
	}
}

func TestIncomeStatementAsReported(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	for _, symbol := range testCaseSymbolList {
		_, err = APIClient.CompanyValuation.IncomeStatementAsReported(objects.RequestIncomeStatementAsReported{
			Symbol: symbol,
			Period: objects.CompanyValuationPeriodAnnual,
			Limit:  testCaseLimit,
		})
		if err != nil {
			t.Fatal(err.Error())
		}
	}
}

func TestBalanceSheetStatementAsReported(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	for _, symbol := range testCaseSymbolList {
		_, err = APIClient.CompanyValuation.BalanceSheetStatementAsReported(objects.RequestBalanceSheetStatementAsReported{
			Symbol: symbol,
			Period: objects.CompanyValuationPeriodAnnual,
			Limit:  testCaseLimit,
		})
		if err != nil {
			t.Fatal(err.Error())
		}
	}
}

func TestCashFlowStatementAsReported(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	for _, symbol := range testCaseSymbolList {
		_, err = APIClient.CompanyValuation.CashFlowStatementAsReported(objects.RequestCashFlowStatementAsReported{
			Symbol: symbol,
			Period: objects.CompanyValuationPeriodAnnual,
			Limit:  testCaseLimit,
		})
		if err != nil {
			t.Fatal(err.Error())
		}
	}
}

func TestFullFinancialStatementAsReported(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	for _, symbol := range testCaseSymbolList {
		_, err = APIClient.CompanyValuation.FullFinancialStatementAsReported(objects.RequestFullFinancialStatementAsReported{
			Symbol: symbol,
			Period: objects.CompanyValuationPeriodAnnual,
			Limit:  testCaseLimit,
		})
		if err != nil {
			t.Fatal(err.Error())
		}
	}
}

func TestStockNews(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.CompanyValuation.StockNews(objects.RequestStockNews{
		SymbolList: testCaseSymbolList,
		Limit:      testCaseLimit,
	})
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestDelstedCompanies(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.CompanyValuation.DelstedCompanies(testCaseLimit)
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestFinancialRatios(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	for _, symbol := range testCaseSymbolList {
		_, err = APIClient.CompanyValuation.FinancialRatios(objects.RequestFinancialRatios{
			Symbol: symbol,
			Period: objects.CompanyValuationPeriodAnnual,
			Limit:  testCaseLimit,
		})
		if err != nil {
			t.Fatal(err.Error())
		}
	}
}

func TestFinancialRatiosTTM(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	for _, symbol := range testCaseSymbolList {
		_, err = APIClient.CompanyValuation.FinancialRatiosTTM(symbol)
		if err != nil {
			t.Fatal(err.Error())
		}
	}
}

func TestKeyMetrics(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	for _, symbol := range testCaseSymbolList {
		_, err = APIClient.CompanyValuation.KeyMetrics(objects.RequestKeyMetrics{
			Symbol: symbol,
			Period: objects.CompanyValuationPeriodQuarter,
			Limit:  testCaseLimit,
		})
		if err != nil {
			t.Fatal(err.Error())
		}
	}
}

func TestKeyMetricsTTM(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	for _, symbol := range testCaseSymbolList {
		_, err = APIClient.CompanyValuation.KeyMetricsTTM(symbol)
		if err != nil {
			t.Fatal(err.Error())
		}
	}
}

func TestSECFilings(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	for _, symbol := range testCaseSymbolList {
		_, err = APIClient.CompanyValuation.SECFilings(objects.RequestSECFilings{
			Symbol: symbol,
			Limit:  testCaseLimit,
		})
		if err != nil {
			t.Fatal(err.Error())
		}
	}
}

func TestEnterpriseValue(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	for _, symbol := range testCaseSymbolList {
		_, err = APIClient.CompanyValuation.EnterpriseValue(objects.RequestEnterpriseValue{
			Symbol: symbol,
			Period: objects.CompanyValuationPeriodAnnual,
			Limit:  testCaseLimit,
		})
		if err != nil {
			t.Fatal(err.Error())
		}
	}
}

func TestFinancialStatementsGrowth(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	for _, symbol := range testCaseSymbolList {
		_, err = APIClient.CompanyValuation.FinancialStatementsGrowth(objects.RequestFinancialStatementsGrowth{
			Symbol: symbol,
			Period: objects.CompanyValuationPeriodQuarter,
			Limit:  testCaseLimit,
		})
		if err != nil {
			t.Fatal(err.Error())
		}
	}
}

func TestDiscountedCashFlow(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	for _, symbol := range testCaseSymbolList {
		_, err = APIClient.CompanyValuation.DiscountedCashFlow(symbol)
		if err != nil {
			t.Fatal(err.Error())
		}
	}
}

func TestDailyDiscountedCashFlow(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	for _, symbol := range testCaseSymbolList {
		_, err = APIClient.CompanyValuation.DailyDiscountedCashFlow(objects.RequestDailyDiscountedCashFlow{
			Symbol: symbol,
			Limit:  testCaseLimit,
		})
		if err != nil {
			t.Fatal(err.Error())
		}
	}
}

func TestHistoryDiscountedCashFlow(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	for _, symbol := range testCaseSymbolList {
		_, err = APIClient.CompanyValuation.HistoryDiscountedCashFlow(objects.RequestHistoryDiscountedCashFlow{
			Symbol: symbol,
			Period: objects.CompanyValuationPeriodAnnual,
			Limit:  testCaseLimit,
		})
		if err != nil {
			t.Fatal(err.Error())
		}
	}
}

func TestRating(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	for _, symbol := range testCaseSymbolList {
		_, err = APIClient.CompanyValuation.Rating(symbol)
		if err != nil {
			t.Fatal(err.Error())
		}
	}
}

func TestDailyHistoryRating(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	for _, symbol := range testCaseSymbolList {
		_, err = APIClient.CompanyValuation.DailyHistoryRating(objects.RequestRating{
			Symbol: symbol,
			Limit:  testCaseLimit,
		})
		if err != nil {
			t.Fatal(err.Error())
		}
	}
}

func TestMarketCapitalization(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	for _, symbol := range testCaseSymbolList {
		_, err = APIClient.CompanyValuation.MarketCapitalization(symbol)
		if err != nil {
			t.Fatal(err.Error())
		}
	}
}

func TestDailyHistoryMarketCapitalization(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	for _, symbol := range testCaseSymbolList {
		_, err = APIClient.CompanyValuation.DailyHistoryMarketCapitalization(objects.RequestMarketCapitalization{
			Symbol: symbol,
			Limit:  testCaseLimit,
		})
		if err != nil {
			t.Fatal(err.Error())
		}
	}
}

func TestAnalystEstimates(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	for _, symbol := range testCaseSymbolList {
		_, err = APIClient.CompanyValuation.AnalystEstimates(objects.RequestAnalystEstimates{
			Symbol: symbol,
		})
		if err != nil {
			t.Fatal(err.Error())
		}
	}
}

func TestGrade(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	for _, symbol := range testCaseSymbolList {
		_, err = APIClient.CompanyValuation.Grade(objects.RequestGrade{
			Symbol: symbol,
			Limit:  testCaseLimit,
		})
		if err != nil {
			t.Fatal(err.Error())
		}
	}
}

func TestAnalystStockRecommendations(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	for _, symbol := range testCaseSymbolList {
		_, err = APIClient.CompanyValuation.AnalystStockRecommendations(objects.RequestAnalystStockRecommendations{
			Symbol: symbol,
			Limit:  testCaseLimit,
		})
		if err != nil {
			t.Fatal(err.Error())
		}
	}
}

func TestPressReleases(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	for _, symbol := range testCaseSymbolList {
		_, err = APIClient.CompanyValuation.PressReleases(objects.RequestPressReleases{
			Symbol: symbol,
			Limit:  testCaseLimit,
		})
		if err != nil {
			t.Fatal(err.Error())
		}
	}
}

func TestStockScreener(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	sec := objects.StockScreenerSectorServices
	_, err = APIClient.CompanyValuation.StockScreener(objects.RequestStockScreener{
		Exchange: []string{string(objects.StockScreenerExchangeNasdaq)},
		Sector:   &sec,
		Limit:    testCaseLimit,
	})
	if err != nil {
		t.Fatal(err.Error())
	}
}
