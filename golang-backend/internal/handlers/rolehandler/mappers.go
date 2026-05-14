package rolehandler

import (
	"github.com/aportela/doneo/internal/browser"
	"github.com/aportela/doneo/internal/domain"
	"github.com/aportela/doneo/internal/handlers"
)

func addRequestToRole(request addRequest) domain.Role {
	return domain.Role{
		Name: request.Name,
	}
}

func updateRequestToRole(request updateRequest) domain.Role {
	return domain.Role{
		ID:   request.Id,
		Name: request.Name,
	}
}

func roleToResponse(role domain.Role) roleResponse {
	return roleResponse{
		ID:   role.ID,
		Name: role.Name,
	}
}

func roleArrayToResponse(roles []domain.Role) []roleResponse {
	roleResponses := []roleResponse{}
	for _, role := range roles {
		roleResponses = append(roleResponses, roleToResponse(role))
	}
	return roleResponses
}

func ToSearchResponse(roles []domain.Role, pager browser.Result) searchResponse {
	return searchResponse{
		Roles: roleArrayToResponse(roles),
		Pager: handlers.PagerResponse{
			Enabled:      pager.ResultsPage > 0,
			CurrentPage:  pager.CurrentPage,
			ResultsPage:  pager.ResultsPage,
			TotalPages:   pager.TotalPages,
			TotalResults: pager.TotalResults,
		},
	}
}
