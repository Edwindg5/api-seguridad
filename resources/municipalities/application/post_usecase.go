// api-seguridad/resources/municipalities/application/post_usecase.go
package application

import (
	"context"
	"errors"
	"api-seguridad/resources/municipalities/domain/entities"
	"api-seguridad/resources/municipalities/domain/repository"
)

type PostMunicipalityUseCase struct {
	municipalityRepo repository.MunicipalityRepository
}

func NewPostMunicipalityUseCase(municipalityRepo repository.MunicipalityRepository) *PostMunicipalityUseCase {
	return &PostMunicipalityUseCase{municipalityRepo: municipalityRepo}
}

func (uc *PostMunicipalityUseCase) Execute(ctx context.Context, municipality *entity.Municipality) error {
	if municipality.Name == "" {
		return errors.New("municipality name is required")
	}

	existingMunicipality, err := uc.municipalityRepo.GetByName(ctx, municipality.Name)
	if err != nil {
		return err
	}
	if existingMunicipality != nil {
		return errors.New("municipality with this name already exists")
	}

	return uc.municipalityRepo.Create(ctx, municipality)
}