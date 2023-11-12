package dto

import (
	"fmt"
)

type Action struct {
	ActionDate string `json:"actionDate"`
	Text       string `json:"text"`
}

type Actions struct {
}

type Amendments struct {
}

type PolicyArea struct {
	name string `json:"policyArea"`
}

type Sponsors struct {
}

type Bill struct {
	Actions
	Congress                int     `json:"congress"`
	Number                  string  `json:"number"`
	LatestAction            *Action `json:"latestAction"`
	OriginChamber           string  `json:"originChamber"`
	OriginChamberCode       string  `json:"originChamberCode, omitempty"`
	Title                   string  `json:"title"`
	Type                    string  `json:"type"`
	UpdateDate              string  `json:"updateDate"`
	UpdateDateIncludingText string  `json:"updateDateIncludingText"`
}

func (b *Bill) ToString() string {
	var billString string

	billString += fmt.Sprintf("")
	return billString
}
