package fmpcloud

import (
	"fmt"
	"strings"
	"time"

	"github.com/gocarina/gocsv"
	jsoniter "github.com/json-iterator/go"
	"github.com/spacecodewor/fmpcloud-go/objects"
)

// Url const for request
const (
	urlAPICompanyValuationRSSFeed                          = "/v3/rss_feed"
	urlAPICompanyValuationEarningCalendar                  = "/v3/earning_calendar"
	urlAPICompanyValuationEarningCalendarConfirmed         = "/v4/earning-calendar-confirmed"
	urlAPICompanyValuationEarningsSurpises                 = "/v3/earnings-surpises/%s"
	urlAPICompanyValuationEarningCallTranscript            = "/v3/earning_call_transcript/%s"
	urlAPICompanyValuationHistoryEarningCalendar           = "/v3/historical/earning_calendar/%s"
	urlAPICompanyValuationIPOCalendar                      = "/v3/ipo_calendar"
	urlAPICompanyValuationIPOCalendarConfirmed             = "/v4/ipo-calendar-confirmed"
	urlAPICompanyValuationIPOCalendarProspectus            = "/v4/ipo-calendar-prospectus"
	urlAPICompanyValuationSplitCalendar                    = "/v3/stock_split_calendar"
	urlAPICompanyValuationDividendCalendar                 = "/v3/stock_dividend_calendar"
	urlAPICompanyValuationInstituionalHolder               = "/v3/institutional-holder/%s"
	urlAPICompanyValuationMutualFundHolder                 = "/v3/mutual-fund-holder/%s"
	urlAPICompanyValuationETFHolder                        = "/v3/etf-holder/%s"
	urlAPICompanyValuationETFStockExposure                 = "/v3/etf-stock-exposure/%s"
	urlAPICompanyValuationETFSectorWeightings              = "/v3/etf-sector-weightings/%s"
	urlAPICompanyValuationETFCountryWeightings             = "/v3/etf-country-weightings/%s"
	urlAPICompanyValuationIncomeStatement                  = "/v3/income-statement/%s"
	urlAPICompanyValuationIncomeStatementGrowth            = "/v3/income-statement-growth/%s"
	urlAPICompanyValuationBulkIncomeStatement              = "/v4/income-statement-bulk"
	urlAPICompanyValuationBulkBalanceSheetStatement        = "/v4/balance-sheet-statement-bulk"
	urlAPICompanyValuationBulkCashFlowStatement            = "/v4/cash-flow-statement-bulk"
	urlAPICompanyValuationBulkRatios                       = "/v4/ratios-bulk"
	urlAPICompanyValuationBulkKeyMetrics                   = "/v4/key-metrics-bulk"
	urlAPICompanyValuationBulkEarningsSurpises             = "/v4/earnings-surprises-bulk"
	urlAPICompanyValuationBulkRating                       = "/v4/rating-bulk"
	urlAPICompanyValuationBulkScores                       = "/v4/scores-bulk"
	urlAPICompanyValuationBalanceSheetStatement            = "/v3/balance-sheet-statement/%s"
	urlAPICompanyValuationBalanceSheetStatementGrowth      = "/v3/balance-sheet-statement-growth/%s"
	urlAPICompanyValuationCashFlowStatement                = "/v3/cash-flow-statement/%s"
	urlAPICompanyValuationCashFlowStatementGrowth          = "/v3/cash-flow-statement-growth/%s"
	urlAPICompanyValuationIncomeStatementAsReported        = "/v3/income-statement-as-reported/%s"
	urlAPICompanyValuationBalanceSheetStatementAsReported  = "/v3/balance-sheet-statement-as-reported/%s"
	urlAPICompanyValuationCashFlowStatementAsReported      = "/v3/cash-flow-statement-as-reported/%s"
	urlAPICompanyValuationFinancialStatementFullAsReported = "/v3/financial-statement-full-as-reported/%s"
	urlAPICompanyValuationFinancialRatios                  = "/v3/ratios/%s"
	urlAPICompanyValuationFinancialRatiosTTM               = "/v3/ratios-ttm/%s"
	urlAPICompanyValuationKeyMetrics                       = "/v3/key-metrics/%s"
	urlAPICompanyValuationKeyMetricsTTM                    = "/v3/key-metrics-ttm/%s"
	urlAPICompanyValuationEnterpriseValues                 = "/v3/enterprise-values/%s"
	urlAPICompanyValuationFinancialGrowth                  = "/v3/financial-growth/%s"
	urlAPICompanyValuationDiscountedCashFlow               = "/v3/discounted-cash-flow/%s"
	urlAPICompanyValuationHistoryDailyDiscountedCashFlow   = "/v3/historical-daily-discounted-cash-flow/%s"
	urlAPICompanyValuationHistoryDiscountedCashFlow        = "/v3/historical-discounted-cash-flow-statement/%s"
	urlAPICompanyValuationRating                           = "/v3/rating/%s"
	urlAPICompanyValuationHistoryRating                    = "/v3/historical-rating/%s"
	urlAPICompanyValuationMarketCapitalization             = "/v3/market-capitalization/%s"
	urlAPICompanyValuationHistoryMarketCapitalization      = "/v3/historical-market-capitalization/%s"
	urlAPICompanyValuationDelistedCompanyList              = "/v3/delisted-companies"
	urlAPICompanyValuationStockNews                        = "/v3/stock_news"
	urlAPICompanyValuationStockScreener                    = "/v3/stock-screener"
	urlAPICompanyValuationAnalystEstimates                 = "/v3/analyst-estimates/%s"
	urlAPICompanyValuationAnalystStockRecommendations      = "/v3/analyst-stock-recommendations/%s"
	urlAPICompanyValuationGrade                            = "/v3/grade/%s"
	urlAPICompanyValuationPressReleases                    = "/v3/press-releases/%s"
	urlAPICompanyValuationFinancialStatementsList          = "/v3/financial-statement-symbol-lists"
	urlAPICompanyValuationEconomicCalendarEventList        = "/v3/economic_calendar_event_list"
	urlAPICompanyValuationEconomicCalendar                 = "/v3/economic_calendar"
	urlAPICompanyValuationHistoryEconomicCalendar          = "/v3/historical/economic_calendar/%s"
	urlAPICompanyValuationSECFillings                      = "/v3/sec_filings/%s"
	urlAPICompanyValuationETFList                          = "/v3/etf/list"
	urlAPICompanyValuationAvailableTradedList              = "/v3/available-traded/list"
	urlAPICompanyValuationCompanyOutlook                   = "/v4/company-outlook"
	urlAPICompanyValuationEmployeeCount                    = "/v4/employee_count"
	urlAPICompanyValuationSocialSentimentTrending          = "/v4/social-sentiment/trending"
	urlAPICompanyValuationSocialSentimentChange            = "/v4/social-sentiments/change"
	urlAPICompanyValuationHistoricalSocialSentiment        = "/v4/historical/social-sentiment"
	urlAPICompanyValuationScore                            = "/v4/score"
	urlAPICompanyValuationSharesFloat                      = "/v4/shares_float"
	urlAPICompanyValuationSharesFloatAll                   = "/v4/shares_float/all"
	urlAPICompanyValuationRatiosTTMBulk                    = "/v4/ratios-ttm-bulk"
	urlAPICompanyValuationDCFBulk                          = "/v4/dcf-bulk"
)

