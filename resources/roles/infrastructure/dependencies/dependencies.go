package dependencies

import (
	"api-seguridad/core/database"
	"api-seguridad/resources/roles/application"
	"api-seguridad/resources/roles/domain/repository"
	"api-seguridad/resources/roles/infrastructure/adapters"
)

var (
	roleRepo     repository.RoleRepository
	roleService  *application.RoleService
)

func InitDependencies() {
	db := database.GetDB()
	
	roleRepo = adapters.NewRoleRepository(db)
	roleService = application.NewRoleService(roleRepo)
}

func GetRoleService() *application.RoleService {
	return roleService
}