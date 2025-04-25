package dependencies

import (
	"api-seguridad/core/database"
	"api-seguridad/resources/chiefs_periods/application"
	"api-seguridad/resources/chiefs_periods/domain/repository"
	"api-seguridad/resources/chiefs_periods/infrastructure/adapters"
)

var (
	chiefsPeriodRepo repository.ChiefsPeriodRepository

	// Use Cases
	createChiefsPeriodUseCase          *application.CreateChiefsPeriodUseCase
	getChiefsPeriodByIDUseCase         *application.GetChiefsPeriodByIDUseCase
	getAllChiefsPeriodsUseCase         *application.GetAllChiefsPeriodsUseCase
	updateChiefsPeriodUseCase          *application.UpdateChiefsPeriodUseCase
	softDeleteChiefsPeriodUseCase      *application.SoftDeleteChiefsPeriodUseCase
	getActiveChiefsPeriodUseCase       *application.GetActiveChiefsPeriodUseCase
	getChiefsPeriodsByDateRangeUseCase *application.GetChiefsPeriodsByDateRangeUseCase
)

func InitChiefsPeriodsDependencies() {
	db := database.GetDB()
	
	// Initialize repository
	chiefsPeriodRepo = adapters.NewChiefsPeriodRepository(db)
	
	// Initialize use cases
	createChiefsPeriodUseCase = application.NewCreateChiefsPeriodUseCase(chiefsPeriodRepo)
	getChiefsPeriodByIDUseCase = application.NewGetChiefsPeriodByIDUseCase(chiefsPeriodRepo)
	getAllChiefsPeriodsUseCase = application.NewGetAllChiefsPeriodsUseCase(chiefsPeriodRepo)
	updateChiefsPeriodUseCase = application.NewUpdateChiefsPeriodUseCase(chiefsPeriodRepo)
	softDeleteChiefsPeriodUseCase = application.NewSoftDeleteChiefsPeriodUseCase(chiefsPeriodRepo)
	getActiveChiefsPeriodUseCase = application.NewGetActiveChiefsPeriodUseCase(chiefsPeriodRepo)
	getChiefsPeriodsByDateRangeUseCase = application.NewGetChiefsPeriodsByDateRangeUseCase(chiefsPeriodRepo)
}

// Getter functions for use cases
func GetCreateChiefsPeriodUseCase() *application.CreateChiefsPeriodUseCase {
	return createChiefsPeriodUseCase
}

func GetChiefsPeriodByIDUseCase() *application.GetChiefsPeriodByIDUseCase {
	return getChiefsPeriodByIDUseCase
}

func GetAllChiefsPeriodsUseCase() *application.GetAllChiefsPeriodsUseCase {
	return getAllChiefsPeriodsUseCase
}

func GetUpdateChiefsPeriodUseCase() *application.UpdateChiefsPeriodUseCase {
	return updateChiefsPeriodUseCase
}

func GetSoftDeleteChiefsPeriodUseCase() *application.SoftDeleteChiefsPeriodUseCase {
	return softDeleteChiefsPeriodUseCase
}

func GetActiveChiefsPeriodUseCase() *application.GetActiveChiefsPeriodUseCase {
	return getActiveChiefsPeriodUseCase
}

func GetChiefsPeriodsByDateRangeUseCase() *application.GetChiefsPeriodsByDateRangeUseCase {
	return getChiefsPeriodsByDateRangeUseCase
}

// Repository getter
func GetChiefsPeriodRepository() repository.ChiefsPeriodRepository {
	return chiefsPeriodRepo
}