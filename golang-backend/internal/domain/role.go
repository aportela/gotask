package domain

type Role struct {
	ID         string
	Name       string
	Permission Permission
}

func (p Permission) HasPermission(v Permission) bool {
	return p&v == v
}

func (p *Permission) AddPermission(v Permission) {
	*p |= v
}

func (p *Permission) RemovePermission(v Permission) {
	*p &^= v
}
