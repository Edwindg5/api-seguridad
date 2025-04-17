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

func (r *MunicipalityRepositoryImpl) Create(ctx context.Context, municipality *entity.Municipality) error {
	return r.db.WithContext(ctx).Create(municipality).Error
}

func (r *MunicipalityRepositoryImpl) GetByID(ctx context.Context, id uint) (*entity.Municipality, error) {
	var municipality entity.Municipality
	err := r.db.WithContext(ctx).Preload("Delegations").First(&municipality, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &municipality, err
}

func (r *MunicipalityRepositoryImpl) GetByName(ctx context.Context, name string) (*entity.Municipality, error) {
	var municipality entity.Municipality
	err := r.db.WithContext(ctx).Where("name = ?", name).First(&municipality).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &municipality, err
}

func (r *MunicipalityRepositoryImpl) Update(ctx context.Context, municipality *entity.Municipality) error {
	return r.db.WithContext(ctx).Save(municipality).Error
}

func (r *MunicipalityRepositoryImpl) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&entity.Municipality{}, id).Error
}

func (r *MunicipalityRepositoryImpl) List(ctx context.Context) ([]*entity.Municipality, error) {
	var municipalities []*entity.Municipality
	err := r.db.WithContext(ctx).Preload("Delegations").Find(&municipalities).Error
	return municipalities, err
}