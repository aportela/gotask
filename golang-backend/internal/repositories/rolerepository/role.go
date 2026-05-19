package rolerepository

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

type RoleRepository interface {
	Add(ctx context.Context, role roleDTO) error
	Update(ctx context.Context, role roleDTO) error
	Delete(ctx context.Context, id string) error
	Get(ctx context.Context, id string) (roleDTO, error)
	Search(ctx context.Context, pager browser.Params, order browser.Order, filter searchFilterDTO) ([]roleDTO, browser.Result, error)
}

type roleRepository struct {
	database database.Database
}

func NewRoleRepository(database database.Database) RoleRepository {
	return &roleRepository{database: database}
}

func (roleRepository *roleRepository) Add(ctx context.Context, role roleDTO) error {
	_, err := roleRepository.database.ExecContext(
		ctx,
		`
            INSERT INTO roles (id, name, permissions_bitmask)
			VALUES (?, ?, ?)
        `,
		role.ID,
		role.Name,
		role.PermissionsBitmask,
	)
	if err != nil {
		fmt.Println(err.Error())
		var sqlErr *sqlite.Error
		if !errors.As(err, &sqlErr) {
			return err
		}
		switch sqlErr.Code() {
		case sqlite3.SQLITE_CONSTRAINT_UNIQUE:
			if strings.Contains(sqlErr.Error(), "roles.name") {
				return &domain.AlreadyExistsError{Field: "name"}
			} else if strings.Contains(sqlErr.Error(), "roles.id") {
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

func (roleRepository *roleRepository) Update(ctx context.Context, role roleDTO) error {
	_, err := roleRepository.database.ExecContext(
		ctx,
		`
            UPDATE roles SET
				name = ?,
				permissions_bitmask = ?
			WHERE id = ?
        `,
		role.Name,
		role.PermissionsBitmask,
		role.ID,
	)
	if err != nil {
		fmt.Println(err.Error())
		var sqlErr *sqlite.Error
		if !errors.As(err, &sqlErr) {
			return err
		}
		switch sqlErr.Code() {
		case sqlite3.SQLITE_CONSTRAINT_UNIQUE:
			if strings.Contains(sqlErr.Error(), "roles.name") {
				return &domain.AlreadyExistsError{Field: "name"}
			}
		case sqlite3.SQLITE_CONSTRAINT_CHECK:
			if strings.Contains(sqlErr.Error(), "length(name)") {
				return &domain.ValidationError{Field: "name"}
			}
		}
	}
	return err
}

func (roleRepository *roleRepository) Delete(ctx context.Context, id string) error {
	_, err := roleRepository.database.ExecContext(
		ctx,
		`
            DELETE FROM roles
			WHERE id = ?
        `,
		id,
	)
	return err
}

func (roleRepository *roleRepository) Get(ctx context.Context, id string) (roleDTO, error) {
	var role roleDTO
	err := roleRepository.database.QueryRowContext(
		ctx,
		`
            SELECT
                R.id, R.name, R.permissions_bitmask
            FROM roles R
            WHERE R.id = ?
        `,
		id).Scan(&role.ID, &role.Name, &role.PermissionsBitmask)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return roleDTO{}, domain.NotFoundError
		}
		return roleDTO{}, err
	}
	return role, err
}

func (roleRepository *roleRepository) Search(ctx context.Context, pager browser.Params, order browser.Order, filter searchFilterDTO) ([]roleDTO, browser.Result, error) {
	var filterArgs []any
	var queryArgs []any
	sqlQuery := `
		SELECT
			R.id, R.name, R.permissions_bitmask
		FROM roles R
	`
	var field string
	switch order.Field {
	case "name":
		field = "R.name COLLATE NOCASE"
	default:
		field = "R.name COLLATE NOCASE"
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
		sqlWhereConditions = append(sqlWhereConditions, "R.name LIKE ?")
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
	rows, err := roleRepository.database.QueryContext(ctx, sqlQuery, queryArgs...)
	if err != nil {
		return nil, browser.Result{}, err
	}
	defer rows.Close()
	roles := make([]roleDTO, 0)
	for rows.Next() {
		var role roleDTO
		if err := rows.Scan(&role.ID, &role.Name, &role.PermissionsBitmask); err != nil {
			return nil, browser.Result{}, err
		}
		roles = append(roles, role)
	}
	if err := rows.Err(); err != nil {
		return nil, browser.Result{}, err
	}

	var totalResults int

	if pager.Enabled() {
		sqlCountQuery := `
			SELECT
				COUNT(*) AS total_roles
			FROM roles R
		`
		sqlCountQuery = fmt.Sprintf("%s %s", sqlCountQuery, sqlWhere)
		err = roleRepository.database.QueryRowContext(
			ctx,
			sqlCountQuery,
			filterArgs...,
		).Scan(&totalResults)

		if err != nil {
			return nil, browser.Result{}, err
		}
	} else {
		totalResults = len(roles)
	}

	return roles, browser.NewResult(pager, totalResults), nil
}
