package pace

import (
	"net/http"

	"github.com/paingha/vider"
)

const baseUrl = "https://client-api.gopace.xyz/api/v1"

var urls = map[string]string{
	"serverauth":        baseUrl + "/app/client-public-auth",
	"getfarecoordinate": baseUrl + "/fare/private-estimate",
	"getfareaddress":    baseUrl + "/fare/private-estimate",
	"trackorder":        baseUrl + "/order/track",
	"createorder":       baseUrl + "/order/private-create",
}

type AuthRequest struct {
	PrivateKey string `json:"private_key"`
}

type GoPaceRequest struct {
	Token string `json:"token"`
}

type InitResponse struct {
	Token string `json:"token"`
}

func (g *GoPaceRequest) Init(auth *AuthRequest) {
	resp, err := vider.Get(&vider.Request{
		URL:    urls["serverauth"],
		Client: &http.Client{},
		Params: vider.Params{
			"headers": {
				"Content-Type": "application/json",
			},
		},
		Data: auth,
	})
	if err != nil {
		panic(err)
	}
	g.Token = resp.Body.(InitResponse).Token
}
