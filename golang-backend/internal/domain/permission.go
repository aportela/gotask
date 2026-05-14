package domain

type Permission uint8

const (
	PermissionCreate Permission = 1 << iota
	PermissionUpdate
	PermissionDelete
	PermissionView
	PermissionList
	PermissionExecute
)
