package projectservice

import (
	"context"
	"fmt"

	"github.com/aportela/doneo/internal/browser"
	"github.com/aportela/doneo/internal/domain"
	"github.com/aportela/doneo/internal/repositories/projectrepository"
)

type ProjectService interface {
	Add(ctx context.Context, Project domain.Project) error
	Update(ctx context.Context, Project domain.Project) error
	Delete(ctx context.Context, id string) error
	Get(ctx context.Context, id string) (domain.Project, error)
	Search(ctx context.Context, pager browser.Params, order browser.Order, filter domain.SearchProjectFilter) ([]domain.Project, browser.Result, error)
}

type projectService struct {
	repository projectrepository.ProjectRepository
}

func NewProjectService(repository projectrepository.ProjectRepository) ProjectService {
	return &projectService{repository: repository}
}

func (s *projectService) Add(ctx context.Context, project domain.Project) error {
	return s.repository.Add(ctx, projectrepository.DomainToDTO(project))
}

func (s *projectService) Update(ctx context.Context, project domain.Project) error {
	return s.repository.Update(ctx, projectrepository.DomainToDTO(project))
}

func (s *projectService) Delete(ctx context.Context, id string) error {
	return s.repository.Delete(ctx, id)
}

func (s *projectService) Get(ctx context.Context, id string) (domain.Project, error) {
	project, err := s.repository.Get(ctx, id)
	if err != nil {
		return projectrepository.DTOToDomain(project), fmt.Errorf("[ProjectService] failed to get project with ID %s: %w", id, err)
	}
	return projectrepository.DTOToDomain(project), nil
}

func (s *projectService) Search(ctx context.Context, pager browser.Params, order browser.Order, filter domain.SearchProjectFilter) ([]domain.Project, browser.Result, error) {
	projects, pagerResult, err := s.repository.Search(ctx, pager, order, projectrepository.DomainFilterToDTO(filter))
	if err != nil {
		return nil, browser.Result{}, fmt.Errorf("[ProjectService] failed to search projects: %w", err)
	}
	return projectrepository.DTOArrayToDomainArray(projects), pagerResult, nil
}
