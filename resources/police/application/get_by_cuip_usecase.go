package application

import (
	"context"
	"errors"
	"api-seguridad/resources/police/domain/entities"
	"api-seguridad/resources/police/domain/repository"
)

type GetPoliceByCUIPUseCase struct {
	repo repository.PoliceRepository
}

func NewGetPoliceByCUIPUseCase(repo repository.PoliceRepository) *GetPoliceByCUIPUseCase {
	return &GetPoliceByCUIPUseCase{repo: repo}
}

func (uc *GetPoliceByCUIPUseCase) Execute(ctx context.Context, cuip string) (*entities.Police, error) {
	if cuip == "" {
		return nil, errors.New("CUIP is required")
	}

	police, err := uc.repo.GetByCUIP(ctx, cuip)
	if err != nil {
		return nil, err
	}

	if police == nil || police.IsDeleted() {
		return nil, errors.New("police not found")
	}

	return police, nil
}