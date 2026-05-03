package domain

type ProjectType struct {
	ID          string
	WorkspaceId string
	Name        string
	HexColor    string
}

type ProjectTypeFilter struct {
	WorkspaceId string
}
