package model

type CongressRes struct {
	Congresses []Congresses `json:"congresses"`
}

type Congresses struct {
	EndYear   string     `json:"endYear"`
	Name      string     `json:"name"`
	Sessions  []Sessions `json:"sessions"`
	StartYear string     `json:"startYear"`
}

type Sessions struct {
	Chamber   string `json:"chamber"`
	EndDate   string `json:"endDate"`
	Number    int    `json:"number"`
	StartDate string `json:"startDate"`
	Type      string `json:"type"`
}

type CongressReqOptions struct {
}
