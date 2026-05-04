package demodatascripts

import (
	"fmt"

	"github.com/aportela/doneo/internal/database"
)

func CreateDemoData(database database.Database) {
	userIds := createUsers(database, 32)
	projectTypeIds := createProjectTypes(database)
	projectPriorityIds := createProjectPriorities(database)
	projectStatusIds := createProjectStatuses(database)
	fmt.Println(userIds)
	fmt.Println(projectTypeIds)
	fmt.Println(projectPriorityIds)
	fmt.Println(projectStatusIds)
	createProjects(database, projectTypeIds, projectPriorityIds, projectStatusIds, userIds, 32)
}
