package projectstatushandler

type addRequest struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	HexColor string `json:"hexColor"`
	Index    int    `json:"index"`
}

type updateRequest struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	HexColor string `json:"hexColor"`
	Index    int    `json:"index"`
}

type projectStatusResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	HexColor string `json:"hexColor"`
	Index    int    `json:"index"`
}

type addResponse struct {
	ProjectStatus projectStatusResponse `json:"projectStatus"`
}

type updateResponse struct {
	ProjectStatus projectStatusResponse `json:"projectStatus"`
}

type getResponse struct {
	ProjectStatus projectStatusResponse `json:"projectStatus"`
}

type searchResponse struct {
	ProjectStatuses []projectStatusResponse `json:"projectStatuses"`
}
