//api-seguridad/resources/municipalities/infrastructure/adapters/municipality_repository_impl.go
package adapters

import (
	"context"
	"errors"

	"api-seguridad/resources/municipalities/domain/entities"
	"api-seguridad/resources/municipalities/domain/repository"

	"gorm.io/gorm"
)

type MunicipalityRepositoryImpl struct {
	db *gorm.DB
}

func NewMunicipalityRepository(db *gorm.DB) repository.MunicipalityRepository {
	return &MunicipalityRepositoryImpl{db: db}
}

func (r *MunicipalityRepositoryImpl) Create(ctx context.Context, municipality *entities.Municipality) error {
	return r.db.WithContext(ctx).Create(municipality).Error
}

func (r *MunicipalityRepositoryImpl) GetByID(ctx context.Context, id uint) (*entities.Municipality, error) {
	var municipality entities.Municipality
	err := r.db.WithContext(ctx).
		Where("id_municipalities = ? AND deleted = ?", id, false).
		First(&municipality).Error
		
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &municipality, err
}

func (r *MunicipalityRepositoryImpl) GetByName(ctx context.Context, name string) (*entities.Municipality, error) {
	var municipality entities.Municipality
	err := r.db.WithContext(ctx).
		Where("name_municipalities = ? AND deleted = ?", name, false).
		First(&municipality).Error
		
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &municipality, err
}

func (r *MunicipalityRepositoryImpl) GetAll(ctx context.Context) ([]*entities.Municipality, error) {
	var municipalities []*entities.Municipality
	err := r.db.WithContext(ctx).
		Where("deleted = ?", false).
		Find(&municipalities).Error
	return municipalities, err
}

func (r *MunicipalityRepositoryImpl) Update(ctx context.Context, municipality *entities.Municipality) error {
	return r.db.WithContext(ctx).Save(municipality).Error
}

func (r *MunicipalityRepositoryImpl) SoftDelete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).
		Model(&entities.Municipality{}).
		Where("id_municipalities = ?", id).
		Update("deleted", true).Error
}