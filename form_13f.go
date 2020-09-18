package fmpcloud

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/spacecodewor/fmpcloud-go/objects"
)

// Url const for request
const (
	urlAPIForm13FList         = "/cik_list"
	urlAPIForm13FSearchByName = "/cik-search/%s"
	urlAPIForm13FGetByCik     = "/cik/%s"
	urlAPIForm13FGetThirteen  = "/form-thirteen/%s"
	urlAPIForm13FCusipMapper  = "/cusip/%s"
)

// Form13F client
type Form13F struct {
	Client *resty.Client
	url    string
	apiKey string
}

// List - 13F List
func (f *Form13F) List() (fList []objects.Form, err error) {
	data, err := f.Client.R().
		SetQueryParams(map[string]string{"apikey": f.apiKey}).
		Get(f.url + urlAPIForm13FList)

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
	data, err := f.Client.R().
		SetQueryParams(map[string]string{"apikey": f.apiKey}).
		Get(f.url + fmt.Sprintf(urlAPIForm13FSearchByName, name))

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
	data, err := f.Client.R().
		SetQueryParams(map[string]string{"apikey": f.apiKey}).
		Get(f.url + fmt.Sprintf(urlAPIForm13FGetByCik, cik))

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
	reqParam := map[string]string{"apikey": c.apiKey}
	if date != nil {
		reqParam["date"] = date.Format("2006-01-02")
	}

	data, err := f.Client.R().
		SetQueryParams(reqParam).
		Get(f.url + fmt.Sprintf(urlAPIForm13FGetThirteen, cik))

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
	data, err := f.Client.R().
		SetQueryParams(map[string]string{"apikey": f.apiKey}).
		Get(f.url + fmt.Sprintf(urlAPIForm13FCusipMapper, cusip))

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data.Body(), &cList)
	if err != nil {
		return nil, err
	}

	return cList, nil
}
