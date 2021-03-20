package fmpcloud

import (
	"testing"
	"time"
)

func TestCOTSymbolList(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfigV4)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.AlternativeData.COTSymbolList()
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestCOTReportListBySymbol(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfigV4)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.AlternativeData.COTReportListBySymbol("ES")
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestCOTReportListByPeriod(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfigV4)
	if err != nil {
		t.Fatal(err.Error())
	}

	from := time.Now().AddDate(0, 0, -10)
	to := time.Now()
	_, err = APIClient.AlternativeData.COTReportListByPeriod(&from, &to)
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestCOTAnalysisListBySymbol(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfigV4)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.AlternativeData.COTAnalysisListBySymbol("ES")
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestCOTAnalysisListByPeriod(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfigV4)
	if err != nil {
		t.Fatal(err.Error())
	}

	from := time.Now().AddDate(0, 0, -10)
	to := time.Now()
	_, err = APIClient.AlternativeData.COTAnalysisListByPeriod(&from, &to)
	if err != nil {
		t.Fatal(err.Error())
	}
}
