// api-seguridad/resources/municipalities/application/municipality_service.go
package application

import (
	"context"
	"api-seguridad/resources/municipalities/domain/entities"
	"api-seguridad/resources/municipalities/domain/repository"
)

type MunicipalityService struct {
	municipalityRepo repository.MunicipalityRepository
}

func NewMunicipalityService(municipalityRepo repository.MunicipalityRepository) *MunicipalityService {
	return &MunicipalityService{municipalityRepo: municipalityRepo}
}

func (s *MunicipalityService) CreateMunicipality(ctx context.Context, municipality *entity.Municipality) error {
	postUC := NewPostMunicipalityUseCase(s.municipalityRepo)
	return postUC.Execute(ctx, municipality)
}

func (s *MunicipalityService) GetMunicipalityByID(ctx context.Context, id uint) (*entity.Municipality, error) {
	return s.municipalityRepo.GetByID(ctx, id)
}

func (s *MunicipalityService) UpdateMunicipality(ctx context.Context, municipality *entity.Municipality) error {
	return s.municipalityRepo.Update(ctx, municipality)
}

func (s *MunicipalityService) DeleteMunicipality(ctx context.Context, id uint) error {
	return s.municipalityRepo.Delete(ctx, id)
}

func (s *MunicipalityService) ListMunicipalities(ctx context.Context) ([]*entity.Municipality, error) {
	return s.municipalityRepo.List(ctx)
}