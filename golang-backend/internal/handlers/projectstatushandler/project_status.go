package projectstatushandler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aportela/doneo/internal/database"
	"github.com/aportela/doneo/internal/domain"
	"github.com/aportela/doneo/internal/handlers"
	"github.com/aportela/doneo/internal/repositories/projectstatusrepository"
	"github.com/aportela/doneo/internal/services/projectstatusservice"
	"github.com/go-chi/chi/v5"
)

type ProjectStatusHandler struct {
	service projectstatusservice.ProjectStatusService
}

func NewProjectStatusHandler(db database.Database) *ProjectStatusHandler {
	projectStatusRepository := projectstatusrepository.NewProjectStatusRepository(db)
	projectStatusService := projectstatusservice.NewProjectStatusService(projectStatusRepository)
	return &ProjectStatusHandler{service: projectStatusService}
}

func (h *ProjectStatusHandler) AddProjectStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var request addProjectStatusRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		handlers.ToHandlerJSONResponse(w, nil, fmt.Errorf("[ProjectStatusHandler] invalid request payload: %w", err))
		return
	}
	projectStatus := mapAddProjectStatusRequestToProjectStatusDomain(request)
	err := h.service.AddProjectStatus(r.Context(), projectStatus)
	if err != nil {
		handlers.ToHandlerJSONResponse(w, nil, fmt.Errorf("[ProjectStatusHandler] failed to add project status with ID %s: %w", request.ID, err))
		return
	}
	handlers.ToHandlerJSONResponse(w, mapProjectStatusDomainToAddProjectStatusResponse(projectStatus), nil, http.StatusCreated)
}

func (h *ProjectStatusHandler) UpdateProjectStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var request updateProjectStatusRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		handlers.ToHandlerJSONResponse(w, nil, fmt.Errorf("[ProjectStatusHandler] invalid request payload: %w", err))
		return
	}
	projectStatus := mapUpdateProjectStatusRequestToProjectStatusDomain(request)
	err := h.service.UpdateProjectStatus(r.Context(), projectStatus)
	if err != nil {
		handlers.ToHandlerJSONResponse(w, nil, fmt.Errorf("[ProjectStatusHandler] failed to update project status with ID %s: %w", projectStatus.ID, err))
		return
	}
	handlers.ToHandlerJSONResponse(w, mapProjectStatusDomainToUpdateProjectStatusResponse(projectStatus), nil)
}

func (h *ProjectStatusHandler) DeleteProjectStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	projectStatusId := chi.URLParam(r, "id")
	err := h.service.DeleteProjectStatus(r.Context(), projectStatusId)
	if err != nil {
		handlers.ToHandlerJSONResponse(w, nil, fmt.Errorf("[ProjectStatusService] failed to delete project status with ID %s: %w", projectStatusId, err))
		return
	}
	handlers.ToHandlerJSONResponse(w, handlers.ToEmptyResponse(), nil)
}

func (h *ProjectStatusHandler) GetProjectStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	projectStatusId := chi.URLParam(r, "id")
	projectStatus, err := h.service.GetProjectStatus(r.Context(), projectStatusId)
	if err != nil {
		if err == domain.ErrNotFound {
			handlers.ToHandlerJSONResponse(w, nil, fmt.Errorf("[ProjectStatusService] not found projectStatus with ID %s: %w", projectStatusId, err))
			return
		} else {
			handlers.ToHandlerJSONResponse(w, nil, fmt.Errorf("[ProjectStatusService] failed to get projectStatus with ID %s: %w", projectStatusId, err))
			return
		}
	}
	handlers.ToHandlerJSONResponse(w, mapProjectStatusDomainToGetProjectStatusResponse(projectStatus), nil)
}

func (h *ProjectStatusHandler) SearchProjectStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var request searchProjectStatusesRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		handlers.ToHandlerJSONResponse(w, nil, fmt.Errorf("[ProjectPriorityHandler] invalid request payload: %w", err))
		return
	}
	projectStatuses, err := h.service.SearchProjectStatuses(r.Context())
	handlers.ToHandlerJSONResponse(w, mapProjectStatusArrayDomainToSearchProjectStatussResponse(projectStatuses), err)
}
