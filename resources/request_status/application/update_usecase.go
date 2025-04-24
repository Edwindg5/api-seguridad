// api-seguridad/resources/request_status/application/update_usecase.go
package application

import (
	"context"
	"errors"
	"time"
	"api-seguridad/resources/request_status/domain/entities"
	"api-seguridad/resources/request_status/domain/repository"
)

type UpdateRequestStatusUseCase struct {
	repo repository.RequestStatusRepository
}

func NewUpdateRequestStatusUseCase(repo repository.RequestStatusRepository) *UpdateRequestStatusUseCase {
	return &UpdateRequestStatusUseCase{repo: repo}
}

func (uc *UpdateRequestStatusUseCase) Execute(ctx context.Context, status *entities.RequestStatus) error {
	if status.ID == 0 {
		return errors.New("invalid status ID")
	}
	if status.Name == "" {
		return errors.New("status name is required")
	}
	if status.UpdatedBy == 0 {
		return errors.New("updater user is required")
	}

	// Verify status exists
	existing, err := uc.repo.GetByID(ctx, status.ID)
	if err != nil {
		return err
	}
	if existing == nil || existing.IsDeleted() {
		return errors.New("status not found")
	}

	// Check name uniqueness if changed
	if existing.Name != status.Name {
		existingWithName, err := uc.repo.GetByName(ctx, status.Name)
		if err != nil {
			return err
		}
		if existingWithName != nil {
			return errors.New("status with this name already exists")
		}
	}

	// Update timestamp
	status.UpdatedAt = time.Now()

	return uc.repo.Update(ctx, status)
}