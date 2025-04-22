package application

import (
	"context"
	"errors"
	"api-seguridad/resources/delegation/domain/entities"
	"api-seguridad/resources/delegation/domain/repository"
)

type GetDelegationByIDUseCase struct {
	repo repository.DelegationRepository
}

func NewGetDelegationByIDUseCase(repo repository.DelegationRepository) *GetDelegationByIDUseCase {
	return &GetDelegationByIDUseCase{repo: repo}
}

func (uc *GetDelegationByIDUseCase) Execute(ctx context.Context, id uint) (*entities.Delegation, error) {
	if id == 0 {
		return nil, errors.New("invalid delegation ID")
	}

	delegation, err := uc.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if delegation == nil || delegation.IsDeleted() {
		return nil, errors.New("delegation not found")
	}

	return delegation, nil
}