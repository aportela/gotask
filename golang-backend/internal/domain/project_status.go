package domain

type ProjectStatus struct {
	ID          string
	WorkspaceId string
	Name        string
	Index       int
	HexColor    string
}

type ProjectStatusFilter struct {
	WorkspaceId string
}
