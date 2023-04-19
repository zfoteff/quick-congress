package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/zfoteff/quick-congress/pkg/model"
)

const (
	BaseURLV3  = "https://api.congress.gov"
	ApiVersion = "/v3/"
)

type CongressClient struct {
	apiKey     string
	baseURL    string
	httpClient *http.Client
}

type CongressClientSuccessRes struct {
	Code uint16      `json:"code"`
	Data interface{} `json:"data"`
}

func (c *CongressClient) NewClient(apiKey string) *CongressClient {
	return &CongressClient{
		baseURL: BaseURLV3 + ApiVersion,
		apiKey:  apiKey,
		httpClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

func (c *CongressClient) sendRequest(req *http.Request, v interface{}) error {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	// Unmarshall into error response
	if res.StatusCode != http.StatusOK {
		var errRes model.CongressesErrorRes
		if err = json.NewDecoder(res.Body).Decode(&errRes); err == nil {
			return errors.New(errRes.CongressErrors.Message)
		}

		return fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}

	// Unmarshall into success response
	if err = json.NewDecoder(res.Body).Decode(&CongressClientSuccessRes{Data: v}); err != nil {
		return err
	}

	return nil
}

func (c *CongressClient) GetCongresses(ctx context.Context, options *model.CongressesReqOptions, query *model.CongressesReqQuery) (*model.CongressesSuccessRes, error) {
	var limit uint16 = 1
	var format string = "json"
	var offset uint16 = 0

	if options != nil {
		limit = options.QueryString.Limit
		format = options.QueryString.Format
		offset = options.QueryString.Offset
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/congress?limit=%d&offset=%d&format=%s%api_key=%s", limit, offset, format, c.apiKey), nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := model.CongressesSuccessRes{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
