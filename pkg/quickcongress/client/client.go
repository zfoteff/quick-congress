package client

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/zfoteff/quick-congress/pkg/quickcongress/model"
)

const (
	BaseURL    = "https://api.congress.gov"
	APIVersion = "v3"
)

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
		log.Fatalf("Could not load application environment variables. Err: %s", goEnvErr)
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
	exists, cachedResponse := c.redisClient.GetCacheValue(req.RequestURI)

	if exists {
		println(cachedResponse)
		return nil
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "applica ption/json")

	response, err := c.httpClient.Do(req)
	if err != nil {
		log.Printf("Error sending HTTP request to congress server. Err: %s", err)
		return err
	}

	defer response.Body.Close()

	// Unmarshall into error response
	if response.StatusCode != http.StatusOK {
		log.Printf("Status: %d. Could not unmarshall response into response object", response.StatusCode)
		var errRes model.CongressError
		if err = json.NewDecoder(response.Body).Decode(&errRes); err == nil {
			log.Printf("Could not unmarshall error response into error object")
			return err
		}

		return err
	}

	// Unmarshall into success response
	if err = json.NewDecoder(response.Body).Decode(res); err != nil {
		return err
	}

	c.redisClient.SetCacheValue(req.RequestURI, res)

	return nil
}
