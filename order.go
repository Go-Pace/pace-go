package pace

import (
	"net/http"

	"github.com/paingha/vider"
)

type CreateOrderRequest struct {
	SenderID              int32   `json:"senderId"`
	SenderName            string  `json:"senderName"`
	SenderEmail           string  `json:"senderEmail"`
	SenderPhone           string  `json:"senderPhone"`
	ReceiverName          string  `json:"receiverName"`
	ReceiverPhone         string  `json:"receiverPhone"`
	PickupAddress         string  `json:"pickupAddress"`
	PickupCoordinateLat   float64 `json:"pickupCoordinateLat"`
	PickupCoordinateLng   float64 `json:"pickupCoordinateLng"`
	DeliveryAddress       string  `json:"deliveryAddress"`
	DeliveryCoordinateLat float64 `json:"deliveryCoordinateLat"`
	DeliveryCoordinateLng float64 `json:"deliveryCoordinateLng"`
	PayOnDelivery         string  `json:"payOnDelivery"`
	PackageWeight         string  `json:"packageWeight"`
}

type CreateOrderResponse struct {
	SenderID              int32   `json:"senderId"`
	SenderName            string  `json:"senderName"`
	SenderEmail           string  `json:"senderEmail"`
	SenderPhone           string  `json:"senderPhone"`
	ReceiverName          string  `json:"receiverName"`
	ReceiverPhone         string  `json:"receiverPhone"`
	PickupAddress         string  `json:"pickupAddress"`
	PickupCoordinateLat   float64 `json:"pickupCoordinateLat"`
	PickupCoordinateLng   float64 `json:"pickupCoordinateLng"`
	DeliveryAddress       string  `json:"deliveryAddress"`
	DeliveryCoordinateLat float64 `json:"deliveryCoordinateLat"`
	DeliveryCoordinateLng float64 `json:"deliveryCoordinateLng"`
	PayOnDelivery         string  `json:"payOnDelivery"`
	PackageWeight         string  `json:"packageWeight"`
	TrackingCode          string  `json:"trackingCode"`
	Status                string  `json:"status"`
	Amount                float64 `json:"amount"`
	Distance              float64 `json:"distance"`
	OrderReference        string  `json:"orderReference"`
}

type TrackOrderRequest struct {
	TrackingCode string `json:"trackingCode"`
}

type TrackOrderResponse struct {
	OrderReference string `json:"orderReference"`
	TrackingCode   string `json:"trackingCode"`
	Status         string `json:"status"`
}

func (g *GoPaceRequest) CreateOrder(createOrderReq *CreateOrderRequest) (CreateOrderResponse, error) {
	resp, err := vider.Post(&vider.Request{
		URL:    urls["createorder"],
		Client: &http.Client{},
		Params: vider.Params{
			"headers": {
				"Content-Type":  "application/json",
				"Authorization": "Bearer " + g.Token,
			},
		},
		Data: createOrderReq,
	})
	if err != nil {
		return CreateOrderResponse{}, err
	}
	return resp.Body.(CreateOrderResponse), nil
}

func (g *GoPaceRequest) TrackOrder(trackOrderReq *TrackOrderRequest) (TrackOrderResponse, error) {
	resp, err := vider.Post(&vider.Request{
		URL:    urls["trackeorder"],
		Client: &http.Client{},
		Params: vider.Params{
			"headers": {
				"Content-Type":  "application/json",
				"Authorization": "Bearer " + g.Token,
			},
		},
		Data: trackOrderReq,
	})
	if err != nil {
		return TrackOrderResponse{}, err
	}
	return resp.Body.(TrackOrderResponse), nil
}
