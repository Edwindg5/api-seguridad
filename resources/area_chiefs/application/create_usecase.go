// api-seguridad/resources/area_chiefs/application/create_usecase.go
package application

import (
	"context"
	"errors"
	"api-seguridad/resources/area_chiefs/domain/entities"
	"api-seguridad/resources/area_chiefs/domain/repository"
	"time"
)

type CreateAreaChiefUseCase struct {
	repo repository.AreaChiefRepository
}

func NewCreateAreaChiefUseCase(repo repository.AreaChiefRepository) *CreateAreaChiefUseCase {
	return &CreateAreaChiefUseCase{repo: repo}
}

func (uc *CreateAreaChiefUseCase) Execute(ctx context.Context, chief *entities.AreaChief) error {
	// Validations
	if chief.Name == "" {
		return errors.New("chief name is required")
	}
	if chief.Position == "" {
		return errors.New("position is required")
	}
	if chief.CreatedBy == 0 {
		return errors.New("creator user is required")
	}

	// Set timestamps
	if chief.CreatedAt.IsZero() {
		chief.CreatedAt = time.Now()
	}
	chief.UpdatedAt = time.Now()

	return uc.repo.Create(ctx, chief)
}