// CompanyValuation client
type CompanyValuation struct {
	Client *HTTPClient
}

// RssFeed - SEC RSS feeds is a very helpful resource for staying current on the most recent financial statements posted on the SEC
func (c *CompanyValuation) RssFeed() (fList []objects.RssFeed, err error) {
	data, err := c.Client.Get(urlAPICompanyValuationRSSFeed, nil)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &fList)
	if err != nil {
		return nil, err
	}

	return fList, nil
}

// EarningCalendar - earning Calendar (between from and to maximum interval can be 3 months)
func (c *CompanyValuation) EarningCalendar(from, to *time.Time) (eList []objects.EarningCalendar, err error) {
	reqParam := make(map[string]string)
	if from != nil {
		reqParam["from"] = from.Format("2006-01-02")
	}

	if to != nil {
		reqParam["to"] = to.Format("2006-01-02")
	}

	data, err := c.Client.Get(urlAPICompanyValuationEarningCalendar, reqParam)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &eList)
	if err != nil {
		return nil, err
	}

	return eList, nil
}

// EarningCalendarConfirmed - Earnings calendar confirmed (between the "from" and "to" parameters the maximum time interval can be 3 months)
func (c *CompanyValuation) EarningCalendarConfirmed(from, to *time.Time) (eList []objects.EarningCalendarConfirmed, err error) {
	reqParam := make(map[string]string)
	if from != nil {
		reqParam["from"] = from.Format("2006-01-02")
	}

	if to != nil {
		reqParam["to"] = to.Format("2006-01-02")
	}

	data, err := c.Client.Get(urlAPICompanyValuationEarningCalendarConfirmed, reqParam)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &eList)
	if err != nil {
		return nil, err
	}

	return eList, nil
}

// EarningSurpriseList ...
func (c *CompanyValuation) EarningSurpriseList(symbol string) (eList []objects.EarningSurprise, err error) {
	data, err := c.Client.Get(fmt.Sprintf(urlAPICompanyValuationEarningsSurpises, symbol), nil)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &eList)
	if err != nil {
		return nil, err
	}

	return eList, nil
}

// EarningCallTranscript - transcript of specific earning
func (c *CompanyValuation) EarningCallTranscript(req objects.RequestEarningCallTranscript) (tList []objects.EarningCallTranscript, err error) {
	data, err := c.Client.Get(
		fmt.Sprintf(urlAPICompanyValuationEarningCallTranscript, req.Symbol),
		map[string]string{
			"quarter": fmt.Sprint(req.Quarter),
			"year":    fmt.Sprint(req.Year),
		})
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &tList)
	if err != nil {
		return nil, err
	}

	return tList, nil
}

// HistoryEarningCalendar - historical earning calendar
func (c *CompanyValuation) HistoryEarningCalendar(symbol string) (eList []objects.EarningCalendar, err error) {
	data, err := c.Client.Get(fmt.Sprintf(urlAPICompanyValuationHistoryEarningCalendar, symbol), nil)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &eList)
	if err != nil {
		return nil, err
	}

	return eList, nil
}

// IPOCalendar - IPO calendar
func (c *CompanyValuation) IPOCalendar(from, to *time.Time) (ipoList []objects.IPOCalendar, err error) {
	reqParam := make(map[string]string)
	if from != nil {
		reqParam["from"] = from.Format("2006-01-02")
	}

	if to != nil {
		reqParam["to"] = to.Format("2006-01-02")
	}

	data, err := c.Client.Get(urlAPICompanyValuationIPOCalendar, reqParam)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &ipoList)
	if err != nil {
		return nil, err
	}

	return ipoList, nil
}

// IPOCalendarConfirmed - IPO Calendar confirmed (between the "from" and "to" parameters the maximum time interval can be 3 months)
func (c *CompanyValuation) IPOCalendarConfirmed(from, to *time.Time) (ipoList []objects.IPOCalendarConfirmed, err error) {
	reqParam := make(map[string]string)
	if from != nil {
		reqParam["from"] = from.Format("2006-01-02")
	}

	if to != nil {
		reqParam["to"] = to.Format("2006-01-02")
	}

	data, err := c.Client.Get(urlAPICompanyValuationIPOCalendarConfirmed, reqParam)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &ipoList)
	if err != nil {
		return nil, err
	}

	return ipoList, nil
}

