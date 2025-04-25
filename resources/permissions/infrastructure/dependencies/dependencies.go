package dependencies

import (
	"api-seguridad/core/database"
	"api-seguridad/resources/permissions/application"
	"api-seguridad/resources/permissions/domain/repository"
	"api-seguridad/resources/permissions/infrastructure/adapters"
)

var (
	permissionRepo repository.PermissionRepository

	// Use Cases
	createPermissionUseCase     *application.CreatePermissionUseCase
	getPermissionByIDUseCase    *application.GetPermissionByIDUseCase
	getAllPermissionsUseCase    *application.GetAllPermissionsUseCase
	updatePermissionUseCase     *application.UpdatePermissionUseCase
	softDeletePermissionUseCase *application.SoftDeletePermissionUseCase
)

func InitPermissionsDependencies() {
	db := database.GetDB()
	
	// Initialize repository
	permissionRepo = adapters.NewPermissionRepository(db)
	
	// Initialize use cases
	createPermissionUseCase = application.NewCreatePermissionUseCase(permissionRepo)
	getPermissionByIDUseCase = application.NewGetPermissionByIDUseCase(permissionRepo)
	getAllPermissionsUseCase = application.NewGetAllPermissionsUseCase(permissionRepo)
	updatePermissionUseCase = application.NewUpdatePermissionUseCase(permissionRepo)
	softDeletePermissionUseCase = application.NewSoftDeletePermissionUseCase(permissionRepo)
}

// Getter functions for use cases
func GetCreatePermissionUseCase() *application.CreatePermissionUseCase {
	return createPermissionUseCase
}

func GetPermissionByIDUseCase() *application.GetPermissionByIDUseCase {
	return getPermissionByIDUseCase
}

func GetAllPermissionsUseCase() *application.GetAllPermissionsUseCase {
	return getAllPermissionsUseCase
}

func GetUpdatePermissionUseCase() *application.UpdatePermissionUseCase {
	return updatePermissionUseCase
}

func GetSoftDeletePermissionUseCase() *application.SoftDeletePermissionUseCase {
	return softDeletePermissionUseCase
}

// Repository getter
func GetPermissionRepository() repository.PermissionRepository {
	return permissionRepo
}