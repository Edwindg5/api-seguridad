//api-seguridad/resources/role_permissions/infrastructure/dependencies/dependencies.go
package dependencies

import (
	"api-seguridad/core/database"
	"api-seguridad/resources/role_permissions/application"
	"api-seguridad/resources/role_permissions/domain/repository"
	"api-seguridad/resources/role_permissions/infrastructure/adapters"
)

var (
	rolePermissionRepo repository.RolePermissionRepository

	// Use Cases
	createRolePermissionUseCase          *application.CreateRolePermissionUseCase
	getRolePermissionByIDUseCase         *application.GetRolePermissionByIDUseCase
	getByRoleAndPermissionUseCase        *application.GetByRoleAndPermissionUseCase
	getAllByRoleUseCase                  *application.GetAllByRoleUseCase
	updateRolePermissionUseCase          *application.UpdateRolePermissionUseCase
	deleteRolePermissionUseCase          *application.DeleteRolePermissionUseCase

)

func InitDependencies() {
	db := database.GetDB()
	
	// Initialize repository
	rolePermissionRepo = adapters.NewRolePermissionRepository(db)
	
	// Initialize use cases
	createRolePermissionUseCase = application.NewCreateRolePermissionUseCase(rolePermissionRepo)
	getRolePermissionByIDUseCase = application.NewGetRolePermissionByIDUseCase(rolePermissionRepo)
	getByRoleAndPermissionUseCase = application.NewGetByRoleAndPermissionUseCase(rolePermissionRepo)
	getAllByRoleUseCase = application.NewGetAllByRoleUseCase(rolePermissionRepo)
	updateRolePermissionUseCase = application.NewUpdateRolePermissionUseCase(rolePermissionRepo)
	deleteRolePermissionUseCase = application.NewDeleteRolePermissionUseCase(rolePermissionRepo)
}

// Getter functions for use cases
func GetCreateRolePermissionUseCase() *application.CreateRolePermissionUseCase {
	return createRolePermissionUseCase
}

func GetRolePermissionByIDUseCase() *application.GetRolePermissionByIDUseCase {
	return getRolePermissionByIDUseCase
}

func GetByRoleAndPermissionUseCase() *application.GetByRoleAndPermissionUseCase {
	return getByRoleAndPermissionUseCase
}

func GetAllByRoleUseCase() *application.GetAllByRoleUseCase {
	return getAllByRoleUseCase
}

func GetUpdateRolePermissionUseCase() *application.UpdateRolePermissionUseCase {
	return updateRolePermissionUseCase
}

func GetDeleteRolePermissionUseCase() *application.DeleteRolePermissionUseCase {
	return deleteRolePermissionUseCase
}

// Repository getter
func GetRolePermissionRepository() repository.RolePermissionRepository {
	return rolePermissionRepo
}