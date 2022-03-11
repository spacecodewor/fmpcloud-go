package fmpcloud

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/spacecodewor/fmpcloud-go/objects"
)

// Url const for request
const (
	urlAPIForm13FList         = "/v3/cik_list"
	urlAPIForm13FSearchByName = "/v3/cik-search/%s"
	urlAPIForm13FGetByCik     = "/v3/cik/%s"
	urlAPIForm13FGetThirteen  = "/v3/form-thirteen/%s"
	urlAPIForm13FCusipMapper  = "/v3/cusip/%s"
)

// Form13F client
type Form13F struct {
	Client *HTTPClient
}

// List - 13F List
func (f *Form13F) List() (fList []objects.Form, err error) {
	data, err := f.Client.Get(urlAPIForm13FList, nil)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data.Body(), &fList)
	if err != nil {
		return nil, err
	}

	return fList, nil
}

// SearchByName - 13F cik search by name
func (f *Form13F) SearchByName(name string) (fList []objects.Form, err error) {
	data, err := f.Client.Get(fmt.Sprintf(urlAPIForm13FSearchByName, name), nil)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data.Body(), &fList)
	if err != nil {
		return nil, err
	}

	return fList, nil
}

// GetCompanyByCIK - 13F get company name by cik
func (f *Form13F) GetCompanyByCIK(cik string) (cList []objects.Form, err error) {
	data, err := f.Client.Get(fmt.Sprintf(urlAPIForm13FGetByCik, cik), nil)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data.Body(), &cList)
	if err != nil {
		return nil, err
	}

	return cList, nil
}

// ThirteenList - 13F
func (f *Form13F) ThirteenList(cik string, date *time.Time) (fList []objects.Thirteen, err error) {
	reqParam := make(map[string]string)
	if date != nil {
		reqParam["date"] = date.Format("2006-01-02")
	}

	data, err := f.Client.Get(fmt.Sprintf(urlAPIForm13FGetThirteen, cik), reqParam)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data.Body(), &fList)
	if err != nil {
		return nil, err
	}

	return fList, nil
}

// CusipMapper - Cusip mapper
func (f *Form13F) CusipMapper(cusip string) (cList []objects.Cusip, err error) {
	data, err := f.Client.Get(fmt.Sprintf(urlAPIForm13FCusipMapper, cusip), nil)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data.Body(), &cList)
	if err != nil {
		return nil, err
	}

	return cList, nil
}
