package model

import (
	"encoding/json"

	"github.com/zfoteff/quick-congress/pkg/quickcongress/model/dto"
)

type BillsSuccessRes struct {
	Bills           []dto.Bill          `json:"bills"`
	Pagination      dto.Pagination      `json:"pagination"`
	RequestMetadata dto.RequestMetadata `json:"request"`
}

type BillsReqOptions struct {
	QueryString BillsRequestQuery
}

type BillsRequestQuery struct {
	Format string
	Limit  uint16
	Offset uint16
}

func (b *BillsSuccessRes) ToString() string {
	var congressessString string

	for _, bill := range b.Bills {
		congressessString += bill.ToString()
	}

	return congressessString
}

func (b *BillsSuccessRes) MarshalBinary() (data []byte, err error) {
	return json.Marshal(b)
}

func (b *BillsSuccessRes) UnmarshalBinary(data []byte) (err error) {
	return json.Unmarshal(data, b)
}
