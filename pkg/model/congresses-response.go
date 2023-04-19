package model

import "github.com/zfoteff/quick-congress/pkg/model/dto"

type CongressesSuccessRes struct {
	Congresses      []dto.Congress      `json:"congresses"`
	Pagination      dto.Pagination      `json:"pagination"`
	RequestMetadata dto.RequestMetadata `json:"request"`
}

type CongressesErrorRes struct {
	CongressErrors CongressError `json:"error"`
}

type CongressError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type CongressesReqQuery struct {
	Format string `json:"format"`
	Limit  uint16 `json:"limit"`
	Offset uint16 `json:"offset"`
	ApiKey string `json:"api_key"`
}

type CongressesReqOptions struct {
	QueryString CongressesReqQuery
}
