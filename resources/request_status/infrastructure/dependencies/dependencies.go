// api-seguridad/resources/request_status/infrastructure/dependencies/dependencies.go
package dependencies

import (
	"api-seguridad/core/database"
	"api-seguridad/resources/request_status/application"
	"api-seguridad/resources/request_status/domain/repository"
	"api-seguridad/resources/request_status/infrastructure/adapters"
)

var (
	requestStatusRepo repository.RequestStatusRepository

	// Use Cases
	createRequestStatusUseCase     *application.CreateRequestStatusUseCase
	getRequestStatusByIDUseCase    *application.GetRequestStatusByIDUseCase
	getAllRequestStatusUseCase     *application.GetAllRequestStatusUseCase
	updateRequestStatusUseCase     *application.UpdateRequestStatusUseCase
	deleteRequestStatusUseCase     *application.DeleteRequestStatusUseCase
)

func InitDependencies() {
	db := database.GetDB()
	
	// Initialize repository
	requestStatusRepo = adapters.NewRequestStatusRepository(db)
	
	// Initialize use cases
	createRequestStatusUseCase = application.NewCreateRequestStatusUseCase(requestStatusRepo)
	getRequestStatusByIDUseCase = application.NewGetRequestStatusByIDUseCase(requestStatusRepo)
	getAllRequestStatusUseCase = application.NewGetAllRequestStatusUseCase(requestStatusRepo)
	updateRequestStatusUseCase = application.NewUpdateRequestStatusUseCase(requestStatusRepo)
	deleteRequestStatusUseCase = application.NewDeleteRequestStatusUseCase(requestStatusRepo)
}

// Getter functions for use cases
func GetCreateRequestStatusUseCase() *application.CreateRequestStatusUseCase {
	return createRequestStatusUseCase
}

func GetRequestStatusByIDUseCase() *application.GetRequestStatusByIDUseCase {
	return getRequestStatusByIDUseCase
}

func GetAllRequestStatusUseCase() *application.GetAllRequestStatusUseCase {
	return getAllRequestStatusUseCase
}

func GetUpdateRequestStatusUseCase() *application.UpdateRequestStatusUseCase {
	return updateRequestStatusUseCase
}

func GetDeleteRequestStatusUseCase() *application.DeleteRequestStatusUseCase {
	return deleteRequestStatusUseCase
}

// Repository getter
func GetRequestStatusRepository() repository.RequestStatusRepository {
	return requestStatusRepo
}