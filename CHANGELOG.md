# Release notes
## 1.0.0
* Release package with support all avalible methods

## 1.0.1
* Add support: New param for form-thirteen endpoint: date

## 1.0.2
**Add support:**
* International Filings for income statements (new endpoint)
* Technical Indicators (new endpoint)

**Fix:**
* 13F list request

## 1.0.3
**Add support:**
* Earning Call Transcript (new endpoint)

## 1.0.4
**Add support:**
* Analyst Estimates (new endpoint)

## 1.0.5
**Add support:**
* New endpoint: Press Releases
* New endpoint: Analyst Stock Recommendations
* New endpoint: Stock Grade

## 1.0.7
**Add support:**
* New endpoint: Financial Statements List
* New endpoint: Economic Calendar
* New website: Formulas

## 1.0.9
**Fix:**
* Error url, for request split calendar (method SplitCalendar)
* Linters error

**Add support:**
* New properties: currency, isin and cusip for: Company Profile
* New properties: dividend, recordDate, paymentDate and declarationDate for: Historical Dividends
* New endpoint: Earnings Surprises

## 1.0.10
**Fix:**
* License

**Add support:**
* New endpoint for: Key Metrics TTM
* New properties: currentRatioTTM, quickRatioTTM, cashRatioTTM, daysOfSalesOutstandingTTM, daysOfInventoryOutstandingTTM and more for: Company TTM ratios

## 1.0.11
**Add support:**
* New endpoint: SEC Filings

## 1.0.12
**Add support:**
* New website: Status
* New properties: From and To (time) for request Candles - historical candles (Stock, Forex, Crypto)

## 1.1.0
**Add support:**
* New endpoint: List of Dow Jones companies
* New property: cik for: Company Profile
* New endpoint: List of Nasdaq 100 companies
* New endpoint: Historical Nasdaq 100
* New endpoint: Historical Dow Jones

**Other:**
* Method renamed SP500CompanyList => IndexConstituentList
* Method renamed HistorySP500CompanyList => HistoryIndexConstituentList

## 1.1.1
**Fix:**
* Index const

## 1.1.2
**Fix:**
* Period const and types renamed (forex, stock, crypto)

## 1.2.0
**Add:**
* New stock method DailyBatch

## 2.0.0
**Update:**
* Add new logic for retries request if response not status Ok
* Change Licinse to MIT

## 2.1.0
**Add:**
* New endpoint: Survivorship Bias Free EOD
* New params for Stock Screener: priceLowerThan and priceMoreThan
* New property reportedCurrency for: Financial Statements
* New API service for custome request for FMP

## 2.2.0
**Add:**
* New endpoint: Stock Insider Trading (Only for API v4)
* New endpoint: Insider Trading RSS Feed (Only for API v4)

**Fix:**
* Merge two pull request from [mariusgrigoriu](https://github.com/mariusgrigoriu) to fix panics

## 2.3.0
**Add:**
* New properties: isEtf and isActivelyTrading and defaultImage for: Company Profile

## 2.4.0
**Add:**
* New endpoint: ETF List
* New endpoint: Tradable Symbols List
* New params for Stock Screener: isEtf, isActivelyTrading and country

## 2.4.1
**Add:**
* New property typeOfOwner for: Stock Insider Trading

**Fix:**
* CryptoDailyCandleList => StockDailyCandleList in StockDailyCandleList

## 2.5.0
**Add:**
* New endpoint: Company Outlook
* New endpoint: COT Trading Symbols List
* New endpoint: Commitments of Traders Analysis
* New endpoint: Commitments of Traders Report
* New property price for: Stock Insider Trading

**Fix:**
* Stock => DailyBatch