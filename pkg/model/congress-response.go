package model

import "github.com/zfoteff/quick-congress/pkg/model/dto"

type CongressSuccessRes struct {
	Congress dto.Congress `json:"congress"`
}

type CongressReqPath struct {
	CongressNumber uint16
}

type CongressReqOptions struct {
	PathParameters CongressReqPath
}
