package attachmentservice

import (
	"context"
	"fmt"

	"github.com/aportela/doneo/internal/domain"
	"github.com/aportela/doneo/internal/repositories/attachmentrepository"
)

type AttachmentService interface {
	AddProjectAttachment(ctx context.Context, projectId string, attachment domain.Attachment) error
	DeleteProjectAttachment(ctx context.Context, projectId string, attachmentId string) error
	GetProjectAttachments(ctx context.Context, projectId string) ([]domain.Attachment, error)
}

type attachmentService struct {
	repository attachmentrepository.AttachmentRepository
}

func NewAttachmentService(repository attachmentrepository.AttachmentRepository) AttachmentService {
	return &attachmentService{repository: repository}
}

func (s *attachmentService) AddProjectAttachment(ctx context.Context, projectId string, attachment domain.Attachment) error {
	return s.repository.AddProjectAttachment(ctx, projectId, attachmentrepository.DomainToDTO(attachment))
}

func (s *attachmentService) DeleteProjectAttachment(ctx context.Context, projectId string, attachmentId string) error {
	return s.repository.DeleteProjectAttachment(ctx, projectId, attachmentId)
}

func (s *attachmentService) GetProjectAttachments(ctx context.Context, projectId string) ([]domain.Attachment, error) {
	attachments, err := s.repository.GetProjectAttachments(ctx, projectId)
	if err != nil {
		return nil, fmt.Errorf("[ProjectTypeService] failed to get project attachments: %w", err)
	}
	return attachmentrepository.DTOArrayToDomainArray(attachments), nil
}
