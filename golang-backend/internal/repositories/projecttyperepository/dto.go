package projecttyperepository

type projectTypeDTO struct {
	ID       string `db:"id"`
	Name     string `db:"name"`
	Index    int    `db:"item_index"`
	HexColor string `db:"item_hex_color"`
}
