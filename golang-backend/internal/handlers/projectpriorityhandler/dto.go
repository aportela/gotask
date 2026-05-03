package projectpriorityhandler

type addProjectPriorityRequest struct {
	ID          string `json:"id"`
	WorkspaceId string `json:"workspaceId"`
	Name        string `json:"name"`
	Index       int    `json:"index"`
	HexColor    string `json:"hexColor"`
}

type updateProjectPriorityRequest struct {
	ID          string `json:"id"`
	WorkspaceId string `json:"workspaceId"`
	Name        string `json:"name"`
	Index       int    `json:"index"`
	HexColor    string `json:"hexColor"`
}

type projectPriorityResponse struct {
	ID          string `json:"id"`
	WorkspaceId string `json:"workspaceId"`
	Name        string `json:"name"`
	Index       int    `json:"index"`
	HexColor    string `json:"hexColor"`
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

type searchProjectPrioritysRequest struct {
	WorkspaceId string `json:"workspaceId"`
}

type searchProjectPrioritysResponse struct {
	ProjectPriorities []projectPriorityResponse `json:"projectPriorities"`
}
