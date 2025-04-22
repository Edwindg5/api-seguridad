package application

import (
	"context"
	"errors"
	"api-seguridad/resources/type_police/domain/entities"
	"api-seguridad/resources/type_police/domain/repository"
)

type UpdateTypePoliceUseCase struct {
	repo repository.TypePoliceRepository
}

func NewUpdateTypePoliceUseCase(repo repository.TypePoliceRepository) *UpdateTypePoliceUseCase {
	return &UpdateTypePoliceUseCase{repo: repo}
}

func (uc *UpdateTypePoliceUseCase) Execute(ctx context.Context, typePolice *entities.TypePolice) error {
	if typePolice.ID == 0 {
		return errors.New("invalid type police ID")
	}
	if typePolice.TitleKindPolice == "" {
		return errors.New("type police title is required")
	}

	existingType, err := uc.repo.GetByID(ctx, typePolice.ID)
	if err != nil {
		return err
	}
	if existingType == nil || existingType.IsDeleted() {
		return errors.New("type police not found")
	}

	return uc.repo.Update(ctx, typePolice)
}