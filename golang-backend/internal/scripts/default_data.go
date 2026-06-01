package scripts

import (
	"context"
	"fmt"
	"time"

	"github.com/aportela/doneo/internal/database"
	"github.com/aportela/doneo/internal/domain"
	"github.com/aportela/doneo/internal/repositories/userrepository"
	"github.com/aportela/doneo/internal/services/userservice"
	"github.com/aportela/doneo/internal/utils"
)

func CreateDefaultData(db database.Database) {
	userRepository := userrepository.NewRepository(db)
	userService := userservice.NewUserService(userRepository)
	permissionsBitmask := domain.PermissionsBitmask(0)
	permissionsBitmask.AddPermission(domain.UserPermissionAdmin)
	err := userService.Add(context.Background(), domain.User{
		UserBase: domain.UserBase{
			ID:   utils.UUID(),
			Name: "administrator",
		},
		Email:              "admin@localhost.localnet",
		Password:           "secret",
		CreatedAt:          time.Now(),
		UpdatedAt:          nil,
		DeletedAt:          nil,
		PermissionsBitmask: permissionsBitmask,
	})
	if err != nil {
		fmt.Printf("Error creating user %s\n", err.Error())
	}
}
