//api-seguridad/resources/police/application/soft_delete_usecase.go
package application

import (
	"context"
	"errors"
	"api-seguridad/resources/police/domain/repository"
)

type SoftDeletePoliceUseCase struct {
	repo repository.PoliceRepository
}

func NewSoftDeletePoliceUseCase(repo repository.PoliceRepository) *SoftDeletePoliceUseCase {
	return &SoftDeletePoliceUseCase{repo: repo}
}

func (uc *SoftDeletePoliceUseCase) Execute(ctx context.Context, id uint) error {
	if id == 0 {
		return errors.New("invalid police ID")
	}

	existing, err := uc.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if existing == nil || existing.IsDeleted() {
		return errors.New("police not found")
	}

	return uc.repo.SoftDelete(ctx, id)
}