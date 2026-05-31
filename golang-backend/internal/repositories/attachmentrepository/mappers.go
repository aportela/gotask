package attachmentrepository

import (
	"time"

	"github.com/aportela/doneo/internal/domain"
)

func DomainToDTO(attachment domain.Attachment) attachmentDTO {
	return attachmentDTO{
		ID:           attachment.ID,
		UserId:       attachment.CreatedBy.ID,
		UserName:     attachment.CreatedBy.Name,
		CreatedAt:    attachment.CreatedAt.UnixMilli(),
		OriginalName: attachment.OriginalName,
		ContentType:  attachment.ContentType,
		Size:         attachment.Size,
	}
}

func DTOToDomain(attachment attachmentDTO) domain.Attachment {
	return domain.Attachment{
		ID: attachment.ID,
		CreatedBy: domain.UserBase{
			ID:   attachment.UserId,
			Name: attachment.UserName,
		},
		CreatedAt:    time.UnixMilli(attachment.CreatedAt),
		OriginalName: attachment.OriginalName,
		ContentType:  attachment.ContentType,
		Size:         attachment.Size,
	}
}

func DTOArrayToDomainArray(attachments []attachmentDTO) []domain.Attachment {
	results := make([]domain.Attachment, 0, len(attachments))
	for _, attachment := range attachments {
		results = append(results, DTOToDomain(attachment))
	}
	return results
}
