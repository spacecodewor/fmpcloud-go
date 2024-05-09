package fmpcloud

import (
	"math"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/time/rate"
)

// APIUrl type for api url
type APIUrl string

// Config for create new API client
type Config struct {
	Logger        *zap.Logger
	HTTPClient    *resty.Client
	APIKey        string
	APIUrl        APIUrl
	Debug         bool
	RateLimiter   *rate.Limiter
	RetryCount    *int
	RetryWaitTime *time.Duration
	Timeout       int
}

// APIClient ...
type APIClient struct {
	Stock              *Stock
	Forex              *Forex
	Form13F            *Form13F
	Crypto             *Crypto
	CompanyValuation   *CompanyValuation
	TechnicalIndicator *TechnicalIndicator
	InsiderTrading     *InsiderTrading
	AlternativeData    *AlternativeData
	Economics          *Economics
	API                *API
	Logger             *zap.Logger
	Debug              bool
}

// Core params
const (
	APIFmpcloudURL              APIUrl = "https://fmpcloud.io/api"
	APIFinancialModelingPrepURL APIUrl = "https://financialmodelingprep.com/api"
	apiDefaultKey                      = "demo"
	apiDefaultTimeout                  = 25
)

// NewAPIClient creates a new API client
func NewAPIClient(cfg Config) (*APIClient, error) {
	APIClient := &APIClient{Logger: cfg.Logger, Debug: cfg.Debug}
	if APIClient.Logger == nil {
		logger, err := createNewLogger()
		if err != nil {
			return nil, errors.Wrap(err, "Error create new zap logger")
		}

		APIClient.Logger = logger
	}

	// Check set timeout param
	if cfg.Timeout == 0 {
		cfg.Timeout = apiDefaultTimeout
	}

	// Init rest client (resty)
	if cfg.HTTPClient == nil {
		cfg.HTTPClient = resty.New()
	}

	cfg.HTTPClient.SetDebug(APIClient.Debug)
	cfg.HTTPClient.SetTimeout(time.Duration(cfg.Timeout) * time.Second)

	// Check set APIUrl param
	if len(cfg.APIUrl) == 0 {
		cfg.APIUrl = APIFinancialModelingPrepURL
	}

	// Check set APIKey param
	if len(cfg.APIKey) == 0 {
		cfg.APIKey = apiDefaultKey
	}

	cfg.HTTPClient.SetHostURL(string(cfg.APIUrl))

	limiter := cfg.RateLimiter
	if limiter == nil {
		limiter = rate.NewLimiter(rate.Limit(math.Inf(0)), 1)
	}

	HTTPClient := &HTTPClient{
		client:      cfg.HTTPClient,
		apiKey:      cfg.APIKey,
		logger:      APIClient.Logger,
		rateLimiter: limiter,
	}

	if cfg.RetryCount != nil {
		HTTPClient.retryCount = cfg.RetryCount
	}

	if cfg.RetryWaitTime != nil {
		HTTPClient.retryWaitTime = cfg.RetryWaitTime
	}

	if HTTPClient.retryCount == nil || *HTTPClient.retryCount == 0 {
		retryCount := 1
		HTTPClient.retryCount = &retryCount
	}

	if HTTPClient.retryWaitTime == nil || *HTTPClient.retryWaitTime == 0 {
		retryWaitTime := 1 * time.Second
		HTTPClient.retryWaitTime = &retryWaitTime
	}

	APIClient.Stock = &Stock{Client: HTTPClient}
	APIClient.Form13F = &Form13F{Client: HTTPClient}
	APIClient.Forex = &Forex{Client: HTTPClient}
	APIClient.Crypto = &Crypto{Client: HTTPClient}
	APIClient.CompanyValuation = &CompanyValuation{Client: HTTPClient}
	APIClient.TechnicalIndicator = &TechnicalIndicator{Client: HTTPClient}
	APIClient.API = &API{Client: HTTPClient}
	APIClient.InsiderTrading = &InsiderTrading{Client: HTTPClient}
	APIClient.AlternativeData = &AlternativeData{Client: HTTPClient}
	APIClient.Economics = &Economics{Client: HTTPClient}

	return APIClient, nil
}

// Create new logger
func createNewLogger() (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder

	logger, err := cfg.Build()
	if err != nil {
		return nil, errors.Wrap(err, "Logger Error: init")
	}

	return logger, nil
}
