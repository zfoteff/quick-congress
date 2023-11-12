package model

import (
	"encoding/json"
	"github.com/zfoteff/quick-congress/pkg/quickcongress/model/dto"
)

type CongressSuccessRes struct {
	Congress dto.Congress `json:"congress"`
}

func (c *CongressSuccessRes) MarshalBinary() (data []byte, err error) {
	return json.Marshal(c)
}

func (c *CongressSuccessRes) UnmarshalBinary(data []byte) (err error) {
	return json.Unmarshal(data, c)
}

type CongressReqPath struct {
	CongressNumber uint16
}

type CongressReqOptions struct {
	PathParameters CongressReqPath
}
