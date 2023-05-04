package dto

type Pagination struct {
	Count   uint16 `json:"count"`
	NextUrl string `json:"next"`
}

type RequestMetadata struct {
	ContentType string `json:"contentType"`
	Format      string `json:"format"`
}
