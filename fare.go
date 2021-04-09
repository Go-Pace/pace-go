package pace

import (
	"net/http"

	"github.com/paingha/vider"
)

type FareEstimateCoordinateRequest struct {
	PickupCoordinateLat   float64 `json:"pickup_lat"`
	PickupCoordinateLng   float64 `json:"pickup_long"`
	DeliveryCoordinateLat float64 `json:"destination_lat"`
	DeliveryCoordinateLng float64 `json:"destination_long"`
}

type FareEstimateAddressRequest struct {
	DeliveryAddress string `json:"delivery_address"`
	PickupAddress   string `json:"pickup_address"`
}

type FareEstimateResponse struct {
	Distance   float64 `json:"distance"`
	Fare       float64 `json:"data"`
	StatusCode string  `json:"statusCode"`
}

func (g *GoPaceRequest) GetEstimateCoordinate(fareReq *FareEstimateCoordinateRequest) (FareEstimateResponse, error) {
	resp, err := vider.Post(&vider.Request{
		URL:    urls["getfarecoordinate"],
		Client: &http.Client{},
		Params: vider.Params{
			"headers": {
				"Content-Type":  "application/json",
				"Authorization": "Bearer " + g.Token,
			},
		},
		Data: fareReq,
	})
	if err != nil {
		return FareEstimateResponse{}, err
	}
	return resp.Body.(FareEstimateResponse), nil
}

func (g *GoPaceRequest) GetEstimateAddress(fareReq *FareEstimateAddressRequest) (FareEstimateResponse, error) {
	resp, err := vider.Post(&vider.Request{
		URL:    urls["getfareaddress"],
		Client: &http.Client{},
		Params: vider.Params{
			"headers": {
				"Content-Type":  "application/json",
				"Authorization": "Bearer " + g.Token,
			},
		},
		Data: fareReq,
	})
	if err != nil {
		return FareEstimateResponse{}, err
	}
	return resp.Body.(FareEstimateResponse), nil
}
