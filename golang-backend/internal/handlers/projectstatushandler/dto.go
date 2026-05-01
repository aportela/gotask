package projectstatushandler

type addProjectStatusRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type updateProjectStatusRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type projectStatusResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
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