// IPOCalendarProspectus - IPO Calendar with prospectus (between the "from" and "to" parameters the maximum time interval can be 3 months)
func (c *CompanyValuation) IPOCalendarProspectus(from, to *time.Time) (ipoList []objects.IPOCalendarProspectus, err error) {
	reqParam := make(map[string]string)
	if from != nil {
		reqParam["from"] = from.Format("2006-01-02")
	}

	if to != nil {
		reqParam["to"] = to.Format("2006-01-02")
	}

	data, err := c.Client.Get(urlAPICompanyValuationIPOCalendarProspectus, reqParam)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &ipoList)
	if err != nil {
		return nil, err
	}

	return ipoList, nil
}

// SplitCalendar - stock split calendar
func (c *CompanyValuation) SplitCalendar(from, to *time.Time) (sList []objects.SplitCalendar, err error) {
	reqParam := make(map[string]string)
	if from != nil {
		reqParam["from"] = from.Format("2006-01-02")
	}

	if to != nil {
		reqParam["to"] = to.Format("2006-01-02")
	}

	data, err := c.Client.Get(urlAPICompanyValuationSplitCalendar, reqParam)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &sList)
	if err != nil {
		return nil, err
	}

	return sList, nil
}

// DividendCalendar - dividend calendar
func (c *CompanyValuation) DividendCalendar(from, to *time.Time) (dList []objects.DividendCalendar, err error) {
	reqParam := make(map[string]string)
	if from != nil {
		reqParam["from"] = from.Format("2006-01-02")
	}

	if to != nil {
		reqParam["to"] = to.Format("2006-01-02")
	}

	data, err := c.Client.Get(urlAPICompanyValuationDividendCalendar, reqParam)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &dList)
	if err != nil {
		return nil, err
	}

	return dList, nil
}

// InstitutionalHolders - institutional holders
func (c *CompanyValuation) InstitutionalHolders(symbol string) (hList []objects.InstitutionalHolder, err error) {
	data, err := c.Client.Get(fmt.Sprintf(urlAPICompanyValuationInstituionalHolder, symbol), nil)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &hList)
	if err != nil {
		return nil, err
	}

	return hList, nil
}

// MutualFundHolders - mutual fund holders
func (c *CompanyValuation) MutualFundHolders(symbol string) (hList []objects.MutualFundHolder, err error) {
	data, err := c.Client.Get(fmt.Sprintf(urlAPICompanyValuationMutualFundHolder, symbol), nil)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &hList)
	if err != nil {
		return nil, err
	}

	return hList, nil
}

// ETFHolders - ETF holders
func (c *CompanyValuation) ETFHolders(symbol string) (hList []objects.ETFHolder, err error) {
	data, err := c.Client.Get(fmt.Sprintf(urlAPICompanyValuationETFHolder, symbol), nil)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &hList)
	if err != nil {
		return nil, err
	}

	return hList, nil
}

// ETFStockExposure - ETF stock exposure
func (c *CompanyValuation) ETFStockExposure(symbol string) (eList []objects.ETFStockExposure, err error) {
	data, err := c.Client.Get(fmt.Sprintf(urlAPICompanyValuationETFStockExposure, symbol), nil)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &eList)
	if err != nil {
		return nil, err
	}

	return eList, nil
}

// ETFSectorWeightings - ETF sector weightings
func (c *CompanyValuation) ETFSectorWeightings(symbol string) (sList []objects.ETFSectorWeighting, err error) {
	data, err := c.Client.Get(fmt.Sprintf(urlAPICompanyValuationETFSectorWeightings, symbol), nil)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &sList)
	if err != nil {
		return nil, err
	}

	return sList, nil
}

// ETFCountryWeightings - ETF country weightings
func (c *CompanyValuation) ETFCountryWeightings(symbol string) (cList []objects.ETFCountryWeighting, err error) {
	data, err := c.Client.Get(fmt.Sprintf(urlAPICompanyValuationETFCountryWeightings, symbol), nil)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &cList)
	if err != nil {
		return nil, err
	}

	return cList, nil
}

// IncomeStatement - income statement
func (c *CompanyValuation) IncomeStatement(req objects.RequestIncomeStatement) (sList []objects.IncomeStatement, err error) {
	reqParam := map[string]string{"limit": fmt.Sprint(req.Limit)}
	if req.Period != objects.CompanyValuationPeriodAnnual {
		reqParam["period"] = string(objects.CompanyValuationPeriodQuarter)
	}

	data, err := c.Client.Get(fmt.Sprintf(urlAPICompanyValuationIncomeStatement, req.Symbol), reqParam)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &sList)
	if err != nil {
		return nil, err
	}

	return sList, nil
}

// IncomeStatementGrowth - income statement growth
func (c *CompanyValuation) IncomeStatementGrowth(req objects.RequestIncomeStatementGrowth) (sList []objects.IncomeStatementGrowth, err error) {
	reqParam := map[string]string{"limit": fmt.Sprint(req.Limit)}
	if req.Period != objects.CompanyValuationPeriodAnnual {
		reqParam["period"] = string(objects.CompanyValuationPeriodQuarter)
	}

	data, err := c.Client.Get(fmt.Sprintf(urlAPICompanyValuationIncomeStatementGrowth, req.Symbol), reqParam)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &sList)
	if err != nil {
		return nil, err
	}

	return sList, nil
}

