package client

import "net/http"
import "github.com/joho/godotenv"
import "log"
import "os"
import "time"
import "encoding/json"
import "github.com/zfoteff/quick-congress/pkg/quickcongress/model"

const (
	BaseURL    = "https://api.congress.gov"
	APIVersion = "v3"
)

type QuickCongressClientInterface interface {
	exchange(req *http.Request, res interface{}) error
	CreateQuickCongressClientRequest()
}

type QuickCongressClient struct {
	apiKey     string
	baseUrl    string
	httpClient *http.Client
}

func NewQuickCongressClient() *QuickCongressClient {
	if goEnvErr := godotenv.Load(".env"); goEnvErr != nil {
		log.Fatalf("Could not load application environment variables. Err: %s", goEnvErr)
	}

	return &QuickCongressClient{
		baseUrl: BaseURL + APIVersion,
		apiKey:  os.Getenv("LIBRARY_OF_CONGRESS_API_KEY"),
		httpClient: &http.Client{
			Transport: nil,
			Timeout:   time.Minute,
		},
	}
}

func (c *CongressClient) exchange(req *http.Request, res interface{}) error {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

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

	return nil
}
