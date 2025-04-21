// api-seguridad/resources/users/infrastructure/dependencies/dependencies.go
package dependencies

import (
	"api-seguridad/core/database"
	"api-seguridad/resources/users/application"
	"api-seguridad/resources/users/domain/repository"
	"api-seguridad/resources/users/infrastructure/adapters"
)

var (
	userRepo repository.UserRepository

	// Casos de uso
	createUserUC       *application.CreateUserUseCase
	getUserByIDUC      *application.GetUserByIDUseCase
	updateUserUC       *application.UpdateUserUseCase
	deleteUserUC       *application.DeleteUserUseCase
	listUsersUC        *application.ListUsersUseCase
	getUserByUsernameUC *application.GetUserByUsernameUseCase
	getUserByEmailUC   *application.GetUserByEmailUseCase
)

func InitDependencies() {
	db := database.GetDB()
	
	// Inicializar repositorio
	userRepo = adapters.NewUserRepository(db)
	
	// Inicializar casos de uso
	createUserUC = application.NewCreateUserUseCase(userRepo)
	getUserByIDUC = application.NewGetUserByIDUseCase(userRepo)
	updateUserUC = application.NewUpdateUserUseCase(userRepo)
	deleteUserUC = application.NewDeleteUserUseCase(userRepo)
	listUsersUC = application.NewListUsersUseCase(userRepo)
	getUserByUsernameUC = application.NewGetUserByUsernameUseCase(userRepo)
	getUserByEmailUC = application.NewGetUserByEmailUseCase(userRepo)
}

// Funciones para obtener casos de uso
func GetCreateUserUseCase() *application.CreateUserUseCase {
	return createUserUC
}

func GetUserByIDUseCase() *application.GetUserByIDUseCase {
	return getUserByIDUC
}

func GetUpdateUserUseCase() *application.UpdateUserUseCase {
	return updateUserUC
}

func GetDeleteUserUseCase() *application.DeleteUserUseCase {
	return deleteUserUC
}

func GetListUsersUseCase() *application.ListUsersUseCase {
	return listUsersUC
}

func GetUserByUsernameUseCase() *application.GetUserByUsernameUseCase {
	return getUserByUsernameUC
}

func GetUserByEmailUseCase() *application.GetUserByEmailUseCase {
	return getUserByEmailUC
}

// Mantenemos esto por compatibilidad si es necesario
func GetUserRepository() repository.UserRepository {
	return userRepo
}