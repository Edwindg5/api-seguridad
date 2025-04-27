//api-seguridad/resources/request_details/application/create_usecase.go
package application

import (
	"context"
	"errors"
	"api-seguridad/resources/request_details/domain/entities"
	"api-seguridad/resources/request_details/domain/repository"
)

type CreateRequestDetailUseCase struct {
	repo repository.RequestDetailRepository
}

func NewCreateRequestDetailUseCase(repo repository.RequestDetailRepository) *CreateRequestDetailUseCase {
	return &CreateRequestDetailUseCase{repo: repo}
}

func (uc *CreateRequestDetailUseCase) Execute(ctx context.Context, detail *entities.RequestDetail) error {
	// Validaciones básicas
	if detail.RequestID == 0 {
		return errors.New("ID de solicitud es requerido")
	}
	if detail.PoliceID == 0 {
		return errors.New("ID de policía es requerido")
	}
	if detail.CreatedBy == 0 {
		return errors.New("usuario creador es requerido")
	}

	return uc.repo.Create(ctx, detail)
}