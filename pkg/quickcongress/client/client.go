package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/zfoteff/quick-congress/bin"
)

const (
	BaseURL    = "https://api.congress.gov"
	APIVersion = "v3"
)

var clientLogger = bin.NewLogger("Client", "client.log")

type QuickCongressClientInterface interface {
	Exchange(url string, req *http.Request, res interface{}) error
}

type QuickCongressClient struct {
	apiKey      string
	BaseURL     string
	transport   *customTimingTransport
	httpClient  *http.Client
	redisClient *QuickCongressRedisCache
}

func NewQuickCongressClient() *QuickCongressClient {
	if goEnvErr := godotenv.Load(".env"); goEnvErr != nil {
		clientLogger.Error("Could not load application environment variables", goEnvErr)
	}

	tp := NewTransport()

	return &QuickCongressClient{
		BaseURL:   BaseURL + APIVersion,
		apiKey:    os.Getenv("LIBRARY_OF_CONGRESS_API_KEY"),
		transport: tp,
		httpClient: &http.Client{
			Transport: tp,
			Timeout:   time.Second * 10,
		},
		redisClient: NewQuickCongressRedisCache(),
	}
}

func (c *QuickCongressClient) GetAPIKey() string {
	return c.apiKey
}

func GetRequestUrl(req *http.Request) string {
	return fmt.Sprintf("%s/%s%s?%s",
		req.URL.Scheme,
		req.URL.Host,
		req.URL.Path,
		req.URL.Query().Encode())
}

func (c *QuickCongressClient) Exchange(url string, req *http.Request, res interface{}) error {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "applicaption/json")

	start := time.Now()
	// Check if cache contains the request and response before invoking client
	if err := c.redisClient.GetCacheValue(url, &res); err == nil {
		clientLogger.Info(fmt.Sprintf("Restored from cache in %s", time.Now().Sub(start).String()))
		return nil
	}

	response, err := c.httpClient.Do(req)
	if err != nil {
		clientLogger.Error("Error sending HTTP request to congress server", err)
		return err
	}

	defer response.Body.Close()

	// Unmarshall error into response if the status code is not 200
	if response.StatusCode != http.StatusOK {
		clientLogger.Warning(fmt.Sprintf("Status: %d", response.StatusCode))
		if err = json.NewDecoder(response.Body).Decode(&res); err == nil {
			clientLogger.Error("Could not unmarshall error response into response object", err)
		}

		return err
	}

	// Unmarshall into success response
	if err = json.NewDecoder(response.Body).Decode(res); err != nil {
		return err
	}

	// Set cache value if it didn't exist before, then return the response
	if err := c.redisClient.SetCacheValue(url, res); err != nil {
		return nil
	}

	clientLogger.Info(fmt.Sprintf("Status: %d. Received response from client in %s", response.StatusCode, time.Now().Sub(start).String()))
	return nil
}
