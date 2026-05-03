package demodatascripts

import (
	"github.com/aportela/doneo/internal/database"
)

func CreateDemoData(database database.Database) {
	userIds := createUsers(database, 32)
	workspaceId := createDefaultWorkspace(database, userIds[0])
	projectTypeIds := createProjectTypes(database, workspaceId)
	projectPriorityIds := createProjectPriorities(database, workspaceId)
	projectStatusIds := createProjectStatuses(database, workspaceId)
	createProjects(database, workspaceId, projectTypeIds, projectPriorityIds, projectStatusIds, userIds, 32)
}
