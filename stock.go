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
	urlAPIStockSymbolList                = "/v3/stock/list"
	urlAPIStockSymbolByExchangeList      = "/v3/symbol/%s"
	urlAPIStockBulkProfile               = "/v4/profile/all"
	urlAPIStockCompanyProfile            = "/v3/profile/%s"
	urlAPIStockCompanyExecutives         = "/v3/key-executives/%s"
	urlAPIStockDividends                 = "/v3/historical-price-full/stock_dividend/%s"
	urlAPIStockSplits                    = "/v3/historical-price-full/stock_split/%s"
	urlAPIStockQuote                     = "/v3/quote/%s"
	urlAPIStockQuoteShot                 = "/v3/quote-short/%s"
	urlAPIStockQuotes                    = "/v3/quotes/%s"
	urlAPIStockSearch                    = "/v3/search"
	urlAPIStockSearchTicker              = "/v3/search-ticker"
	urlAPIStockSearchName                = "/v3/search-name"
	urlAPIStockCandles                   = "/v3/historical-chart/%s/%s"
	urlAPIStockDaily                     = "/v3/historical-price-full/%s"
	urlAPIStockSP500List                 = "/v3/sp500_constituent"
	urlAPIStockHistorySP500List          = "/v3/historical/sp500_constituent"
	urlAPIStockDowJonesList              = "/v3/dowjones_constituent"
	urlAPIStockHistoryDowJonesList       = "/v3/historical/dowjones_constituent"
	urlAPIStockNasdaqList                = "/v3/nasdaq_constituent"
	urlAPIStockHistoryNasdaqList         = "/v3/historical/nasdaq_constituent"
	urlAPIStockEODCandles                = "/v3/batch-request-end-of-day-prices"
	urlAPIStockEODBatchCandles           = "/v3/batch-request-end-of-day-prices/%s"
	urlAPIStockEODBatchPrices            = "/v4/batch-request-end-of-day-prices"
	urlAPIStockMarketHours               = "/v3/market-hours"
	urlAPIStockActives                   = "/v3/actives"
	urlAPIStockLosers                    = "/v3/losers"
	urlAPIStockGainers                   = "/v3/gainers"
	urlAPIStockSectorsPerformance        = "/v3/sectors-performance"
	urlAPIStockOTCRealTimePrice          = "/v3/otc/real-time-price/%s"
	urlAPIStockHistorySectorsPerformance = "/v3/historical-sectors-performance"
	urlAPIStockPeers                     = "/v4/stock_peers"
	urlAPIStockBulkPeers                 = "/v4/stock_peers_bulk"
	urlAPIStockCompanyCoreInformation    = "/v4/company-core-information"
	urlAPIStockSurvivorshipBiasFree      = "/v4/historical-price-full/%s/%s"
	urlAPIStockPriceChange               = "/v3/stock-price-change/%s"
	urlAPIStockPriceChangeBatch          = "/v3/stock-price-change/%s"
)

// Stock client
type Stock struct {
	Client *HTTPClient
}

// QuoteShort - real-time single quote short
func (s *Stock) QuoteShort(symbol string) (qList []objects.StockQuoteShot, err error) {
	data, err := s.Client.Get(fmt.Sprintf(urlAPIStockQuoteShot, symbol), nil)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &qList)
	if err != nil {
		return nil, err
	}

	return qList, nil
}

// Quote - real-time single quote
func (s *Stock) Quote(symbol string) (qList []objects.StockQuote, err error) {
	data, err := s.Client.Get(fmt.Sprintf(urlAPIStockQuote, symbol), nil)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &qList)
	if err != nil {
		return nil, err
	}

	return qList, nil
}

// BatchQuote - real-time batch quote
func (s *Stock) BatchQuote(symbolList []string) (qList []objects.StockQuote, err error) {
	data, err := s.Client.Get(fmt.Sprintf(urlAPIStockQuote, strings.Join(symbolList, ",")), nil)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &qList)
	if err != nil {
		return nil, err
	}

	return qList, nil
}

// QuoteByExchange - real-time single quote
func (s *Stock) QuoteByExchange(exchange objects.StockSearch) (qList []objects.StockQuote, err error) {
	data, err := s.Client.Get(fmt.Sprintf(urlAPIStockQuotes, exchange.String()), nil)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &qList)
	if err != nil {
		return nil, err
	}

	return qList, nil
}

// Search - Search via ticker and company name
func (s *Stock) Search(req objects.RequestStockSearch) (sList []objects.StockSymbol, err error) {
	reqParam := map[string]string{
		"limit": fmt.Sprint(req.Limit),
		"query": req.Query,
	}
	if req.Exchange != nil {
		reqParam["exchange"] = req.Exchange.String()
	}

	data, err := s.Client.Get(urlAPIStockSearch, reqParam)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &sList)
	if err != nil {
		return nil, err
	}

	return sList, nil
}

