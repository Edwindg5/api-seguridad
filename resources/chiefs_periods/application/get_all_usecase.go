package application

import (
	"context"
	"api-seguridad/resources/chiefs_periods/domain/entities"
	"api-seguridad/resources/chiefs_periods/domain/repository"
)

type GetAllChiefsPeriodsUseCase struct {
	repo repository.ChiefsPeriodRepository
}

func NewGetAllChiefsPeriodsUseCase(repo repository.ChiefsPeriodRepository) *GetAllChiefsPeriodsUseCase {
	return &GetAllChiefsPeriodsUseCase{repo: repo}
}

func (uc *GetAllChiefsPeriodsUseCase) Execute(ctx context.Context) ([]*entities.ChiefsPeriod, error) {
	periods, err := uc.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	// Filtrar registros eliminados
	var activePeriods []*entities.ChiefsPeriod
	for _, p := range periods {
		if !p.IsDeleted() {
			activePeriods = append(activePeriods, p)
		}
	}

	return activePeriods, nil
}