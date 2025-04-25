// api-seguridad/resources/request/application/update_usecase.go
package application

import (
	"context"
	"errors"
	"api-seguridad/resources/request/domain/entities"
	"api-seguridad/resources/request/domain/repository"
	"time"
)

type UpdateRequestUseCase struct {
	repo repository.RequestRepository
}

func NewUpdateRequestUseCase(repo repository.RequestRepository) *UpdateRequestUseCase {
	return &UpdateRequestUseCase{repo: repo}
}

// Nuevo método para obtener request por ID
func (uc *UpdateRequestUseCase) GetByID(ctx context.Context, id uint) (*entities.Request, error) {
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

func (uc *UpdateRequestUseCase) Execute(ctx context.Context, request *entities.Request) error {
	if request.ID == 0 {
		return errors.New("invalid request ID")
	}
	if request.OfficeNumber == "" {
		return errors.New("office number is required")
	}
	if request.UpdatedBy == 0 {
		return errors.New("updater user is required")
	}

	// Verificar que existe
	_, err := uc.GetByID(ctx, request.ID)
	if err != nil {
		return err
	}

	// Actualizar timestamp
	request.UpdatedAt = time.Now()

	return uc.repo.Update(ctx, request)
}