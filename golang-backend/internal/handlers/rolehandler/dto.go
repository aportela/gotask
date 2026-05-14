package rolehandler

import "github.com/aportela/doneo/internal/handlers"

type addRequest struct {
	Name string `json:"name"`
}

type updateRequest struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type searchRequest struct {
	Pager handlers.PagerRequest `json:"pager"`
	Order handlers.OrderRequest `json:"order"`
}

type roleResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type searchResponse struct {
	Roles []roleResponse         `json:"roles"`
	Pager handlers.PagerResponse `json:"pager"`
}
