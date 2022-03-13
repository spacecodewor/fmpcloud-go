package fmpcloud

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/spacecodewor/fmpcloud-go/objects"
)

// Url const for request
const (
	urlAPIAlternativeDataCommitmentOfTradersReportPeriod         = "/v4/commitment_of_traders_report"
	urlAPIAlternativeDataCommitmentOfTradersReportSymbol         = "/v4/commitment_of_traders_report/%s"
	urlAPIAlternativeDataCommitmentOfTradersReportPeriodAnalysis = "/v4/commitment_of_traders_report_analysis"
	urlAPIAlternativeDataCommitmentOfTradersReportSymbolAnalysis = "/v4/commitment_of_traders_report_analysis/%s"
	urlAPIAlternativeDataCommitmentOfTradersReportList           = "/v4/commitment_of_traders_report/list"
)

// AlternativeData client
type AlternativeData struct {
	Client *HTTPClient
}

// COTSymbolList - COT Trading Symbols List
func (a *AlternativeData) COTSymbolList() (sList []objects.COTSymbol, err error) {
	data, err := a.Client.Get(urlAPIAlternativeDataCommitmentOfTradersReportList, nil)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data.Body(), &sList)
	if err != nil {
		return nil, err
	}

	return sList, nil
}

// COTReportBySymbol - List of reports for specific symbol
func (a *AlternativeData) COTReportListBySymbol(symbol string) (rList []objects.COTReport, err error) {
	data, err := a.Client.Get(fmt.Sprintf(urlAPIAlternativeDataCommitmentOfTradersReportSymbol, symbol), nil)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data.Body(), &rList)
	if err != nil {
		return nil, err
	}

	return rList, nil
}

// COTReportListByPeriod - List of reports for period of time
func (a *AlternativeData) COTReportListByPeriod(from, to *time.Time) (rList []objects.COTReport, err error) {
	reqParam := make(map[string]string)
	if from != nil {
		reqParam["from"] = from.Format("2006-01-02")
	}

	if to != nil {
		reqParam["to"] = to.Format("2006-01-02")
	}

	data, err := a.Client.Get(urlAPIAlternativeDataCommitmentOfTradersReportPeriod, reqParam)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data.Body(), &rList)
	if err != nil {
		return nil, err
	}

	return rList, nil
}

// COTAnalysisListBySymbol - Analysis of reports for trading symbol
func (a *AlternativeData) COTAnalysisListBySymbol(symbol string) (aList []objects.COTAnalysis, err error) {
	data, err := a.Client.Get(fmt.Sprintf(urlAPIAlternativeDataCommitmentOfTradersReportSymbolAnalysis, symbol), nil)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data.Body(), &aList)
	if err != nil {
		return nil, err
	}

	return aList, nil
}

// COTAnalysisListByPeriod - Analysis of reports for time period
func (a *AlternativeData) COTAnalysisListByPeriod(from, to *time.Time) (aList []objects.COTAnalysis, err error) {
	reqParam := make(map[string]string)
	if from != nil {
		reqParam["from"] = from.Format("2006-01-02")
	}

	if to != nil {
		reqParam["to"] = to.Format("2006-01-02")
	}

	data, err := a.Client.Get(urlAPIAlternativeDataCommitmentOfTradersReportPeriodAnalysis, reqParam)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data.Body(), &aList)
	if err != nil {
		return nil, err
	}

	return aList, nil
}
