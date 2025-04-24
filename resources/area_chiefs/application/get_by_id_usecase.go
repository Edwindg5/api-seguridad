// api-seguridad/resources/area_chiefs/application/get_by_id_usecase.go
package application

import (
	"context"
	"errors"
	"api-seguridad/resources/area_chiefs/domain/entities"
	"api-seguridad/resources/area_chiefs/domain/repository"
)

type GetAreaChiefByIDUseCase struct {
	repo repository.AreaChiefRepository
}

func NewGetAreaChiefByIDUseCase(repo repository.AreaChiefRepository) *GetAreaChiefByIDUseCase {
	return &GetAreaChiefByIDUseCase{repo: repo}
}

func (uc *GetAreaChiefByIDUseCase) Execute(ctx context.Context, id uint) (*entities.AreaChief, error) {
	if id == 0 {
		return nil, errors.New("invalid chief ID")
	}

	chief, err := uc.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if chief == nil || chief.IsDeleted() {
		return nil, errors.New("area chief not found")
	}

	return chief, nil
}