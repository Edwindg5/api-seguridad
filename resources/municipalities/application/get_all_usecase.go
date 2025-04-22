package application

import (
	"context"
	"api-seguridad/resources/municipalities/domain/entities"
	"api-seguridad/resources/municipalities/domain/repository"
)

type GetAllMunicipalitiesUseCase struct {
	repo repository.MunicipalityRepository
}

func NewGetAllMunicipalitiesUseCase(repo repository.MunicipalityRepository) *GetAllMunicipalitiesUseCase {
	return &GetAllMunicipalitiesUseCase{repo: repo}
}

func (uc *GetAllMunicipalitiesUseCase) Execute(ctx context.Context) ([]*entities.Municipality, error) {
	municipalities, err := uc.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	// Filter out deleted municipalities if not already filtered in repository
	var activeMunicipalities []*entities.Municipality
	for _, m := range municipalities {
		if !m.IsDeleted() {
			activeMunicipalities = append(activeMunicipalities, m)
		}
	}

	return activeMunicipalities, nil
}