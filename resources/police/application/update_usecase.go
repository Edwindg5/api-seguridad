package application

import (
	"context"
	"errors"
	"api-seguridad/resources/police/domain/entities"
	"api-seguridad/resources/police/domain/repository"
)

type UpdatePoliceUseCase struct {
	repo repository.PoliceRepository
}

func NewUpdatePoliceUseCase(repo repository.PoliceRepository) *UpdatePoliceUseCase {
	return &UpdatePoliceUseCase{repo: repo}
}

func (uc *UpdatePoliceUseCase) Execute(ctx context.Context, police *entities.Police) error {
	if police.ID == 0 {
		return errors.New("invalid police ID")
	}
	if police.Name == "" || police.PaternalName == "" {
		return errors.New("name and paternal lastname are required")
	}

	existing, err := uc.repo.GetByID(ctx, police.ID)
	if err != nil {
		return err
	}
	if existing == nil || existing.IsDeleted() {
		return errors.New("police not found")
	}

	return uc.repo.Update(ctx, police)
}