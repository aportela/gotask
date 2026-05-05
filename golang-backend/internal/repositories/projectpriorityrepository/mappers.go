package projectpriorityrepository

import (
	"github.com/aportela/doneo/internal/domain"
)

func ProjectPriorityToDTO(projectPriority domain.ProjectPriority) projectPriorityDTO {
	return projectPriorityDTO{
		ID:       projectPriority.ID,
		Name:     projectPriority.Name,
		HexColor: projectPriority.HexColor,
		Index:    projectPriority.Index,
	}
}

func DTOToProjectPriority(projectPriority projectPriorityDTO) domain.ProjectPriority {
	return domain.ProjectPriority{
		ID:       projectPriority.ID,
		Name:     projectPriority.Name,
		HexColor: projectPriority.HexColor,
		Index:    projectPriority.Index,
	}
}

func ToProjectPriorityArray(projectPriorities []projectPriorityDTO) []domain.ProjectPriority {
	results := make([]domain.ProjectPriority, 0, len(projectPriorities))
	for _, projectPriority := range projectPriorities {
		results = append(results, DTOToProjectPriority(projectPriority))
	}
	return results
}
