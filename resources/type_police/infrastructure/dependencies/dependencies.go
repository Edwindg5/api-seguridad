//api-seguridad/resources/type_police/infrastructure/dependencies/dependencies.go
package dependencies

import (
	"api-seguridad/core/database"
	"api-seguridad/resources/type_police/application"
	"api-seguridad/resources/type_police/domain/repository"
	"api-seguridad/resources/type_police/infrastructure/adapters"
)

var (
	typePoliceRepo repository.TypePoliceRepository

	// Use Cases
	createTypePoliceUseCase     *application.CreateTypePoliceUseCase
	getTypePoliceByIDUseCase    *application.GetTypePoliceByIDUseCase
	getAllTypePoliceUseCase     *application.GetAllTypePoliceUseCase
	updateTypePoliceUseCase     *application.UpdateTypePoliceUseCase
	softDeleteTypePoliceUseCase *application.SoftDeleteTypePoliceUseCase
)

func InitDependencies() {
	db := database.GetDB()
	
	// Initialize repository
	typePoliceRepo = adapters.NewTypePoliceRepository(db)
	
	// Initialize use cases
	createTypePoliceUseCase = application.NewCreateTypePoliceUseCase(typePoliceRepo)
	getTypePoliceByIDUseCase = application.NewGetTypePoliceByIDUseCase(typePoliceRepo)
	getAllTypePoliceUseCase = application.NewGetAllTypePoliceUseCase(typePoliceRepo)
	updateTypePoliceUseCase = application.NewUpdateTypePoliceUseCase(typePoliceRepo)
	softDeleteTypePoliceUseCase = application.NewSoftDeleteTypePoliceUseCase(typePoliceRepo)
}

// Getter functions for use cases
func GetCreateTypePoliceUseCase() *application.CreateTypePoliceUseCase {
	return createTypePoliceUseCase
}

func GetTypePoliceByIDUseCase() *application.GetTypePoliceByIDUseCase {
	return getTypePoliceByIDUseCase
}

func GetAllTypePoliceUseCase() *application.GetAllTypePoliceUseCase {
	return getAllTypePoliceUseCase
}

func GetUpdateTypePoliceUseCase() *application.UpdateTypePoliceUseCase {
	return updateTypePoliceUseCase
}

func GetSoftDeleteTypePoliceUseCase() *application.SoftDeleteTypePoliceUseCase {
	return softDeleteTypePoliceUseCase
}

// Repository getter
func GetTypePoliceRepository() repository.TypePoliceRepository {
	return typePoliceRepo
}