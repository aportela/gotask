package projectpermissionhandler

import (
	"github.com/aportela/doneo/internal/handlers/rolehandler"
	"github.com/aportela/doneo/internal/handlers/userhandler"
)

type addRequest struct {
	UserId string `json:"userId"`
	RoleId string `json:"roleId"`
}

type ProjectPermissionResponse struct {
	ID   string                       `json:"id"`
	User userhandler.UserBaseResponse `json:"user"`
	Role rolehandler.RoleResponse     `json:"role"`
}

type searchResponse struct {
	ProjectPermissions []ProjectPermissionResponse `json:"projectPermissions"`
}
