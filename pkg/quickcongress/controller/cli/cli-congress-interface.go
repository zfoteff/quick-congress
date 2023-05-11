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

	return string(response.ToString())
}
