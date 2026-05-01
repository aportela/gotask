package projectpriorityhandler

type addProjectPriorityRequest struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Index    int    `json:"index"`
	HexColor string `json:"hexColor"`
}

type updateProjectPriorityRequest struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Index    int    `json:"index"`
	HexColor string `json:"hexColor"`
}

type projectPriorityResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Index    int    `json:"index"`
	HexColor string `json:"hexColor"`
}

type addProjectPriorityResponse struct {
	ProjectPriority projectPriorityResponse `json:"projectPriority"`
}

type updateProjectPriorityResponse struct {
	ProjectPriority projectPriorityResponse `json:"projectPriority"`
}

type getProjectPriorityResponse struct {
	ProjectPriority projectPriorityResponse `json:"projectPriority"`
}

type searchProjectPrioritysResponse struct {
	ProjectPriorities []projectPriorityResponse `json:"projectPriorities"`
}
