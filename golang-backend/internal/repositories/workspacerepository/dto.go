package workspacerepository

type workspaceDTO struct {
	ID          string  `db:"id"`
	Name        string  `db:"name"`
	Description *string `db:"description"`
	CreatorId   string  `db:"creator_id"`
	CreatorName string  `db:"creator_name"`
	CreatedAt   int64   `db:"created_at"`
	UpdatedAt   *int64  `db:"updated_at"`
}
