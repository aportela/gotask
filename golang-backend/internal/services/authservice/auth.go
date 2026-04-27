package authservice

import (
	"context"
	"time"

	"github.com/aportela/doneo/internal/domain"
	"github.com/aportela/doneo/internal/repositories/userrepository"
	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type Token struct {
	Token     string
	ExpiresAt int64
}

type AuthService interface {
	SignUp(ctx context.Context, user domain.User) error
	SignIn(ctx context.Context, user domain.User) (Token, Token, error)
}

type authService struct {
	repository                 userrepository.UserRepository
	secretKey                  string
	accessTokenExpirationDays  int
	refreshTokenExpirationDays int
}

func NewAuthService(repository userrepository.UserRepository, secretKey string, accessTokenExpirationDays int, refreshTokenExpirationDays int) AuthService {
	return &authService{
		repository:                 repository,
		secretKey:                  secretKey,
		accessTokenExpirationDays:  accessTokenExpirationDays,
		refreshTokenExpirationDays: refreshTokenExpirationDays,
	}
}

func (s *authService) SignUp(ctx context.Context, user domain.User) error {
	user.ID = func() string { u, _ := uuid.NewV7(); return u.String() }()
	user.CreatedAt = time.Now().UnixMilli()
	return s.repository.Add(ctx, userrepository.MapUserDomainToUserDTO(user))
}

func (s *authService) SignIn(ctx context.Context, user domain.User) (Token, Token, error) {
	credentialUser, err := s.repository.GetByEmailForVerifyCredentials(ctx, user.Email, *user.Password)
	if err != nil {
		return Token{}, Token{}, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(*credentialUser.PasswordHash), []byte(*user.Password))
	if err != nil {
		return Token{}, Token{}, domain.ErrInvalidCredentials
	}
	var accessToken, refreshToken Token
	accessToken.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()        // expires in 1 hour
	accessToken.ExpiresAt = time.Now().Add(365 * 24 * time.Hour).Unix() // expires in 365 days

	accessToken.Token, err = s.generateJWT(user, accessToken.ExpiresAt)
	if err != nil {
		return Token{}, Token{}, err
	}
	refreshToken.Token, err = s.generateJWT(user, accessToken.ExpiresAt)
	if err != nil {
		return Token{}, Token{}, err
	}
	return accessToken, refreshToken, nil
}

func (s *authService) generateJWT(user domain.User, expiration int64) (string, error) {
	role := "user"
	if user.IsSuperUser {
		role = "administrator"
	}
	claims := jwt.MapClaims{
		"sub":  user.ID,
		"exp":  expiration,
		"iat":  time.Now().Unix(),
		"role": role,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
