package fmpcloud

import (
	"testing"

	"github.com/spacecodewor/fmpcloud-go/objects"
)

func TestIndicators(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.TechnicalIndicator.Indicators(objects.RequestIndicators{
		Resolution: objects.TechnicalIndicatorResolutionDaily,
		Indicator:  objects.TechnicalIndicatorTypeSMA,
		Timeperiod: 10,
		Symbol:     "AAPL",
	})
	if err != nil {
		t.Fatal(err.Error())
	}
}
