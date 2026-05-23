package taskstatusrepository

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

type ProjectStatusRepository interface {
	Add(ctx context.Context, taskStatus taskStatusDTO) error
	Update(ctx context.Context, taskStatus taskStatusDTO) error
	Delete(ctx context.Context, id string) error
	Get(ctx context.Context, id string) (taskStatusDTO, error)
	Search(ctx context.Context, pager browser.Params, order browser.Order, filter searchFilterDTO) ([]taskStatusDTO, browser.Result, error)
}

type taskStatusRepository struct {
	database database.Database
}

func NewTaskStatusRepository(database database.Database) ProjectStatusRepository {
	return &taskStatusRepository{database: database}
}

func (taskStatusRepository *taskStatusRepository) Add(ctx context.Context, taskStatus taskStatusDTO) error {
	_, err := taskStatusRepository.database.ExecContext(
		ctx,
		`
            INSERT INTO task_statuses (id, name, item_hex_color)
			VALUES (?, ?, ?)
        `,
		taskStatus.ID,
		taskStatus.Name,
		taskStatus.HexColor,
	)
	if err != nil {
		fmt.Println(err.Error())
		var sqlErr *sqlite.Error
		if !errors.As(err, &sqlErr) {
			return err
		}
		switch sqlErr.Code() {
		case sqlite3.SQLITE_CONSTRAINT_UNIQUE:
			if strings.Contains(sqlErr.Error(), "task_statuses.name") {
				return &domain.AlreadyExistsError{Field: "name"}
			} else if strings.Contains(sqlErr.Error(), "task_statuses.id") {
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

func (taskStatusRepository *taskStatusRepository) Update(ctx context.Context, taskStatus taskStatusDTO) error {
	_, err := taskStatusRepository.database.ExecContext(
		ctx,
		`
            UPDATE task_statuses SET
				name = ?,
				item_hex_color = ?
			WHERE id = ?
        `,
		taskStatus.Name,
		taskStatus.HexColor,
		taskStatus.ID,
	)
	if err != nil {
		fmt.Println(err.Error())
		var sqlErr *sqlite.Error
		if !errors.As(err, &sqlErr) {
			return err
		}
		switch sqlErr.Code() {
		case sqlite3.SQLITE_CONSTRAINT_UNIQUE:
			if strings.Contains(sqlErr.Error(), "task_statuses.name") {
				return &domain.AlreadyExistsError{Field: "name"}
			} else if strings.Contains(sqlErr.Error(), "task_statuses.id") {
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

func (taskStatusRepository *taskStatusRepository) Delete(ctx context.Context, id string) error {
	_, err := taskStatusRepository.database.ExecContext(
		ctx,
		`
            DELETE FROM task_statuses
			WHERE id = ?
        `,
		id,
	)
	return err
}

func (taskStatusRepository *taskStatusRepository) Get(ctx context.Context, id string) (taskStatusDTO, error) {
	var taskStatus taskStatusDTO
	err := taskStatusRepository.database.QueryRowContext(
		ctx,
		`
            SELECT
                TS.id, TS.name, TS.item_hex_color
            FROM task_statuses TS
            WHERE TS.id = ?
        `,
		id).Scan(&taskStatus.ID, &taskStatus.Name, &taskStatus.HexColor)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return taskStatusDTO{}, domain.NotFoundError
		}
		return taskStatusDTO{}, err
	}
	return taskStatus, err
}

func (taskStatusRepository *taskStatusRepository) Search(ctx context.Context, pager browser.Params, order browser.Order, filter searchFilterDTO) ([]taskStatusDTO, browser.Result, error) {
	var filterArgs []any
	var queryArgs []any
	sqlQuery := `
			SELECT
				TS.id, TS.name, TS.item_hex_color
			FROM task_statuses TS
    `
	var field string
	switch order.Field {
	case "name":
		field = "TS.name COLLATE NOCASE"
	default:
		field = "TS.name COLLATE NOCASE"
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
		sqlWhereConditions = append(sqlWhereConditions, "TS.name LIKE ?")
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
	rows, err := taskStatusRepository.database.QueryContext(ctx, sqlQuery, queryArgs...)
	if err != nil {
		return nil, browser.Result{}, err
	}
	defer rows.Close()
	taskStatuses := make([]taskStatusDTO, 0)
	for rows.Next() {
		var taskStatus taskStatusDTO
		if err := rows.Scan(
			&taskStatus.ID, &taskStatus.Name, &taskStatus.HexColor,
		); err != nil {
			return nil, browser.Result{}, err
		}
		taskStatuses = append(taskStatuses, taskStatus)
	}
	if err := rows.Err(); err != nil {
		return nil, browser.Result{}, err
	}

	var totalResults int

	if pager.Enabled() {
		sqlCountQuery := `
			SELECT
				COUNT(*) AS total_task_statuses
			FROM task_statuses TS
		`
		sqlCountQuery = fmt.Sprintf("%s %s", sqlCountQuery, sqlWhere)
		err = taskStatusRepository.database.QueryRowContext(
			ctx,
			sqlCountQuery,
			filterArgs...,
		).Scan(&totalResults)

		if err != nil {
			return nil, browser.Result{}, err
		}
	} else {
		totalResults = len(taskStatuses)
	}

	return taskStatuses, browser.NewResult(pager, totalResults), nil
}