// BalanceSheetStatement - balance sheet statement
func (c *CompanyValuation) BalanceSheetStatement(req objects.RequestBalanceSheetStatement) (sList []objects.BalanceSheetStatement, err error) {
	reqParam := map[string]string{"limit": fmt.Sprint(req.Limit)}
	if req.Period != objects.CompanyValuationPeriodAnnual {
		reqParam["period"] = string(objects.CompanyValuationPeriodQuarter)
	}

	data, err := c.Client.Get(fmt.Sprintf(urlAPICompanyValuationBalanceSheetStatement, req.Symbol), reqParam)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &sList)
	if err != nil {
		return nil, err
	}

	return sList, nil
}

// BalanceSheetStatementGrowth - balance sheet statement growth
func (c *CompanyValuation) BalanceSheetStatementGrowth(req objects.RequestBalanceSheetStatementGrowth) (sList []objects.BalanceSheetStatementGrowth, err error) {
	reqParam := map[string]string{"limit": fmt.Sprint(req.Limit)}
	if req.Period != objects.CompanyValuationPeriodAnnual {
		reqParam["period"] = string(objects.CompanyValuationPeriodQuarter)
	}

	data, err := c.Client.Get(fmt.Sprintf(urlAPICompanyValuationBalanceSheetStatementGrowth, req.Symbol), reqParam)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &sList)
	if err != nil {
		return nil, err
	}

	return sList, nil
}

// CashFlowStatement - cash flow statement
func (c *CompanyValuation) CashFlowStatement(req objects.RequestCashFlowStatement) (sList []objects.CashFlowStatement, err error) {
	reqParam := map[string]string{"limit": fmt.Sprint(req.Limit)}
	if req.Period != objects.CompanyValuationPeriodAnnual {
		reqParam["period"] = string(objects.CompanyValuationPeriodQuarter)
	}

	data, err := c.Client.Get(fmt.Sprintf(urlAPICompanyValuationCashFlowStatement, req.Symbol), reqParam)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &sList)
	if err != nil {
		return nil, err
	}

	return sList, nil
}

// CashFlowStatementGrowth - cash flow statement growth
func (c *CompanyValuation) CashFlowStatementGrowth(req objects.RequestCashFlowStatementGrowth) (sList []objects.CashFlowStatementGrowth, err error) {
	reqParam := map[string]string{"limit": fmt.Sprint(req.Limit)}
	if req.Period != objects.CompanyValuationPeriodAnnual {
		reqParam["period"] = string(objects.CompanyValuationPeriodQuarter)
	}

	data, err := c.Client.Get(fmt.Sprintf(urlAPICompanyValuationCashFlowStatementGrowth, req.Symbol), reqParam)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &sList)
	if err != nil {
		return nil, err
	}

	return sList, nil
}

// IncomeStatementAsReported - income statement AS REPORTED
func (c *CompanyValuation) IncomeStatementAsReported(req objects.RequestIncomeStatementAsReported) (sList []objects.IncomeStatementAsReported, err error) {
	reqParam := map[string]string{"limit": fmt.Sprint(req.Limit)}
	if req.Period != objects.CompanyValuationPeriodAnnual {
		reqParam["period"] = string(objects.CompanyValuationPeriodQuarter)
	}

	data, err := c.Client.Get(fmt.Sprintf(urlAPICompanyValuationIncomeStatementAsReported, req.Symbol), reqParam)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &sList)
	if err != nil {
		return nil, err
	}

	return sList, nil
}

// BalanceSheetStatementAsReported - balance sheet statement AS REPORTED
func (c *CompanyValuation) BalanceSheetStatementAsReported(req objects.RequestBalanceSheetStatementAsReported) (sList []objects.BalanceSheetStatementAsReported, err error) {
	reqParam := map[string]string{"limit": fmt.Sprint(req.Limit)}
	if req.Period != objects.CompanyValuationPeriodAnnual {
		reqParam["period"] = string(objects.CompanyValuationPeriodQuarter)
	}

	data, err := c.Client.Get(fmt.Sprintf(urlAPICompanyValuationBalanceSheetStatementAsReported, req.Symbol), reqParam)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &sList)
	if err != nil {
		return nil, err
	}

	return sList, nil
}

// CashFlowStatementAsReported - cash flow statement AS REPORTED
func (c *CompanyValuation) CashFlowStatementAsReported(req objects.RequestCashFlowStatementAsReported) (sList []objects.CashFlowStatementAsReported, err error) {
	reqParam := map[string]string{"limit": fmt.Sprint(req.Limit)}
	if req.Period != objects.CompanyValuationPeriodAnnual {
		reqParam["period"] = string(objects.CompanyValuationPeriodQuarter)
	}

	data, err := c.Client.Get(fmt.Sprintf(urlAPICompanyValuationCashFlowStatementAsReported, req.Symbol), reqParam)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &sList)
	if err != nil {
		return nil, err
	}

	return sList, nil
}

// FullFinancialStatementAsReported - full financial statement AS REPORTED
func (c *CompanyValuation) FullFinancialStatementAsReported(req objects.RequestFullFinancialStatementAsReported) (sList []objects.FullFinancialStatementAsReported, err error) {
	reqParam := map[string]string{"limit": fmt.Sprint(req.Limit)}
	if req.Period != objects.CompanyValuationPeriodAnnual {
		reqParam["period"] = string(objects.CompanyValuationPeriodQuarter)
	}

	data, err := c.Client.Get(fmt.Sprintf(urlAPICompanyValuationFinancialStatementFullAsReported, req.Symbol), reqParam)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &sList)
	if err != nil {
		return nil, err
	}

	return sList, nil
}

