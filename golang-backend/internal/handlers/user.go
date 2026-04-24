package handlers

import (
	"database/sql"
	"net/http"

	"github.com/aportela/gotask/internal/models"
	"github.com/aportela/gotask/internal/services"
	"github.com/aportela/gotask/internal/utils"
	"github.com/go-chi/chi/v5"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) AddUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	w.Header().Set("Content-Type", "application/json")
	project := models.Project{}
	err := h.service.AddUser(ctx, project)
	if err != nil {
		utils.ToJSONResponse(w, http.StatusInternalServerError, map[string]string{
			"debugErrorMessage": err.Error(),
		})
		return
	}
	utils.ToJSONResponse(w, http.StatusOK, project)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	w.Header().Set("Content-Type", "application/json")
	project := models.Project{}
	err := h.service.UpdateUser(ctx, project)
	if err != nil {
		if err == sql.ErrNoRows {
			utils.ToJSONResponse(w, http.StatusNotFound, map[string]string{
				"debugErrorMessage": err.Error(),
			})
			return
		} else {
			utils.ToJSONResponse(w, http.StatusInternalServerError, map[string]string{
				"debugErrorMessage": err.Error(),
			})
			return
		}
	}
	utils.ToJSONResponse(w, http.StatusOK, project)

}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	w.Header().Set("Content-Type", "application/json")
	projectId := chi.URLParam(r, "id")
	err := h.service.DeleteUser(ctx, projectId)
	if err != nil {
		if err == sql.ErrNoRows {
			utils.ToJSONResponse(w, http.StatusNotFound, map[string]string{
				"debugErrorMessage": err.Error(),
			})
			return
		} else {
			utils.ToJSONResponse(w, http.StatusInternalServerError, map[string]string{
				"debugErrorMessage": err.Error(),
			})
			return
		}
	}
	utils.ToJSONResponse(w, http.StatusOK, nil)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	w.Header().Set("Content-Type", "application/json")
	projectId := chi.URLParam(r, "id")
	project, err := h.service.GetProject(ctx, projectId)
	if err != nil {
		if err == sql.ErrNoRows {
			utils.ToJSONResponse(w, http.StatusNotFound, map[string]string{
				"debugErrorMessage": err.Error(),
			})
			return
		} else {
			utils.ToJSONResponse(w, http.StatusInternalServerError, map[string]string{
				"debugErrorMessage": err.Error(),
			})
			return
		}
	}
	utils.ToJSONResponse(w, http.StatusOK, project)
}

func (h *UserHandler) SearchProjects(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	w.Header().Set("Content-Type", "application/json")
	projects, err := h.service.SearchProjects(ctx)
	if err != nil {
		utils.ToJSONResponse(w, http.StatusInternalServerError, map[string]string{
			"debugErrorMessage": err.Error(),
		})
		return
	}
	utils.ToJSONResponse(w, http.StatusOK, projects)
}
