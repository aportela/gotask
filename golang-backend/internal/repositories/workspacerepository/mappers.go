package workspacerepository

import (
	"github.com/aportela/doneo/internal/domain"
)

func MapWorkspaceDomainToWorkspaceDTO(workspace domain.Workspace) workspaceDTO {
	return workspaceDTO{
		ID:          workspace.ID,
		Name:        workspace.Name,
		Description: workspace.Description,
		HexColor:    workspace.HexColor,
		CreatorId:   workspace.CreatedBy.ID,
		CreatorName: workspace.CreatedBy.Name,
		CreatedAt:   workspace.CreatedAt,
		UpdatedAt:   workspace.UpdatedAt,
		DeletedAt:   workspace.DeletedAt,
	}
}

func MapWorkspaceDTOToWorkspaceDomain(workspace workspaceDTO) domain.Workspace {
	return domain.Workspace{
		ID:          workspace.ID,
		Name:        workspace.Name,
		Description: workspace.Description,
		HexColor:    workspace.HexColor,
		CreatedBy: domain.UserBase{
			ID:   workspace.CreatorId,
			Name: workspace.CreatorName,
		},
		CreatedAt: workspace.CreatedAt,
		UpdatedAt: workspace.UpdatedAt,
		DeletedAt: workspace.DeletedAt,
	}
}

func MapWorkspaceArrayDTOToWorkspaceArrayDomain(workspaces []workspaceDTO) []domain.Workspace {
	results := []domain.Workspace{}
	for _, workspace := range workspaces {
		results = append(results, MapWorkspaceDTOToWorkspaceDomain(workspace))
	}
	return results
}