// FinancialRatios - financial ratios
func (c *CompanyValuation) FinancialRatios(req objects.RequestFinancialRatios) (rList []objects.FinancialRatios, err error) {
	reqParam := map[string]string{"limit": fmt.Sprint(req.Limit)}
	if req.Period != objects.CompanyValuationPeriodAnnual {
		reqParam["period"] = string(objects.CompanyValuationPeriodQuarter)
	}

	data, err := c.Client.Get(fmt.Sprintf(urlAPICompanyValuationFinancialRatios, req.Symbol), reqParam)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &rList)
	if err != nil {
		return nil, err
	}

	return rList, nil
}

// FinancialRatiosTTM - financial ratios TTM
func (c *CompanyValuation) FinancialRatiosTTM(symbol string) (rList []objects.FinancialRatiosTTM, err error) {
	data, err := c.Client.Get(fmt.Sprintf(urlAPICompanyValuationFinancialRatiosTTM, symbol), nil)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &rList)
	if err != nil {
		return nil, err
	}

	return rList, nil
}

// KeyMetrics - key metrics
func (c *CompanyValuation) KeyMetrics(req objects.RequestKeyMetrics) (mList []objects.KeyMetrics, err error) {
	reqParam := map[string]string{"limit": fmt.Sprint(req.Limit)}
	if req.Period != objects.CompanyValuationPeriodAnnual {
		reqParam["period"] = string(objects.CompanyValuationPeriodQuarter)
	}

	data, err := c.Client.Get(fmt.Sprintf(urlAPICompanyValuationKeyMetrics, req.Symbol), reqParam)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &mList)
	if err != nil {
		return nil, err
	}

	return mList, nil
}

// KeyMetricsTTM - key metrics ttm
func (c *CompanyValuation) KeyMetricsTTM(symbol string) (mList []objects.KeyMetricsTTM, err error) {
	data, err := c.Client.Get(fmt.Sprintf(urlAPICompanyValuationKeyMetricsTTM, symbol), nil)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &mList)
	if err != nil {
		return nil, err
	}

	return mList, nil
}

// EnterpriseValue - enterprise value
func (c *CompanyValuation) EnterpriseValue(req objects.RequestEnterpriseValue) (vList []objects.EnterpriseValue, err error) {
	reqParam := map[string]string{"limit": fmt.Sprint(req.Limit)}
	if req.Period != objects.CompanyValuationPeriodAnnual {
		reqParam["period"] = string(objects.CompanyValuationPeriodQuarter)
	}

	data, err := c.Client.Get(fmt.Sprintf(urlAPICompanyValuationEnterpriseValues, req.Symbol), reqParam)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &vList)
	if err != nil {
		return nil, err
	}

	return vList, nil
}

// FinancialStatementsGrowth - financial statements growth
func (c *CompanyValuation) FinancialStatementsGrowth(req objects.RequestFinancialStatementsGrowth) (vList []objects.FinancialStatementsGrowth, err error) {
	reqParam := map[string]string{"limit": fmt.Sprint(req.Limit)}
	if req.Period != objects.CompanyValuationPeriodAnnual {
		reqParam["period"] = string(objects.CompanyValuationPeriodQuarter)
	}

	data, err := c.Client.Get(fmt.Sprintf(urlAPICompanyValuationFinancialGrowth, req.Symbol), reqParam)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &vList)
	if err != nil {
		return nil, err
	}

	return vList, nil
}

// DiscountedCashFlow - discounted cash flow value
func (c *CompanyValuation) DiscountedCashFlow(symbol string) (vList []objects.DiscountedCashFlow, err error) {
	data, err := c.Client.Get(fmt.Sprintf(urlAPICompanyValuationDiscountedCashFlow, symbol), nil)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &vList)
	if err != nil {
		return nil, err
	}

	return vList, nil
}

// DailyDiscountedCashFlow - daily historical DCF
func (c *CompanyValuation) DailyDiscountedCashFlow(req objects.RequestDailyDiscountedCashFlow) (vList []objects.DailyDiscountedCashFlow, err error) {
	data, err := c.Client.Get(
		fmt.Sprintf(urlAPICompanyValuationHistoryDailyDiscountedCashFlow, req.Symbol),
		map[string]string{"limit": fmt.Sprint(req.Limit)},
	)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &vList)
	if err != nil {
		return nil, err
	}

	return vList, nil
}

// HistoryDiscountedCashFlow - history DCF
func (c *CompanyValuation) HistoryDiscountedCashFlow(req objects.RequestHistoryDiscountedCashFlow) (vList []objects.HistoryDiscountedCashFlow, err error) {
	reqParam := map[string]string{"limit": fmt.Sprint(req.Limit)}
	if req.Period != objects.CompanyValuationPeriodAnnual {
		reqParam["period"] = string(objects.CompanyValuationPeriodQuarter)
	}

	data, err := c.Client.Get(fmt.Sprintf(urlAPICompanyValuationHistoryDiscountedCashFlow, req.Symbol), reqParam)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &vList)
	if err != nil {
		return nil, err
	}

	return vList, nil
}

// Rating - get rating by symbol
func (c *CompanyValuation) Rating(symbol string) (rList []objects.Rating, err error) {
	data, err := c.Client.Get(fmt.Sprintf(urlAPICompanyValuationRating, symbol), nil)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &rList)
	if err != nil {
		return nil, err
	}

	return rList, nil
}

// DailyHistoryRating - daily historical rating
func (c *CompanyValuation) DailyHistoryRating(req objects.RequestRating) (rList []objects.Rating, err error) {
	data, err := c.Client.Get(
		fmt.Sprintf(urlAPICompanyValuationHistoryRating, req.Symbol),
		map[string]string{"limit": fmt.Sprint(req.Limit)},
	)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &rList)
	if err != nil {
		return nil, err
	}

	return rList, nil
}

