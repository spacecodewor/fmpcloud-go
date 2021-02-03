package fmpcloud

import (
	"testing"
)

func TestCall(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.API.Call("test", nil)
	if err != nil {
		t.Fatal(err.Error())
	}
}
