package application

import (
	"context"
	"errors"
	"api-seguridad/resources/request_details/domain/repository"
)

type SoftDeleteRequestDetailUseCase struct {
	repo repository.RequestDetailRepository
}

func NewSoftDeleteRequestDetailUseCase(repo repository.RequestDetailRepository) *SoftDeleteRequestDetailUseCase {
	return &SoftDeleteRequestDetailUseCase{repo: repo}
}

func (uc *SoftDeleteRequestDetailUseCase) Execute(ctx context.Context, id uint, deletedBy uint) error {
	if id == 0 {
		return errors.New("ID inválido")
	}
	if deletedBy == 0 {
		return errors.New("usuario que elimina es requerido")
	}

	existing, err := uc.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if existing == nil || existing.IsDeleted() {
		return errors.New("detalle de solicitud no encontrado")
	}

	// Actualizar el campo updated_by antes de eliminar
	existing.SetUpdatedBy(deletedBy)
	if err := uc.repo.Update(ctx, existing); err != nil {
		return err
	}

	return uc.repo.Delete(ctx, id)
}