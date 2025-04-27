//api-seguridad/resources/request_details/application/update_usecase.go
package application

import (
	"context"
	"errors"
	"api-seguridad/resources/request_details/domain/entities"
	"api-seguridad/resources/request_details/domain/repository"
)

type UpdateRequestDetailUseCase struct {
	repo repository.RequestDetailRepository
}

func NewUpdateRequestDetailUseCase(repo repository.RequestDetailRepository) *UpdateRequestDetailUseCase {
	return &UpdateRequestDetailUseCase{repo: repo}
}

func (uc *UpdateRequestDetailUseCase) Execute(ctx context.Context, detail *entities.RequestDetail) error {
	if detail.ID == 0 {
		return errors.New("ID inválido")
	}

	existing, err := uc.repo.GetByID(ctx, detail.ID)
	if err != nil {
		return err
	}
	if existing == nil || existing.IsDeleted() {
		return errors.New("detalle de solicitud no encontrado")
	}

	// Validar que no se modifiquen las claves foráneas
	if existing.RequestID != detail.RequestID || existing.PoliceID != detail.PoliceID {
		return errors.New("no se pueden modificar los IDs de solicitud o policía")
	}

	return uc.repo.Update(ctx, detail)
}