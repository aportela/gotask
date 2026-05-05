package projecttypehandler

import (
	"github.com/aportela/doneo/internal/domain"
)

func addRequestToProjectType(request addRequest) domain.ProjectType {
	return domain.ProjectType{
		Name:     request.Name,
		HexColor: request.HexColor,
	}
}

func updateRequestToProjectType(request updateRequest) domain.ProjectType {
	return domain.ProjectType{
		Name:     request.Name,
		HexColor: request.HexColor,
	}
}

func projectTypeToResponse(projectType domain.ProjectType) projectTypeResponse {
	return projectTypeResponse{
		ID:       projectType.ID,
		Name:     projectType.Name,
		HexColor: projectType.HexColor,
	}
}

func projectTypeToAddResponse(projectType domain.ProjectType) addResponse {
	return addResponse{
		ProjectType: projectTypeToResponse(projectType),
	}
}

func projectTypeToUpdateResponse(projectType domain.ProjectType) updateResponse {
	return updateResponse{
		ProjectType: projectTypeToResponse(projectType),
	}
}

func projectTypeToGetResponse(projectType domain.ProjectType) getResponse {
	return getResponse{
		ProjectType: projectTypeToResponse(projectType),
	}
}

func projectTypeArrayToResponse(projectTypes []domain.ProjectType) []projectTypeResponse {
	projectTypeResponses := []projectTypeResponse{}
	for _, projectType := range projectTypes {
		projectTypeResponses = append(projectTypeResponses, projectTypeToResponse(projectType))
	}
	return projectTypeResponses
}

func projectTypeArrayToSearchResponse(users []domain.ProjectType) searchResponse {
	return searchResponse{
		ProjectTypes: projectTypeArrayToResponse(users),
	}
}
