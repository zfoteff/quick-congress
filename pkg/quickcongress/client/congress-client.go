package client

import (
	"fmt"

	"github.com/zfoteff/quick-congress/pkg/quickcongress/model"
)

type CongressClient struct {
	client *QuickCongressClient
}

func NewCongressClient() *CongressClient {
	return &CongressClient{client: NewQuickCongressClient()}
}

func NewCongressClientFromSource(client *QuickCongressClient) *CongressClient {
	return &CongressClient{client: client}
}

func (c *CongressClient) GetCongress(options *model.CongressReqOptions) (*model.CongressSuccessRes, error) {
	var congressNumber uint16 = 1

	if options != nil {
		congressNumber = options.PathParameters.CongressNumber
	}

	builder := NewRequestBuilder()
	req := builder.BaseUrl(BaseURL).APIVersion(APIVersion).Path(
		fmt.Sprintf("congress/%d",
			congressNumber)).APIKey(c.client.GetAPIKey()).build()

	res := model.CongressSuccessRes{}
	if err := c.client.Exchange(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *CongressClient) GetCongresses(options *model.CongressesReqOptions) (*model.CongressesSuccessRes, error) {
	var limit uint16 = 1
	var format string = "json"
	var offset uint16 = 0

	if options != nil {
		limit = options.QueryString.Limit
		format = options.QueryString.Format
		offset = options.QueryString.Offset
	}

	builder := NewRequestBuilder()
	req := builder.BaseUrl(BaseURL).APIVersion(APIVersion).Path("congress").QueryString(
		fmt.Sprintf("limit=%d&offset=%d&format=%s",
			limit,
			offset,
			format)).APIKey(c.client.GetAPIKey()).build()

	res := model.CongressesSuccessRes{}
	if err := c.client.Exchange(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
