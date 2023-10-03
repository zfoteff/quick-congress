package model

import (
	"encoding/json"

	"github.com/zfoteff/quick-congress/pkg/quickcongress/model/dto"
)

type CongressesSuccessRes struct {
	Congresses      []dto.Congress      `json:"congresses"`
	Pagination      dto.Pagination      `json:"pagination"`
	RequestMetadata dto.RequestMetadata `json:"request"`
}

func (c *CongressesSuccessRes) ToString() string {
	var congressessString string

	for _, congress := range c.Congresses {
		congressessString += congress.ToString()
	}

	return congressessString
}

func (c *CongressesSuccessRes) MarshalBinary() ([]byte, error) {
	return json.Marshal(c)
}

func (c *CongressesSuccessRes) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, c)
}

type CongressesErrorRes struct {
	CongressErrors CongressError `json:"error"`
}

type CongressError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type CongressesReqOptions struct {
	QueryString CongressesReqQuery
}

type CongressesReqQuery struct {
	Format string
	Limit  uint16
	Offset uint16
}
