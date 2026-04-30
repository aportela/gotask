package workspaceservice

import (
	"context"
	"fmt"

	"github.com/aportela/doneo/internal/domain"
	"github.com/aportela/doneo/internal/repositories/workspacerepository"
	"github.com/aportela/doneo/internal/utils"
)

type WorkspaceService interface {
	AddWorkspace(ctx context.Context, user domain.Workspace) error
	SearchWorkspaces(ctx context.Context) ([]domain.Workspace, error)
}

type workspaceService struct {
	repository workspacerepository.WorkspaceRepository
}

func NewWorkspaceService(repository workspacerepository.WorkspaceRepository) WorkspaceService {
	return &workspaceService{
		repository: repository,
	}
}

func (s *workspaceService) AddWorkspace(ctx context.Context, workspace domain.Workspace) error {
	workspace.CreatedAt = utils.CurrentMSTimestamp()
	if err := s.repository.Add(ctx, workspacerepository.MapWorkspaceDomainToWorkspaceDTO(workspace)); err != nil {
		return fmt.Errorf("[WorkspaceService] failed to add workspace with ID %s: %w", workspace.ID, err)
	}
	return nil
}

func (s *workspaceService) SearchWorkspaces(ctx context.Context) ([]domain.Workspace, error) {
	repositories, err := s.repository.Search(ctx)
	if err != nil {
		return nil, fmt.Errorf("[WorkspaceService] failed to search workspaces: %w", err)
	}
	return workspacerepository.MapWorkspaceArrayDTOToWorkspaceArrayDomain(repositories), nil
}
