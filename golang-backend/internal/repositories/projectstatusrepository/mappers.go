package projectstatusrepository

import (
	"github.com/aportela/doneo/internal/domain"
)

func DomainToDTO(projectStatus domain.ProjectStatus) ProjectStatusDTO {
	return ProjectStatusDTO{
		ID:       projectStatus.ID,
		Name:     projectStatus.Name,
		HexColor: projectStatus.HexColor,
	}
}

func DTOToDomain(projectStatus ProjectStatusDTO) domain.ProjectStatus {
	return domain.ProjectStatus{
		ID:       projectStatus.ID,
		Name:     projectStatus.Name,
		HexColor: projectStatus.HexColor,
	}
}

func DTOArrayToDomainArray(projectStatuses []ProjectStatusDTO) []domain.ProjectStatus {
	results := make([]domain.ProjectStatus, 0, len(projectStatuses))
	for _, projectStatus := range projectStatuses {
		results = append(results, DTOToDomain(projectStatus))
	}
	return results
}

func DomainFilterToDTO(filter domain.SearchProjectStatusesFilter) searchFilterDTO {
	return searchFilterDTO{
		Name: filter.Name,
	}
}
