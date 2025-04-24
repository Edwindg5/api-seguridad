// api-seguridad/resources/request_status/application/get_all_usecase.go
package application

import (
	"context"
	"api-seguridad/resources/request_status/domain/entities"
	"api-seguridad/resources/request_status/domain/repository"
)

type GetAllRequestStatusUseCase struct {
	repo repository.RequestStatusRepository
}

func NewGetAllRequestStatusUseCase(repo repository.RequestStatusRepository) *GetAllRequestStatusUseCase {
	return &GetAllRequestStatusUseCase{repo: repo}
}

func (uc *GetAllRequestStatusUseCase) Execute(ctx context.Context) ([]*entities.RequestStatus, error) {
	statusList, err := uc.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	// Filter out deleted records
	var activeStatuses []*entities.RequestStatus
	for _, s := range statusList {
		if !s.IsDeleted() {
			activeStatuses = append(activeStatuses, s)
		}
	}

	return activeStatuses, nil
}