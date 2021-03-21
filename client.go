package fmpcloud

import (
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// APIUrl type for api url
type APIUrl string

// Config for create new API client
type Config struct {
	Logger        *zap.Logger
	HTTPClient    *resty.Client
	Version       string
	APIKey        string
	APIUrl        APIUrl
	Debug         bool
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
	API                *API
	Logger             *zap.Logger
	Debug              bool
}

// Core params
const (
	APIFmpcloudURL              APIUrl = "https://fmpcloud.io/api/%s"
	APIFinancialModelingPrepURL APIUrl = "https://financialmodelingprep.com/api/%s"
	apiDefaultVersion                  = "v3"
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

	// Check set Version param
	if len(cfg.Version) == 0 {
		cfg.Version = apiDefaultVersion
	}

	// Check set APIKey param
	if len(cfg.APIKey) == 0 {
		cfg.APIKey = apiDefaultKey
	}

	cfg.HTTPClient.SetHostURL(fmt.Sprintf(string(cfg.APIUrl), cfg.Version))

	HTTPClient := &HTTPClient{
		client: cfg.HTTPClient,
		apiKey: cfg.APIKey,
		logger: APIClient.Logger,
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

	// Only for API v4
	APIClient.InsiderTrading = &InsiderTrading{Client: HTTPClient}
	APIClient.AlternativeData = &AlternativeData{Client: HTTPClient}

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
