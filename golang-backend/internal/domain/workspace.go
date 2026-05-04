package domain

type Workspace struct {
	ID          string
	Name        string
	Description *string
	HexColor    string
	CreatedBy   UserBase
	CreatedAt   int64
	UpdatedAt   *int64
	DeletedAt   *int64
}
