package projectpriorityrepository

type projectPriorityDTO struct {
	ID          string `db:"id"`
	WorkspaceId string `db:"workspace_id"`
	Name        string `db:"name"`
	Index       int    `db:"item_index"`
	HexColor    string `db:"item_hex_color"`
}

type projectPriorityFilterDTO struct {
	WorkspaceId string
}
