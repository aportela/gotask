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
	"modernc.org/sqlite"
	sqlite3 "modernc.org/sqlite/lib"
)

type ProjectTypeRepository interface {
	Add(ctx context.Context, projectType ProjectTypeDTO) error
	Update(ctx context.Context, projectType ProjectTypeDTO) error
	Get(ctx context.Context, id string) (ProjectTypeDTO, error)
	Delete(ctx context.Context, id string) error
	Search(ctx context.Context, pager browser.Params, order browser.Order, filter searchFilterDTO) ([]ProjectTypeDTO, browser.Result, error)
}

type projectTypeRepository struct {
	database database.Database
}

func NewProjectTypeRepository(database database.Database) ProjectTypeRepository {
	return &projectTypeRepository{database: database}
}

func (projectTypeRepository *projectTypeRepository) Add(ctx context.Context, projectType ProjectTypeDTO) error {
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
	if err != nil {
		fmt.Println(err.Error())
		var sqlErr *sqlite.Error
		if !errors.As(err, &sqlErr) {
			return err
		}
		switch sqlErr.Code() {
		case sqlite3.SQLITE_CONSTRAINT_UNIQUE:
			if strings.Contains(sqlErr.Error(), "project_types.name") {
				return &domain.AlreadyExistsError{Field: "name"}
			} else if strings.Contains(sqlErr.Error(), "project_types.id") {
				return &domain.AlreadyExistsError{Field: "id"}
			}
		case sqlite3.SQLITE_CONSTRAINT_PRIMARYKEY:
			return &domain.ValidationError{Field: "id"}
		case sqlite3.SQLITE_CONSTRAINT_CHECK:
			if strings.Contains(sqlErr.Error(), "length(name)") {
				return &domain.ValidationError{Field: "name"}
			} else if strings.Contains(sqlErr.Error(), "length(id)") {
				return &domain.ValidationError{Field: "id"}
			}
		}
	}
	return err
}

func (projectTypeRepository *projectTypeRepository) Update(ctx context.Context, projectType ProjectTypeDTO) error {
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
	if err != nil {
		fmt.Println(err.Error())
		var sqlErr *sqlite.Error
		if !errors.As(err, &sqlErr) {
			return err
		}
		switch sqlErr.Code() {
		case sqlite3.SQLITE_CONSTRAINT_UNIQUE:
			if strings.Contains(sqlErr.Error(), "project_types.name") {
				return &domain.AlreadyExistsError{Field: "name"}
			} else if strings.Contains(sqlErr.Error(), "project_types.id") {
				return &domain.AlreadyExistsError{Field: "id"}
			}
		case sqlite3.SQLITE_CONSTRAINT_PRIMARYKEY:
			return &domain.ValidationError{Field: "id"}
		case sqlite3.SQLITE_CONSTRAINT_CHECK:
			if strings.Contains(sqlErr.Error(), "length(name)") {
				return &domain.ValidationError{Field: "name"}
			} else if strings.Contains(sqlErr.Error(), "length(id)") {
				return &domain.ValidationError{Field: "id"}
			}
		}
	}
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

func (projectTypeRepository *projectTypeRepository) Get(ctx context.Context, id string) (ProjectTypeDTO, error) {
	var projectType ProjectTypeDTO
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
			return ProjectTypeDTO{}, domain.NotFoundError
		}
		return ProjectTypeDTO{}, err
	}
	return projectType, err
}

func (projectTypeRepository *projectTypeRepository) Search(ctx context.Context, pager browser.Params, order browser.Order, filter searchFilterDTO) ([]ProjectTypeDTO, browser.Result, error) {
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
	if filter.Name != nil && len(*filter.Name) > 0 {
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
	projectTypes := make([]ProjectTypeDTO, 0)
	for rows.Next() {
		var projectType ProjectTypeDTO
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
