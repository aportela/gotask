package notehandler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aportela/doneo/internal/database"
	"github.com/aportela/doneo/internal/handlers"
	"github.com/aportela/doneo/internal/repositories/noterepository"
	"github.com/aportela/doneo/internal/services/noteservice"
	"github.com/aportela/doneo/internal/utils"
	"github.com/go-chi/chi/v5"
)

type NoteHandler struct {
	service noteservice.NoteService
}

func NewNoteHandler(db database.Database) *NoteHandler {
	// TODO: rename to repository/service
	noteRepository := noterepository.NewNoteRepository(db)
	noteService := noteservice.NewNoteService(noteRepository)
	return &NoteHandler{service: noteService}
}

func (h *NoteHandler) AddProjectNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var request addRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		handlers.ToHandlerJSONResponse(w, nil, fmt.Errorf("[NoteHandler] invalid request payload: %w", err))
		return
	}
	note := addRequestToDomain(request)
	note.ID = utils.UUID()
	projectId := chi.URLParam(r, "id")

	err := h.service.AddProjectNote(r.Context(), projectId, note)
	if err != nil {
		handlers.ToHandlerJSONResponse(w, nil, fmt.Errorf("[NoteHandler] failed to add note: %w", err))
		return
	}
	handlers.ToHandlerJSONResponse(w, domainToResponse(note), nil, http.StatusCreated)
}

func (h *NoteHandler) DeleteProjectNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	noteId := chi.URLParam(r, "note_id")
	err := h.service.DeleteProjectNote(r.Context(), noteId)
	if err != nil {
		handlers.ToHandlerJSONResponse(w, nil, fmt.Errorf("[NoteHandler] failed to delete note: %w", err))
		return
	}
	handlers.ToHandlerJSONResponse(w, handlers.ToEmptyResponse(), nil)
}

func (h *NoteHandler) SearchProjectNotes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	projectId := chi.URLParam(r, "id")
	projectPermissions, err := h.service.SearchProjectNotes(r.Context(), projectId)
	handlers.ToHandlerJSONResponse(w, toSearchResponse(projectPermissions), err)
}
