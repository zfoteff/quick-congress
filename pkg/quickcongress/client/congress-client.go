package client

import (
	"context"
	"fmt"
	"github.com/zfoteff/quick-congress/pkg/quickcongress/model"
	"net/http"
)

type CongressClient struct {
	client *QuickCongressClient
}

type CongressClientSuccessRes struct {
	Code       uint16
	Congresses interface{} `json:"data"`
}

type CongressClientErrorRes struct {
	Code  uint16
	Error interface{} `json:"error"`
}

func NewCongressClient() *CongressClient {
	return &CongressClient{client: NewQuickCongressClient()}
}

func NewCongressClientFromSource(client *QuickCongressClient) *CongressClient {
	return &CongressClient{client: client}
}

func (c *CongressClient) GetCongress(ctx context.Context, options *model.CongressReqOptions) (*model.CongressSuccessRes, error) {
	var congressNumber uint16 = 1

	if options != nil {
		congressNumber = options.PathParameters.CongressNumber
	}

	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("%s/%s/congress/%d?api_key=%s",
			BaseURL,
			APIVersion,
			congressNumber,
		nil)

	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := model.CongressSuccessRes{}
	if err := c.exchange(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *CongressClient) GetCongresses(ctx context.Context, options *model.CongressesReqOptions) (*model.CongressesSuccessRes, error) {
	var limit uint16 = 1
	var format string = "json"
	var offset uint16 = 0

	if options != nil {
		limit = options.QueryString.Limit
		format = options.QueryString.Format
		offset = options.QueryString.Offset
	}

	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("%s/congress?limit=%d&offset=%d&format=%s&api_key=%s",
			c.baseURL,
			limit,
			offset,
			format,
			c.apiKey),
		nil)

	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := model.CongressesSuccessRes{}
	if err := c.exchange(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
