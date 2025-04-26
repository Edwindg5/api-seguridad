package application

import (
	"context"
	"errors"
	"api-seguridad/resources/request_details/domain/entities"
	"api-seguridad/resources/request_details/domain/repository"
)

type GetByRequestIDUseCase struct {
	repo repository.RequestDetailRepository
}

func NewGetByRequestIDUseCase(repo repository.RequestDetailRepository) *GetByRequestIDUseCase {
	return &GetByRequestIDUseCase{repo: repo}
}

func (uc *GetByRequestIDUseCase) Execute(ctx context.Context, requestID uint) ([]*entities.RequestDetail, error) {
	if requestID == 0 {
		return nil, errors.New("ID de solicitud es requerido")
	}

	details, err := uc.repo.GetByRequestID(ctx, requestID)
	if err != nil {
		return nil, err
	}

	// Filtrar registros eliminados (doble verificación)
	var activeDetails []*entities.RequestDetail
	for _, d := range details {
		if !d.IsDeleted() {
			activeDetails = append(activeDetails, d)
		}
	}

	return activeDetails, nil
}