// SearchTiker - Search only via ticker
func (s *Stock) SearchTiker(req objects.RequestStockSearch) (sList []objects.StockSymbol, err error) {
	reqParam := map[string]string{
		"limit": fmt.Sprint(req.Limit),
		"query": req.Query,
	}
	if req.Exchange != nil {
		reqParam["exchange"] = req.Exchange.String()
	}

	data, err := s.Client.Get(urlAPIStockSearchTicker, reqParam)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &sList)
	if err != nil {
		return nil, err
	}

	return sList, nil
}

// SearchByName - Search only via company name
func (s *Stock) SearchByName(req objects.RequestStockSearch) (sList []objects.StockSymbol, err error) {
	reqParam := map[string]string{
		"limit": fmt.Sprint(req.Limit),
		"query": req.Query,
	}
	if req.Exchange != nil {
		reqParam["exchange"] = req.Exchange.String()
	}

	data, err := s.Client.Get(urlAPIStockSearchName, reqParam)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &sList)
	if err != nil {
		return nil, err
	}

	return sList, nil
}

// BulkProfile - get all available profiles
func (s *Stock) BulkProfile() (companyProfile []objects.StockCompanyProfile, err error) {
	data, err := s.Client.Get(urlAPIStockBulkProfile, nil)
	if err != nil {
		return
	}

	err = gocsv.UnmarshalBytes(data.Body(), &companyProfile)
	return
}

// CompanyProfile - get general information of a company. You can query by symbol.
func (s *Stock) CompanyProfile(symbol string) (companyProfile []objects.StockCompanyProfile, err error) {
	data, err := s.Client.Get(fmt.Sprintf(urlAPIStockCompanyProfile, symbol), nil)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &companyProfile)
	if err != nil {
		return nil, err
	}

	return companyProfile, nil
}

// Peers - Stock peers based on sector, exchange and market cap
func (s *Stock) Peers(symbol string) (pList []objects.StockPeers, err error) {
	data, err := s.Client.Get(urlAPIStockPeers, map[string]string{"symbol": symbol})
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &pList)
	if err != nil {
		return nil, err
	}

	return pList, nil
}

// BulkPeers - Stock peers for all symbols with profile CSV
func (s *Stock) BulkPeers() (pList []objects.StockBulkPeers, err error) {
	data, err := s.Client.Get(urlAPIStockBulkPeers, nil)
	if err != nil {
		return nil, err
	}

	err = gocsv.UnmarshalBytes(data.Body(), &pList)
	if err != nil {
		return nil, err
	}

	return pList, nil
}

// CompanyCoreInformation - Company core information
func (s *Stock) CompanyCoreInformation(symbol string) (company []objects.CompanyCoreInformation, err error) {
	data, err := s.Client.Get(urlAPIStockCompanyCoreInformation, map[string]string{"symbol": symbol})
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &company)
	if err != nil {
		return nil, err
	}

	return company, nil
}

// CompanyExecutive - get a list of company's executives and members of the Board.
func (s *Stock) CompanyExecutive(symbol string) (companyProfile []objects.CompanyExecutive, err error) {
	data, err := s.Client.Get(fmt.Sprintf(urlAPIStockCompanyExecutives, symbol), nil)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &companyProfile)
	if err != nil {
		return nil, err
	}

	return companyProfile, nil
}

// Candles - historical candles
func (s *Stock) Candles(req objects.RequestStockCandleList) (cList []objects.StockCandle, err error) {
	reqParam := make(map[string]string)
	if req.From != nil {
		reqParam["from"] = req.From.Format("2006-01-02")
	}

	if req.To != nil {
		reqParam["to"] = req.To.Format("2006-01-02")
	}

	data, err := s.Client.Get(fmt.Sprintf(urlAPIStockCandles, req.Period, req.Symbol), reqParam)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &cList)
	if err != nil {
		return nil, err
	}

	return cList, nil
}

// DailyLine - daily line
func (s *Stock) DailyLine(symbol string, serieType objects.StockSerieType) (cList *objects.StockDailyLineList, err error) {
	data, err := s.Client.Get(fmt.Sprintf(urlAPIStockDaily, symbol), map[string]string{"serietype": string(serieType)})
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &cList)
	if err != nil {
		return nil, err
	}

	return cList, nil
}

