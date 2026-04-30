package projecttyperepository

import (
	"github.com/aportela/doneo/internal/domain"
)

func MapProyectTypeDomainToProyectTypeDTO(projectType domain.ProjectType) proyectTypeDTO {
	return proyectTypeDTO{
		ID:   projectType.ID,
		Name: projectType.Name,
	}
}

func MapProjectTypeDTOToProjectTypeDomain(projectType proyectTypeDTO) domain.ProjectType {
	return domain.ProjectType{
		ID:   projectType.ID,
		Name: projectType.Name,
	}
}

func MapProjectTypeArrayDTOToProjectTypeArrayDomain(projectTypes []proyectTypeDTO) []domain.ProjectType {
	var results []domain.ProjectType
	for _, projectType := range projectTypes {
		results = append(results, MapProjectTypeDTOToProjectTypeDomain(projectType))
	}
	return results
}
