package cli

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/zfoteff/quick-congress/pkg/quickcongress/client"
	"github.com/zfoteff/quick-congress/pkg/quickcongress/model"
)

// Get current Congress session from Congress endpoint, using a preset request.
func GetCurrentCongressSession() (current_session string) {
	goEnvErr := godotenv.Load(".env")
	client := client.NewCongressClient(os.Getenv("LIBRARY_OF_CONGRESS_API_KEY"))

	if goEnvErr != nil {
		log.Fatalf("Some error occured. Err: %s", goEnvErr)
	}

	request_query := &model.CongressesReqQuery{
		Format: "json",
		Limit:  1,
		Offset: 0,
	}

	response, error := client.GetCongresses(context.TODO(), &model.CongressesReqOptions{QueryString: *request_query})

	if error != nil {
		panic(error)
	}

	return response.ToString()
}

// Get Congress session by number from Congress endpoint.
func GetCongressSession(session uint16) (selected_session string) {
	goEnvErr := godotenv.Load(".env")
	client := client.NewCongressClient(os.Getenv("LIBRARY_OF_CONGRESS_API_KEY"))

	if goEnvErr != nil {
		log.Fatalf("Some error occured. Err: %s", goEnvErr)
	}

	request_query := &model.CongressReqPath{
		CongressNumber: session,
	}

	response, error := client.GetCongress(context.TODO(), &model.CongressReqOptions{PathParameters: *request_query})

	if error != nil {
		panic(error)
	}

	return response.Congress.ToString()
}

// Get Congress sessions by limit and offset from Congress endpoint.
func GetCongressSessions(sessions uint16, offset uint16) (session string) {
	goEnvErr := godotenv.Load(".env")
	client := client.NewCongressClient(os.Getenv("LIBRARY_OF_CONGRESS_API_KEY"))

	if goEnvErr != nil {
		log.Fatalf("Some error occured. Err: %s", goEnvErr)
	}

	request_query := &model.CongressesReqQuery{
		Format: "json",
		Limit:  sessions,
		Offset: 0,
	}

	response, error := client.GetCongresses(context.TODO(), &model.CongressesReqOptions{QueryString: *request_query})

	if error != nil {
		panic(error)
	}

	return response.ToString()
}
