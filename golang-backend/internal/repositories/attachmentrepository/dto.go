package attachmentrepository

type attachmentDTO struct {
	ID           string `db:"id"`
	UserId       string `db:"user_id"`
	UserName     string `db:"user_name"`
	CreatedAt    int64  `db:"created_at"`
	OriginalName string `db:"original_name"`
	ContentType  string `db:"content_type"`
	Size         uint32 `db:"size"`
}
