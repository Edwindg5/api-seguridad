// api-seguridad/resources/police/infrastructure/dependencies/dependencies.go
package dependencies

import (
	"api-seguridad/core/database"
	"api-seguridad/resources/police/application"
	"api-seguridad/resources/police/domain/repository"
	"api-seguridad/resources/police/infrastructure/adapters"
)

var (
	policeRepo repository.PoliceRepository

	// Use Cases
	createPoliceUseCase     *application.CreatePoliceUseCase
	getPoliceByIDUseCase    *application.GetPoliceByIDUseCase
	getAllPoliceUseCase     *application.GetAllPoliceUseCase
	updatePoliceUseCase     *application.UpdatePoliceUseCase
	softDeletePoliceUseCase *application.SoftDeletePoliceUseCase
	getPoliceByCUIPUseCase  *application.GetPoliceByCUIPUseCase
	searchPoliceByNameUseCase *application.SearchPoliceByNameUseCase
)

func InitDependencies() {
	db := database.GetDB()
	
	// Initialize repository
	policeRepo = adapters.NewPoliceRepository(db)
	
	// Initialize use cases
	createPoliceUseCase = application.NewCreatePoliceUseCase(policeRepo)
	getPoliceByIDUseCase = application.NewGetPoliceByIDUseCase(policeRepo)
	getAllPoliceUseCase = application.NewGetAllPoliceUseCase(policeRepo)
	updatePoliceUseCase = application.NewUpdatePoliceUseCase(policeRepo)
	softDeletePoliceUseCase = application.NewSoftDeletePoliceUseCase(policeRepo)
	getPoliceByCUIPUseCase = application.NewGetPoliceByCUIPUseCase(policeRepo)
	searchPoliceByNameUseCase = application.NewSearchPoliceByNameUseCase(policeRepo)
}

// Getter functions for use cases
func GetCreatePoliceUseCase() *application.CreatePoliceUseCase {
	return createPoliceUseCase
}

func GetPoliceByIDUseCase() *application.GetPoliceByIDUseCase {
	return getPoliceByIDUseCase
}

func GetAllPoliceUseCase() *application.GetAllPoliceUseCase {
	return getAllPoliceUseCase
}

func GetUpdatePoliceUseCase() *application.UpdatePoliceUseCase {
	return updatePoliceUseCase
}

func GetSoftDeletePoliceUseCase() *application.SoftDeletePoliceUseCase {
	return softDeletePoliceUseCase
}

func GetPoliceByCUIPUseCase() *application.GetPoliceByCUIPUseCase {
	return getPoliceByCUIPUseCase
}

func GetSearchPoliceByNameUseCase() *application.SearchPoliceByNameUseCase {
	return searchPoliceByNameUseCase
}

// Repository getter
func GetPoliceRepository() repository.PoliceRepository {
	return policeRepo
}