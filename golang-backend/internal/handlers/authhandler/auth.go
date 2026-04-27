package authhandler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/aportela/doneo/internal/database"
	"github.com/aportela/doneo/internal/handlers"
	"github.com/aportela/doneo/internal/repositories/userrepository"
	"github.com/aportela/doneo/internal/services/authservice"
	"github.com/aportela/doneo/internal/utils"
)

type AuthHandler struct {
	service authservice.AuthService
}

func NewAuthHandler(db database.Database, secretKey string, accessTokenExpirationDays int, refreshTokenExpirationDays int) *AuthHandler {
	userRepository := userrepository.NewUserRepository(db)
	authService := authservice.NewAuthService(userRepository, secretKey, accessTokenExpirationDays, refreshTokenExpirationDays)
	return &AuthHandler{service: authService}
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
	accessToken, refreshToken, err := h.service.SignIn(r.Context(), mapSigninRequestToUserDomain(request))
	if err != nil {
		handlers.ToHandlerJSONResponse(w, nil, fmt.Errorf("[AuthHandler] failed to signin with email %s: %w", request.Email, err))
		return
	}
	cookie := http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken.Token,
		Path:     "/api/auth/renew_access_token",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		Expires:  utils.MSTimestampToTime(refreshToken.ExpiresAt),
	}
	http.SetCookie(w, &cookie)
	utils.ToJSONResponse(w, http.StatusOK,
		SuccessSignInResponse{
			AccessToken:  TokenResponse{Token: accessToken.Token, ExpiresAt: accessToken.ExpiresAt},
			RefreshToken: TokenResponse{Token: refreshToken.Token, ExpiresAt: refreshToken.ExpiresAt},
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
}
