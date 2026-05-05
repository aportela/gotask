package projectstatusrepository

import (
	"github.com/aportela/doneo/internal/domain"
)

func ProjectStatusToDTO(projectStatus domain.ProjectStatus) projectStatusDTO {
	return projectStatusDTO{
		ID:       projectStatus.ID,
		Name:     projectStatus.Name,
		HexColor: projectStatus.HexColor,
		Index:    projectStatus.Index,
	}
}

func DTOToProjectStatus(projectStatus projectStatusDTO) domain.ProjectStatus {
	return domain.ProjectStatus{
		ID:       projectStatus.ID,
		Name:     projectStatus.Name,
		HexColor: projectStatus.HexColor,
		Index:    projectStatus.Index,
	}
}

func ToProjectStatusArray(projectStatuses []projectStatusDTO) []domain.ProjectStatus {
	results := make([]domain.ProjectStatus, 0, len(projectStatuses))
	for _, projectStatus := range projectStatuses {
		results = append(results, DTOToProjectStatus(projectStatus))
	}
	return results
}
