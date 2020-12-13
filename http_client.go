package fmpcloud

import (
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
)

// HTTPClient ...
type HTTPClient struct {
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

		if response.StatusCode() != http.StatusOK {
			time.Sleep(*h.retryWaitTime)
			retries++

			continue
		}

		if err == nil {
			break
		}
	}

	return response, err
}
