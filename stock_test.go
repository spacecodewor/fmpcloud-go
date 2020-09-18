package fmpcloud

import (
	"testing"
	"time"

	"github.com/spacecodewor/fmpcloud-go/objects"
)

func TestStockQuoteShort(t *testing.T) {
	APIClient, err := NewAPIClient(Config{Debug: true})
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.Stock.QuoteShort("AAPL")
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestStockQuote(t *testing.T) {
	APIClient, err := NewAPIClient(Config{Debug: true})
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.Stock.Quote("AAPL")
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestStockQuoteByExchange(t *testing.T) {
	APIClient, err := NewAPIClient(Config{Debug: true})
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.Stock.QuoteByExchange(objects.StockSearchNasdaq)
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestStockBatchQuote(t *testing.T) {
	APIClient, err := NewAPIClient(Config{Debug: true})
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.Stock.BatchQuote([]string{"AAPL", "ADBE", "FB"})
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestStockSearch(t *testing.T) {
	APIClient, err := NewAPIClient(Config{Debug: true})
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.Stock.Search(objects.RequestStockSearch{
		Query:    "AAPL",
		Limit:    100,
		Exchange: objects.StockSearchCrypto,
	})
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestStockCompanyProfile(t *testing.T) {
	APIClient, err := NewAPIClient(Config{Debug: true})
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.Stock.CompanyProfile("NVDA")
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestStockCompanyExecutive(t *testing.T) {
	APIClient, err := NewAPIClient(Config{Debug: true})
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.Stock.CompanyExecutive("AAPL")
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestStockCandles(t *testing.T) {
	APIClient, err := NewAPIClient(Config{Debug: true})
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.Stock.Candles(objects.RequestStockCandleList{
		Symbol: "AAPL",
		Period: objects.StockCnadlePeriod1Min,
	})
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestStockDailyLine(t *testing.T) {
	APIClient, err := NewAPIClient(Config{Debug: true})
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.Stock.DailyLine("AAPL", objects.StockSerieTypeLine)
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestStockDailyChangeAndVolume(t *testing.T) {
	APIClient, err := NewAPIClient(Config{Debug: true})
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.Stock.DailyChangeAndVolume("AAPL")
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestStockDailySpecificPeriod(t *testing.T) {
	APIClient, err := NewAPIClient(Config{Debug: true})
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.Stock.DailySpecificPeriod("AAPL", time.Now().AddDate(0, 0, -10), time.Now())
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestStockDailyLastNDays(t *testing.T) {
	APIClient, err := NewAPIClient(Config{Debug: true})
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.Stock.DailyLastNDays("AAPL", 10)
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestStockDividends(t *testing.T) {
	APIClient, err := NewAPIClient(Config{Debug: true})
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.Stock.Dividends("AAPL")
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestStockSplits(t *testing.T) {
	APIClient, err := NewAPIClient(Config{Debug: true})
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.Stock.Splits("AAPL")
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestStockSymbolList(t *testing.T) {
	APIClient, err := NewAPIClient(Config{Debug: true})
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.Stock.SymbolList(objects.StockSymbolAll)
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestStockSP500CompanyList(t *testing.T) {
	APIClient, err := NewAPIClient(Config{Debug: true})
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.Stock.SP500CompanyList()
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestStockHistorySP500CompanyList(t *testing.T) {
	APIClient, err := NewAPIClient(Config{Debug: true})
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.Stock.HistorySP500CompanyList()
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestStockEODCandleList(t *testing.T) {
	APIClient, err := NewAPIClient(Config{Debug: true})
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.Stock.EODCandleList(time.Now().AddDate(0, 0, -15))
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestStockBatchEODCandleList(t *testing.T) {
	APIClient, err := NewAPIClient(Config{Debug: true})
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.Stock.BatchEODCandleList([]string{"AAPL"}, time.Now().AddDate(0, 0, -15))
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestStockExchangeTradingHours(t *testing.T) {
	APIClient, err := NewAPIClient(Config{Debug: true})
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.Stock.ExchangeTradingHours()
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestStockActives(t *testing.T) {
	APIClient, err := NewAPIClient(Config{Debug: true})
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.Stock.Actives()
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestStockLosers(t *testing.T) {
	APIClient, err := NewAPIClient(Config{Debug: true})
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.Stock.Losers()
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestStockGainers(t *testing.T) {
	APIClient, err := NewAPIClient(Config{Debug: true})
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.Stock.Gainers()
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestStockSectorPerformance(t *testing.T) {
	APIClient, err := NewAPIClient(Config{Debug: true})
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.Stock.SectorPerformance()
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestStockHistorySectorPerformance(t *testing.T) {
	APIClient, err := NewAPIClient(Config{Debug: true})
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.Stock.HistorySectorPerformance()
	if err != nil {
		t.Fatal(err.Error())
	}
}
