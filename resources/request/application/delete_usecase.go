// api-seguridad/resources/request/application/delete_usecase.go
package application

import (
	"context"
	"errors"
	"api-seguridad/resources/request/domain/repository"
)

type DeleteRequestUseCase struct {
	repo repository.RequestRepository
}

func NewDeleteRequestUseCase(repo repository.RequestRepository) *DeleteRequestUseCase {
	return &DeleteRequestUseCase{repo: repo}
}

func (uc *DeleteRequestUseCase) Execute(ctx context.Context, id uint) error {
	if id == 0 {
		return errors.New("invalid request ID")
	}

	// Verify request exists
	existing, err := uc.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if existing == nil || existing.IsDeleted() {
		return errors.New("request not found")
	}

	return uc.repo.Delete(ctx, id)
}