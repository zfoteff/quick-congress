package model

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

err = godotenv.Load(".env")

const (
	ApiKey = os.Getenv("API_KEY")
)

type CongressRes struct {
	congresses []Congresses `json:"congresses"`
}

type Congresses struct {
	endYear   string     `json:"endYear"`
	name      string     `json:"name"`
	sessions  []Sessions `json:"sessions"`
	startYear string     `json:"startYear"`
}

type Sessions struct {
	chamber   string `json:"chamber"`
	endDate   string `json:"endDate"`
	number    int    `json:"number"`
	startDate string `json:"startDate"`
	Type      string `json:"type"`
}

type CongressReqQuery struct {
	format string `json:"format"`
	limit  int    `json:"limit"`
	offset int    `json:"offset"`
	apiKey string `json:"api_key"`
}

type CongressReqParameters struct {
	congressNumber int
}

type CongressReqOptions struct {
	PathParams CongressReqParameters
	QueryString CongressReqQuery
}

func (c *http.Client) GetCongresses(ctx context.Context, options *CongressReqOptions, query *CongressReqQuery) (*CongressRes, error) {
	limit := 1
	format := "json"
	offset := 0

	if options == nil {

	}

	if options != nil {
		limit = options.QueryString.limit
		format = options.QueryString.format
		offset = options.QueryString.offset
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/congress?limit=%d&offset=%d&format=%s%api_key=%s"), nil)
	if err != nill {
		return nil, err
	}

	return 
}