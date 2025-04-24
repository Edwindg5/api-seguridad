//api-seguridad/resources/type_police/application/create_usecase.go
package application

import (
	"context"
	"errors"
	"api-seguridad/resources/type_police/domain/entities"
	"api-seguridad/resources/type_police/domain/repository"
)

type CreateTypePoliceUseCase struct {
	repo repository.TypePoliceRepository
}

func NewCreateTypePoliceUseCase(repo repository.TypePoliceRepository) *CreateTypePoliceUseCase {
	return &CreateTypePoliceUseCase{repo: repo}
}

func (uc *CreateTypePoliceUseCase) Execute(ctx context.Context, typePolice *entities.TypePolice) error {
	if typePolice.TitleKindPolice == "" {
		return errors.New("type police title is required")
	}
	return uc.repo.Create(ctx, typePolice)
}