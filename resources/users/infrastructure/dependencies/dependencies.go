package dependencies

import (
	"api-seguridad/core/database"
	"api-seguridad/resources/users/application"
	"api-seguridad/resources/users/domain/repository"
	"api-seguridad/resources/users/infrastructure/adapters"
)

var (
	userRepo     repository.UserRepository
	userService  *application.UserService
)

func InitDependencies() {
	db := database.GetDB()
	
	// Inicializar repositorios
	userRepo = adapters.NewUserRepository(db)
	
	// Inicializar servicios
	userService = application.NewUserService(userRepo)
}

func GetUserService() *application.UserService {
	return userService
}