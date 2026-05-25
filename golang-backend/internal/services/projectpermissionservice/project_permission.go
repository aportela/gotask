package projectpermissionservice

import (
	"context"
	"fmt"

	"github.com/aportela/doneo/internal/domain"
	"github.com/aportela/doneo/internal/repositories/projectpermissionrepository"
)

type ProjectPermissionService interface {
	Add(ctx context.Context, permissionId string, projectId string, userId string, roleId string) error
	Delete(ctx context.Context, permissionId string) error
	Search(ctx context.Context, projectId string) ([]domain.ProjectPermission, error)
}

type projectPermissionService struct {
	repository projectpermissionrepository.ProjectPermissionRepository
}

func NewProjectPermissionService(repository projectpermissionrepository.ProjectPermissionRepository) ProjectPermissionService {
	return &projectPermissionService{repository: repository}
}

func (s *projectPermissionService) Add(ctx context.Context, permissionId string, projectId string, userId string, roleId string) error {
	return s.repository.Add(ctx, permissionId, projectId, userId, roleId)
}

func (s *projectPermissionService) Delete(ctx context.Context, permissionId string) error {
	return s.repository.Delete(ctx, permissionId)
}

func (s *projectPermissionService) Search(ctx context.Context, projectId string) ([]domain.ProjectPermission, error) {
	projectPermissions, err := s.repository.Search(ctx, projectId)
	if err != nil {
		return nil, fmt.Errorf("[ProjectTypeService] failed to get project permissions: %w", err)
	}
	return projectpermissionrepository.DTOArrayToDomainArray(projectPermissions), nil
}
