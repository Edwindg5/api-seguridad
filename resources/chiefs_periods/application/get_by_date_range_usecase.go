//api-seguridad/resources/chiefs_periods/application/get_by_date_range_usecase.go
package application

import (
	"context"
	"errors"
	"time"
	"api-seguridad/resources/chiefs_periods/domain/entities"
	"api-seguridad/resources/chiefs_periods/domain/repository"
)

type GetChiefsPeriodsByDateRangeUseCase struct {
	repo repository.ChiefsPeriodRepository
}

func NewGetChiefsPeriodsByDateRangeUseCase(repo repository.ChiefsPeriodRepository) *GetChiefsPeriodsByDateRangeUseCase {
	return &GetChiefsPeriodsByDateRangeUseCase{repo: repo}
}

func (uc *GetChiefsPeriodsByDateRangeUseCase) Execute(ctx context.Context, start, end time.Time) ([]*entities.ChiefsPeriod, error) {
    // Validaciones adicionales
    if start.IsZero() || end.IsZero() {
        return nil, errors.New("dates cannot be zero")
    }
    
    // Asegurar que las fechas estén en UTC
    start = start.UTC()
    end = end.UTC()
    
    return uc.repo.GetPeriodsByDateRange(ctx, start, end)
}