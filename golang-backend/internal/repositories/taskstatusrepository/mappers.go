package taskstatusrepository

import (
	"github.com/aportela/doneo/internal/domain"
)

func DomainToDTO(taskStatus domain.TaskStatus) taskStatusDTO {
	return taskStatusDTO{
		ID:       taskStatus.ID,
		Name:     taskStatus.Name,
		HexColor: taskStatus.HexColor,
	}
}

func DTOToDomain(taskStatus taskStatusDTO) domain.TaskStatus {
	return domain.TaskStatus{
		ID:       taskStatus.ID,
		Name:     taskStatus.Name,
		HexColor: taskStatus.HexColor,
	}
}

func DTOArrayToDomainArray(taskStatuses []taskStatusDTO) []domain.TaskStatus {
	results := make([]domain.TaskStatus, 0, len(taskStatuses))
	for _, projectStatus := range taskStatuses {
		results = append(results, DTOToDomain(projectStatus))
	}
	return results
}

func DomainFilterToDTO(filter domain.SearchTaskStatusesFilter) searchFilterDTO {
	return searchFilterDTO{
		Name: filter.Name,
	}
}
