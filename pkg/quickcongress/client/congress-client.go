package client

import (
	"fmt"
	"net/http"

	"github.com/zfoteff/quick-congress/bin"
	"github.com/zfoteff/quick-congress/pkg/quickcongress/model"
)

var congressClientLogger = bin.NewLogger("Congress Client", "congress-client.log")

type CongressClient struct {
	client *QuickCongressClient
}

func NewCongressClient() *CongressClient {
	return &CongressClient{client: NewQuickCongressClient()}
}

func NewCongressClientFromSource(client *QuickCongressClient) *CongressClient {
	return &CongressClient{client: client}
}

func getRequestUrl(req *http.Request) string {
	return fmt.Sprintf("%s/%s%s?%s", req.URL.Scheme, req.URL.Host, req.URL.Path, req.URL.Query().Encode())
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

	req := NewRequestBuilder().BaseUrl(BaseURL).APIVersion(APIVersion).Path("congress").QueryString(
		fmt.Sprintf("limit=%d&offset=%d&format=%s",
			limit,
			offset,
			format)).APIKey(c.client.GetAPIKey()).build()

	res := model.CongressesSuccessRes{}

	url := getRequestUrl(req)
	exists, value := c.client.redisClient.GetCacheValue(url)
	if exists {
		congressClientLogger.Debug(value)
		return &res, nil
	}

	if err := c.client.Exchange(req, &res); err != nil {
		congressClientLogger.Error("Error retrieving congress", err)
		return nil, err
	}

	if err := c.client.redisClient.SetCacheValue(url, res); err != nil {
		return nil, err
	}

	return &res, nil
}
