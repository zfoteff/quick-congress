package model

import (
	"encoding/json"

	"github.com/zfoteff/quick-congress/pkg/quickcongress/model/dto"
)

type BillSuccessRes struct {
	Bill dto.Bill `json:"bill"`
}

func (b *BillSuccessRes) MarshalBinary() (data []byte, err error) {
	return json.Marshal(b)
}

func 
