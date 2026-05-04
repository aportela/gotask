package workspacehandler

type addWorkspaceRequest struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	HexColor string `json:"hexColor"`
}

type workspaceResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	HexColor string `json:"hexColor"`
}

type addWorkspaceResponse struct {
	Workspace workspaceResponse `json:"workspace"`
}

type searchWorkspacesResponse struct {
	Workspaces []workspaceResponse `json:"workspaces"`
}
