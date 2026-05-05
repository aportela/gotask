package projectstatushandler

import (
	"github.com/aportela/doneo/internal/domain"
)

func addRequestToProjectStatus(request addRequest) domain.ProjectStatus {
	return domain.ProjectStatus{
		ID:       request.ID,
		Name:     request.Name,
		HexColor: request.HexColor,
		Index:    request.Index,
	}
}

func updateRequestToProjectStatus(request updateRequest) domain.ProjectStatus {
	return domain.ProjectStatus{
		ID:       request.ID,
		Name:     request.Name,
		HexColor: request.HexColor,
		Index:    request.Index,
	}
}

func mapProjectStatusDomainToResponse(projectStatus domain.ProjectStatus) projectStatusResponse {
	return projectStatusResponse{
		ID:       projectStatus.ID,
		Name:     projectStatus.Name,
		HexColor: projectStatus.HexColor,
		Index:    projectStatus.Index,
	}
}

func projectStatusToAddResponse(projectStatus domain.ProjectStatus) addResponse {
	return addResponse{
		ProjectStatus: mapProjectStatusDomainToResponse(projectStatus),
	}
}

func projectStatusToUpdateResponse(projectStatus domain.ProjectStatus) updateResponse {
	return updateResponse{
		ProjectStatus: mapProjectStatusDomainToResponse(projectStatus),
	}
}

func projectStatusToGetResponse(projectStatus domain.ProjectStatus) getResponse {
	return getResponse{
		ProjectStatus: mapProjectStatusDomainToResponse(projectStatus),
	}
}

func projectStatusArrayToResponse(projectStatuses []domain.ProjectStatus) []projectStatusResponse {
	projectStatusResponses := []projectStatusResponse{}
	for _, projectStatus := range projectStatuses {
		projectStatusResponses = append(projectStatusResponses, mapProjectStatusDomainToResponse(projectStatus))
	}
	return projectStatusResponses
}

func projectStatusArrayToSearchResponse(users []domain.ProjectStatus) searchResponse {
	return searchResponse{
		ProjectStatuses: projectStatusArrayToResponse(users),
	}
}
