package application

import (
	"context"
	"errors"
	"api-seguridad/resources/type_police/domain/repository"
)

type SoftDeleteTypePoliceUseCase struct {
	repo repository.TypePoliceRepository
}

func NewSoftDeleteTypePoliceUseCase(repo repository.TypePoliceRepository) *SoftDeleteTypePoliceUseCase {
	return &SoftDeleteTypePoliceUseCase{repo: repo}
}

func (uc *SoftDeleteTypePoliceUseCase) Execute(ctx context.Context, id uint) error {
	if id == 0 {
		return errors.New("invalid type police ID")
	}

	// Verificar que existe antes de borrar
	typePolice, err := uc.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if typePolice == nil || typePolice.IsDeleted() {
		return errors.New("type police not found")
	}

	return uc.repo.SoftDelete(ctx, id)
}