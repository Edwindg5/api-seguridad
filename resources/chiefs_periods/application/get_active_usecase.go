//api-seguridad/resources/chiefs_periods/application/get_active_usecase.go
package application

import (
	"context"
	"api-seguridad/resources/chiefs_periods/domain/entities"
	"api-seguridad/resources/chiefs_periods/domain/repository"
)

type GetActiveChiefsPeriodUseCase struct {
	repo repository.ChiefsPeriodRepository
}

func NewGetActiveChiefsPeriodUseCase(repo repository.ChiefsPeriodRepository) *GetActiveChiefsPeriodUseCase {
	return &GetActiveChiefsPeriodUseCase{repo: repo}
}

func (uc *GetActiveChiefsPeriodUseCase) Execute(ctx context.Context) (*entities.ChiefsPeriod, error) {
	period, err := uc.repo.GetActivePeriod(ctx)
	if err != nil {
		return nil, err
	}
	if period == nil {
		return nil, nil // Retornar nil sin error para indicar "no encontrado"
	}
	return period, nil
}