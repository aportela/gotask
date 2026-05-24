package projecthandler

import (
	"github.com/aportela/doneo/internal/handlers"
	"github.com/aportela/doneo/internal/handlers/projectpriorityhandler"
	"github.com/aportela/doneo/internal/handlers/projectstatushandler"
	"github.com/aportela/doneo/internal/handlers/projecttypehandler"
	"github.com/aportela/doneo/internal/handlers/userhandler"
)

type addRequest struct {
	ID      string `json:"id"`
	Key     string `json:"key"`
	Summary string `json:"summary"`
}

type updateRequest struct {
	ID      string `json:"id"`
	Key     string `json:"key"`
	Summary string `json:"summary"`
}

type filterRequest struct {
	Key *string `json:"key"`
}

type searchRequest struct {
	Pager  handlers.PagerRequest `json:"pager"`
	Order  handlers.OrderRequest `json:"order"`
	Filter *filterRequest        `json:"filter"`
}

type projectResponse struct {
	ID          string                                         `json:"id"`
	Key         string                                         `json:"key"`
	Summary     string                                         `json:"summary"`
	Description string                                         `json:"description"`
	CreatedBy   userhandler.UserBaseResponse                   `json:"createdBy"`
	CreatedAt   int64                                          `json:"createdAt"`
	Type        projecttypehandler.ProjectTypeResponse         `json:"type"`
	Priority    projectpriorityhandler.ProjectPriorityResponse `json:"priority"`
	Status      projectstatushandler.ProjectStatusResponse     `json:"status"`
}

type searchProjectsResponse struct {
	Projects []projectResponse      `json:"projects"`
	Pager    handlers.PagerResponse `json:"pager"`
}
