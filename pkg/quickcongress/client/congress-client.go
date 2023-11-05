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
	return fmt.Sprintf("%s/%s%s?%s",
		req.URL.Scheme,
		req.URL.Host,
		req.URL.Path,
		req.URL.Query().Encode())
}

func (c *CongressClient) GetCongress(options *model.CongressReqOptions) (*model.CongressSuccessRes, error) {
	var congressNumber uint16 = 1
	res := model.CongressSuccessRes{}

	if options != nil {
		congressNumber = options.PathParameters.CongressNumber
	}

	req := NewRequestBuilder().BaseUrl(BaseURL).APIVersion(APIVersion).Path(
		fmt.Sprintf("congress/%d",
			congressNumber)).APIKey(c.client.GetAPIKey()).build()
	url := getRequestUrl(req)

	// Check if cache contains the
	if err := c.client.redisClient.GetCacheValue(url, &res); err == nil {
		return &res, nil
	}

	// Send request to client if key dne in cache
	if err := c.client.Exchange(req, &res); err != nil {
		congressClientLogger.Errorf("Error retrieving congress", err)
		return nil, err
	}

	// Set cache value if it didn't exist before, then return the response
	if err := c.client.redisClient.SetCacheValue(url, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *CongressClient) GetCongresses(options *model.CongressesReqOptions) (*model.CongressesSuccessRes, error) {
	var limit uint16 = 1
	var format string = "json"
	var offset uint16 = 0
	res := model.CongressesSuccessRes{}

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
	url := getRequestUrl(req)

	// Check if cache contains value for the request url before request is sent
	if err := c.client.redisClient.GetCacheValue(url, &res); err == nil {
		return &res, nil
	}

	// Send request to client if key dne in cache
	if err := c.client.Exchange(req, &res); err != nil {
		congressClientLogger.Errorf("Error retrieving congresses", err)
		return nil, err
	}

	// Set cache value if it didn't exist before, then return the response
	if err := c.client.redisClient.SetCacheValue(url, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
