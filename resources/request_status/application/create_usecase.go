// api-seguridad/resources/request_status/application/create_usecase.go
package application

import (
	"context"
	"errors"
	"api-seguridad/resources/request_status/domain/entities"
	"api-seguridad/resources/request_status/domain/repository"
	"time"
)

type CreateRequestStatusUseCase struct {
	repo repository.RequestStatusRepository
}

func NewCreateRequestStatusUseCase(repo repository.RequestStatusRepository) *CreateRequestStatusUseCase {
	return &CreateRequestStatusUseCase{repo: repo}
}

func (uc *CreateRequestStatusUseCase) Execute(ctx context.Context, status *entities.RequestStatus) error {
	// Validations
	if status.Name == "" {
		return errors.New("status name is required")
	}
	if status.CreatedBy == 0 {
		return errors.New("creator user is required")
	}

	// Check name uniqueness
	existing, err := uc.repo.GetByName(ctx, status.Name)
	if err != nil {
		return err
	}
	if existing != nil {
		return errors.New("status with this name already exists")
	}

	// Set timestamps
	if status.CreatedAt.IsZero() {
		status.CreatedAt = time.Now()
	}
	status.UpdatedAt = time.Now()

	return uc.repo.Create(ctx, status)
}