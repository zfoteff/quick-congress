package dto

type CongressSession struct {
	Chamber   string `json:"chamber"`
	EndDate   string `json:"endDate"`
	Number    uint16 `json:"number"`
	StartDate string `json:"startDate"`
	Type      string `json:"type"`
}
