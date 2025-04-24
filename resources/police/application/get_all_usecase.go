//api-seguridad/resources/police/application/get_all_usecase.go
package application

import (
	"context"
	"api-seguridad/resources/police/domain/entities"
	"api-seguridad/resources/police/domain/repository"
)

type GetAllPoliceUseCase struct {
	repo repository.PoliceRepository
}

func NewGetAllPoliceUseCase(repo repository.PoliceRepository) *GetAllPoliceUseCase {
	return &GetAllPoliceUseCase{repo: repo}
}

func (uc *GetAllPoliceUseCase) Execute(ctx context.Context) ([]*entities.Police, error) {
	policeList, err := uc.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	// Filter out deleted records
	var activePolice []*entities.Police
	for _, p := range policeList {
		if !p.IsDeleted() {
			activePolice = append(activePolice, p)
		}
	}

	return activePolice, nil
}