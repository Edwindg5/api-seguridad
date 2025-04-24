//api-seguridad/resources/police/application/get_by_id_usecase.go
package application

import (
	"context"
	"errors"
	"api-seguridad/resources/police/domain/entities"
	"api-seguridad/resources/police/domain/repository"
)

type GetPoliceByIDUseCase struct {
	repo repository.PoliceRepository
}

func NewGetPoliceByIDUseCase(repo repository.PoliceRepository) *GetPoliceByIDUseCase {
	return &GetPoliceByIDUseCase{repo: repo}
}

func (uc *GetPoliceByIDUseCase) Execute(ctx context.Context, id uint) (*entities.Police, error) {
	if id == 0 {
		return nil, errors.New("invalid police ID")
	}

	police, err := uc.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if police == nil || police.IsDeleted() {
		return nil, errors.New("police not found")
	}

	return police, nil
}