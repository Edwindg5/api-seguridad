// api-seguridad/resources/request_status/application/delete_usecase.go
package application

import (
	"context"
	"errors"
	"api-seguridad/resources/request_status/domain/repository"
)

type DeleteRequestStatusUseCase struct {
	repo repository.RequestStatusRepository
}

func NewDeleteRequestStatusUseCase(repo repository.RequestStatusRepository) *DeleteRequestStatusUseCase {
	return &DeleteRequestStatusUseCase{repo: repo}
}

func (uc *DeleteRequestStatusUseCase) Execute(ctx context.Context, id uint) error {
	if id == 0 {
		return errors.New("invalid status ID")
	}

	existing, err := uc.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if existing == nil || existing.IsDeleted() {
		return errors.New("status not found")
	}

	return uc.repo.Delete(ctx, id)
}