package domain

import (
	"time"
)

type UserBase struct {
	ID   string
	Name string
}

type User struct {
	UserBase
	Email              string
	Password           string
	PasswordHash       string
	CreatedAt          time.Time
	UpdatedAt          *time.Time
	DeletedAt          *time.Time
	PermissionsBitmask PermissionsBitmask
}

func (u *User) IsActive() bool {
	return u.DeletedAt == nil
}

func (u *User) IsAdmin() bool {
	return u.PermissionsBitmask.HasPermission(UserPermissionAdmin)
}

type SearchUsersFilter struct {
	Name                        *string
	Email                       *string
	RequiredPermissionsBitmask  *PermissionsBitmask
	ForbiddenPermissionsBitmask *PermissionsBitmask
	CreatedAt                   *TimestampFilter
	UpdatedAt                   *TimestampFilter
	DeletedAt                   *TimestampFilter
}
