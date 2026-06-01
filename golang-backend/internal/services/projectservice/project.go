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

func NewService(repository projectrepository.ProjectRepository) ProjectService {
	return &projectService{repository: repository}
}

func (service *projectService) Add(ctx context.Context, project domain.Project) error {
	return service.repository.Add(ctx, projectrepository.DomainToDTO(project))
}

func (service *projectService) Update(ctx context.Context, project domain.Project) error {
	return service.repository.Update(ctx, projectrepository.DomainToDTO(project))
}

func (service *projectService) Delete(ctx context.Context, id string) error {
	return service.repository.Delete(ctx, id)
}

func (service *projectService) Get(ctx context.Context, id string) (domain.Project, error) {
	project, err := service.repository.Get(ctx, id)
	if err != nil {
		return projectrepository.DTOToDomain(project), fmt.Errorf("[ProjectService] failed to get project with ID %s: %w", id, err)
	}
	return projectrepository.DTOToDomain(project), nil
}

func (service *projectService) Search(ctx context.Context, pager browser.Params, order browser.Order, filter domain.SearchProjectFilter) ([]domain.Project, browser.Result, error) {
	projects, pagerResult, err := service.repository.Search(ctx, pager, order, projectrepository.DomainFilterToDTO(filter))
	if err != nil {
		return nil, browser.Result{}, fmt.Errorf("[ProjectService] failed to search projects: %w", err)
	}
	return projectrepository.DTOArrayToDomainArray(projects), pagerResult, nil
}
