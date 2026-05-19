package projecttyperepository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/aportela/doneo/internal/browser"
	"github.com/aportela/doneo/internal/database"
	"github.com/aportela/doneo/internal/domain"
)

type ProjectTypeRepository interface {
	Add(ctx context.Context, projectType projectTypeDTO) error
	Update(ctx context.Context, projectType projectTypeDTO) error
	Get(ctx context.Context, id string) (projectTypeDTO, error)
	Delete(ctx context.Context, id string) error
	Search(ctx context.Context, pager browser.Params, order browser.Order, filter searchFilterDTO) ([]projectTypeDTO, browser.Result, error)
}

type projectTypeRepository struct {
	database database.Database
}

func NewProjectTypeRepository(database database.Database) ProjectTypeRepository {
	return &projectTypeRepository{database: database}
}

func (projectTypeRepository *projectTypeRepository) Add(ctx context.Context, projectType projectTypeDTO) error {
	_, err := projectTypeRepository.database.ExecContext(
		ctx,
		`
            INSERT INTO project_types (id, name, item_hex_color)
			VALUES (?, ?, ?)
        `,
		projectType.ID,
		projectType.Name,
		projectType.HexColor,
	)
	return err
}

func (projectTypeRepository *projectTypeRepository) Update(ctx context.Context, projectType projectTypeDTO) error {
	_, err := projectTypeRepository.database.ExecContext(
		ctx,
		`
            UPDATE project_types SET
				name = ?,
				item_hex_color = ?
			WHERE id = ?
        `,
		projectType.Name,
		projectType.HexColor,
		projectType.ID,
	)
	return err
}

func (projectTypeRepository *projectTypeRepository) Delete(ctx context.Context, id string) error {
	_, err := projectTypeRepository.database.ExecContext(
		ctx,
		`
            DELETE FROM project_types
			WHERE id = ?
        `,
		id,
	)
	return err
}

func (projectTypeRepository *projectTypeRepository) Get(ctx context.Context, id string) (projectTypeDTO, error) {
	var projectType projectTypeDTO
	err := projectTypeRepository.database.QueryRowContext(
		ctx,
		`
            SELECT
                PT.id, PT.name, PT.item_hex_color
            FROM project_types PT
            WHERE PT.id = ?
        `,
		id).Scan(&projectType.ID, &projectType.Name, &projectType.HexColor)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return projectTypeDTO{}, domain.NotFoundError
		}
		return projectTypeDTO{}, err
	}
	return projectType, err
}

func (projectTypeRepository *projectTypeRepository) Search(ctx context.Context, pager browser.Params, order browser.Order, filter searchFilterDTO) ([]projectTypeDTO, browser.Result, error) {
	var filterArgs []any
	var queryArgs []any
	sqlQuery := `
			SELECT
				PT.id, PT.name, PT.item_hex_color
			FROM project_types PT
    `
	var field string
	switch order.Field {
	case "name":
		field = "PT.name COLLATE NOCASE"
	default:
		field = "PT.name COLLATE NOCASE"
	}
	var sort string
	switch order.Sort {
	case "DESC":
		sort = "DESC"
	case "ASC":
		sort = "ASC"
	default:
		sort = "ASC"
	}
	sqlOrder := fmt.Sprintf(" ORDER BY %s %s ", field, sort)
	sqlWhere := ""
	var sqlWhereConditions []string
	if filter.Name != nil {
		sqlWhereConditions = append(sqlWhereConditions, "PT.name LIKE ?")
		filterArgs = append(filterArgs, "%"+*filter.Name+"%")
	}

	if len(sqlWhereConditions) > 0 {
		sqlWhere = " WHERE " + strings.Join(sqlWhereConditions, " AND ")
	}
	queryArgs = append(queryArgs, filterArgs...)
	var sqlLimit string
	if pager.Enabled() {
		sqlLimit = " LIMIT ? OFFSET ? "
		queryArgs = append(queryArgs, pager.Limit(), pager.Offset())
	} else {
		sqlLimit = ""
	}
	sqlQuery = fmt.Sprintf("%s %s %s %s ", sqlQuery, sqlWhere, sqlOrder, sqlLimit)
	rows, err := projectTypeRepository.database.QueryContext(ctx, sqlQuery, queryArgs...)
	if err != nil {
		return nil, browser.Result{}, err
	}
	defer rows.Close()
	projectTypes := make([]projectTypeDTO, 0)
	for rows.Next() {
		var projectType projectTypeDTO
		if err := rows.Scan(
			&projectType.ID, &projectType.Name, &projectType.HexColor,
		); err != nil {
			return nil, browser.Result{}, err
		}
		projectTypes = append(projectTypes, projectType)
	}
	if err := rows.Err(); err != nil {
		return nil, browser.Result{}, err
	}

	var totalResults int

	if pager.Enabled() {
		sqlCountQuery := `
			SELECT
				COUNT(*) AS total_project_types
			FROM project_types PT
		`
		sqlCountQuery = fmt.Sprintf("%s %s", sqlCountQuery, sqlWhere)
		err = projectTypeRepository.database.QueryRowContext(
			ctx,
			sqlCountQuery,
			filterArgs...,
		).Scan(&totalResults)

		if err != nil {
			return nil, browser.Result{}, err
		}
	} else {
		totalResults = len(projectTypes)
	}

	return projectTypes, browser.NewResult(pager, totalResults), nil
}
