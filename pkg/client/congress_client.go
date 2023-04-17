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
	BaseURL    string
	apiKey     string
	HTTPClient *http.Client
}

type clientResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

func NewClient(appKey string) *Client {
	return &Client{
		BaseURL: BaseURLV3 + ApiVersion,
		apiKey:  appKey,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

func GetFullCongressSessions() {

}
