package provider

import (
	"github.com/zfoteff/quick-congress/pkg/quickcongress/client"
)

/*
 * Creates & formats requests to be sent to the client
 */
type BillProvider struct {
	client client.BillClient
}

func NewBillProvider(client *client.BillClient) BillProvider {
	return BillProvider{client: *client}
}