// DailyChangeAndVolume - daily candle change and volume
func (s *Stock) DailyChangeAndVolume(symbol string) (cList *objects.StockDailyCandleList, err error) {
	data, err := s.Client.Get(fmt.Sprintf(urlAPIStockDaily, symbol), nil)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &cList)
	if err != nil {
		return nil, err
	}

	return cList, nil
}

// DailySpecificPeriod - daily candle list by specific period
func (s *Stock) DailySpecificPeriod(symbol string, from time.Time, to time.Time) (cList *objects.StockDailyCandleList, err error) {
	data, err := s.Client.Get(
		fmt.Sprintf(urlAPIStockDaily, symbol),
		map[string]string{
			"from": from.Format("2006-01-02"),
			"to":   to.Format("2006-01-02"),
		})
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &cList)
	if err != nil {
		return nil, err
	}

	return cList, nil
}

// DailyLastNDays - daily candle list last N days
func (s *Stock) DailyLastNDays(symbol string, days int) (cList *objects.StockDailyCandleList, err error) {
	data, err := s.Client.Get(fmt.Sprintf(urlAPIStockDaily, symbol), map[string]string{"timeseries": fmt.Sprint(days)})
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &cList)
	if err != nil {
		return nil, err
	}

	return cList, nil
}

// DailyBatch - daily candle list
func (s *Stock) DailyBatch(symbolList []string, from *time.Time, to *time.Time) (cList []objects.StockBatchData, err error) {
	reqParam := make(map[string]string)
	if from != nil {
		reqParam["from"] = from.Format("2006-01-02")
	}

	if to != nil {
		reqParam["to"] = to.Format("2006-01-02")
	}

	data, err := s.Client.Get(fmt.Sprintf(urlAPIStockDaily, strings.Join(symbolList, ",")), reqParam)
	if err != nil {
		return nil, err
	}

	if len(symbolList) > 1 {
		var resp objects.StockBatchDaily
		err = jsoniter.Unmarshal(data.Body(), &resp)
		if err != nil {
			return nil, err
		}

		cList = resp.Data
	}

	if len(symbolList) == 1 {
		var resp objects.StockBatchData
		err = jsoniter.Unmarshal(data.Body(), &resp)
		if err != nil {
			return nil, err
		}

		cList = append(cList, resp)
	}

	return cList, nil
}

// Dividends - stock dividends
func (s *Stock) Dividends(symbol string) (dList *objects.StockDividends, err error) {
	data, err := s.Client.Get(fmt.Sprintf(urlAPIStockDividends, symbol), nil)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &dList)
	if err != nil {
		return nil, err
	}

	return dList, nil
}

// Splits - stock splits
func (s *Stock) Splits(symbol string) (sList *objects.StockSplit, err error) {
	data, err := s.Client.Get(fmt.Sprintf(urlAPIStockSplits, symbol), nil)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &sList)
	if err != nil {
		return nil, err
	}

	return sList, nil
}

// AvalibleSymbolsByExchange - symbol list by exchange
func (s *Stock) AvalibleSymbolsByExchange(exchange objects.StockSymbolExchange) (sList []objects.StockSymbol, err error) {
	data, err := s.Client.Get(fmt.Sprintf(urlAPIStockSymbolByExchangeList, exchange.String()), nil)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &sList)
	if err != nil {
		return nil, err
	}

	return sList, nil
}

// AvalibleSymbols - all avalible symbol list
func (s *Stock) AvalibleSymbols() (sList []objects.StockSymbolList, err error) {
	data, err := s.Client.Get(urlAPIStockSymbolList, nil)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &sList)
	if err != nil {
		return nil, err
	}

	return sList, nil
}

// IndexConstituentList - list of index companies (SP500, Nasdaq, DJ)
func (s *Stock) IndexConstituentList(index objects.Index) (sList []objects.IndexSymbol, err error) {
	var endpoint string
	switch index {
	case objects.IndexSP500:
		endpoint = urlAPIStockSP500List
	case objects.IndexDowJones:
		endpoint = urlAPIStockDowJonesList
	case objects.IndexNasdaq100:
		endpoint = urlAPIStockNasdaqList
	}

	data, err := s.Client.Get(endpoint, nil)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &sList)
	if err != nil {
		return nil, err
	}

	return sList, nil
}

// HistoryIndexConstituentList - historical index companies list (SP500, Nasdaq, DJ)
func (s *Stock) HistoryIndexConstituentList(index objects.Index) (sList []objects.HistoryIndexSymbol, err error) {
	var endpoint string
	switch index {
	case objects.IndexSP500:
		endpoint = urlAPIStockHistorySP500List
	case objects.IndexDowJones:
		endpoint = urlAPIStockHistoryDowJonesList
	case objects.IndexNasdaq100:
		endpoint = urlAPIStockHistoryNasdaqList
	}

	data, err := s.Client.Get(endpoint, nil)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &sList)
	if err != nil {
		return nil, err
	}

	return sList, nil
}

