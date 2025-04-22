package application

import (
	"context"
	"errors"
	"api-seguridad/resources/municipalities/domain/repository"
)

type SoftDeleteMunicipalityUseCase struct {
	repo repository.MunicipalityRepository
}

func NewSoftDeleteMunicipalityUseCase(repo repository.MunicipalityRepository) *SoftDeleteMunicipalityUseCase {
	return &SoftDeleteMunicipalityUseCase{repo: repo}
}

func (uc *SoftDeleteMunicipalityUseCase) Execute(ctx context.Context, id uint) error {
	if id == 0 {
		return errors.New("invalid municipality ID")
	}

	// Verify municipality exists before deleting
	municipality, err := uc.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if municipality == nil || municipality.IsDeleted() {
		return errors.New("municipality not found")
	}

	return uc.repo.SoftDelete(ctx, id)
}