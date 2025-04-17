package dependencies

import (
	"api-seguridad/core/database"
	"api-seguridad/resources/request/application"
	"api-seguridad/resources/request/domain/repository"
	"api-seguridad/resources/request/infrastructure/adapters"
)

var (
	requestRepo     repository.RequestRepository
	requestService  *application.RequestService
)

func InitDependencies() {
	db := database.GetDB()
	
	requestRepo = adapters.NewRequestRepository(db)
	requestService = application.NewRequestService(requestRepo)
}

func GetRequestService() *application.RequestService {
	return requestService
}