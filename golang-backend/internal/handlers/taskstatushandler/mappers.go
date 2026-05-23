package taskstatushandler

import (
	"github.com/aportela/doneo/internal/browser"
	"github.com/aportela/doneo/internal/domain"
	"github.com/aportela/doneo/internal/handlers"
)

func addRequestToDomain(request addRequest) domain.TaskStatus {
	return domain.TaskStatus{
		ID:       request.ID,
		Name:     request.Name,
		HexColor: request.HexColor,
	}
}

func updateRequestToDomain(request updateRequest) domain.TaskStatus {
	return domain.TaskStatus{
		ID:       request.ID,
		Name:     request.Name,
		HexColor: request.HexColor,
	}
}

func domainToResponse(taskStatus domain.TaskStatus) taskStatusResponse {
	return taskStatusResponse{
		ID:       taskStatus.ID,
		Name:     taskStatus.Name,
		HexColor: taskStatus.HexColor,
	}
}

func domainArrayToResponseArray(taskStatuses []domain.TaskStatus) []taskStatusResponse {
	projectStatusResponses := []taskStatusResponse{}
	for _, projectStatus := range taskStatuses {
		projectStatusResponses = append(projectStatusResponses, domainToResponse(projectStatus))
	}
	return projectStatusResponses
}

func toSearchResponse(taskStatuses []domain.TaskStatus, pager browser.Result) searchResponse {
	return searchResponse{
		TaskStatuses: domainArrayToResponseArray(taskStatuses),
		Pager: handlers.PagerResponse{
			Enabled:      pager.ResultsPage > 0,
			CurrentPage:  pager.CurrentPage,
			ResultsPage:  pager.ResultsPage,
			TotalPages:   pager.TotalPages,
			TotalResults: pager.TotalResults,
		},
	}
}
