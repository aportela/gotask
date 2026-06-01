package projectstatusservice

import (
	"context"
	"fmt"

	"github.com/aportela/doneo/internal/browser"
	"github.com/aportela/doneo/internal/domain"
	"github.com/aportela/doneo/internal/repositories/projectstatusrepository"
)

type ProjectStatusService interface {
	Add(ctx context.Context, projectStatus domain.ProjectStatus) error
	Update(ctx context.Context, projectStatus domain.ProjectStatus) error
	Delete(ctx context.Context, id string) error
	Get(ctx context.Context, id string) (domain.ProjectStatus, error)
	Search(ctx context.Context, pager browser.Params, order browser.Order, filter domain.SearchProjectStatusesFilter) ([]domain.ProjectStatus, browser.Result, error)
}

type projectStatusService struct {
	repository projectstatusrepository.ProjectStatusRepository
}

func NewService(repository projectstatusrepository.ProjectStatusRepository) ProjectStatusService {
	return &projectStatusService{repository: repository}
}

func (service *projectStatusService) Add(ctx context.Context, projectStatus domain.ProjectStatus) error {
	return service.repository.Add(ctx, projectstatusrepository.DomainToDTO(projectStatus))
}

func (service *projectStatusService) Update(ctx context.Context, projectStatus domain.ProjectStatus) error {
	return service.repository.Update(ctx, projectstatusrepository.DomainToDTO(projectStatus))
}

func (service *projectStatusService) Delete(ctx context.Context, id string) error {
	return service.repository.Delete(ctx, id)
}

func (service *projectStatusService) Get(ctx context.Context, id string) (domain.ProjectStatus, error) {
	projectStatus, err := service.repository.Get(ctx, id)
	if err != nil {
		return projectstatusrepository.DTOToDomain(projectStatus), fmt.Errorf("[ProjectStatusService] failed to get project status with ID %s: %w", id, err)
	}
	return projectstatusrepository.DTOToDomain(projectStatus), nil
}

func (service *projectStatusService) Search(ctx context.Context, pager browser.Params, order browser.Order, filter domain.SearchProjectStatusesFilter) ([]domain.ProjectStatus, browser.Result, error) {
	projectStatuses, pagerResult, err := service.repository.Search(ctx, pager, order, projectstatusrepository.DomainFilterToDTO(filter))
	if err != nil {
		return nil, browser.Result{}, fmt.Errorf("[ProjectStatusService] failed to search project statuses: %w", err)
	}
	return projectstatusrepository.DTOArrayToDomainArray(projectStatuses), pagerResult, nil
}
