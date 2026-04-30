package workspacerepository

import (
	"context"
	"database/sql"

	"github.com/aportela/doneo/internal/database"
	"github.com/aportela/doneo/internal/utils"
)

type WorkspaceRepository interface {
	Add(ctx context.Context, workspace workspaceDTO) error
	Search(ctx context.Context) ([]workspaceDTO, error)
}

type workspaceRepository struct {
	database database.Database
}

func NewWorkspaceRepository(database database.Database) WorkspaceRepository {
	return &workspaceRepository{database: database}
}

func (workspaceRepository *workspaceRepository) Add(ctx context.Context, workspace workspaceDTO) error {

	_, err := workspaceRepository.database.ExecContext(
		ctx,
		`
            INSERT INTO workspaces (id, name, description, creator_id, created_at, updated_at)
			VALUES (?, ?, ?, ?, ?, NULL)
        `,
		workspace.ID,
		workspace.Name,
		workspace.Description,
		workspace.CreatorId,
		workspace.CreatedAt,
	)
	return err
}

func (workspaceRepository *workspaceRepository) Search(ctx context.Context) ([]workspaceDTO, error) {
	rows, err := workspaceRepository.database.QueryContext(
		ctx,
		`
			SELECT
				W.id, W.name, W.description, W.creator_id, U.name AS creator_name, W.created_at, W.updated_at
			FROM workspaces W
			INNER JOIN users U ON U.id = W.creator_id
        `,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var workspaces []workspaceDTO
	for rows.Next() {
		var workspace workspaceDTO
		var description *string
		var updatedAt sql.NullInt64
		if err := rows.Scan(&workspace.ID, &workspace.Name, &description, &workspace.CreatorId, &workspace.CreatorName, &workspace.CreatedAt, &updatedAt); err != nil {
			return nil, err
		}

		workspace.UpdatedAt = utils.SQLInt64Ptr(updatedAt)
		if description != nil {
			workspace.Description = description
		}
		workspaces = append(workspaces, workspace)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return workspaces, nil
}
