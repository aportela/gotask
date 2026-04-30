package projectrepository

type proyectDTO struct {
	ID          string  `db:"id"`
	Key         string  `db:"key"`
	Summary     string  `db:"summary"`
	Description *string `json:"description"`
	CreatorId   string  `db:"creator_id"`
	CreatorName string  `db:"creator_name"`
	CreatedAt   int64   `db:"created_at"`
	UpdatedAt   *int64  `db:"updated_at"`
	StartedAt   *int64  `db:"started_at"`
	FinishedAt  *int64  `db:"finished_at"`
	DueAt       *int64  `db:"due_at"`
}
