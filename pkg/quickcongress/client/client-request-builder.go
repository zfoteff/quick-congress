package client

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

type HttpMethod string

const (
	Get  HttpMethod = "GET"
	Post HttpMethod = "POST"
)

type Builder interface {
	build() *http.Request
}

type RequestBuilder struct {
	baseURL    string
	apiKey     string
	apiVersion string
	path       string
	query      string
	method     HttpMethod
}

func NewRequestBuilder() RequestBuilder {
	builder := RequestBuilder{}
	builder.clear()
	return builder
}

func (b *RequestBuilder) clear() {
	b.baseURL = BaseURL
	b.apiKey = ""
	b.apiVersion = APIVersion
	b.path = ""
	b.query = ""
	b.method = Get
}

func (b *RequestBuilder) BaseUrl(base_url string) *RequestBuilder {
	b.baseURL = base_url
	return b
}

func (b *RequestBuilder) APIKey(api_key string) *RequestBuilder {
	b.apiKey = api_key
	return b
}

func (b *RequestBuilder) APIVersion(api_version string) *RequestBuilder {
	b.apiVersion = api_version
	return b
}

func (b *RequestBuilder) Path(path string) *RequestBuilder {
	b.path = path
	return b
}

func (b *RequestBuilder) QueryString(query_value string) *RequestBuilder {
	b.query = query_value
	return b
}

func (b *RequestBuilder) QueryStringFromMap(query_values map[string]string) *RequestBuilder {
	return b
}

func (b *RequestBuilder) Method(http_method HttpMethod) *RequestBuilder {
	b.method = http_method
	return b
}

func (b *RequestBuilder) build() *http.Request {
	request, err := http.NewRequestWithContext(
		context.TODO(),
		string(b.method),
		fmt.Sprintf("%s/%s/%s?api_key=%s&%s",
			b.baseURL,
			b.apiVersion,
			b.path,
			b.apiKey,
			b.query,
		),
		nil,
	)

	if err != nil {
		log.Fatal(err)
	}

	b.clear()
	return request
}
