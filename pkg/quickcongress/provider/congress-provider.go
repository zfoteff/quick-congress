package provider

import (
	"log"

	"github.com/zfoteff/quick-congress/pkg/quickcongress/client"
	"github.com/zfoteff/quick-congress/pkg/quickcongress/model"
)

/*
*Creates & formats requests to be sent to the client
 */
type CongressProvider struct {
	client client.CongressClient
}

func NewCongressProvider(client *client.CongressClient) CongressProvider {
	return CongressProvider{client: *client}
}

func (p *CongressProvider) GetCurrentCongress() model.CongressesSuccessRes {
	var request *model.CongressesReqOptions = &model.CongressesReqOptions{
		QueryString: model.CongressesReqQuery{
			Format: "json",
			Limit:  1,
			Offset: 0,
		}}

	response, error := p.client.GetCongresses(request)

	if error != nil {
		log.Fatal(error)
	}

	return *response
}

func (p *CongressProvider) GetCongress(session uint16) model.CongressSuccessRes {
	var request *model.CongressReqOptions = &model.CongressReqOptions{
		PathParameters: model.CongressReqPath{
			CongressNumber: session,
		},
	}

	response, error := p.client.GetCongress(request)

	if error != nil {
		log.Fatal(error)
	}

	return *response
}

func (p *CongressProvider) GetCongresses(limit uint16, offset uint16) model.CongressesSuccessRes {
	var request *model.CongressesReqOptions = &model.CongressesReqOptions{
		QueryString: model.CongressesReqQuery{
			Format: "json",
			Limit:  limit,
			Offset: offset,
		},
	}

	response, error := p.client.GetCongresses(request)

	if error != nil {
		log.Fatal(error)
	}

	return *response
}
