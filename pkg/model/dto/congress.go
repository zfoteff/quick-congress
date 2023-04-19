package dto

type Congress struct {
	EndYear   string            `json:"endYear"`
	Name      string            `json:"name"`
	Sessions  []CongressSession `json:"sessions"`
	StartYear string            `json:"startYear"`
	Url       string            `json:"url"`
}
