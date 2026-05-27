package noterepository

import (
	"time"

	"github.com/aportela/doneo/internal/domain"
	"github.com/aportela/doneo/internal/utils"
)

func DomainToDTO(note domain.Note) noteDTO {
	return noteDTO{
		ID:        note.ID,
		UserId:    note.User.ID,
		UserName:  note.User.Name,
		CreatedAt: note.CreatedAt.UnixMilli(),
		UpdatedAt: utils.TimePtrToSQLNullInt64(note.UpdatedAt),
		Body:      note.Body,
	}
}

func DTOToDomain(note noteDTO) domain.Note {
	return domain.Note{
		ID: note.ID,
		User: domain.UserBase{
			ID:   note.UserId,
			Name: note.UserName,
		},
		CreatedAt: time.UnixMilli(note.CreatedAt),
		UpdatedAt: utils.SQLNullInt64ToTimePtr(note.UpdatedAt),
		Body:      note.Body,
	}
}

func DTOArrayToDomainArray(note []noteDTO) []domain.Note {
	results := make([]domain.Note, 0, len(note))
	for _, note := range note {
		results = append(results, DTOToDomain(note))
	}
	return results
}
