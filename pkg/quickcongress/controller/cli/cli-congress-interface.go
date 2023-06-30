package cli

import (
	"context"

	"github.com/zfoteff/quick-congress/pkg/quickcongress/client"
	"github.com/zfoteff/quick-congress/pkg/quickcongress/model"
)

func GetCurrentCongressSession(client *client.CongressClient, ctx context.Context) (current_session string) {
	request_query := &model.CongressesReqQuery{
		Format: "json",
		Limit:  1,
		Offset: 0,
	}

	response, error := client.GetCongresses(ctx, &model.CongressesReqOptions{QueryString: *request_query})

	if error != nil {
		panic(error)
	}

	return response.ToString()
}

func GetCongressSession(client *client.CongressClient, ctx context.Context, session uint16) (selected_session string) {
	request_query := &model.CongressReqPath{
		CongressNumber: session,
	}

	response, error := client.GetCongress(ctx, &model.CongressReqOptions{PathParameters: *request_query})

	if error != nil {
		panic(error)
	}

	return response.Congress.ToString()
}

func GetCongressSessions(client *client.CongressClient, ctx context.Context, sessions uint16, offset uint16) (session string) {
	request_query := &model.CongressesReqQuery{
		Format: "json",
		Limit:  sessions,
		Offset: 0,
	}

	response, error := client.GetCongresses(ctx, &model.CongressesReqOptions{QueryString: *request_query})

	if error != nil {
		panic(error)
	}

	return response.ToString()
}
