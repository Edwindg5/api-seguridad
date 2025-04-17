package dependencies

import (
	"api-seguridad/core/database"
	"api-seguridad/resources/police/application"
	"api-seguridad/resources/police/domain/repository"
	"api-seguridad/resources/police/infrastructure/adapters"
)

var (
	policeRepo     repository.PoliceRepository
	policeService  *application.PoliceService
)

func InitDependencies() {
	db := database.GetDB()
	
	policeRepo = adapters.NewPoliceRepository(db)
	policeService = application.NewPoliceService(policeRepo)
}

func GetPoliceService() *application.PoliceService {
	return policeService
}