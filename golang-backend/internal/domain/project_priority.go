package domain

type ProjectPriority struct {
	ID          string
	WorkspaceId string
	Name        string
	Index       int
	HexColor    string
}

type ProjectPriorityFilter struct {
	WorkspaceId string
}
