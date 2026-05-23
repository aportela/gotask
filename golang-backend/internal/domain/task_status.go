package domain

type TaskStatus struct {
	ID       string
	Name     string
	HexColor string
}

type SearchTaskStatusesFilter struct {
	Name *string
}
