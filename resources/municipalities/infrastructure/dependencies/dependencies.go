//api-seguridad/resources/municipalities/infrastructure/dependencies/dependencies.go
package dependencies

import (
	"api-seguridad/core/database"
	"api-seguridad/resources/municipalities/application"
	"api-seguridad/resources/municipalities/domain/repository"
	"api-seguridad/resources/municipalities/infrastructure/adapters"
)

var (
	municipalityRepo     repository.MunicipalityRepository
	municipalityService  *application.MunicipalityService
)

func InitDependencies() {
	db := database.GetDB()
	
	municipalityRepo = adapters.NewMunicipalityRepository(db)
	municipalityService = application.NewMunicipalityService(municipalityRepo)
}

func GetMunicipalityService() *application.MunicipalityService {
	return municipalityService
}