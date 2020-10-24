package fmpcloud

import (
	"testing"
	"time"

	"github.com/spacecodewor/fmpcloud-go/objects"
)

func TestForexAvalibleSymbols(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.Forex.AvalibleSymbols()
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestForexQuotes(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.Forex.Quotes()
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestForexListSymbolsAndQuotes(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.Forex.ListSymbolsAndQuotes()
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestForexCandles(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.Forex.Candles(
		objects.RequestForexCandleList{
			Period: objects.ForexCandlePeriod1Min,
			Symbol: "JPYUSD",
		})

	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestForexDailyLine(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.Forex.DailyLine("JPYUSD", objects.ForexSerieTypeLine)
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestForexDailyChangeAndVolume(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.Forex.DailyChangeAndVolume("JPYUSD")
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestForexDailySpecificPeriod(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.Forex.DailySpecificPeriod("JPYUSD", time.Now().AddDate(0, 0, -10), time.Now())
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestForexDailyLastNDays(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.Forex.DailyLastNDays("JPYUSD", 30)

	if err != nil {
		t.Fatal(err.Error())
	}
}
