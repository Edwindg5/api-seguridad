package adapters

import (
	"context"
	"errors"
	
	"api-seguridad/resources/request/domain/entities"
	"api-seguridad/resources/request/domain/repository"

	"gorm.io/gorm"
)

type RequestRepositoryImpl struct {
	db *gorm.DB
}

func NewRequestRepository(db *gorm.DB) repository.RequestRepository {
	return &RequestRepositoryImpl{db: db}
}

func (r *RequestRepositoryImpl) Create(ctx context.Context, request *entity.Request) error {
	return r.db.WithContext(ctx).Create(request).Error
}

func (r *RequestRepositoryImpl) GetByID(ctx context.Context, id uint) (*entity.Request, error) {
	var request entity.Request
	err := r.db.WithContext(ctx).
		Preload("Municipality").
		Preload("Status").
		Preload("Police").
		First(&request, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &request, err
}

func (r *RequestRepositoryImpl) Update(ctx context.Context, request *entity.Request) error {
	return r.db.WithContext(ctx).Save(request).Error
}

func (r *RequestRepositoryImpl) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&entity.Request{}, id).Error
}

func (r *RequestRepositoryImpl) List(ctx context.Context) ([]*entity.Request, error) {
	var requests []*entity.Request
	err := r.db.WithContext(ctx).
		Preload("Municipality").
		Preload("Status").
		Preload("Police").
		Find(&requests).Error
	return requests, err
}