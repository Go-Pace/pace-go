package pace_test

import (
	"os"
	"testing"

	"github.com/gopace/pace-go"
)

var paceRequest pace.GoPaceRequest

func TestMain(m *testing.M) {
	paceRequest.Init(&pace.AuthRequest{
		PrivateKey: "hello",
	})
	os.Exit(m.Run())
}

func TestGetEstimateCoordinate(t *testing.T) {
	response, err := paceRequest.GetEstimateCoordinate(&pace.FareEstimateCoordinateRequest{})
	if err != nil {
		t.Error(err)
	}
	if response.Distance == 0 {
		t.Error("Distance cannot be empty")
	}
	if response.Fare == 0 {
		t.Error("Fare cannot be empty")
	}
	if response.StatusCode != "200" {
		t.Error("HTTP request was not successful")
	}
}

func TestGetEstimateAddress(t *testing.T) {
	response, err := paceRequest.GetEstimateAddress(&pace.FareEstimateAddressRequest{})
	if err != nil {
		t.Error(err)
	}
	if response.Distance == 0 {
		t.Error("Distance cannot be empty")
	}
	if response.Fare == 0 {
		t.Error("Fare cannot be empty")
	}
	if response.StatusCode != "200" {
		t.Error("HTTP request was not successful")
	}
}
