// api-seguridad/resources/request/infrastructure/dependencies/dependencies.go
package dependencies

import (
	"api-seguridad/core/database"
	"api-seguridad/resources/request/application"
	"api-seguridad/resources/request/domain/repository"
	"api-seguridad/resources/request/infrastructure/adapters"
)

var (
	requestRepo repository.RequestRepository

	// Use Cases
	createRequestUseCase          *application.CreateRequestUseCase
	getRequestByIDUseCase         *application.GetRequestByIDUseCase
	updateRequestUseCase          *application.UpdateRequestUseCase
	deleteRequestUseCase          *application.DeleteRequestUseCase
	getRequestsByStatusUseCase    *application.GetRequestsByStatusUseCase
	getRequestsByMunicipalityUseCase *application.GetRequestsByMunicipalityUseCase
	getAllRequestsUseCase         *application.GetAllRequestsUseCase
)

func InitDependencies() {
	db := database.GetDB()
	
	// Initialize repository
	requestRepo = adapters.NewRequestRepository(db)
	
	// Initialize use cases
	createRequestUseCase = application.NewCreateRequestUseCase(requestRepo)
	getRequestByIDUseCase = application.NewGetRequestByIDUseCase(requestRepo)
	updateRequestUseCase = application.NewUpdateRequestUseCase(requestRepo)
	deleteRequestUseCase = application.NewDeleteRequestUseCase(requestRepo)
	getRequestsByStatusUseCase = application.NewGetRequestsByStatusUseCase(requestRepo)
	getRequestsByMunicipalityUseCase = application.NewGetRequestsByMunicipalityUseCase(requestRepo)
	getAllRequestsUseCase = application.NewGetAllRequestsUseCase(requestRepo)
}

// Getter functions for use cases
func GetCreateRequestUseCase() *application.CreateRequestUseCase {
	return createRequestUseCase
}

func GetRequestByIDUseCase() *application.GetRequestByIDUseCase {
	return getRequestByIDUseCase
}

func GetUpdateRequestUseCase() *application.UpdateRequestUseCase {
	return updateRequestUseCase
}

func GetDeleteRequestUseCase() *application.DeleteRequestUseCase {
	return deleteRequestUseCase
}

func GetRequestsByStatusUseCase() *application.GetRequestsByStatusUseCase {
	return getRequestsByStatusUseCase
}

func GetRequestsByMunicipalityUseCase() *application.GetRequestsByMunicipalityUseCase {
	return getRequestsByMunicipalityUseCase
}

func GetGetAllRequestsUseCase() *application.GetAllRequestsUseCase {
	return getAllRequestsUseCase
}

// Repository getter
func GetRequestRepository() repository.RequestRepository {
	return requestRepo
}