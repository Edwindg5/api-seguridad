package application

import (
	"context"
	"api-seguridad/resources/type_police/domain/entities"
	"api-seguridad/resources/type_police/domain/repository"
)

type GetAllTypePoliceUseCase struct {
	repo repository.TypePoliceRepository
}

func NewGetAllTypePoliceUseCase(repo repository.TypePoliceRepository) *GetAllTypePoliceUseCase {
	return &GetAllTypePoliceUseCase{repo: repo}
}

func (uc *GetAllTypePoliceUseCase) Execute(ctx context.Context) ([]*entities.TypePolice, error) {
	typesPolice, err := uc.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	// Filtro adicional por si acaso
	var activeTypes []*entities.TypePolice
	for _, tp := range typesPolice {
		if !tp.IsDeleted() {
			activeTypes = append(activeTypes, tp)
		}
	}

	return activeTypes, nil
}