// EODCandleList - all stocks Batch EOD stock price
func (s *Stock) EODCandleList(date time.Time) (sList []objects.StockEODCandle, err error) {
	data, err := s.Client.Get(urlAPIStockEODCandles, map[string]string{"date": date.Format("2006-01-02")})
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &sList)
	if err != nil {
		return nil, err
	}

	return sList, nil
}

// BatchEODCandleList - specific Stocks Batch EOD stock prices
func (s *Stock) BatchEODCandleList(symbolList []string, date time.Time) (sList []objects.StockEODCandle, err error) {
	data, err := s.Client.Get(
		fmt.Sprintf(urlAPIStockEODBatchCandles, strings.Join(symbolList, ",")),
		map[string]string{
			"date": date.Format("2006-01-02"),
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

// PriceChangeBatch - Price percentage change for multiple timeframes
func (s *Stock) PriceChange(symbol string) (sList []objects.StockPriceChange, err error) {
	data, err := s.Client.Get(fmt.Sprintf(urlAPIStockPriceChange, symbol), nil)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &sList)
	if err != nil {
		return nil, err
	}

	return sList, nil
}

// PriceChangeBatch - Multiple companies price percentage change
func (s *Stock) PriceChangeBatch(symbolList []string) (sList []objects.StockPriceChange, err error) {
	data, err := s.Client.Get(fmt.Sprintf(urlAPIStockPriceChangeBatch, strings.Join(symbolList, ",")), nil)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &sList)
	if err != nil {
		return nil, err
	}

	return sList, nil
}

// EODBatchPrices ...
func (s *Stock) EODBatchPrices() (sList []objects.StockEODCandle, err error) {
	data, err := s.Client.Get(urlAPIStockEODBatchPrices, nil)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &sList)
	if err != nil {
		return nil, err
	}

	return sList, nil
}

// ExchangeTradingHours - stock market trading hours
func (s *Stock) ExchangeTradingHours() (eList []objects.Exchange, err error) {
	data, err := s.Client.Get(urlAPIStockMarketHours, nil)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &eList)
	if err != nil {
		return nil, err
	}

	return eList, nil
}

// Actives - stock market top active
func (s *Stock) Actives() (aList []objects.Active, err error) {
	data, err := s.Client.Get(urlAPIStockActives, nil)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &aList)
	if err != nil {
		return nil, err
	}

	return aList, nil
}

// Losers - stock market top losers
func (s *Stock) Losers() (lList []objects.Loser, err error) {
	data, err := s.Client.Get(urlAPIStockLosers, nil)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &lList)
	if err != nil {
		return nil, err
	}

	return lList, nil
}

// Gainers - stock market top gainers
func (s *Stock) Gainers() (gList []objects.Gainer, err error) {
	data, err := s.Client.Get(urlAPIStockGainers, nil)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &gList)
	if err != nil {
		return nil, err
	}

	return gList, nil
}

// SectorPerformance - stock market sector performance
func (s *Stock) SectorPerformance() (eList []objects.Sector, err error) {
	data, err := s.Client.Get(urlAPIStockSectorsPerformance, nil)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &eList)
	if err != nil {
		return nil, err
	}

	return eList, nil
}

// HistorySectorPerformance - historical stock market sector performance
func (s *Stock) HistorySectorPerformance() (eList []objects.HistorySector, err error) {
	data, err := s.Client.Get(urlAPIStockHistorySectorsPerformance, nil)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &eList)
	if err != nil {
		return nil, err
	}

	return eList, nil
}

// SurvivorshipBiasFree - Survivorship Bias Free end of day (only for api v4)
func (s *Stock) SurvivorshipBiasFree(symbol string, date time.Time) (sBias *objects.SurvivorshipBiasFree, err error) {
	data, err := s.Client.Get(fmt.Sprintf(urlAPIStockSurvivorshipBiasFree, symbol, date.Format("2006-01-02")), nil)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &sBias)
	if err != nil {
		return nil, err
	}

	return sBias, nil
}

// OTCRealTimePrice - Prices of OTC companies
func (s *Stock) OTCRealTimePrice(symbolList []string) (pList *objects.OTCRealTimePrice, err error) {
	data, err := s.Client.Get(fmt.Sprintf(urlAPIStockOTCRealTimePrice, strings.Join(symbolList, ",")), nil)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &pList)
	if err != nil {
		return nil, err
	}

	return pList, nil
}
