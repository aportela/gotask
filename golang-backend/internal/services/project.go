package services

import (
	"context"

	"github.com/aportela/gotask/internal/models"
	"github.com/aportela/gotask/internal/repositories"
)

type ProjectService struct {
	repository *repositories.ProjectRepository
}

func NewProjectService(repository *repositories.ProjectRepository) *ProjectService {
	return &ProjectService{
		repository: repository,
	}
}

func (s *ProjectService) GetProject(ctx context.Context, id string) (models.Project, error) {
	return s.repository.Get(ctx, id)
}

func (s *ProjectService) SearchProjects(ctx context.Context) ([]models.Project, error) {
	return s.repository.Search(ctx)
}
