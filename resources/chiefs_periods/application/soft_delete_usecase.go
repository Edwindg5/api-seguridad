package application

import (
	"context"
	"errors"
	"api-seguridad/resources/chiefs_periods/domain/repository"
)

type SoftDeleteChiefsPeriodUseCase struct {
	repo repository.ChiefsPeriodRepository
}

func NewSoftDeleteChiefsPeriodUseCase(repo repository.ChiefsPeriodRepository) *SoftDeleteChiefsPeriodUseCase {
	return &SoftDeleteChiefsPeriodUseCase{repo: repo}
}

func (uc *SoftDeleteChiefsPeriodUseCase) Execute(ctx context.Context, id uint) error {
	if id == 0 {
		return errors.New("invalid period ID")
	}

	existing, err := uc.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if existing == nil || existing.IsDeleted() {
		return errors.New("period not found")
	}

	return uc.repo.SoftDelete(ctx, id)
}