package domain

type ProjectType struct {
	ID       string
	Name     string
	HexColor string
}

type SearchProjectTypesFilter struct {
	Name *string
}
