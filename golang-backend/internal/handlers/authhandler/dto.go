package authhandler

type signInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type signUpRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TokenResponse struct {
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expiresAt"`
}

type SignInResponse struct {
	AccessToken  TokenResponse `json:"accessToken"`
	RefreshToken TokenResponse `json:"refreshToken"`
}

type RenewAccessTokenResponse struct {
	AccessToken TokenResponse `json:"accessToken"`
}
