package userrepository

type userDTO struct {
	ID           string  `db:"id"`
	Name         string  `db:"name"`
	Email        string  `db:"email"`
	PasswordHash *string `db:"password_hash"`
	CreatedAt    int64   `db:"created_at"`
	UpdatedAt    *int64  `db:"updated_at"`
	IsSuperUser  bool    `db:"is_super_user"`
}
