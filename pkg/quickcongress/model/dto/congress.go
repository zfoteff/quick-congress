package dto

import "fmt"

type Congress struct {
	Name       string            `json:"name"`
	Number     uint16            `json:"number"`
	Sessions   []CongressSession `json:"sessions"`
	StartYear  string            `json:"startYear"`
	UpdateDate string            `json:"updateDate,omitempty"`
	EndYear    string            `json:"endYear,omitempty"`
	Url        string            `json:"url,omitempty"`
}

func (c *Congress) ToString() string {
	var congressString, sessionString string

	for _, session := range c.Sessions {
		sessionString += session.ToString()
	}

	congressString += fmt.Sprintf("%s (%s - %s):\n%s", c.Name, c.StartYear, c.EndYear, sessionString)
	return congressString
}
