//api-seguridad/resources/type_police/application/get_by_id_usecase.go
package application

import (
	"context"
	"errors"
	"api-seguridad/resources/type_police/domain/entities"
	"api-seguridad/resources/type_police/domain/repository"
)

type GetTypePoliceByIDUseCase struct {
	repo repository.TypePoliceRepository
}

func NewGetTypePoliceByIDUseCase(repo repository.TypePoliceRepository) *GetTypePoliceByIDUseCase {
	return &GetTypePoliceByIDUseCase{repo: repo}
}

func (uc *GetTypePoliceByIDUseCase) Execute(ctx context.Context, id uint) (*entities.TypePolice, error) {
	if id == 0 {
		return nil, errors.New("invalid type police ID")
	}

	typePolice, err := uc.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if typePolice == nil || typePolice.IsDeleted() {
		return nil, errors.New("type police not found")
	}

	return typePolice, nil
}