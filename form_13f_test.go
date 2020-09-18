package fmpcloud

import "testing"

func TestForm13FList(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.Form13F.List()
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestForm13FSearchByName(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.Form13F.SearchByName("BlackRock")
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestForm13FGetCompanyByCIK(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.Form13F.GetCompanyByCIK("0001067983")
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestForm13FThirteenList(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.Form13F.ThirteenList("0001067983")
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestForm13FCusipMapper(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.Form13F.CusipMapper("57636Q104")
	if err != nil {
		t.Fatal(err.Error())
	}
}
