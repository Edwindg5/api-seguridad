//api-seguridad/resources/roles/infrastructure/dependencies/dependencies.go
package dependencies

import (
	"api-seguridad/core/database"
	"api-seguridad/resources/roles/application"
	"api-seguridad/resources/roles/domain/repository"
	"api-seguridad/resources/roles/infrastructure/adapters"
)

var (
	roleRepo repository.RoleRepository

	// Casos de uso
	createRoleUseCase     *application.CreateRoleUseCase
	getRoleByIDUseCase    *application.GetRoleByIDUseCase
	getAllRolesUseCase    *application.GetAllRolesUseCase
	updateRoleUseCase     *application.UpdateRoleUseCase
	softDeleteRoleUseCase *application.SoftDeleteRoleUseCase
)

func InitDependencies() {
	db := database.GetDB()
	
	// Inicializar repositorio
	roleRepo = adapters.NewRoleRepository(db)
	
	// Inicializar casos de uso
	createRoleUseCase = application.NewCreateRoleUseCase(roleRepo)
	getRoleByIDUseCase = application.NewGetRoleByIDUseCase(roleRepo)
	getAllRolesUseCase = application.NewGetAllRolesUseCase(roleRepo)
	updateRoleUseCase = application.NewUpdateRoleUseCase(roleRepo)
	softDeleteRoleUseCase = application.NewSoftDeleteRoleUseCase(roleRepo)
}

// Funciones para obtener los casos de uso
func GetCreateRoleUseCase() *application.CreateRoleUseCase {
	return createRoleUseCase
}

func GetRoleByIDUseCase() *application.GetRoleByIDUseCase {
	return getRoleByIDUseCase
}

func GetAllRolesUseCase() *application.GetAllRolesUseCase {
	return getAllRolesUseCase
}

func GetUpdateRoleUseCase() *application.UpdateRoleUseCase {
	return updateRoleUseCase
}

func GetSoftDeleteRoleUseCase() *application.SoftDeleteRoleUseCase {
	return softDeleteRoleUseCase
}