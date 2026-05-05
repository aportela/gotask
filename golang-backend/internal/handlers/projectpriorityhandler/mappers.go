package projectpriorityhandler

import (
	"github.com/aportela/doneo/internal/domain"
)

func addRequestToProjectPriority(request addRequest) domain.ProjectPriority {
	return domain.ProjectPriority{
		ID:       request.ID,
		Name:     request.Name,
		HexColor: request.HexColor,
		Index:    request.Index,
	}
}

func updateRequestToProjectPriority(request updateRequest) domain.ProjectPriority {
	return domain.ProjectPriority{
		ID:       request.ID,
		Name:     request.Name,
		HexColor: request.HexColor,
		Index:    request.Index,
	}
}

func projectPriorityToResponse(projectPriority domain.ProjectPriority) projectPriorityResponse {
	return projectPriorityResponse{
		ID:       projectPriority.ID,
		Name:     projectPriority.Name,
		HexColor: projectPriority.HexColor,
		Index:    projectPriority.Index,
	}
}

func projectPriorityToAddResponse(projectPriority domain.ProjectPriority) addResponse {
	return addResponse{
		ProjectPriority: projectPriorityToResponse(projectPriority),
	}
}

func projectPriorityToUpdateResponse(projectPriority domain.ProjectPriority) updateResponse {
	return updateResponse{
		ProjectPriority: projectPriorityToResponse(projectPriority),
	}
}

func projectPriorityToGetResponse(projectPriority domain.ProjectPriority) getResponse {
	return getResponse{
		ProjectPriority: projectPriorityToResponse(projectPriority),
	}
}

func projectPriorityArrayToResponse(projectPriorities []domain.ProjectPriority) []projectPriorityResponse {
	projectPriorityResponses := []projectPriorityResponse{}
	for _, projectPriority := range projectPriorities {
		projectPriorityResponses = append(projectPriorityResponses, projectPriorityToResponse(projectPriority))
	}
	return projectPriorityResponses
}

func projectPriorityArrayToSearchResponse(projectPriorities []domain.ProjectPriority) searchResponse {
	return searchResponse{
		ProjectPriorities: projectPriorityArrayToResponse(projectPriorities),
	}
}
