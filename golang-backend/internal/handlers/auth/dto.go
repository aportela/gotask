package handlers

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"passsword"`
}

type SuccessSignInResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
