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
	userRepository := userrepository.NewUserRepository(db)
	userService := userservice.NewUserService(userRepository)
	err := userService.Add(context.Background(), domain.User{
		UserBase: domain.UserBase{
			ID:   utils.UUID(),
			Name: "administrator",
		},
		Email:       "admin@localhost.localnet",
		Password:    "secret",
		CreatedAt:   time.Now(),
		UpdatedAt:   nil,
		DeletedAt:   nil,
		IsSuperUser: true,
	})
	if err != nil {
		fmt.Printf("Error creating user %s\n", err.Error())
	}
}
