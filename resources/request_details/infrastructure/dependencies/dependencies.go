//api-seguridad/resources/request_details/infrastructure/dependencies/dependencies.go
package dependencies

import (
	"api-seguridad/core/database"
	"api-seguridad/resources/request_details/application"
	"api-seguridad/resources/request_details/domain/repository"
	"api-seguridad/resources/request_details/infrastructure/adapters"
)

var (
	requestDetailRepo repository.RequestDetailRepository

	// Use Cases
	createRequestDetailUseCase       *application.CreateRequestDetailUseCase
	getRequestDetailByIDUseCase      *application.GetRequestDetailByIDUseCase
	getByRequestIDUseCase            *application.GetByRequestIDUseCase
	getByPoliceIDUseCase             *application.GetByPoliceIDUseCase
	updateRequestDetailUseCase       *application.UpdateRequestDetailUseCase
	softDeleteRequestDetailUseCase   *application.SoftDeleteRequestDetailUseCase
)

func InitRequestDetailsDependencies() {
	db := database.GetDB()
	
	// Initialize repository
	requestDetailRepo = adapters.NewRequestDetailRepository(db)
	
	// Initialize use cases
	createRequestDetailUseCase = application.NewCreateRequestDetailUseCase(requestDetailRepo)
	getRequestDetailByIDUseCase = application.NewGetRequestDetailByIDUseCase(requestDetailRepo)
	getByRequestIDUseCase = application.NewGetByRequestIDUseCase(requestDetailRepo)
	getByPoliceIDUseCase = application.NewGetByPoliceIDUseCase(requestDetailRepo)
	updateRequestDetailUseCase = application.NewUpdateRequestDetailUseCase(requestDetailRepo)
	softDeleteRequestDetailUseCase = application.NewSoftDeleteRequestDetailUseCase(requestDetailRepo)
}

// Getter functions for use cases
func GetCreateRequestDetailUseCase() *application.CreateRequestDetailUseCase {
	return createRequestDetailUseCase
}

func GetRequestDetailByIDUseCase() *application.GetRequestDetailByIDUseCase {
	return getRequestDetailByIDUseCase
}

func GetByRequestIDUseCase() *application.GetByRequestIDUseCase {
	return getByRequestIDUseCase
}

func GetByPoliceIDUseCase() *application.GetByPoliceIDUseCase {
	return getByPoliceIDUseCase
}

func GetUpdateRequestDetailUseCase() *application.UpdateRequestDetailUseCase {
	return updateRequestDetailUseCase
}

func GetSoftDeleteRequestDetailUseCase() *application.SoftDeleteRequestDetailUseCase {
	return softDeleteRequestDetailUseCase
}

// Repository getter
func GetRequestDetailRepository() repository.RequestDetailRepository {
	return requestDetailRepo
}