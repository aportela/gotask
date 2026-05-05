package projecttypehandler

type addRequest struct {
	Name     string `json:"name"`
	HexColor string `json:"hexColor"`
}

type updateRequest struct {
	Name     string `json:"name"`
	HexColor string `json:"hexColor"`
}

type projectTypeResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	HexColor string `json:"hexColor"`
}

type addResponse struct {
	ProjectType projectTypeResponse `json:"projectType"`
}

type updateResponse struct {
	ProjectType projectTypeResponse `json:"projectType"`
}

type getResponse struct {
	ProjectType projectTypeResponse `json:"projectType"`
}

type searchResponse struct {
	ProjectTypes []projectTypeResponse `json:"projectTypes"`
}
