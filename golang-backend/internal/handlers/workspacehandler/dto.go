package workspacehandler

type addWorkspaceRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type workspaceResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type addWorkspaceResponse struct {
	Workspace workspaceResponse `json:"workspace"`
}

type searchWorkspacesResponse struct {
	Workspaces []workspaceResponse `json:"workspaces"`
}
