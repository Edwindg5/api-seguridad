//api-seguridad/resources/request_details/application/update_usecase.go
package application

import (
	"context"
	"errors"
	"api-seguridad/resources/request_details/domain/entities"
	"api-seguridad/resources/request_details/domain/repository"
)

type UpdateRequestDetailDTO struct {
	ID                 uint
	Active             bool
	Census             bool
	Located            bool
	Register           bool
	Approved           bool
	Comments           string
	MunicipalityActive bool
	UpdaterID          uint
}

type UpdateRequestDetailUseCase struct {
	repo repository.RequestDetailRepository
}

func NewUpdateRequestDetailUseCase(repo repository.RequestDetailRepository) *UpdateRequestDetailUseCase {
	return &UpdateRequestDetailUseCase{repo: repo}
}

func (uc *UpdateRequestDetailUseCase) ExecuteWithDTO(ctx context.Context, dto *UpdateRequestDetailDTO) (*entities.RequestDetail, error) {
	if dto.ID == 0 {
		return nil, errors.New("ID inválido")
	}

	existing, err := uc.repo.GetByID(ctx, dto.ID)
	if err != nil {
		return nil, err
	}
	if existing == nil || existing.IsDeleted() {
		return nil, errors.New("detalle de solicitud no encontrado")
	}

	// Aplicar actualizaciones
	existing.Active = dto.Active
	existing.Census = dto.Census
	existing.Located = dto.Located
	existing.Register = dto.Register
	existing.Approved = dto.Approved
	existing.Comments = dto.Comments
	existing.MunicipalityActive = dto.MunicipalityActive
	existing.UpdatedBy = dto.UpdaterID

	if err := uc.repo.Update(ctx, existing); err != nil {
		return nil, err
	}

	return existing, nil
}

// Mantener el método Execute original para compatibilidad
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

	if existing.RequestID != detail.RequestID || existing.PoliceID != detail.PoliceID {
		return errors.New("no se pueden modificar los IDs de solicitud o policía")
	}

	return uc.repo.Update(ctx, detail)
}