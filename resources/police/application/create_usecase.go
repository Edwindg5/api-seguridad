//api-seguridad/resources/police/application/create_usecase.go
package application

import (
	"context"
	"errors"
	"api-seguridad/resources/police/domain/entities"
	"api-seguridad/resources/police/domain/repository"
)

type CreatePoliceUseCase struct {
	repo repository.PoliceRepository
}

func NewCreatePoliceUseCase(repo repository.PoliceRepository) *CreatePoliceUseCase {
	return &CreatePoliceUseCase{repo: repo}
}

func (uc *CreatePoliceUseCase) Execute(ctx context.Context, police *entities.Police) error {
    // Validaciones básicas
    if police.Name == "" || police.PaternalName == "" {
        return errors.New("name and paternal lastname are required")
    }
    if police.CUIP == "" {
        return errors.New("CUIP is required")
    }
    if police.Sex != "M" && police.Sex != "F" {
        return errors.New("invalid sex, must be M or F")
    }

    // Verificar que el usuario creador exista
    if police.CreatedBy == 0 {
        return errors.New("creator user does not exist")
    }

    // Check CUIP uniqueness
    existing, err := uc.repo.GetByCUIP(ctx, police.CUIP)
    if err != nil {
        return err
    }
    if existing != nil {
        return errors.New("police with this CUIP already exists")
    }

    return uc.repo.Create(ctx, police)
}