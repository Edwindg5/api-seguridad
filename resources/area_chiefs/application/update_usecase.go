// api-seguridad/resources/area_chiefs/application/update_usecase.go
package application

import (
	"context"
	"errors"
	"time"
	"api-seguridad/resources/area_chiefs/domain/entities"
	"api-seguridad/resources/area_chiefs/domain/repository"
)

type UpdateAreaChiefUseCase struct {
	repo repository.AreaChiefRepository
}

func NewUpdateAreaChiefUseCase(repo repository.AreaChiefRepository) *UpdateAreaChiefUseCase {
	return &UpdateAreaChiefUseCase{repo: repo}
}

func (uc *UpdateAreaChiefUseCase) Execute(ctx context.Context, chief *entities.AreaChief) error {
	if chief.ID == 0 {
		return errors.New("invalid chief ID")
	}
	if chief.Name == "" {
		return errors.New("chief name is required")
	}
	if chief.UpdatedBy == 0 {
		return errors.New("updater user is required")
	}

	// Verify chief exists
	existing, err := uc.repo.GetByID(ctx, chief.ID)
	if err != nil {
		return err
	}
	if existing == nil || existing.IsDeleted() {
		return errors.New("area chief not found")
	}

	// Update timestamp
	chief.UpdatedAt = time.Now()

	return uc.repo.Update(ctx, chief)
}