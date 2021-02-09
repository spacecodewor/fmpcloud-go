package fmpcloud

import (
	"testing"

	"github.com/spacecodewor/fmpcloud-go/objects"
)

func TestInsiderTradingList(t *testing.T) {
	testCaseAPIConfig.Version = "v4"
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.InsiderTrading.List(objects.RequestInsiderTrading{
		Symbol: "AAPL",
		Limit:  25,
	})
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestInsiderTradingRSSFeed(t *testing.T) {
	testCaseAPIConfig.Version = "v4"
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.InsiderTrading.RSSFeed(25)
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestInsiderTradingMapperCikCompany(t *testing.T) {
	testCaseAPIConfig.Version = "v4"
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.InsiderTrading.MapperCikCompany("AAPL")
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestInsiderTradingMapperCikName(t *testing.T) {
	testCaseAPIConfig.Version = "v4"
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.InsiderTrading.MapperCikName(nil)
	if err != nil {
		t.Fatal(err.Error())
	}
}
