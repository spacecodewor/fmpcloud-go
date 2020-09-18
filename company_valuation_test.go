package fmpcloud

import (
	"testing"

	"github.com/spacecodewor/fmpcloud-go/objects"
)

// Symbol list for check request
var checkSymbolList = []string{"AAPL", "GM", "NVDA", "TSLA", "ADBE", "JPM", "BAC"}

func TestRssFeed(t *testing.T) {
	APIClient, err := NewAPIClient(Config{Debug: true})
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.CompanyValuation.RssFeed()
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestEarningCalendar(t *testing.T) {
	APIClient, err := NewAPIClient(Config{Debug: true})
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.CompanyValuation.EarningCalendar(nil, nil)
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestHistoryEarningCalendar(t *testing.T) {
	APIClient, err := NewAPIClient(Config{Debug: true})
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.CompanyValuation.HistoryEarningCalendar("AAPL")
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestIPOCalendar(t *testing.T) {
	APIClient, err := NewAPIClient(Config{Debug: true})
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.CompanyValuation.IPOCalendar(nil, nil)
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestSplitCalendar(t *testing.T) {
	APIClient, err := NewAPIClient(Config{Debug: true})
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.CompanyValuation.SplitCalendar(nil, nil)
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestInstitutionalHolders(t *testing.T) {
	APIClient, err := NewAPIClient(Config{Debug: true})
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.CompanyValuation.InstitutionalHolders("AAPL")
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestMutualFundHolders(t *testing.T) {
	APIClient, err := NewAPIClient(Config{Debug: true})
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.CompanyValuation.MutualFundHolders("AAPL")
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestETFHolders(t *testing.T) {
	APIClient, err := NewAPIClient(Config{Debug: true})
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.CompanyValuation.ETFHolders("QQQ")
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestETFSectorWeightings(t *testing.T) {
	APIClient, err := NewAPIClient(Config{Debug: true})
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.CompanyValuation.ETFSectorWeightings("QQQ")
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestETFCountryWeightings(t *testing.T) {
	APIClient, err := NewAPIClient(Config{Debug: true})
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.CompanyValuation.ETFCountryWeightings("QQQ")
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestIncomeStatement(t *testing.T) {
	APIClient, err := NewAPIClient(Config{Debug: true})
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.CompanyValuation.IncomeStatement(objects.RequestIncomeStatement{
		Symbol: "AAPL",
		Period: objects.CompanyValuationPeriodAnnual,
		Limit:  50,
	})
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestIncomeStatementGrowth(t *testing.T) {
	APIClient, err := NewAPIClient(Config{Debug: true})
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.CompanyValuation.IncomeStatementGrowth(objects.RequestIncomeStatementGrowth{
		Symbol: "AAPL",
		Period: objects.CompanyValuationPeriodAnnual,
		Limit:  50,
	})
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestBalanceSheetStatement(t *testing.T) {
	APIClient, err := NewAPIClient(Config{Debug: true})
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.CompanyValuation.BalanceSheetStatement(objects.RequestBalanceSheetStatement{
		Symbol: "AAPL",
		Period: objects.CompanyValuationPeriodAnnual,
		Limit:  50,
	})
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestBalanceSheetStatementGrowth(t *testing.T) {
	APIClient, err := NewAPIClient(Config{Debug: true})
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.CompanyValuation.BalanceSheetStatementGrowth(objects.RequestBalanceSheetStatementGrowth{
		Symbol: "AAPL",
		Period: objects.CompanyValuationPeriodAnnual,
		Limit:  50,
	})
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestCashFlowStatement(t *testing.T) {
	APIClient, err := NewAPIClient(Config{Debug: true})
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.CompanyValuation.CashFlowStatement(objects.RequestCashFlowStatement{
		Symbol: "AAPL",
		Period: objects.CompanyValuationPeriodAnnual,
		Limit:  50,
	})
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestCashFlowStatementGrowth(t *testing.T) {
	APIClient, err := NewAPIClient(Config{Debug: true})
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.CompanyValuation.CashFlowStatementGrowth(objects.RequestCashFlowStatementGrowth{
		Symbol: "ADBE",
		Period: objects.CompanyValuationPeriodAnnual,
		Limit:  50,
	})
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestIncomeStatementAsReported(t *testing.T) {
	APIClient, err := NewAPIClient(Config{Debug: true})
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.CompanyValuation.IncomeStatementAsReported(objects.RequestIncomeStatementAsReported{
		Symbol: "GM",
		Period: objects.CompanyValuationPeriodAnnual,
		Limit:  3,
	})
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestBalanceSheetStatementAsReported(t *testing.T) {
	APIClient, err := NewAPIClient(Config{Debug: true})
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.CompanyValuation.BalanceSheetStatementAsReported(objects.RequestBalanceSheetStatementAsReported{
		Symbol: "TSLA",
		Period: objects.CompanyValuationPeriodAnnual,
		Limit:  3,
	})
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestCashFlowStatementAsReported(t *testing.T) {
	APIClient, err := NewAPIClient(Config{Debug: true})
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.CompanyValuation.CashFlowStatementAsReported(objects.RequestCashFlowStatementAsReported{
		Symbol: "NVDA",
		Period: objects.CompanyValuationPeriodAnnual,
		Limit:  3,
	})
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestFullFinancialStatementAsReported(t *testing.T) {
	APIClient, err := NewAPIClient(Config{Debug: true})
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.CompanyValuation.FullFinancialStatementAsReported(objects.RequestFullFinancialStatementAsReported{
		Symbol: "AAPL",
		Period: objects.CompanyValuationPeriodAnnual,
		Limit:  3,
	})
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestStockNews(t *testing.T) {
	APIClient, err := NewAPIClient(Config{Debug: true})
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.CompanyValuation.StockNews(objects.RequestStockNews{
		SymbolList: []string{"AAPL", "FB"},
		Limit:      50,
	})
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestDelstedCompanies(t *testing.T) {
	APIClient, err := NewAPIClient(Config{Debug: true})
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.CompanyValuation.DelstedCompanies(50)
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestFinancialRatios(t *testing.T) {
	APIClient, err := NewAPIClient(Config{Debug: true})
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.CompanyValuation.FinancialRatios(objects.RequestFinancialRatios{
		Symbol: "AAPL",
		Period: objects.CompanyValuationPeriodAnnual,
		Limit:  50,
	})
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestFinancialRatiosTTM(t *testing.T) {
	APIClient, err := NewAPIClient(Config{Debug: true})
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.CompanyValuation.FinancialRatiosTTM("AAPL")
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestKeyMetrics(t *testing.T) {
	APIClient, err := NewAPIClient(Config{Debug: true})
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.CompanyValuation.KeyMetrics(objects.RequestKeyMetrics{
		Symbol: "AAPL",
		Period: objects.CompanyValuationPeriodQuarter,
		Limit:  50,
	})
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestEnterpriseValue(t *testing.T) {
	APIClient, err := NewAPIClient(Config{Debug: true})
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.CompanyValuation.EnterpriseValue(objects.RequestEnterpriseValue{
		Symbol: "AAPL",
		Period: objects.CompanyValuationPeriodAnnual,
		Limit:  50,
	})
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestFinancialStatementsGrowth(t *testing.T) {
	APIClient, err := NewAPIClient(Config{Debug: true})
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.CompanyValuation.FinancialStatementsGrowth(objects.RequestFinancialStatementsGrowth{
		Symbol: "AAPL",
		Period: objects.CompanyValuationPeriodQuarter,
		Limit:  50,
	})
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestDiscountedCashFlow(t *testing.T) {
	APIClient, err := NewAPIClient(Config{Debug: true})
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.CompanyValuation.DiscountedCashFlow("AAPL")
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestDailyDiscountedCashFlow(t *testing.T) {
	APIClient, err := NewAPIClient(Config{Debug: true})
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.CompanyValuation.DailyDiscountedCashFlow(objects.RequestDailyDiscountedCashFlow{
		Symbol: "AAPL",
		Limit:  50,
	})
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestHistoryDiscountedCashFlow(t *testing.T) {
	APIClient, err := NewAPIClient(Config{Debug: true})
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.CompanyValuation.HistoryDiscountedCashFlow(objects.RequestHistoryDiscountedCashFlow{
		Symbol: "AAPL",
		Period: objects.CompanyValuationPeriodAnnual,
		Limit:  50,
	})
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestRating(t *testing.T) {
	APIClient, err := NewAPIClient(Config{Debug: true})
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.CompanyValuation.Rating("AAPL")
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestDailyHistoryRating(t *testing.T) {
	APIClient, err := NewAPIClient(Config{Debug: true})
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.CompanyValuation.DailyHistoryRating(objects.RequestRating{
		Symbol: "AAPL",
		Limit:  50,
	})
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestMarketCapitalization(t *testing.T) {
	APIClient, err := NewAPIClient(Config{Debug: true})
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.CompanyValuation.MarketCapitalization("AAPL")
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestDailyHistoryMarketCapitalization(t *testing.T) {
	APIClient, err := NewAPIClient(Config{Debug: true})
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.CompanyValuation.DailyHistoryMarketCapitalization(objects.RequestMarketCapitalization{
		Symbol: "AAPL",
		Limit:  50,
	})
	if err != nil {
		t.Fatal(err.Error())
	}
}
