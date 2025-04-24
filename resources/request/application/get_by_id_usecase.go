// api-seguridad/resources/request/application/get_by_id_usecase.go
package application

import (
	"context"
	"errors"
	"api-seguridad/resources/request/domain/entities"
	"api-seguridad/resources/request/domain/repository"
)

type GetRequestByIDUseCase struct {
	repo repository.RequestRepository
}

func NewGetRequestByIDUseCase(repo repository.RequestRepository) *GetRequestByIDUseCase {
	return &GetRequestByIDUseCase{repo: repo}
}

func (uc *GetRequestByIDUseCase) Execute(ctx context.Context, id uint) (*entities.Request, error) {
	if id == 0 {
		return nil, errors.New("invalid request ID")
	}

	request, err := uc.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if request == nil || request.IsDeleted() {
		return nil, errors.New("request not found")
	}

	return request, nil
}