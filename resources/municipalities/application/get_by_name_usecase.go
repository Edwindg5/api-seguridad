//api-seguridad/resources/municipalities/application/get_by_name_usecase.go
package application

import (
	"context"
	"errors"
	"api-seguridad/resources/municipalities/domain/entities"
	"api-seguridad/resources/municipalities/domain/repository"
)

type GetMunicipalityByNameUseCase struct {
	repo repository.MunicipalityRepository
}

func NewGetMunicipalityByNameUseCase(repo repository.MunicipalityRepository) *GetMunicipalityByNameUseCase {
	return &GetMunicipalityByNameUseCase{repo: repo}
}

func (uc *GetMunicipalityByNameUseCase) Execute(ctx context.Context, name string) (*entities.Municipality, error) {
	if name == "" {
		return nil, errors.New("municipality name is required")
	}

	municipality, err := uc.repo.GetByName(ctx, name)
	if err != nil {
		return nil, err
	}

	if municipality == nil || municipality.IsDeleted() {
		return nil, errors.New("municipality not found")
	}

	return municipality, nil
}