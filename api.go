package fmpcloud

// API client
type API struct {
	Client *HTTPClient
}

// Call ...
func (c *API) Call(endpoint string, requestParam map[string]string) (resp []byte, err error) {
	data, err := c.Client.Get(endpoint, requestParam)
	if err != nil {
		return nil, err
	}

	return data.Body(), nil
}
