package application

import (
	"context"
	"api-seguridad/resources/delegation/domain/entities"
	"api-seguridad/resources/delegation/domain/repository"
)

type GetAllDelegationsUseCase struct {
	repo repository.DelegationRepository
}

func NewGetAllDelegationsUseCase(repo repository.DelegationRepository) *GetAllDelegationsUseCase {
	return &GetAllDelegationsUseCase{repo: repo}
}

func (uc *GetAllDelegationsUseCase) Execute(ctx context.Context) ([]*entities.Delegation, error) {
	delegations, err := uc.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	// Filtro adicional por si acaso
	var activeDelegations []*entities.Delegation
	for _, d := range delegations {
		if !d.IsDeleted() {
			activeDelegations = append(activeDelegations, d)
		}
	}

	return activeDelegations, nil
}