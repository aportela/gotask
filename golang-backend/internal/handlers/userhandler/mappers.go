package userhandler

import (
	"github.com/aportela/doneo/internal/browser"
	"github.com/aportela/doneo/internal/domain"
	"github.com/aportela/doneo/internal/handlers"
	"github.com/aportela/doneo/internal/utils"
)

func permissionsToDomainPermissionsBitmask(permissions userPermissions) domain.PermissionsBitmask {
	var permissionsBitmask domain.PermissionsBitmask
	if permissions.IsSuperUser {
		permissionsBitmask.AddPermission(domain.UserPermissionAdmin)
	}
	return permissionsBitmask
}

func addRequestToUser(request addRequest) domain.User {
	return domain.User{
		UserBase:           domain.UserBase{Name: request.Name},
		Email:              request.Email,
		Password:           request.Password,
		PermissionsBitmask: permissionsToDomainPermissionsBitmask(request.Permissions),
	}
}

func updateRequestToUser(request updateRequest) domain.User {
	user := domain.User{
		UserBase:           domain.UserBase{Name: request.Name},
		Email:              request.Email,
		PermissionsBitmask: permissionsToDomainPermissionsBitmask(request.Permissions),
	}
	if request.Password != nil {
		user.Password = *request.Password
	}
	return user
}

func domainPermissionsBitmaskToPermissions(permissionsBitmask domain.PermissionsBitmask) userPermissions {
	return userPermissions{
		IsSuperUser: permissionsBitmask.HasPermission(domain.UserPermissionAdmin),
	}
}

func userToResponse(user domain.User) userResponse {
	return userResponse{
		ID:          user.ID,
		Name:        user.Name,
		Email:       user.Email,
		CreatedAt:   user.CreatedAt.UnixMilli(),
		UpdatedAt:   utils.TimePtrToInt64Ptr(user.UpdatedAt),
		DeletedAt:   utils.TimePtrToInt64Ptr(user.DeletedAt),
		Permissions: domainPermissionsBitmaskToPermissions(user.PermissionsBitmask),
		AvatarURL:   user.AvatarURL,
	}
}

func userArrayToResponse(users []domain.User) []userResponse {
	userResponses := []userResponse{}
	for _, user := range users {
		userResponses = append(userResponses, userToResponse(user))
	}
	return userResponses
}

func ToSearchResponse(users []domain.User, pager browser.Result) searchResponse {
	return searchResponse{
		Users: userArrayToResponse(users),
		Pager: handlers.PagerResponse{
			Enabled:      pager.ResultsPage > 0,
			CurrentPage:  pager.CurrentPage,
			ResultsPage:  pager.ResultsPage,
			TotalPages:   pager.TotalPages,
			TotalResults: pager.TotalResults,
		},
	}
}
