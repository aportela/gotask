package domain

type PermissionsBitmask uint64

func (p PermissionsBitmask) HasPermission(v PermissionsBitmask) bool {
	return p&v == v
}

func (p PermissionsBitmask) HasAny(v PermissionsBitmask) bool {
	return p&v != 0
}

func (p *PermissionsBitmask) AddPermission(v PermissionsBitmask) {
	*p |= v
}

func (p *PermissionsBitmask) RemovePermission(v PermissionsBitmask) {
	*p &^= v
}

func (p *PermissionsBitmask) TogglePermission(v PermissionsBitmask) {
	*p ^= v
}

func (p *PermissionsBitmask) Clear() {
	*p = 0
}

// user permissions
const (
	UserPermissionAdmin = 1 << iota
)

// TODO: new
// PermissionUpdateProject
// PermissionDeleteProject
// PermissionViewProject
// PermissionAddTask
// PermissionUpdateTask
// PermissionDeleteTask
// PermissionViewTask
// PermissionExecuteReport

// app permissions
const (
	PermissionCreate PermissionsBitmask = 1 << iota
	PermissionUpdate
	PermissionDelete
	PermissionView
	PermissionList
	PermissionExecute
)
