package projectstatushandler

type addProjectStatusRequest struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Index    int    `json:"index"`
	HexColor string `json:"hexColor"`
}

type updateProjectStatusRequest struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Index    int    `json:"index"`
	HexColor string `json:"hexColor"`
}

type projectStatusResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Index    int    `json:"index"`
	HexColor string `json:"hexColor"`
}

type addProjectStatusResponse struct {
	ProjectStatus projectStatusResponse `json:"projectStatus"`
}

type updateProjectStatusResponse struct {
	ProjectStatus projectStatusResponse `json:"projectStatus"`
}

type getProjectStatusResponse struct {
	ProjectStatus projectStatusResponse `json:"projectStatus"`
}

type searchProjectStatussResponse struct {
	ProjectStatuses []projectStatusResponse `json:"projectStatuses"`
}
