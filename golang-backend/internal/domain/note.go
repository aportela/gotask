package domain

import "time"

type Note struct {
	ID        string
	User      UserBase
	CreatedAt time.Time
	UpdatedAt *time.Time
	Body      string
}
