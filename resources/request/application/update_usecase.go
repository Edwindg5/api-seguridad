// api-seguridad/resources/request/application/update_usecase.go
package application

import (
	"context"
	"errors"
	"api-seguridad/resources/request/domain/entities"
	"api-seguridad/resources/request/domain/repository"
	"time"
)

type UpdateRequestUseCase struct {
	repo repository.RequestRepository
}

func NewUpdateRequestUseCase(repo repository.RequestRepository) *UpdateRequestUseCase {
	return &UpdateRequestUseCase{repo: repo}
}

func (uc *UpdateRequestUseCase) Execute(ctx context.Context, request *entities.Request) error {
	if request.ID == 0 {
		return errors.New("invalid request ID")
	}
	if request.OfficeNumber == "" {
		return errors.New("office number is required")
	}
	if request.UpdatedBy == 0 {
		return errors.New("updater user is required")
	}

	// Verify request exists
	existing, err := uc.repo.GetByID(ctx, request.ID)
	if err != nil {
		return err
	}
	if existing == nil || existing.IsDeleted() {
		return errors.New("request not found")
	}

	// Update timestamp
	request.UpdatedAt = time.Now()

	return uc.repo.Update(ctx, request)
}