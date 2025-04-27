//api-seguridad/resources/chiefs_periods/application/get_by_id_usecase.go
package application

import (
	"context"
	"errors"
	"api-seguridad/resources/chiefs_periods/domain/entities"
	"api-seguridad/resources/chiefs_periods/domain/repository"
)

type GetChiefsPeriodByIDUseCase struct {
	repo repository.ChiefsPeriodRepository
}

func NewGetChiefsPeriodByIDUseCase(repo repository.ChiefsPeriodRepository) *GetChiefsPeriodByIDUseCase {
	return &GetChiefsPeriodByIDUseCase{repo: repo}
}

func (uc *GetChiefsPeriodByIDUseCase) Execute(ctx context.Context, id uint) (*entities.ChiefsPeriod, error) {
	if id == 0 {
		return nil, errors.New("invalid period ID")
	}

	period, err := uc.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if period == nil || period.IsDeleted() {
		return nil, errors.New("period not found")
	}

	return period, nil
}