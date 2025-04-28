//api-seguridad/resources/municipalities/application/post_usecase.go
package application

import (
	"context"
	"errors"
	"api-seguridad/resources/municipalities/domain/entities"
	"api-seguridad/resources/municipalities/domain/repository"
)

type PostMunicipalityUseCase struct {
	repo repository.MunicipalityRepository
}

func NewPostMunicipalityUseCase(repo repository.MunicipalityRepository) *PostMunicipalityUseCase {
	return &PostMunicipalityUseCase{repo: repo}
}

func (uc *PostMunicipalityUseCase) Execute(ctx context.Context, municipality *entities.Municipality) error {
	if municipality.Name == "" {
		return errors.New("municipality name is required")
	}

	existingMunicipality, err := uc.repo.GetByName(ctx, municipality.Name)
	if err != nil {
		return err
	}
	if existingMunicipality != nil && !existingMunicipality.IsDeleted() {
		return errors.New("municipality with this name already exists")
	}

	return uc.repo.Create(ctx, municipality)
}

func (uc *PostMunicipalityUseCase) GetByID(ctx context.Context, id uint) (*entities.Municipality, error) {
	return uc.repo.GetByID(ctx, id)
}