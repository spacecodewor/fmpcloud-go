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
	Logger  *zap.Logger
	Version string
	APIKey  string
	APIUrl  APIUrl
	Debug   bool
	Timeout int
}

// APIClient ...
type APIClient struct {
	Stock              *Stock
	Forex              *Forex
	Form13F            *Form13F
	Crypto             *Crypto
	CompanyValuation   *CompanyValuation
	TechnicalIndicator *TechnicalIndicator
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
	restClient := resty.New()
	restClient.SetDebug(APIClient.Debug)
	restClient.SetTimeout(time.Duration(cfg.Timeout) * time.Second)

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

	url := fmt.Sprintf(string(cfg.APIUrl), cfg.Version)
	APIClient.Stock = &Stock{
		Client: restClient,
		url:    url,
		apiKey: cfg.APIKey,
	}

	APIClient.Form13F = &Form13F{
		Client: restClient,
		url:    url,
		apiKey: cfg.APIKey,
	}

	APIClient.Forex = &Forex{
		Client: restClient,
		url:    url,
		apiKey: cfg.APIKey,
	}

	APIClient.Crypto = &Crypto{
		Client: restClient,
		url:    url,
		apiKey: cfg.APIKey,
	}

	APIClient.CompanyValuation = &CompanyValuation{
		Client: restClient,
		url:    url,
		apiKey: cfg.APIKey,
	}

	APIClient.TechnicalIndicator = &TechnicalIndicator{
		Client: restClient,
		url:    url,
		apiKey: cfg.APIKey,
	}

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
