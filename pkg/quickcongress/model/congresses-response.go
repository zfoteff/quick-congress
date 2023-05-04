package model

import "github.com/zfoteff/quick-congress/pkg/quickcongress/model/dto"

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
	Format string
	Limit  uint16
	Offset uint16
}

type CongressesReqOptions struct {
	QueryString CongressesReqQuery
}