// MarketCapitalization - market capitalization
func (c *CompanyValuation) MarketCapitalization(symbol string) (rList []objects.MarketCapitalization, err error) {
	data, err := c.Client.Get(fmt.Sprintf(urlAPICompanyValuationMarketCapitalization, symbol), nil)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &rList)
	if err != nil {
		return nil, err
	}

	return rList, nil
}

// DailyHistoryMarketCapitalization - daily historical market capitalization
func (c *CompanyValuation) DailyHistoryMarketCapitalization(req objects.RequestMarketCapitalization) (rList []objects.MarketCapitalization, err error) {
	data, err := c.Client.Get(
		fmt.Sprintf(urlAPICompanyValuationHistoryMarketCapitalization, req.Symbol),
		map[string]string{"limit": fmt.Sprint(req.Limit)},
	)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &rList)
	if err != nil {
		return nil, err
	}

	return rList, nil
}

// StockScreener - stock screener
func (c *CompanyValuation) StockScreener(req objects.RequestStockScreener) (sList []objects.StockScreener, err error) {
	reqParam := map[string]string{"limit": fmt.Sprint(req.Limit)}
	if len(req.Exchange) != 0 {
		reqParam["exchange"] = strings.Join(req.Exchange, ",")
	}

	if req.Industry != nil {
		reqParam["industry"] = string(*req.Industry)
	}

	if req.Sector != nil {
		reqParam["sector"] = string(*req.Sector)
	}

	if req.MarketCapMoreThan != nil {
		reqParam["marketCapMoreThan"] = fmt.Sprint(*req.MarketCapMoreThan)
	}

	if req.MarketCapLowerThan != nil {
		reqParam["marketCapLowerThan"] = fmt.Sprint(*req.MarketCapLowerThan)
	}

	if req.VolumeMoreThan != nil {
		reqParam["volumeMoreThan"] = fmt.Sprint(*req.VolumeMoreThan)
	}

	if req.VolumeLowerThan != nil {
		reqParam["volumeLowerThan"] = fmt.Sprint(*req.VolumeLowerThan)
	}

	if req.BetaMoreThan != nil {
		reqParam["betaMoreThan"] = fmt.Sprint(*req.BetaMoreThan)
	}

	if req.BetaLowerThan != nil {
		reqParam["betaLowerThan"] = fmt.Sprint(*req.BetaLowerThan)
	}

	if req.DividendMoreThan != nil {
		reqParam["dividendMoreThan"] = fmt.Sprint(*req.DividendMoreThan)
	}

	if req.DividendLowerThan != nil {
		reqParam["dividendLowerThan"] = fmt.Sprint(*req.DividendLowerThan)
	}

	if req.PriceMoreThan != nil {
		reqParam["priceMoreThan"] = fmt.Sprint(*req.PriceMoreThan)
	}

	if req.PriceLowerThan != nil {
		reqParam["priceLowerThan"] = fmt.Sprint(*req.PriceLowerThan)
	}

	if req.IsETF != nil {
		reqParam["isEtf"] = fmt.Sprint(*req.IsETF)
	}

	if req.IsActivelyTrading != nil {
		reqParam["isActivelyTrading"] = fmt.Sprint(*req.IsActivelyTrading)
	}

	if req.Country != nil {
		reqParam["country"] = *req.Country
	}

	data, err := c.Client.Get(urlAPICompanyValuationStockScreener, reqParam)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &sList)
	if err != nil {
		return nil, err
	}

	return sList, nil
}

// DelstedCompanies - delsted companies
func (c *CompanyValuation) DelstedCompanies(limit int64) (cList []objects.DelstedCompany, err error) {
	data, err := c.Client.Get(
		urlAPICompanyValuationDelistedCompanyList,
		map[string]string{"limit": fmt.Sprint(limit)},
	)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &cList)
	if err != nil {
		return nil, err
	}

	return cList, nil
}

// StockNews - stock news
func (c *CompanyValuation) StockNews(req objects.RequestStockNews) (vList []objects.StockNews, err error) {
	reqParam := map[string]string{"limit": fmt.Sprint(req.Limit)}
	if len(req.SymbolList) != 0 {
		reqParam["tickers"] = strings.Join(req.SymbolList, ",")
	}

	data, err := c.Client.Get(urlAPICompanyValuationStockNews, reqParam)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &vList)
	if err != nil {
		return nil, err
	}

	return vList, nil
}

// AnalystEstimates - analyst estimates of a stock (Annual || Quarter)
func (c *CompanyValuation) AnalystEstimates(req objects.RequestAnalystEstimates) (vList []objects.AnalystEstimates, err error) {
	reqParam := make(map[string]string)
	if req.Period != objects.CompanyValuationPeriodAnnual {
		reqParam["period"] = string(objects.CompanyValuationPeriodQuarter)
	}

	data, err := c.Client.Get(fmt.Sprintf(urlAPICompanyValuationAnalystEstimates, req.Symbol), reqParam)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &vList)
	if err != nil {
		return nil, err
	}

	return vList, nil
}

// Grade - stock grade from analysts
func (c *CompanyValuation) Grade(req objects.RequestGrade) (gList []objects.Grade, err error) {
	data, err := c.Client.Get(
		fmt.Sprintf(urlAPICompanyValuationGrade, req.Symbol),
		map[string]string{"limit": fmt.Sprint(req.Limit)},
	)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &gList)
	if err != nil {
		return nil, err
	}

	return gList, nil
}

// AnalystStockRecommendations - monthly stock analyst ratings
func (c *CompanyValuation) AnalystStockRecommendations(req objects.RequestAnalystStockRecommendations) (rList []objects.AnalystStockRecommendations, err error) {
	data, err := c.Client.Get(
		fmt.Sprintf(urlAPICompanyValuationAnalystStockRecommendations, req.Symbol),
		map[string]string{"limit": fmt.Sprint(req.Limit)},
	)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &rList)
	if err != nil {
		return nil, err
	}

	return rList, nil
}

