//api-seguridad/resources/municipalities/application/update_usecase.go
package application

import (
	"context"
	"errors"
	"api-seguridad/resources/municipalities/domain/entities"
	"api-seguridad/resources/municipalities/domain/repository"
)

type UpdateMunicipalityUseCase struct {
	repo repository.MunicipalityRepository
}

func NewUpdateMunicipalityUseCase(repo repository.MunicipalityRepository) *UpdateMunicipalityUseCase {
	return &UpdateMunicipalityUseCase{repo: repo}
}

func (uc *UpdateMunicipalityUseCase) Execute(ctx context.Context, municipality *entities.Municipality) error {
	if municipality.ID == 0 {
		return errors.New("invalid municipality ID")
	}
	if municipality.Name == "" {
		return errors.New("municipality name is required")
	}

	existingMunicipality, err := uc.repo.GetByID(ctx, municipality.ID)
	if err != nil {
		return err
	}
	if existingMunicipality == nil || existingMunicipality.IsDeleted() {
		return errors.New("municipality not found")
	}

	// Check if name is being changed to one that already exists
	if existingMunicipality.Name != municipality.Name {
		municipalityWithName, err := uc.repo.GetByName(ctx, municipality.Name)
		if err != nil {
			return err
		}
		if municipalityWithName != nil {
			return errors.New("municipality with this name already exists")
		}
	}

	return uc.repo.Update(ctx, municipality)
}