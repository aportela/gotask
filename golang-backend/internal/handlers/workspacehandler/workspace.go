package workspacehandler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aportela/doneo/internal/database"
	"github.com/aportela/doneo/internal/handlers"
	"github.com/aportela/doneo/internal/repositories/workspacerepository"
	"github.com/aportela/doneo/internal/services/workspaceservice"
)

type WorkspaceHandler struct {
	service workspaceservice.WorkspaceService
}

func NewWorkspaceHandler(db database.Database) *WorkspaceHandler {
	workspaceRepository := workspacerepository.NewWorkspaceRepository(db)
	workspaceService := workspaceservice.NewWorkspaceService(workspaceRepository)
	return &WorkspaceHandler{service: workspaceService}
}

func (h *WorkspaceHandler) AddWorkspace(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var request addWorkspaceRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		handlers.ToHandlerJSONResponse(w, nil, fmt.Errorf("[WorkspaceHandler] invalid request payload: %w", err))
		return
	}
	workspace := mapAddWorkspaceRequestToWorkspaceDomain(request)
	err := h.service.AddWorkspace(r.Context(), workspace)
	if err != nil {
		handlers.ToHandlerJSONResponse(w, nil, fmt.Errorf("[WorkspaceHandler] failed to add workspace with ID %s: %w", request.ID, err))
		return
	}
	handlers.ToHandlerJSONResponse(w, mapWorkspaceDomainToAddWorkspaceResponse(workspace), nil, http.StatusCreated)
}

func (h *WorkspaceHandler) SearchWorkspaces(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	workspaces, err := h.service.SearchWorkspaces(r.Context())
	handlers.ToHandlerJSONResponse(w, mapWorkspaceArrayDomainToSearchWorkspacesResponse(workspaces), err)
}
