package dto

import "fmt"

type CongressSession struct {
	Chamber   string `json:"chamber"`
	Number    uint16 `json:"number"`
	Type      string `json:"type"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate,omitempty"`
}

func (c *CongressSession) ToString() string {
	return fmt.Sprintf("\t[%s] (%s - %s)\t%s\n", c.Type, c.StartDate, c.EndDate, c.Chamber)
}
