package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/aportela/gotask/internal/models"
)

type TasksResponse struct {
	Success bool          `json:"success"`
	Tasks   []models.Task `json:"tasks"`
}

func SearchTasksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	resp := ProjectsResponse{
		Success: true,
		Projects: []models.Project{
			{ID: 1, Name: "Task 1"},
			{ID: 2, Name: "Task 2"},
		},
	}

	json.NewEncoder(w).Encode(resp)

}
