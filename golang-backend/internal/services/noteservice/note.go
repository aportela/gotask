package noteservice

import (
	"context"
	"fmt"

	"github.com/aportela/doneo/internal/domain"
	"github.com/aportela/doneo/internal/repositories/noterepository"
)

type NoteService interface {
	AddProjectNote(ctx context.Context, projectId string, note domain.Note) error
	UpdateProjectNote(ctx context.Context, note domain.Note) error
	DeleteProjectNote(ctx context.Context, projectId string) error
	GetProjectNotes(ctx context.Context, projectId string) ([]domain.Note, error)
}

type noteService struct {
	repository noterepository.NoteRepository
}

func NewNoteService(repository noterepository.NoteRepository) NoteService {
	return &noteService{repository: repository}
}

func (service *noteService) AddProjectNote(ctx context.Context, projectId string, note domain.Note) error {
	return service.repository.AddProjectNote(ctx, projectId, noterepository.DomainToDTO(note))
}

func (service *noteService) UpdateProjectNote(ctx context.Context, note domain.Note) error {
	return service.repository.UpdateProjectNote(ctx, noterepository.DomainToDTO(note))
}

func (service *noteService) DeleteProjectNote(ctx context.Context, projectId string) error {
	return service.repository.DeleteProjectNote(ctx, projectId)
}

func (service *noteService) GetProjectNotes(ctx context.Context, projectId string) ([]domain.Note, error) {
	notes, err := service.repository.GetProjectNotes(ctx, projectId)
	if err != nil {
		return nil, fmt.Errorf("[ProjectTypeService] failed to get project permissions: %w", err)
	}
	return noterepository.DTOArrayToDomainArray(notes), nil
}
