//api-seguridad/resources/municipalities/application/get_by_id_usecase.go
package application

import (
	"context"
	"errors"
	"api-seguridad/resources/municipalities/domain/entities"
	"api-seguridad/resources/municipalities/domain/repository"
)

type GetMunicipalityByIDUseCase struct {
	repo repository.MunicipalityRepository
}

func NewGetMunicipalityByIDUseCase(repo repository.MunicipalityRepository) *GetMunicipalityByIDUseCase {
	return &GetMunicipalityByIDUseCase{repo: repo}
}

func (uc *GetMunicipalityByIDUseCase) Execute(ctx context.Context, id uint) (*entities.Municipality, error) {
	if id == 0 {
		return nil, errors.New("invalid municipality ID")
	}

	municipality, err := uc.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if municipality == nil || municipality.IsDeleted() {
		return nil, errors.New("municipality not found")
	}

	return municipality, nil
}