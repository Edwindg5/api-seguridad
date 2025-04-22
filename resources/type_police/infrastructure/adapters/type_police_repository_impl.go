package adapters

import (
	"context"
	"errors"

	"api-seguridad/resources/type_police/domain/entities"
	"api-seguridad/resources/type_police/domain/repository"

	"gorm.io/gorm"
)

type TypePoliceRepositoryImpl struct {
	db *gorm.DB
}

func NewTypePoliceRepository(db *gorm.DB) repository.TypePoliceRepository {
	return &TypePoliceRepositoryImpl{db: db}
}

func (r *TypePoliceRepositoryImpl) Create(ctx context.Context, typePolice *entities.TypePolice) error {
	return r.db.WithContext(ctx).Create(typePolice).Error
}

func (r *TypePoliceRepositoryImpl) GetByID(ctx context.Context, id uint) (*entities.TypePolice, error) {
	var typePolice entities.TypePolice
	err := r.db.WithContext(ctx).
		Where("id_type_police = ? AND deleted = ?", id, false).
		Preload("CreatedByUser").
		Preload("UpdatedByUser").
		First(&typePolice).Error
		
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &typePolice, err
}

func (r *TypePoliceRepositoryImpl) GetAll(ctx context.Context) ([]*entities.TypePolice, error) {
	var typesPolice []*entities.TypePolice
	err := r.db.WithContext(ctx).
		Where("deleted = ?", false).
		Preload("CreatedByUser").
		Preload("UpdatedByUser").
		Find(&typesPolice).Error
	return typesPolice, err
}

func (r *TypePoliceRepositoryImpl) Update(ctx context.Context, typePolice *entities.TypePolice) error {
	return r.db.WithContext(ctx).Save(typePolice).Error
}

func (r *TypePoliceRepositoryImpl) SoftDelete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).
		Model(&entities.TypePolice{}).
		Where("id_type_police = ?", id).
		Update("deleted", true).Error
}