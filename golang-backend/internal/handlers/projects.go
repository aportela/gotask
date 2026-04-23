package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/aportela/gotask/internal/models"
)

type ProjectsResponse struct {
	Success  bool             `json:"success"`
	Projects []models.Project `json:"projects"`
}

func SearchProjectsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	resp := ProjectsResponse{
		Success: true,
		Projects: []models.Project{
			{ID: 1, Summary: "Project 1"},
			{ID: 2, Summary: "Proyect 2"},
		},
	}

	json.NewEncoder(w).Encode(resp)

}
