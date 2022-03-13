package fmpcloud

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/spacecodewor/fmpcloud-go/objects"
)

// Url const for request
const (
	urlAPIInsiderTrading                 = "/v4/insider-trading"
	urlAPIInsiderTradingTransactionType  = "/v4/insider-trading-transaction-type"
	urlAPIInsiderTradingRSSFeed          = "/v4/insider-trading-rss-feed"
	urlAPIInsiderTradingMapperCikName    = "/v4/mapper-cik-name"
	urlAPIInsiderTradingMapperCikCompany = "/v4/mapper-cik-company/%s"
)

// InsiderTrading client
type InsiderTrading struct {
	Client *HTTPClient
}

// The federal securities laws require certain individuals
// (such as officers, directors, and those that hold more than 10% of any class of a company’s securities, together we’ll call, “insiders”)
// to report purchases, sales, and holdings of their company’s securities by filing Forms 3, 4, and 5.

// List - Stock insider trading list
func (i *InsiderTrading) List(req objects.RequestInsiderTrading) (iList []objects.InsiderTrading, err error) {
	if req.Limit == 0 {
		req.Limit = 100
	}

	reqData := map[string]string{"limit": fmt.Sprint(req.Limit)}
	if len(req.Symbol) != 0 {
		reqData["symbol"] = req.Symbol
	}

	if len(req.CompanyCik) != 0 {
		reqData["companyCik"] = req.CompanyCik
	}

	if len(req.ReportingCik) != 0 {
		reqData["reportingCik"] = req.ReportingCik
	}

	if len(req.TransactionType) != 0 {
		reqData["transactionType"] = strings.Join(req.TransactionType, ",")
	}

	data, err := i.Client.Get(urlAPIInsiderTrading, reqData)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data.Body(), &iList)
	if err != nil {
		return nil, err
	}

	return iList, nil
}

// RSSFeed - RSS Feed of form 3,4 and 5
func (i *InsiderTrading) RSSFeed(limit int64) (iList []objects.InsiderTradingRSSFeed, err error) {
	if limit == 0 {
		limit = 100
	}

	data, err := i.Client.Get(urlAPIInsiderTradingRSSFeed, map[string]string{"limit": fmt.Sprint(limit)})
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data.Body(), &iList)
	if err != nil {
		return nil, err
	}

	return iList, nil
}

// TransactionType - list
func (i *InsiderTrading) TransactionType() (tList []string, err error) {
	data, err := i.Client.Get(urlAPIInsiderTradingTransactionType, nil)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data.Body(), &tList)
	if err != nil {
		return nil, err
	}

	return tList, nil
}

// MapperCikCompany - Company CIK mapper
func (i *InsiderTrading) MapperCikCompany(symbol string) (iList []objects.InsiderTradingMapperCikCompany, err error) {
	data, err := i.Client.Get(fmt.Sprintf(urlAPIInsiderTradingMapperCikCompany, symbol), nil)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data.Body(), &iList)
	if err != nil {
		return nil, err
	}

	return iList, nil
}

// MapperCikName - List with names and their CIK
func (i *InsiderTrading) MapperCikName(name *string) (iList []objects.InsiderTradingMapperCikName, err error) {
	reqData := map[string]string{}
	if name != nil {
		reqData["name"] = *name
	}

	data, err := i.Client.Get(urlAPIInsiderTradingMapperCikName, reqData)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data.Body(), &iList)
	if err != nil {
		return nil, err
	}

	return iList, nil
}
