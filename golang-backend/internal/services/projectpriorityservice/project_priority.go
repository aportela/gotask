package projectpriorityservice

import (
	"context"
	"fmt"

	"github.com/aportela/doneo/internal/browser"
	"github.com/aportela/doneo/internal/domain"
	"github.com/aportela/doneo/internal/repositories/projectpriorityrepository"
)

type ProjectPriorityService interface {
	Add(ctx context.Context, projectPriority domain.ProjectPriority) error
	Update(ctx context.Context, projectPriority domain.ProjectPriority) error
	Delete(ctx context.Context, id string) error
	Get(ctx context.Context, id string) (domain.ProjectPriority, error)
	Search(ctx context.Context, pager browser.Params, order browser.Order, filter domain.SearchProjectPrioritiesFilter) ([]domain.ProjectPriority, browser.Result, error)
}

type projectPriorityService struct {
	repository projectpriorityrepository.ProjectPriorityRepository
}

func NewService(repository projectpriorityrepository.ProjectPriorityRepository) ProjectPriorityService {
	return &projectPriorityService{repository: repository}
}

func (service *projectPriorityService) Add(ctx context.Context, projectPriority domain.ProjectPriority) error {
	return service.repository.Add(ctx, projectpriorityrepository.DomainToDTO(projectPriority))
}

func (service *projectPriorityService) Update(ctx context.Context, projectPriority domain.ProjectPriority) error {
	return service.repository.Update(ctx, projectpriorityrepository.DomainToDTO(projectPriority))
}

func (service *projectPriorityService) Delete(ctx context.Context, id string) error {
	return service.repository.Delete(ctx, id)
}

func (service *projectPriorityService) Get(ctx context.Context, id string) (domain.ProjectPriority, error) {
	projectPriority, err := service.repository.Get(ctx, id)
	if err != nil {
		return projectpriorityrepository.DTOToDomain(projectPriority), fmt.Errorf("[ProjectPriorityService] failed to get project priority with ID %s: %w", id, err)
	}
	return projectpriorityrepository.DTOToDomain(projectPriority), nil
}

func (service *projectPriorityService) Search(ctx context.Context, pager browser.Params, order browser.Order, filter domain.SearchProjectPrioritiesFilter) ([]domain.ProjectPriority, browser.Result, error) {
	projectPriorities, pagerResult, err := service.repository.Search(ctx, pager, order, projectpriorityrepository.DomainFilterToDTO(filter))
	if err != nil {
		return nil, browser.Result{}, fmt.Errorf("[ProjectPriorityService] failed to search project priorities: %w", err)
	}
	return projectpriorityrepository.DTOArrayToDomainArray(projectPriorities), pagerResult, nil
}
