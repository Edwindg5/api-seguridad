//api-seguridad/resources/request_details/application/get_by_id_usecase.go
package application

import (
	"context"
	"errors"
	"api-seguridad/resources/request_details/domain/entities"
	"api-seguridad/resources/request_details/domain/repository"
)

type GetRequestDetailByIDUseCase struct {
	repo repository.RequestDetailRepository
}

func NewGetRequestDetailByIDUseCase(repo repository.RequestDetailRepository) *GetRequestDetailByIDUseCase {
	return &GetRequestDetailByIDUseCase{repo: repo}
}

func (uc *GetRequestDetailByIDUseCase) Execute(ctx context.Context, id uint) (*entities.RequestDetail, error) {
	if id == 0 {
		return nil, errors.New("ID inválido")
	}

	detail, err := uc.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if detail == nil || detail.IsDeleted() {
		return nil, errors.New("detalle de solicitud no encontrado")
	}

	return detail, nil
}