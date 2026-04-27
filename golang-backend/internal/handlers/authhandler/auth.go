package authhandler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/aportela/doneo/internal/database"
	"github.com/aportela/doneo/internal/domain"
	"github.com/aportela/doneo/internal/handlers"
	"github.com/aportela/doneo/internal/repositories/userrepository"
	"github.com/aportela/doneo/internal/services/authservice"
	"github.com/aportela/doneo/internal/utils"
	"github.com/golang-jwt/jwt/v4"
)

type Token struct {
	Token     string
	ExpiresAt time.Time
}

func generateToken(user domain.User, expiresAt time.Time, secretKey string) (Token, error) {
	role := "user"
	if user.IsSuperUser {
		role = "administrator"
	}
	claims := jwt.MapClaims{
		"sub":  user.ID,
		"exp":  expiresAt,
		"iat":  time.Now(),
		"role": role,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return Token{}, err
	}
	return Token{Token: signedToken, ExpiresAt: expiresAt}, nil
}

func verifyToken(tokenString string, secretKey string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if exp, ok := claims["exp"].(float64); ok {
			if time.Now().Unix() > int64(exp) {
				return "", errors.New("token has expired")
			}
		} else {
			return "", errors.New("exp claim is missing or invalid")
		}
		sub, ok := claims["sub"].(string)
		if !ok {
			return "", errors.New("sub claim is missing or invalid")
		}
		return sub, nil
	}
	return "", errors.New("invalid token")
}

type AuthHandler struct {
	service                    authservice.AuthService
	secretKey                  string
	accessTokenExpirationHours int
	refreshTokenExpirationDays int
}

func NewAuthHandler(db database.Database, secretKey string, accessTokenExpirationHours int, refreshTokenExpirationDays int) *AuthHandler {
	userRepository := userrepository.NewUserRepository(db)
	authService := authservice.NewAuthService(userRepository)
	return &AuthHandler{service: authService, secretKey: secretKey, accessTokenExpirationHours: accessTokenExpirationHours, refreshTokenExpirationDays: refreshTokenExpirationDays}
}

func (h *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var request signUpRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		handlers.ToHandlerJSONResponse(w, nil, fmt.Errorf("[AuthHandler] invalid request payload: %w", err))
		return
	}
	err := h.service.SignUp(r.Context(), mapSignUpRequestToUserDomain(request))
	if err != nil {
		handlers.ToHandlerJSONResponse(w, nil, fmt.Errorf("[AuthHandler] failed to register user with email %s: %w", request.Email, err))
		return
	}
	handlers.ToHandlerJSONResponse(w, handlers.ToEmptyResponse(), nil, http.StatusCreated)
}

func (h *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	var request signInRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		handlers.ToHandlerJSONResponse(w, nil, fmt.Errorf("[AuthHandler] invalid request payload: %w", err))
		return
	}
	user, err := h.service.SignIn(r.Context(), mapSigninRequestToUserDomain(request))
	if err != nil {
		handlers.ToHandlerJSONResponse(w, nil, fmt.Errorf("[AuthHandler] failed to signin with email %s: %w", request.Email, err))
		return
	}

	accessToken, err := generateToken(user, time.Now().Add(time.Duration(h.accessTokenExpirationHours)*time.Hour), h.secretKey)
	if err != nil {
		handlers.ToHandlerJSONResponse(w, nil, fmt.Errorf("[AuthHandler] failed to generate access token: %w", err))
		return
	}
	refreshToken, err := generateToken(user, time.Now().Add(time.Duration(h.refreshTokenExpirationDays)*24*time.Hour), h.secretKey)
	if err != nil {
		handlers.ToHandlerJSONResponse(w, nil, fmt.Errorf("[AuthHandler] failed to generate refresh token: %w", err))
		return
	}
	cookie := http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken.Token,
		Path:     "/api/auth/renew_access_token",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		Expires:  refreshToken.ExpiresAt,
	}
	http.SetCookie(w, &cookie)
	utils.ToJSONResponse(w, http.StatusOK,
		SignInResponse{
			AccessToken:  TokenResponse{Token: accessToken.Token, ExpiresAt: accessToken.ExpiresAt.UnixMilli()},
			RefreshToken: TokenResponse{Token: refreshToken.Token, ExpiresAt: refreshToken.ExpiresAt.UnixMilli()},
		},
	)
}

func (h *AuthHandler) SignOut(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     "refresh_token",
		Value:    "",
		Path:     "/api/auth/renew_access_token",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		Expires:  time.Now().Add(-time.Hour),
	}
	http.SetCookie(w, &cookie)
	utils.ToJSONResponse(w, http.StatusOK, handlers.ToEmptyResponse())
}

func (h *AuthHandler) RenewAccessToken(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("refresh_token")
	// TODO: refresh token in request ?
	if err != nil {
		http.Error(w, "No refresh token cookie found", http.StatusUnauthorized)
		return
	}
	var refreshToken Token
	refreshToken.Token = cookie.Value
	userId, err := verifyToken(refreshToken.Token, h.secretKey)
	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}
	user := domain.User{
		UserBase: domain.UserBase{ID: userId},
	}
	accessToken, err := generateToken(user, time.Now().Add(time.Duration(h.accessTokenExpirationHours)*time.Hour), h.secretKey)
	if err != nil {
		handlers.ToHandlerJSONResponse(w, nil, fmt.Errorf("[AuthHandler] failed to generate access token: %w", err))
		return
	}
	utils.ToJSONResponse(w, http.StatusOK,
		RenewAccessTokenResponse{
			AccessToken: TokenResponse{Token: accessToken.Token, ExpiresAt: accessToken.ExpiresAt.UnixMilli()},
		},
	)
}
