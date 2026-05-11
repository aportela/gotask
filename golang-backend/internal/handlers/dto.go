package handlers

type EmptyResponse struct {
}

type PagerRequest struct {
	CurrentPage int `json:"currentPage"`
	ResultsPage int `json:"resultsPage"`
}

type PagerResponse struct {
	Enabled      bool `json:"enabled"`
	CurrentPage  int  `json:"currentPage"`
	ResultsPage  int  `json:"resultsPage"`
	TotalPages   int  `json:"totalPages"`
	TotalResults int  `json:"totalResults"`
}

const (
	SortAsc  SortOrder = "ASC"
	SortDesc SortOrder = "DESC"
)

type SortOrder string

func (o SortOrder) IsValid() bool {
	return o == SortAsc || o == SortDesc
}

type OrderRequest struct {
	Field string    `json:"field"`
	Sort  SortOrder `json:"sort"`
}
