package authservice

import (
	"context"
	"time"

	"github.com/aportela/doneo/internal/domain"
	"github.com/aportela/doneo/internal/repositories/userrepository"
	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	SignUp(ctx context.Context, user domain.User) error
	SignIn(ctx context.Context, user domain.User) (domain.User, error)
}

type authService struct {
	repository userrepository.UserRepository
}

func NewAuthService(repository userrepository.UserRepository) AuthService {
	return &authService{
		repository: repository,
	}
}

func (s *authService) SignUp(ctx context.Context, user domain.User) error {
	user.ID = func() string { u, _ := uuid.NewV7(); return u.String() }()
	user.CreatedAt = time.Now().UnixMilli()
	return s.repository.Add(ctx, userrepository.MapUserDomainToUserDTO(user))
}

func (s *authService) SignIn(ctx context.Context, user domain.User) (domain.User, error) {
	credentialUser, err := s.repository.GetByEmailForVerifyCredentials(ctx, user.Email, *user.Password)
	if err != nil {
		return userrepository.MapUserDTOToUserDomain(credentialUser), err
	}
	err = bcrypt.CompareHashAndPassword([]byte(*credentialUser.PasswordHash), []byte(*user.Password))
	if err != nil {
		return userrepository.MapUserDTOToUserDomain(credentialUser), domain.ErrInvalidCredentials
	}
	return userrepository.MapUserDTOToUserDomain(credentialUser), nil
}
