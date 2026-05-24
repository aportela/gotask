package projecttyperepository

import (
	"github.com/aportela/doneo/internal/domain"
)

func DomainToDTO(projectType domain.ProjectType) ProjectTypeDTO {
	return ProjectTypeDTO{
		ID:       projectType.ID,
		Name:     projectType.Name,
		HexColor: projectType.HexColor,
	}
}

func DTOToDomain(projectType ProjectTypeDTO) domain.ProjectType {
	return domain.ProjectType{
		ID:       projectType.ID,
		Name:     projectType.Name,
		HexColor: projectType.HexColor,
	}
}

func DTOArrayToDomainArray(projectTypes []ProjectTypeDTO) []domain.ProjectType {
	results := make([]domain.ProjectType, 0, len(projectTypes))
	for _, projectType := range projectTypes {
		results = append(results, DTOToDomain(projectType))
	}
	return results
}

func DomainFilterToDTO(filter domain.SearchProjectTypesFilter) searchFilterDTO {
	return searchFilterDTO{
		Name: filter.Name,
	}
}
