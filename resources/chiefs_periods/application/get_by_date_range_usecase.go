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
	if start.IsZero() || end.IsZero() {
		return nil, errors.New("both start and end dates are required")
	}
	if start.After(end) {
		return nil, errors.New("start date cannot be after end date")
	}

	return uc.repo.GetPeriodsByDateRange(ctx, start, end)
}