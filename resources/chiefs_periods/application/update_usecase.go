//api-seguridad/resources/chiefs_periods/application/update_usecase.go
package application

import (
	"context"
	"errors"
	"time"
	"api-seguridad/resources/chiefs_periods/domain/entities"
	"api-seguridad/resources/chiefs_periods/domain/repository"
)

type UpdateChiefsPeriodUseCase struct {
	repo repository.ChiefsPeriodRepository
}

func NewUpdateChiefsPeriodUseCase(repo repository.ChiefsPeriodRepository) *UpdateChiefsPeriodUseCase {
	return &UpdateChiefsPeriodUseCase{repo: repo}
}

func (uc *UpdateChiefsPeriodUseCase) Execute(ctx context.Context, period *entities.ChiefsPeriod) error {
	if period.ID == 0 {
		return errors.New("invalid period ID")
	}
	if period.CeoChiefID == 0 || period.LegalChiefID == 0 {
		return errors.New("both CEO and Legal chief IDs are required")
	}

	existing, err := uc.repo.GetByID(ctx, period.ID)
	if err != nil {
		return err
	}
	if existing == nil || existing.IsDeleted() {
		return errors.New("period not found")
	}

	// Preservar el created_by del registro existente
	period.CreatedBy = existing.CreatedBy

	// Validar fechas
	if !period.EndDate.IsZero() && period.StartDate.After(period.EndDate) {
		return errors.New("start date cannot be after end date")
	}

	// Si se está activando, verificar que no haya otro activo
	if period.PeriodActive && !existing.PeriodActive {
		activePeriod, err := uc.repo.GetActivePeriod(ctx)
		if err != nil {
			return err
		}
		if activePeriod != nil && activePeriod.ID != period.ID {
			return errors.New("there is already an active period")
		}
	}

	period.UpdatedAt = time.Now()
	return uc.repo.Update(ctx, period)
}