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

var clientLogger = bin.NewLogger("Congress Client", "client.log")

type QuickCongressClientInterface interface {
	Exchange(req *http.Request, res interface{}) error
}

type QuickCongressClient struct {
	apiKey      string
	BaseURL     string
	httpClient  *http.Client
	redisClient *QuickCongressRedisClient
}

func NewQuickCongressClient() *QuickCongressClient {
	if goEnvErr := godotenv.Load(".env"); goEnvErr != nil {
		clientLogger.Error("Could not load application environment variables", goEnvErr)
	}

	return &QuickCongressClient{
		BaseURL: BaseURL + APIVersion,
		apiKey:  os.Getenv("LIBRARY_OF_CONGRESS_API_KEY"),
		httpClient: &http.Client{
			Transport: nil,
			Timeout:   time.Minute,
		},
		redisClient: NewRedisClient(),
	}
}

func (c *QuickCongressClient) GetAPIKey() string {
	return c.apiKey
}

func (c *QuickCongressClient) Exchange(req *http.Request, res interface{}) error {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "applicaption/json")

	response, err := c.httpClient.Do(req)
	if err != nil {
		congressClientLogger.Error("Error sending HTTP request to congress server", err)
		return err
	}

	defer response.Body.Close()

	// Unmarshall error into response if the status code is not 200
	if response.StatusCode != http.StatusOK {
		congressClientLogger.Warning(fmt.Sprintf("Status: %d", response.StatusCode))
		if err = json.NewDecoder(response.Body).Decode(&res); err == nil {
			congressClientLogger.Error("Could not unmarshall error response into response object", err)
		}

		return err
	}

	// Unmarshall into success response
	if err = json.NewDecoder(response.Body).Decode(res); err != nil {
		return err
	}

	congressClientLogger.Info(fmt.Sprintf("Status: %d. Recieved response from client in seconds", response.StatusCode))
	return nil
}
