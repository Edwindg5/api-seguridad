package application

import (
	"context"
	"errors"
	"api-seguridad/resources/chiefs_periods/domain/entities"
	"api-seguridad/resources/chiefs_periods/domain/repository"
)

type CreateChiefsPeriodUseCase struct {
	repo repository.ChiefsPeriodRepository
}

func NewCreateChiefsPeriodUseCase(repo repository.ChiefsPeriodRepository) *CreateChiefsPeriodUseCase {
	return &CreateChiefsPeriodUseCase{repo: repo}
}

func (uc *CreateChiefsPeriodUseCase) Execute(ctx context.Context, period *entities.ChiefsPeriod) error {
	// Validaciones básicas
	if period.CeoChiefID == 0 || period.LegalChiefID == 0 {
		return errors.New("both CEO and Legal chief IDs are required")
	}
	if period.StartDate.IsZero() {
		return errors.New("start date is required")
	}
	if !period.EndDate.IsZero() && period.StartDate.After(period.EndDate) {
		return errors.New("start date cannot be after end date")
	}

	// Verificar que el usuario creador exista
	if period.CreatedBy == 0 {
		return errors.New("creator user ID is required")
	}

	// Verificar que no haya un periodo activo si este lo está
	if period.PeriodActive {
		activePeriod, err := uc.repo.GetActivePeriod(ctx)
		if err != nil {
			return err
		}
		if activePeriod != nil {
			return errors.New("there is already an active period")
		}
	}

	return uc.repo.Create(ctx, period)
}