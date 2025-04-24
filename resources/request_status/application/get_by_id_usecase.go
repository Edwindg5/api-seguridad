// api-seguridad/resources/request_status/application/get_by_id_usecase.go
package application

import (
	"context"
	"errors"
	"api-seguridad/resources/request_status/domain/entities"
	"api-seguridad/resources/request_status/domain/repository"
)

type GetRequestStatusByIDUseCase struct {
	repo repository.RequestStatusRepository
}

func NewGetRequestStatusByIDUseCase(repo repository.RequestStatusRepository) *GetRequestStatusByIDUseCase {
	return &GetRequestStatusByIDUseCase{repo: repo}
}

func (uc *GetRequestStatusByIDUseCase) Execute(ctx context.Context, id uint) (*entities.RequestStatus, error) {
	if id == 0 {
		return nil, errors.New("invalid status ID")
	}

	status, err := uc.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if status == nil || status.IsDeleted() {
		return nil, errors.New("status not found")
	}

	return status, nil
}