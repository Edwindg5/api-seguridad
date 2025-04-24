// api-seguridad/resources/area_chiefs/infrastructure/dependencies/dependencies.go
package dependencies

import (
	"api-seguridad/core/database"
	"api-seguridad/resources/area_chiefs/application"
	"api-seguridad/resources/area_chiefs/domain/repository"
	"api-seguridad/resources/area_chiefs/infrastructure/adapters"
)

var (
	areaChiefRepo repository.AreaChiefRepository

	// Use Cases
	createAreaChiefUseCase    *application.CreateAreaChiefUseCase
	getAreaChiefByIDUseCase   *application.GetAreaChiefByIDUseCase
	getAllAreaChiefsUseCase   *application.GetAllAreaChiefsUseCase
	updateAreaChiefUseCase    *application.UpdateAreaChiefUseCase
	deleteAreaChiefUseCase    *application.DeleteAreaChiefUseCase
)

func InitDependencies() {
	db := database.GetDB()
	
	// Initialize repository
	areaChiefRepo = adapters.NewAreaChiefRepository(db)
	
	// Initialize use cases
	createAreaChiefUseCase = application.NewCreateAreaChiefUseCase(areaChiefRepo)
	getAreaChiefByIDUseCase = application.NewGetAreaChiefByIDUseCase(areaChiefRepo)
	getAllAreaChiefsUseCase = application.NewGetAllAreaChiefsUseCase(areaChiefRepo)
	updateAreaChiefUseCase = application.NewUpdateAreaChiefUseCase(areaChiefRepo)
	deleteAreaChiefUseCase = application.NewDeleteAreaChiefUseCase(areaChiefRepo)
}

// Getters for use cases
func GetCreateAreaChiefUseCase() *application.CreateAreaChiefUseCase {
	return createAreaChiefUseCase
}

func GetAreaChiefByIDUseCase() *application.GetAreaChiefByIDUseCase {
	return getAreaChiefByIDUseCase
}

func GetAllAreaChiefsUseCase() *application.GetAllAreaChiefsUseCase {
	return getAllAreaChiefsUseCase
}

func GetUpdateAreaChiefUseCase() *application.UpdateAreaChiefUseCase {
	return updateAreaChiefUseCase
}

func GetDeleteAreaChiefUseCase() *application.DeleteAreaChiefUseCase {
	return deleteAreaChiefUseCase
}

// Repository getter (for testing or extended use)
func GetAreaChiefRepository() repository.AreaChiefRepository {
	return areaChiefRepo
}