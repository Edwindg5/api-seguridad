// api-seguridad/resources/request/application/get_by_municipality_usecase.go
package application

import (
	"context"
	"errors"
	"api-seguridad/resources/request/domain/entities"
	"api-seguridad/resources/request/domain/repository"
)

type GetRequestsByMunicipalityUseCase struct {
	repo repository.RequestRepository
}

func NewGetRequestsByMunicipalityUseCase(repo repository.RequestRepository) *GetRequestsByMunicipalityUseCase {
	return &GetRequestsByMunicipalityUseCase{repo: repo}
}

func (uc *GetRequestsByMunicipalityUseCase) Execute(ctx context.Context, municipalityID uint) ([]*entities.Request, error) {
	if municipalityID == 0 {
		return nil, errors.New("municipality ID is required")
	}

	requests, err := uc.repo.GetByMunicipality(ctx, municipalityID)
	if err != nil {
		return nil, err
	}

	// Filter out deleted requests
	var activeRequests []*entities.Request
	for _, r := range requests {
		if !r.IsDeleted() {
			activeRequests = append(activeRequests, r)
		}
	}

	return activeRequests, nil
}