// PressReleases - stock press releases
func (c *CompanyValuation) PressReleases(req objects.RequestPressReleases) (prList []objects.PressReleases, err error) {
	data, err := c.Client.Get(
		fmt.Sprintf(urlAPICompanyValuationPressReleases, req.Symbol),
		map[string]string{"limit": fmt.Sprint(req.Limit)},
	)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &prList)
	if err != nil {
		return nil, err
	}

	return prList, nil
}

// FinancialStatementList - List of symbols that have financial statements
func (c *CompanyValuation) FinancialStatementList() (fsList []string, err error) {
	data, err := c.Client.Get(urlAPICompanyValuationFinancialStatementsList, nil)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &fsList)
	if err != nil {
		return nil, err
	}

	return fsList, nil
}

// EconomicCalendar - Economic Calendar for time period
func (c *CompanyValuation) EconomicCalendar(req objects.RequestEconomicCalendar) (eList []objects.EconomicCalendar, err error) {
	reqParam := make(map[string]string)
	if req.From != nil {
		reqParam["from"] = req.From.Format("2006-01-02")
	}

	if req.To != nil {
		reqParam["to"] = req.To.Format("2006-01-02")
	}

	data, err := c.Client.Get(urlAPICompanyValuationEconomicCalendar, reqParam)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &eList)
	if err != nil {
		return nil, err
	}

	return eList, nil
}

// SECFilings - SEC Filings with type or no (10-K || 4 || 8-K || 424B2 || FWP || SC 13G/A || SD || 10-Q || PX14A6G || DEFA14A || DEF 14A || 8-A12B || CERT || 25 ...)
func (c *CompanyValuation) SECFilings(req objects.RequestSECFilings) (eList []objects.SECFiling, err error) {
	reqParam := map[string]string{"limit": fmt.Sprint(req.Limit)}
	if req.Type != nil {
		reqParam["type"] = *req.Type
	}

	data, err := c.Client.Get(fmt.Sprintf(urlAPICompanyValuationSECFillings, req.Symbol), reqParam)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &eList)
	if err != nil {
		return nil, err
	}

	return eList, nil
}

// HistoryEconomicCalendar - Economic calendar event list
func (c *CompanyValuation) HistoryEconomicCalendar(req objects.RequestHistoryEconomicCalendar) (hList []objects.HistoryEconomicCalendar, err error) {
	data, err := c.Client.Get(fmt.Sprintf(urlAPICompanyValuationHistoryEconomicCalendar, req.Event), map[string]string{"country": req.Country})
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &hList)
	if err != nil {
		return nil, err
	}

	return hList, nil
}

// EconomicCalendarEventList - Example of historical consumer sentiment in U.S. (take event name and country from event list endpoint)
func (c *CompanyValuation) EconomicCalendarEventList() (eList []objects.EconomicCalendarEventList, err error) {
	data, err := c.Client.Get(urlAPICompanyValuationEconomicCalendarEventList, nil)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &eList)
	if err != nil {
		return nil, err
	}

	return eList, nil
}

// ETFList - All ETF symbols
func (c *CompanyValuation) ETFList() (fList []objects.ETF, err error) {
	data, err := c.Client.Get(urlAPICompanyValuationETFList, nil)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &fList)
	if err != nil {
		return nil, err
	}

	return fList, nil
}

// AvailableTradedList - All tradable Symbols
func (c *CompanyValuation) AvailableTradedList() (fList []objects.AvailableTraded, err error) {
	data, err := c.Client.Get(urlAPICompanyValuationAvailableTradedList, nil)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &fList)
	if err != nil {
		return nil, err
	}

	return fList, nil
}

// CompanyOutlook - Company Outlook
func (c *CompanyValuation) CompanyOutlook(symbol string) (co *objects.CompanyOutlook, err error) {
	data, err := c.Client.Get(urlAPICompanyValuationCompanyOutlook, map[string]string{"symbol": symbol})
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &co)
	if err != nil {
		return nil, err
	}

	return co, nil
}

// EmployeeCount - Historical number of employees
func (c *CompanyValuation) EmployeeCount(symbol string) (eList *objects.EmployeeCount, err error) {
	data, err := c.Client.Get(urlAPICompanyValuationEmployeeCount, map[string]string{"symbol": symbol})
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &eList)
	if err != nil {
		return nil, err
	}

	return eList, nil
}

// SocialSentimentTrending - Trending social sentiment
func (c *CompanyValuation) SocialSentimentTrending(tType, source string) (sList []objects.SocialSentiment, err error) {
	req := map[string]string{}
	if len(tType) != 0 {
		req["type"] = tType
	}

	if len(source) != 0 {
		req["source"] = source
	}

	data, err := c.Client.Get(urlAPICompanyValuationSocialSentimentTrending, req)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &sList)
	if err != nil {
		return nil, err
	}

	return sList, nil
}

// SocialSentimentChange - Biggest changes in social sentiment per type and source
func (c *CompanyValuation) SocialSentimentChange(tType, source string) (sList []objects.SocialSentimentChange, err error) {
	req := map[string]string{}
	if len(tType) != 0 {
		req["type"] = tType
	}

	if len(source) != 0 {
		req["source"] = source
	}

	data, err := c.Client.Get(urlAPICompanyValuationSocialSentimentChange, req)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &sList)
	if err != nil {
		return nil, err
	}

	return sList, nil
}

// HistoricalSocialSentiment - Historical Social Media sentiment for stock (time in UTC)
func (c *CompanyValuation) HistoricalSocialSentiment(symbol string) (sList []objects.SocialSentiment, err error) {
	data, err := c.Client.Get(urlAPICompanyValuationHistoricalSocialSentiment, map[string]string{"symbol": symbol})
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &sList)
	if err != nil {
		return nil, err
	}

	return sList, nil
}

