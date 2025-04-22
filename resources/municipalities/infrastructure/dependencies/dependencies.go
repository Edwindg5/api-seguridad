//api-seguridad/resources/municipalities/infrastructure/dependencies/dependencies.go
package dependencies

import (
	"api-seguridad/core/database"
	"api-seguridad/resources/municipalities/application"
	"api-seguridad/resources/municipalities/domain/repository"
	"api-seguridad/resources/municipalities/infrastructure/adapters"
)

var (
	municipalityRepo repository.MunicipalityRepository

	// Casos de uso
	createUseCase      *application.PostMunicipalityUseCase
	getByIDUseCase     *application.GetMunicipalityByIDUseCase
	getByNameUseCase   *application.GetMunicipalityByNameUseCase
	getAllUseCase      *application.GetAllMunicipalitiesUseCase
	updateUseCase      *application.UpdateMunicipalityUseCase
	softDeleteUseCase  *application.SoftDeleteMunicipalityUseCase
)

func InitDependencies() {
	db := database.GetDB()
	
	// Inicializar repositorio
	municipalityRepo = adapters.NewMunicipalityRepository(db)
	
	// Inicializar casos de uso
	createUseCase = application.NewPostMunicipalityUseCase(municipalityRepo)
	getByIDUseCase = application.NewGetMunicipalityByIDUseCase(municipalityRepo)
	getByNameUseCase = application.NewGetMunicipalityByNameUseCase(municipalityRepo)
	getAllUseCase = application.NewGetAllMunicipalitiesUseCase(municipalityRepo)
	updateUseCase = application.NewUpdateMunicipalityUseCase(municipalityRepo)
	softDeleteUseCase = application.NewSoftDeleteMunicipalityUseCase(municipalityRepo)
}

// Funciones para obtener los casos de uso
func GetCreateUseCase() *application.PostMunicipalityUseCase {
	return createUseCase
}

func GetByIDUseCase() *application.GetMunicipalityByIDUseCase {
	return getByIDUseCase
}

func GetByNameUseCase() *application.GetMunicipalityByNameUseCase {
	return getByNameUseCase
}

func GetAllUseCase() *application.GetAllMunicipalitiesUseCase {
	return getAllUseCase
}

func GetUpdateUseCase() *application.UpdateMunicipalityUseCase {
	return updateUseCase
}

func GetSoftDeleteUseCase() *application.SoftDeleteMunicipalityUseCase {
	return softDeleteUseCase
}

// Función para obtener el repositorio (si es necesario)
func GetMunicipalityRepository() repository.MunicipalityRepository {
	return municipalityRepo
}