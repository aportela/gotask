package domain

type RoleBase struct {
	ID   string
	Name string
}

type Role struct {
	RoleBase
	PermissionsBitmask PermissionsBitmask
}

type SearchRolesFilter struct {
	Name                        *string
	RequiredPermissionsBitmask  *PermissionsBitmask
	ForbiddenPermissionsBitmask *PermissionsBitmask
}
