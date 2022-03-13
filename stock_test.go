package fmpcloud

import (
	"testing"
	"time"

	"github.com/spacecodewor/fmpcloud-go/objects"
)

func TestStockQuoteShort(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	for _, symbol := range testCaseSymbolList {
		_, err = APIClient.Stock.QuoteShort(symbol)
		if err != nil {
			t.Fatal(err.Error())
		}
	}
}

func TestStockQuote(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	for _, symbol := range testCaseSymbolList {
		_, err = APIClient.Stock.Quote(symbol)
		if err != nil {
			t.Fatal(err.Error())
		}
	}
}

func TestStockQuoteByExchange(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.Stock.QuoteByExchange(objects.StockSearchNasdaq)
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestStockBatchQuote(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.Stock.BatchQuote(testCaseSymbolList)
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestStockSearch(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	for _, symbol := range testCaseSymbolList {
		_, err = APIClient.Stock.Search(objects.RequestStockSearch{
			Query: symbol,
			Limit: testCaseLimit,
		})
		if err != nil {
			t.Fatal(err.Error())
		}
	}
}

func TestStockBulkProfile(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	profiles, err := APIClient.Stock.BulkProfile()
	if err != nil {
		t.Fatal(err.Error())
	}

	t.Log(profiles)
}

func TestStockBulkPeers(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	peers, err := APIClient.Stock.BulkPeers()
	if err != nil {
		t.Fatal(err.Error())
	}

	t.Log(peers)
}

func TestStockCompanyProfile(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	for _, symbol := range testCaseSymbolList {
		_, err = APIClient.Stock.CompanyProfile(symbol)
		if err != nil {
			t.Log(err.Error())
			continue
		}
	}
}

func TestStockCompanyExecutive(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	for _, symbol := range testCaseSymbolList {
		_, err = APIClient.Stock.CompanyExecutive(symbol)
		if err != nil {
			t.Fatal(err.Error())
		}
	}
}

func TestStockCandles(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	for _, symbol := range testCaseSymbolList {
		_, err = APIClient.Stock.Candles(objects.RequestStockCandleList{
			Symbol: symbol,
			Period: objects.StockCandlePeriod1Min,
		})
		if err != nil {
			t.Fatal(err.Error())
		}
	}
}

func TestStockDailyLine(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	for _, symbol := range testCaseSymbolList {
		_, err = APIClient.Stock.DailyLine(symbol, objects.StockSerieTypeLine)
		if err != nil {
			t.Fatal(err.Error())
		}
	}
}

func TestDailyBatch(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	from := time.Now().AddDate(0, 0, -10)
	to := time.Now()
	_, err = APIClient.Stock.DailyBatch(testCaseSingleSymbol, &from, &to)
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestStockDailyChangeAndVolume(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	for _, symbol := range testCaseSymbolList {
		_, err = APIClient.Stock.DailyChangeAndVolume(symbol)
		if err != nil {
			t.Fatal(err.Error())
		}
	}
}

func TestStockDailySpecificPeriod(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	for _, symbol := range testCaseSymbolList {
		_, err = APIClient.Stock.DailySpecificPeriod(symbol, time.Now().AddDate(0, 0, -10), time.Now())
		if err != nil {
			t.Fatal(err.Error())
		}
	}
}

func TestStockDailyLastNDays(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	for _, symbol := range testCaseSymbolList {
		_, err = APIClient.Stock.DailyLastNDays(symbol, 10)
		if err != nil {
			t.Fatal(err.Error())
		}
	}
}

func TestStockDividends(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	for _, symbol := range testCaseSymbolList {
		_, err = APIClient.Stock.Dividends(symbol)
		if err != nil {
			t.Fatal(err.Error())
		}
	}
}

func TestStockSplits(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	for _, symbol := range testCaseSymbolList {
		_, err = APIClient.Stock.Splits(symbol)
		if err != nil {
			t.Fatal(err.Error())
		}
	}
}

func TestStockSymbolList(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.Stock.AvalibleSymbols()
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestAvalibleSymbolsByExchange(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.Stock.AvalibleSymbolsByExchange(objects.StockSymbolNasdaq)
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestIndexConstituentList(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.Stock.IndexConstituentList(objects.IndexSP500)
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestHistoryIndexConstituentList(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.Stock.HistoryIndexConstituentList(objects.IndexSP500)
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestStockEODCandleList(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.Stock.EODCandleList(time.Now().AddDate(0, 0, -15))
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestStockBatchEODCandleList(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.Stock.BatchEODCandleList(testCaseSymbolList, time.Now().AddDate(0, 0, -15))
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestStockExchangeTradingHours(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.Stock.ExchangeTradingHours()
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestStockActives(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.Stock.Actives()
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestStockLosers(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.Stock.Losers()
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestStockGainers(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.Stock.Gainers()
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestStockSectorPerformance(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.Stock.SectorPerformance()
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestStockHistorySectorPerformance(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.Stock.HistorySectorPerformance()
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestSurvivorshipBiasFree(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.Stock.SurvivorshipBiasFree("ABI", time.Date(2005, 01, 03, 0, 0, 0, 0, time.UTC))
	if err != nil {
		t.Fatal(err.Error())
	}
}
