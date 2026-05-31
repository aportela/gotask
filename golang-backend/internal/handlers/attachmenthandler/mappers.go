package attachmenthandler

import (
	"github.com/aportela/doneo/internal/domain"
	"github.com/aportela/doneo/internal/handlers/userhandler"
)

func domainToResponse(attachment domain.Attachment) AttachmentResponse {
	return AttachmentResponse{
		ID:          attachment.ID,
		CreatedBy:   userhandler.BaseDomainToBaseResponse(attachment.CreatedBy),
		CreatedAt:   attachment.CreatedAt.UnixMilli(),
		Filename:    attachment.OriginalName,
		Size:        attachment.Size,
		ContentType: attachment.ContentType,
	}
}
