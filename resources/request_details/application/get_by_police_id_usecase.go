//api-seguridad/resources/request_details/application/get_by_police_id_usecase.go
package application

import (
	"context"
	"errors"
	"api-seguridad/resources/request_details/domain/entities"
	"api-seguridad/resources/request_details/domain/repository"
)

type GetByPoliceIDUseCase struct {
	repo repository.RequestDetailRepository
}

func NewGetByPoliceIDUseCase(repo repository.RequestDetailRepository) *GetByPoliceIDUseCase {
	return &GetByPoliceIDUseCase{repo: repo}
}

func (uc *GetByPoliceIDUseCase) Execute(ctx context.Context, policeID uint) ([]*entities.RequestDetail, error) {
	if policeID == 0 {
		return nil, errors.New("ID de policía es requerido")
	}

	details, err := uc.repo.GetByPoliceID(ctx, policeID)
	if err != nil {
		return nil, err
	}

	// Filtrar registros eliminados
	var activeDetails []*entities.RequestDetail
	for _, d := range details {
		if !d.IsDeleted() {
			activeDetails = append(activeDetails, d)
		}
	}

	return activeDetails, nil
}