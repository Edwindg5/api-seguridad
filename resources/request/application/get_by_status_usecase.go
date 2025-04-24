// api-seguridad/resources/request/application/get_by_status_usecase.go
package application

import (
	"context"
	"errors"
	"api-seguridad/resources/request/domain/entities"
	"api-seguridad/resources/request/domain/repository"
)

type GetRequestsByStatusUseCase struct {
	repo repository.RequestRepository
}

func NewGetRequestsByStatusUseCase(repo repository.RequestRepository) *GetRequestsByStatusUseCase {
	return &GetRequestsByStatusUseCase{repo: repo}
}

func (uc *GetRequestsByStatusUseCase) Execute(ctx context.Context, statusID uint) ([]*entities.Request, error) {
	if statusID == 0 {
		return nil, errors.New("status ID is required")
	}

	requests, err := uc.repo.GetByStatus(ctx, statusID)
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