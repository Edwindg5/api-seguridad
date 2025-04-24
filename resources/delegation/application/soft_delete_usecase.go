//api-seguridad/resources/delegation/application/soft_delete_usecase.go
package application

import (
	"context"
	"errors"
	"api-seguridad/resources/delegation/domain/repository"
)

type SoftDeleteDelegationUseCase struct {
	repo repository.DelegationRepository
}

func NewSoftDeleteDelegationUseCase(repo repository.DelegationRepository) *SoftDeleteDelegationUseCase {
	return &SoftDeleteDelegationUseCase{repo: repo}
}

func (uc *SoftDeleteDelegationUseCase) Execute(ctx context.Context, id uint) error {
	if id == 0 {
		return errors.New("invalid delegation ID")
	}

	// Verificar que existe antes de borrar
	delegation, err := uc.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if delegation == nil || delegation.IsDeleted() {
		return errors.New("delegation not found")
	}

	return uc.repo.SoftDelete(ctx, id)
}