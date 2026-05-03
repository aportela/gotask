package projecttyperepository

type projectTypeDTO struct {
	ID          string `db:"id"`
	WorkspaceId string `db:"workspace_id"`
	Name        string `db:"name"`
	HexColor    string `db:"item_hex_color"`
}

type projectTypeFilterDTO struct {
	WorkspaceId string
}
