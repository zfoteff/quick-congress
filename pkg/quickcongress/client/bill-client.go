package client

import ()

type BillClient struct {
	client *QuickCongressClient
}

func NewBillClient() *BillClient {
	return &BillClient{client: NewQuickCongressClient()}
}

func NewBillClientFromSource(client *QuickCongressClient) *BillClient {
	return &BillClient{client: client}
}
