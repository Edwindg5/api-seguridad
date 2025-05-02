// api-seguridad/resources/area_chiefs/application/update_usecase.go
package application

import (
	"context"
	"errors"
	"time"

	"api-seguridad/resources/area_chiefs/domain/repository"
)


type AreaChiefUpdate struct {
	ID        uint
	Name      string
	Position  string
	Type      string
	Signature string 
	UpdatedBy uint
}

type UpdateAreaChiefUseCase struct {
	repo repository.AreaChiefRepository
}

func NewUpdateAreaChiefUseCase(repo repository.AreaChiefRepository) *UpdateAreaChiefUseCase {
	return &UpdateAreaChiefUseCase{repo: repo}
}

func (uc *UpdateAreaChiefUseCase) Execute(ctx context.Context, update *AreaChiefUpdate) error {
	if update.ID == 0 {
		return errors.New("invalid chief ID")
	}
	if update.Name == "" {
		return errors.New("chief name is required")
	}
	if update.UpdatedBy == 0 {
		return errors.New("updater user is required")
	}

	// Verify chief exists
	existing, err := uc.repo.GetByID(ctx, update.ID)
	if err != nil {
		return err
	}
	if existing == nil || existing.IsDeleted() {
		return errors.New("area chief not found")
	}


	existing.Name = update.Name
	existing.Position = update.Position
	existing.Type = update.Type
	existing.UpdatedBy = update.UpdatedBy
	existing.UpdatedAt = time.Now()

	if update.Signature != "" {
		existing.SignaturePath = update.Signature
	}

	return uc.repo.Update(ctx, existing)
}