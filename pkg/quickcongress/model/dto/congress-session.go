package dto

type CongressSession struct {
	Chamber   string `json:"chamber"`
	Number    uint16 `json:"number"`
	Type      string `json:"type"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate,omitempty"`
}
