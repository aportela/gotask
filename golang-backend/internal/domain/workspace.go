package domain

type Workspace struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description *string  `json:"description"`
	CreatedBy   UserBase `json:"createdBy"`
	CreatedAt   int64    `json:"createdAt"`
	UpdatedAt   *int64   `json:"updatedAt"`
}
