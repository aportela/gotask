package authservice

import (
	"context"

	"github.com/aportela/doneo/internal/domain"
	"github.com/aportela/doneo/internal/repositories/userrepository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	SignIn(ctx context.Context, user domain.User) (domain.User, error)
	GetUserInfo(ctx context.Context, userId string) (domain.User, error)
}

type authService struct {
	repository userrepository.UserRepository
}

func NewService(repository userrepository.UserRepository) AuthService {
	return &authService{
		repository: repository,
	}
}

func (service *authService) SignIn(ctx context.Context, user domain.User) (domain.User, error) {
	credentialUser, err := service.repository.GetByEmailForVerifyCredentials(ctx, user.Email, user.Password)
	if err != nil {
		return domain.User{}, err
	}
	if user.DeletedAt != nil {
		return domain.User{}, domain.DeletedError
	}
	err = bcrypt.CompareHashAndPassword([]byte(credentialUser.PasswordHash), []byte(user.Password))
	if err != nil {
		return domain.User{}, domain.InvalidCredentialsError
	}
	return userrepository.DTOToDomain(credentialUser), nil
}

func (service *authService) GetUserInfo(ctx context.Context, userId string) (domain.User, error) {
	user, err := service.repository.Get(ctx, userId)
	if err != nil {
		return domain.User{}, err
	}
	if !user.DeletedAt.Valid {
		// TODO: FAIL
		//return domain.User{}, domain.ErrDeleted
	}
	return userrepository.DTOToDomain(user), nil
}
