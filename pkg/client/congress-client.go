package client

import (
	"net/http"
	"time"
)

const (
	BaseURLV3  = "https://api.congress.gov"
	ApiVersion = "/v3/"
)

type Client struct {
	apiKey     string
	baseURL    string
	httpClient *http.Client
}

type ClientResponse struct {
	code int         `json:"code"`
	data interface{} `json:"data"`
}

func NewClient(apiKey string) *Client {
	return &Client{
		baseURL: BaseURLV3 + ApiVersion,
		apiKey:  apiKey,
		httpClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

func GetFullCongressSessions() {

}
