//api-seguridad/resources/delegation/application/create_usecase.go
package application

import (
	"context"
	"errors"
	"api-seguridad/resources/delegation/domain/entities"
	"api-seguridad/resources/delegation/domain/repository"
)

type CreateDelegationUseCase struct {
	repo repository.DelegationRepository
}

func NewCreateDelegationUseCase(repo repository.DelegationRepository) *CreateDelegationUseCase {
	return &CreateDelegationUseCase{repo: repo}
}

func (uc *CreateDelegationUseCase) Execute(ctx context.Context, delegation *entities.Delegation) error {
	if delegation.Name == "" {
		return errors.New("delegation name is required")
	}

	return uc.repo.Create(ctx, delegation)
}