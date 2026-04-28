package authservice

import (
	"context"

	"github.com/aportela/doneo/internal/domain"
	"github.com/aportela/doneo/internal/repositories/userrepository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
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
