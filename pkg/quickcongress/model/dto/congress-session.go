package dto

type CongressSession struct {
	Number    uint16 `json:"number"`
	Chamber   string `json:"chamber"`
	Type      string `json:"type"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}
