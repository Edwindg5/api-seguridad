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

// Nuevo método para obtener por ID
func (uc *UpdateRequestStatusUseCase) GetByID(ctx context.Context, id uint) (*entities.RequestStatus, error) {
	if id == 0 {
		return nil, errors.New("invalid status ID")
	}

	status, err := uc.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if status == nil || status.IsDeleted() {
		return nil, errors.New("status not found")
	}

	return status, nil
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

	// Verificar existencia
	existing, err := uc.GetByID(ctx, status.ID)
	if err != nil {
		return err
	}

	// Verificar unicidad del nombre si cambió
	if existing.Name != status.Name {
		existingWithName, err := uc.repo.GetByName(ctx, status.Name)
		if err != nil {
			return err
		}
		if existingWithName != nil {
			return errors.New("status with this name already exists")
		}
	}

	// Actualizar timestamp
	status.UpdatedAt = time.Now()

	return uc.repo.Update(ctx, status)
}