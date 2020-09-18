package fmpcloud

import (
	"testing"
	"time"

	"github.com/spacecodewor/fmpcloud-go/objects"
)

func TestCryptoAvalibleSymbols(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.Crypto.AvalibleSymbols()
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestCryptoQuotes(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.Crypto.Quotes()
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestCryptoCandles(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.Crypto.Candles(
		objects.RequestCryptoCandleList{
			Period: objects.CryptoCandlePeriod1Min,
			Symbol: "BTCUSD",
		})

	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestCryptoDailyLine(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.Crypto.DailyLine("BTCUSD", objects.CryptoSerieTypeLine)
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestCryptoDailyChangeAndVolume(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.Crypto.DailyChangeAndVolume("BTCUSD")
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestCryptoDailySpecificPeriod(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.Crypto.DailySpecificPeriod("BTCUSD", time.Now().AddDate(0, 0, -10), time.Now())
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestCryptoDailyLastNDays(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.Crypto.DailyLastNDays("BTCUSD", 30)

	if err != nil {
		t.Fatal(err.Error())
	}
}
