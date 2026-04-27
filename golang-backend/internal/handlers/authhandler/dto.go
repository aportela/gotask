package authhandler


type signInRequest struct {
	Email    string `json:"email"`
	Password string `json:"passsword"`
}

type signUpRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"passsword"`
}

type TokenResponse {
	Token string `json:"token"`
	ExpiresAt int64 `json:"expiresAt"`
}

type SuccessSignInResponse struct {
	AccessToken  TokenResponse `json:"accessToken"`
	RefreshToken TokenResponse `json:"refreshToken"`
}
