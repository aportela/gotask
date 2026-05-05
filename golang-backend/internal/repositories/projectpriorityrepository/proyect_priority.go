package projectpriorityrepository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/aportela/doneo/internal/database"
	"github.com/aportela/doneo/internal/domain"
)

type ProjectPriorityRepository interface {
	Add(ctx context.Context, projectPriority projectPriorityDTO) error
	Update(ctx context.Context, projectPriority projectPriorityDTO) error
	Get(ctx context.Context, id string) (projectPriorityDTO, error)
	Delete(ctx context.Context, id string) error
	Search(ctx context.Context) ([]projectPriorityDTO, error)
}

type projectPriorityRepository struct {
	database database.Database
}

func NewProjectPriorityRepository(database database.Database) ProjectPriorityRepository {
	return &projectPriorityRepository{database: database}
}

func (projectPriorityRepository *projectPriorityRepository) Add(ctx context.Context, projectPriority projectPriorityDTO) error {
	_, err := projectPriorityRepository.database.ExecContext(
		ctx,
		`
            INSERT INTO project_priorities (id, name, item_hex_color, item_index)
			VALUES (?, ?, ?, ?)
        `,
		projectPriority.ID,
		projectPriority.Name,
		projectPriority.HexColor,
		projectPriority.Index,
	)
	return err
}

func (projectPriorityRepository *projectPriorityRepository) Update(ctx context.Context, projectPriority projectPriorityDTO) error {
	_, err := projectPriorityRepository.database.ExecContext(
		ctx,
		`
            UPDATE project_priorities SET
				name = ?,
				item_hex_color = ?,
				item_index = ?
			WHERE id = ?
        `,
		projectPriority.Name,
		projectPriority.HexColor,
		projectPriority.Index,
		projectPriority.ID,
	)
	return err
}

func (projectPriorityRepository *projectPriorityRepository) Delete(ctx context.Context, id string) error {
	_, err := projectPriorityRepository.database.ExecContext(
		ctx,
		`
            DELETE FROM project_priorities
			WHERE id = ?
        `,
		id,
	)
	return err
}

func (projectPriorityRepository *projectPriorityRepository) Get(ctx context.Context, id string) (projectPriorityDTO, error) {
	var projectPriority projectPriorityDTO
	err := projectPriorityRepository.database.QueryRowContext(
		ctx,
		`
            SELECT
                PP.id, PP.name, PP.item_hex_color,  PP.item_index
            FROM project_priorities PP
            WHERE PP.id = ?
        `,
		id).Scan(&projectPriority.ID, &projectPriority.Name, &projectPriority.HexColor, &projectPriority.Index)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return projectPriorityDTO{}, domain.ErrNotFound
		}
		return projectPriorityDTO{}, err
	}
	return projectPriority, err
}

func (projectPriorityRepository *projectPriorityRepository) Search(ctx context.Context) ([]projectPriorityDTO, error) {
	rows, err := projectPriorityRepository.database.QueryContext(
		ctx,
		`
			SELECT
				PP.id, PP.name, PP.item_hex_color,  PP.item_index
			FROM project_priorities PP
			ORDER BY PP.item_index, PP.name
        `,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var proyecPriorities []projectPriorityDTO
	for rows.Next() {
		var projectPriority projectPriorityDTO
		if err := rows.Scan(
			&projectPriority.ID, &projectPriority.Name, &projectPriority.HexColor, &projectPriority.Index,
		); err != nil {
			return nil, err
		}
		proyecPriorities = append(proyecPriorities, projectPriority)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return proyecPriorities, nil
}
