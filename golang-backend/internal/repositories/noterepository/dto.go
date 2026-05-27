package noterepository

import "database/sql"

type noteDTO struct {
	ID        string        `db:"id"`
	UserId    string        `db:"user_id"`
	UserName  string        `db:"user_name"`
	CreatedAt int64         `db:"created_at"`
	UpdatedAt sql.NullInt64 `db:"updated_at"`
	Body      string        `db:"body"`
}
