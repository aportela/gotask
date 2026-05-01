package projectstatushandler

import (
	"github.com/aportela/doneo/internal/domain"
)

func mapAddProjectStatusRequestToProjectStatusDomain(request addProjectStatusRequest) domain.ProjectStatus {
	return domain.ProjectStatus{
		ID:   request.ID,
		Name: request.Name,
	}
}

func mapUpdateProjectStatusRequestToProjectStatusDomain(request updateProjectStatusRequest) domain.ProjectStatus {
	return domain.ProjectStatus{
		ID:   request.ID,
		Name: request.Name,
	}
}

func mapProjectStatusDomainToProjectStatusResponse(projectStatus domain.ProjectStatus) projectStatusResponse {
	return projectStatusResponse{
		ID:   projectStatus.ID,
		Name: projectStatus.Name,
	}
}

func mapProjectStatusDomainToAddProjectStatusResponse(projectStatus domain.ProjectStatus) addProjectStatusResponse {
	return addProjectStatusResponse{
		ProjectStatus: mapProjectStatusDomainToProjectStatusResponse(projectStatus),
	}
}

func mapProjectStatusDomainToUpdateProjectStatusResponse(projectStatus domain.ProjectStatus) updateProjectStatusResponse {
	return updateProjectStatusResponse{
		ProjectStatus: mapProjectStatusDomainToProjectStatusResponse(projectStatus),
	}
}

func mapProjectStatusDomainToGetProjectStatusResponse(projectStatus domain.ProjectStatus) getProjectStatusResponse {
	return getProjectStatusResponse{
		ProjectStatus: mapProjectStatusDomainToProjectStatusResponse(projectStatus),
	}
}

func mapProjectStatusArrayDomainToProjectStatusArrayResponse(projectPriorities []domain.ProjectStatus) []projectStatusResponse {
	var projectStatusResponses []projectStatusResponse
	for _, projectStatus := range projectPriorities {
		projectStatusResponses = append(projectStatusResponses, mapProjectStatusDomainToProjectStatusResponse(projectStatus))
	}
	return projectStatusResponses
}

func mapProjectStatusArrayDomainToSearchProjectStatussResponse(users []domain.ProjectStatus) searchProjectStatussResponse {
	return searchProjectStatussResponse{
		ProjectStatuses: mapProjectStatusArrayDomainToProjectStatusArrayResponse(users),
	}
}