// Score - Stock Financial scores
func (c *CompanyValuation) Score(symbol string) (sList []objects.Score, err error) {
	data, err := c.Client.Get(urlAPICompanyValuationScore, map[string]string{"symbol": symbol})
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &sList)
	if err != nil {
		return nil, err
	}

	return sList, nil
}

// BulkScores - Stock Financial scores (bulk)
func (c *CompanyValuation) BulkScores() (sList []objects.Score, err error) {
	data, err := c.Client.Get(urlAPICompanyValuationBulkScores, nil)
	if err != nil {
		return nil, err
	}

	err = gocsv.UnmarshalBytes(data.Body(), &sList)
	if err != nil {
		return nil, err
	}

	return sList, nil
}

// BulkIncomeStatement ...
func (c *CompanyValuation) BulkIncomeStatement(year int, period string) (pList []objects.IncomeStatement, err error) {
	data, err := c.Client.Get(
		urlAPICompanyValuationBulkIncomeStatement,
		map[string]string{
			"year":   fmt.Sprint(year),
			"period": period,
		})
	if err != nil {
		return nil, err
	}

	err = gocsv.UnmarshalBytes(data.Body(), &pList)
	if err != nil {
		return nil, err
	}

	return pList, nil
}

// BulkBalanceSheetStatement ...
func (c *CompanyValuation) BulkBalanceSheetStatement(year int, period string) (sList []objects.BalanceSheetStatement, err error) {
	data, err := c.Client.Get(
		urlAPICompanyValuationBulkBalanceSheetStatement,
		map[string]string{
			"year":   fmt.Sprint(year),
			"period": period,
		})
	if err != nil {
		return nil, err
	}

	err = gocsv.UnmarshalBytes(data.Body(), &sList)
	if err != nil {
		return nil, err
	}

	return sList, nil
}

// BulkCashFlowStatement ...
func (c *CompanyValuation) BulkCashFlowStatement(year int, period string) (sList []objects.CashFlowStatement, err error) {
	data, err := c.Client.Get(
		urlAPICompanyValuationBulkCashFlowStatement,
		map[string]string{
			"year":   fmt.Sprint(year),
			"period": period,
		})
	if err != nil {
		return nil, err
	}

	err = gocsv.UnmarshalBytes(data.Body(), &sList)
	if err != nil {
		return nil, err
	}

	return sList, nil
}

// BulkCashFlowStatement ...
func (c *CompanyValuation) BulkKeyMetrics(year int, period string) (sList []objects.KeyMetrics, err error) {
	data, err := c.Client.Get(
		urlAPICompanyValuationBulkKeyMetrics,
		map[string]string{
			"year":   fmt.Sprint(year),
			"period": period,
		})
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &sList)
	if err != nil {
		return nil, err
	}

	return sList, nil
}

// BulkRatios ...
func (c *CompanyValuation) BulkRatios(year int, period string) (sList []objects.FinancialRatios, err error) {
	data, err := c.Client.Get(
		urlAPICompanyValuationBulkRatios,
		map[string]string{
			"year":   fmt.Sprint(year),
			"period": period,
		})
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &sList)
	if err != nil {
		return nil, err
	}

	return sList, nil
}

// BulkEarningsSurpises ...
func (c *CompanyValuation) BulkEarningsSurpises(year int) (sList []objects.EarningSurprise, err error) {
	data, err := c.Client.Get(urlAPICompanyValuationBulkEarningsSurpises, map[string]string{"year": fmt.Sprint(year)})
	if err != nil {
		return nil, err
	}

	err = gocsv.UnmarshalBytes(data.Body(), &sList)
	if err != nil {
		return nil, err
	}

	return sList, nil
}

// BulkRating ...
func (c *CompanyValuation) BulkRating() (sList []objects.Rating, err error) {
	data, err := c.Client.Get(urlAPICompanyValuationBulkRating, nil)
	if err != nil {
		return nil, err
	}

	err = gocsv.UnmarshalBytes(data.Body(), &sList)
	if err != nil {
		return nil, err
	}

	return sList, nil
}

// RatiosTTMBulk ...
func (c *CompanyValuation) RatiosTTMBulk() (rList []objects.FinancialRatiosTTM, err error) {
	data, err := c.Client.Get(urlAPICompanyValuationRatiosTTMBulk, nil)
	if err != nil {
		return nil, err
	}

	err = gocsv.UnmarshalBytes(data.Body(), &rList)
	if err != nil {
		return nil, err
	}

	return rList, nil
}

// DCFBulk ...
func (c *CompanyValuation) DCFBulk() (dList []objects.DailyDiscountedCashFlow, err error) {
	data, err := c.Client.Get(urlAPICompanyValuationDCFBulk, nil)
	if err != nil {
		return nil, err
	}

	err = gocsv.UnmarshalBytes(data.Body(), &dList)
	if err != nil {
		return nil, err
	}

	return dList, nil
}

// SharesFloatAll - All latest shares float available
func (c *CompanyValuation) SharesFloatAll() (sList []objects.SharesFloat, err error) {
	data, err := c.Client.Get(urlAPICompanyValuationSharesFloatAll, nil)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &sList)
	if err != nil {
		return nil, err
	}

	return sList, nil
}

// SharesFloat - Shares float for symbol
func (c *CompanyValuation) SharesFloat(symbol string) (ipoList []objects.IPOCalendar, err error) {
	data, err := c.Client.Get(
		urlAPICompanyValuationSharesFloat,
		map[string]string{
			"symbol": symbol,
		},
	)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &ipoList)
	if err != nil {
		return nil, err
	}

	return ipoList, nil
}
