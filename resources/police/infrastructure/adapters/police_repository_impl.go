package adapters

import (
	"context"
	"errors"

	"api-seguridad/resources/police/domain/entities"
	"api-seguridad/resources/police/domain/repository"

	"gorm.io/gorm"
)

type PoliceRepositoryImpl struct {
	db *gorm.DB
}

func NewPoliceRepository(db *gorm.DB) repository.PoliceRepository {
	return &PoliceRepositoryImpl{db: db}
}

func (r *PoliceRepositoryImpl) Create(ctx context.Context, police *entity.Police) error {
	return r.db.WithContext(ctx).Create(police).Error
}

func (r *PoliceRepositoryImpl) GetByID(ctx context.Context, id uint) (*entity.Police, error) {
	var police entity.Police
	err := r.db.WithContext(ctx).Preload("TypePolice").First(&police, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &police, err
}

func (r *PoliceRepositoryImpl) GetByCUIP(ctx context.Context, cuip string) (*entity.Police, error) {
	var police entity.Police
	err := r.db.WithContext(ctx).Where("cuip = ?", cuip).First(&police).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &police, err
}

func (r *PoliceRepositoryImpl) Update(ctx context.Context, police *entity.Police) error {
	return r.db.WithContext(ctx).Save(police).Error
}

func (r *PoliceRepositoryImpl) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&entity.Police{}, id).Error
}

func (r *PoliceRepositoryImpl) List(ctx context.Context) ([]*entity.Police, error) {
	var police []*entity.Police
	err := r.db.WithContext(ctx).Preload("TypePolice").Find(&police).Error
	return police, err
}