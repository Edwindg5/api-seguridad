// api-seguridad/resources/delegation/infrastructure/dependencies/dependencies.go
package dependencies

import (
	"api-seguridad/core/database"
	"api-seguridad/resources/delegation/application"
	"api-seguridad/resources/delegation/domain/repository"
	"api-seguridad/resources/delegation/infrastructure/adapters"
)

var (
	delegationRepo repository.DelegationRepository

	// Use Cases
	createDelegationUseCase     *application.CreateDelegationUseCase
	getDelegationByIDUseCase    *application.GetDelegationByIDUseCase
	getAllDelegationsUseCase    *application.GetAllDelegationsUseCase
	updateDelegationUseCase     *application.UpdateDelegationUseCase
	softDeleteDelegationUseCase *application.SoftDeleteDelegationUseCase
)

func InitDependencies() {
	db := database.GetDB()
	
	// Initialize repository
	delegationRepo = adapters.NewDelegationRepository(db)
	
	// Initialize use cases
	createDelegationUseCase = application.NewCreateDelegationUseCase(delegationRepo)
	getDelegationByIDUseCase = application.NewGetDelegationByIDUseCase(delegationRepo)
	getAllDelegationsUseCase = application.NewGetAllDelegationsUseCase(delegationRepo)
	updateDelegationUseCase = application.NewUpdateDelegationUseCase(delegationRepo)
	softDeleteDelegationUseCase = application.NewSoftDeleteDelegationUseCase(delegationRepo)
}

// Getter functions for use cases
func GetCreateDelegationUseCase() *application.CreateDelegationUseCase {
	return createDelegationUseCase
}

func GetDelegationByIDUseCase() *application.GetDelegationByIDUseCase {
	return getDelegationByIDUseCase
}

func GetAllDelegationsUseCase() *application.GetAllDelegationsUseCase {
	return getAllDelegationsUseCase
}

func GetUpdateDelegationUseCase() *application.UpdateDelegationUseCase {
	return updateDelegationUseCase
}

func GetSoftDeleteDelegationUseCase() *application.SoftDeleteDelegationUseCase {
	return softDeleteDelegationUseCase
}

// Repository getter
func GetDelegationRepository() repository.DelegationRepository {
	return delegationRepo
}