package authhandler

import (
	"github.com/aportela/doneo/internal/domain"
)

func mapSigninRequestToUserDomain(request signInRequest) domain.User {
	return domain.User{
		Email:    request.Email,
		Password: &request.Password,
	}
}

func mapSignUpRequestToUserDomain(request signUpRequest) domain.User {
	return domain.User{
		UserBase: domain.UserBase{Name: request.Name},
		Email:    request.Email,
		Password: &request.Password,
	}
}
