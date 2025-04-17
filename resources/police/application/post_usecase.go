package application

import (
	"context"
	"errors"
	"api-seguridad/resources/police/domain/entities"
	"api-seguridad/resources/police/domain/repository"
)

type PostPoliceUseCase struct {
	policeRepo repository.PoliceRepository
}

func NewPostPoliceUseCase(policeRepo repository.PoliceRepository) *PostPoliceUseCase {
	return &PostPoliceUseCase{policeRepo: policeRepo}
}

func (uc *PostPoliceUseCase) Execute(ctx context.Context, police *entity.Police) error {
	if police.Name == "" {
		return errors.New("name is required")
	}
	if police.PaternalName == "" {
		return errors.New("paternal lastname is required")
	}
	if police.TypePoliceID == 0 {
		return errors.New("type_police_id is required")
	}
	if police.CUIP == "" {
		return errors.New("CUIP is required")
	}

	existingPolice, err := uc.policeRepo.GetByCUIP(ctx, police.CUIP)
	if err != nil {
		return err
	}
	if existingPolice != nil {
		return errors.New("police with this CUIP already exists")
	}

	return uc.policeRepo.Create(ctx, police)
}