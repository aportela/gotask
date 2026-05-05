package projectpriorityhandler

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

type projectPriorityResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	HexColor string `json:"hexColor"`
	Index    int    `json:"index"`
}

type addResponse struct {
	ProjectPriority projectPriorityResponse `json:"projectPriority"`
}

type updateResponse struct {
	ProjectPriority projectPriorityResponse `json:"projectPriority"`
}

type getResponse struct {
	ProjectPriority projectPriorityResponse `json:"projectPriority"`
}

type searchResponse struct {
	ProjectPriorities []projectPriorityResponse `json:"projectPriorities"`
}
