//api-seguridad/resources/delegation/application/update_usecase.go
package application

import (
	"context"
	"errors"
	"api-seguridad/resources/delegation/domain/entities"
	"api-seguridad/resources/delegation/domain/repository"
)

type UpdateDelegationUseCase struct {
	repo repository.DelegationRepository
}

func NewUpdateDelegationUseCase(repo repository.DelegationRepository) *UpdateDelegationUseCase {
	return &UpdateDelegationUseCase{repo: repo}
}

func (uc *UpdateDelegationUseCase) GetExistingDelegation(ctx context.Context, id uint) (*entities.Delegation, error) {
	return uc.repo.GetByID(ctx, id)
}

func (uc *UpdateDelegationUseCase) Execute(ctx context.Context, delegation *entities.Delegation) error {
	if delegation.ID == 0 {
		return errors.New("invalid delegation ID")
	}
	if delegation.Name == "" {
		return errors.New("delegation name is required")
	}

	existingDelegation, err := uc.repo.GetByID(ctx, delegation.ID)
	if err != nil {
		return err
	}
	if existingDelegation == nil || existingDelegation.IsDeleted() {
		return errors.New("delegation not found")
	}

	delegation.SetCreatedBy(existingDelegation.GetCreatedBy())
	delegation.SetCreatedAt(existingDelegation.GetCreatedAt())

	return uc.repo.Update(ctx, delegation)
}