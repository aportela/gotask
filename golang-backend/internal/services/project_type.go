package services

import (
	"context"

	"github.com/aportela/doneo/internal/domain"
	"github.com/aportela/doneo/internal/repositories/projecttyperepository"
)

type ProjectTypeService struct {
	repository *projecttyperepository.ProjectTypeRepository
}

func NewProjectTypeService(repository *projecttyperepository.ProjectTypeRepository) *ProjectTypeService {
	return &ProjectTypeService{
		repository: repository,
	}
}

func (s *ProjectTypeService) AddProjectType(ctx context.Context, project domain.ProjectType) error {
	return s.repository.Add(ctx, project)
}

func (s *ProjectTypeService) UpdateProjectType(ctx context.Context, project domain.ProjectType) error {
	return s.repository.Update(ctx, project)
}

func (s *ProjectTypeService) DeleteProjectType(ctx context.Context, id string) error {
	return s.repository.Delete(ctx, id)
}

func (s *ProjectTypeService) SearchProjectTypes(ctx context.Context) ([]domain.ProjectType, error) {
	return s.repository.Search(ctx)
}
