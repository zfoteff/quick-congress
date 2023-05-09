package dto

type Congress struct {
	Name       string            `json:"name"`
	Number     uint16            `json:"number"`
	Sessions   []CongressSession `json:"sessions"`
	StartYear  string            `json:"startYear"`
	UpdateDate string            `json:"updateDate,omitempty"`
	EndYear    string            `json:"endYear,omitempty"`
	Url        string            `json:"url,omitempty"`
}
