// api-seguridad/resources/request/application/get_all_usecase.go
package application

import (
	"context"
	"api-seguridad/resources/request/domain/entities"
	"api-seguridad/resources/request/domain/repository"
)

type GetAllRequestsUseCase struct {
	repo repository.RequestRepository
}

func NewGetAllRequestsUseCase(repo repository.RequestRepository) *GetAllRequestsUseCase {
	return &GetAllRequestsUseCase{repo: repo}
}

func (uc *GetAllRequestsUseCase) Execute(ctx context.Context) ([]*entities.Request, error) {
	requests, err := uc.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	// Opcional: Filtrar requests eliminados si no se hace en el repositorio
	var activeRequests []*entities.Request
	for _, req := range requests {
		if req != nil && !req.IsDeleted() {
			activeRequests = append(activeRequests, req)
		}
	}

	return activeRequests, nil
}