package taskstatusservice

import (
	"context"
	"fmt"

	"github.com/aportela/doneo/internal/browser"
	"github.com/aportela/doneo/internal/domain"
	"github.com/aportela/doneo/internal/repositories/taskstatusrepository"
)

type TaskStatusService interface {
	Add(ctx context.Context, projectStatus domain.TaskStatus) error
	Update(ctx context.Context, projectStatus domain.TaskStatus) error
	Delete(ctx context.Context, id string) error
	Get(ctx context.Context, id string) (domain.TaskStatus, error)
	Search(ctx context.Context, pager browser.Params, order browser.Order, filter domain.SearchTaskStatusesFilter) ([]domain.TaskStatus, browser.Result, error)
}

type taskStatusService struct {
	repository taskstatusrepository.ProjectStatusRepository
}

func NewService(repository taskstatusrepository.ProjectStatusRepository) TaskStatusService {
	return &taskStatusService{repository: repository}
}

func (service *taskStatusService) Add(ctx context.Context, taskStatus domain.TaskStatus) error {
	return service.repository.Add(ctx, taskstatusrepository.DomainToDTO(taskStatus))
}

func (service *taskStatusService) Update(ctx context.Context, taskStatus domain.TaskStatus) error {
	return service.repository.Update(ctx, taskstatusrepository.DomainToDTO(taskStatus))
}

func (service *taskStatusService) Delete(ctx context.Context, id string) error {
	return service.repository.Delete(ctx, id)
}

func (service *taskStatusService) Get(ctx context.Context, id string) (domain.TaskStatus, error) {
	taskStatus, err := service.repository.Get(ctx, id)
	if err != nil {
		return taskstatusrepository.DTOToDomain(taskStatus), fmt.Errorf("[TaskStatusService] failed to get task status with ID %s: %w", id, err)
	}
	return taskstatusrepository.DTOToDomain(taskStatus), nil
}

func (service *taskStatusService) Search(ctx context.Context, pager browser.Params, order browser.Order, filter domain.SearchTaskStatusesFilter) ([]domain.TaskStatus, browser.Result, error) {
	taskStatuses, pagerResult, err := service.repository.Search(ctx, pager, order, taskstatusrepository.DomainFilterToDTO(filter))
	if err != nil {
		return nil, browser.Result{}, fmt.Errorf("[TaskStatusService] failed to search task statuses: %w", err)
	}
	return taskstatusrepository.DTOArrayToDomainArray(taskStatuses), pagerResult, nil
}
