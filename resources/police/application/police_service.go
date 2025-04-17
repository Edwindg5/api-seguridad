package application

import (
	"context"
	"api-seguridad/resources/police/domain/entities"
	"api-seguridad/resources/police/domain/repository"
)

type PoliceService struct {
	policeRepo repository.PoliceRepository
}

func NewPoliceService(policeRepo repository.PoliceRepository) *PoliceService {
	return &PoliceService{policeRepo: policeRepo}
}

func (s *PoliceService) CreatePolice(ctx context.Context, police *entity.Police) error {
	postUC := NewPostPoliceUseCase(s.policeRepo)
	return postUC.Execute(ctx, police)
}

func (s *PoliceService) GetPoliceByID(ctx context.Context, id uint) (*entity.Police, error) {
	return s.policeRepo.GetByID(ctx, id)
}

func (s *PoliceService) UpdatePolice(ctx context.Context, police *entity.Police) error {
	return s.policeRepo.Update(ctx, police)
}

func (s *PoliceService) DeletePolice(ctx context.Context, id uint) error {
	return s.policeRepo.Delete(ctx, id)
}

func (s *PoliceService) ListPolice(ctx context.Context) ([]*entity.Police, error) {
	return s.policeRepo.List(ctx)
}