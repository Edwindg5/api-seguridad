// api-seguridad/resources/area_chiefs/application/get_all_usecase.go
package application

import (
	"context"
	"api-seguridad/resources/area_chiefs/domain/entities"
	"api-seguridad/resources/area_chiefs/domain/repository"
)

type GetAllAreaChiefsUseCase struct {
	repo repository.AreaChiefRepository
}

func NewGetAllAreaChiefsUseCase(repo repository.AreaChiefRepository) *GetAllAreaChiefsUseCase {
	return &GetAllAreaChiefsUseCase{repo: repo}
}

func (uc *GetAllAreaChiefsUseCase) Execute(ctx context.Context) ([]*entities.AreaChief, error) {
	chiefs, err := uc.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	// Filter out deleted records
	var activeChiefs []*entities.AreaChief
	for _, c := range chiefs {
		if !c.IsDeleted() {
			activeChiefs = append(activeChiefs, c)
		}
	}

	return activeChiefs, nil
}