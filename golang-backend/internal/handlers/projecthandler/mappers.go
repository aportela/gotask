package projecthandler

import (
	"github.com/aportela/doneo/internal/browser"
	"github.com/aportela/doneo/internal/domain"
	"github.com/aportela/doneo/internal/handlers"
	"github.com/aportela/doneo/internal/handlers/projectpriorityhandler"
	"github.com/aportela/doneo/internal/handlers/projectstatushandler"
	"github.com/aportela/doneo/internal/handlers/projecttypehandler"
	"github.com/aportela/doneo/internal/handlers/userhandler"
)

func addRequestToDomain(request addRequest) domain.Project {
	return domain.Project{
		ID:      request.ID,
		Key:     request.Key,
		Summary: request.Summary,
	}
}

func updateRequestToDomain(request updateRequest) domain.Project {
	return domain.Project{
		ID:      request.ID,
		Key:     request.Key,
		Summary: request.Summary,
	}
}

func DomainToResponse(project domain.Project) projectResponse {
	return projectResponse{
		ID:          project.ID,
		Key:         project.Key,
		Summary:     project.Summary,
		Description: *project.Description,
		CreatedBy:   userhandler.BaseDomainToBaseResponse(project.CreatedBy),
		CreatedAt:   project.CreatedAt,
		Type:        projecttypehandler.DomainToResponse(project.Type),
		Priority:    projectpriorityhandler.DomainToResponse(project.Priority),
		Status:      projectstatushandler.DomainToResponse(project.Status),
	}
}

func domainArrayToResponseArray(projects []domain.Project) []projectResponse {
	projectResponses := []projectResponse{}
	for _, project := range projects {
		projectResponses = append(projectResponses, DomainToResponse(project))
	}
	return projectResponses
}

func toSearchResponse(users []domain.Project, pager browser.Result) searchProjectsResponse {
	return searchProjectsResponse{
		Projects: domainArrayToResponseArray(users),
		Pager: handlers.PagerResponse{
			Enabled:      pager.ResultsPage > 0,
			CurrentPage:  pager.CurrentPage,
			ResultsPage:  pager.ResultsPage,
			TotalPages:   pager.TotalPages,
			TotalResults: pager.TotalResults,
		},
	}
}
