package projectpriorityservice

import (
	"context"
	"fmt"

	"github.com/aportela/doneo/internal/domain"
	"github.com/aportela/doneo/internal/repositories/projectpriorityrepository"
)

type ProjectPriorityService interface {
	AddProjectPriority(ctx context.Context, projectPriority domain.ProjectPriority) error
	UpdateProjectPriority(ctx context.Context, projectPriority domain.ProjectPriority) error
	DeleteProjectPriority(ctx context.Context, id string) error
	GetProjectPriority(ctx context.Context, id string) (domain.ProjectPriority, error)
	SearchProjectPriorities(ctx context.Context, filter domain.ProjectPriorityFilter) ([]domain.ProjectPriority, error)
}

type projectPriorityService struct {
	repository projectpriorityrepository.ProjectPriorityRepository
}

func NewProjectPriorityService(repository projectpriorityrepository.ProjectPriorityRepository) ProjectPriorityService {
	return &projectPriorityService{repository: repository}
}

func (s *projectPriorityService) AddProjectPriority(ctx context.Context, projectPriority domain.ProjectPriority) error {
	return s.repository.Add(ctx, projectpriorityrepository.MapProjectPriorityDomainToProjectPriorityDTO(projectPriority))
}

func (s *projectPriorityService) UpdateProjectPriority(ctx context.Context, projectPriority domain.ProjectPriority) error {
	return s.repository.Update(ctx, projectpriorityrepository.MapProjectPriorityDomainToProjectPriorityDTO(projectPriority))
}

func (s *projectPriorityService) DeleteProjectPriority(ctx context.Context, id string) error {
	return s.repository.Delete(ctx, id)
}

func (s *projectPriorityService) GetProjectPriority(ctx context.Context, id string) (domain.ProjectPriority, error) {
	projectPriority, err := s.repository.Get(ctx, id)
	if err != nil {
		return projectpriorityrepository.MapProjectPriorityDTOToProjectPriorityDomain(projectPriority), fmt.Errorf("[ProjectPriorityService] failed to get project priority with ID %s: %w", id, err)
	}
	return projectpriorityrepository.MapProjectPriorityDTOToProjectPriorityDomain(projectPriority), nil
}

func (s *projectPriorityService) SearchProjectPriorities(ctx context.Context, filter domain.ProjectPriorityFilter) ([]domain.ProjectPriority, error) {
	projectPriorities, err := s.repository.Search(ctx, projectpriorityrepository.MapProjectPriorityFilterDomainToProjectPriorityFilterDTO(filter))
	if err != nil {
		return nil, fmt.Errorf("[ProjectPriorityService] failed to search project priorities: %w", err)
	}
	return projectpriorityrepository.MapProjectPriorityArrayDTOToProjectPriorityArrayDomain(projectPriorities), nil
}
