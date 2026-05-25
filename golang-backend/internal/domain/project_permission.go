package domain

type ProjectPermission struct {
	ID   string
	User UserBase
	Role Role
}
