package fmpcloud

import (
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// HTTPClient ...
type HTTPClient struct {
	logger        *zap.Logger
	client        *resty.Client
	apiKey        string
	retryCount    *int
	retryWaitTime *time.Duration
}

// Create new logger
func (h *HTTPClient) Get(endpoint string, data map[string]string) (response *resty.Response, err error) {
	if data == nil {
		data = make(map[string]string)
	}

	data["apikey"] = h.apiKey

	retries := 0
	for retries < *h.retryCount {
		response, err = h.client.R().
			SetQueryParams(data).
			Get(endpoint)

		if err != nil || response.StatusCode() != http.StatusOK {
			time.Sleep(*h.retryWaitTime)
			retries++

			// response is not valid when there is an error
			var errOrStatusCode zapcore.Field
			if err != nil {
				errOrStatusCode = zap.Error(err)
			} else {
				errOrStatusCode = zap.Int("statusCode", response.StatusCode())
			}

			h.logger.Info(
				"Retry request.",
				zap.Int("retries", retries),
				errOrStatusCode,
				zap.String("endpoint", endpoint),
				zap.Any("data", data),
			)

			continue
		}

		if err == nil {
			break
		}
	}

	return response, err
}
