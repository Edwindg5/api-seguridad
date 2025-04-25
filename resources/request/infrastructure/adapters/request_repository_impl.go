// api-seguridad/resources/request/infrastructure/adapters/request_repository_impl.go
package adapters

import (
	"api-seguridad/resources/request/domain/entities"
	"api-seguridad/resources/request/domain/repository"
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
)

type RequestRepositoryImpl struct {
	db *gorm.DB
}

func NewRequestRepository(db *gorm.DB) repository.RequestRepository {
	return &RequestRepositoryImpl{db: db}
}

func (r *RequestRepositoryImpl) Create(ctx context.Context, request *entities.Request) error {
	// Ensure audit fields are set
	if request.CreatedAt.IsZero() {
		request.CreatedAt = time.Now()
	}
	if request.UpdatedAt.IsZero() {
		request.UpdatedAt = time.Now()
	}
	
	return r.db.WithContext(ctx).Create(request).Error
}

func (r *RequestRepositoryImpl) GetByID(ctx context.Context, id uint) (*entities.Request, error) {
    var request entities.Request
    err := r.db.WithContext(ctx).
        Where("id_request = ? AND deleted = ?", id, false).
        Preload("Municipalities").
        Preload("Status").
        Preload("CeoChief").
        Preload("LegalChief").
        Preload("CreatedByUser").
        Preload("UpdatedByUser").
        First(&request).Error
        
    if errors.Is(err, gorm.ErrRecordNotFound) {
        return nil, nil
    }
    return &request, err
}

func (r *RequestRepositoryImpl) Update(ctx context.Context, request *entities.Request) error {
	request.UpdatedAt = time.Now()
	return r.db.WithContext(ctx).Save(request).Error
}

func (r *RequestRepositoryImpl) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).
		Model(&entities.Request{}).
		Where("id_request = ?", id).
		Updates(map[string]interface{}{
			"deleted":    true,
			"updated_at": time.Now(),
		}).Error
}

func (r *RequestRepositoryImpl) GetByStatus(ctx context.Context, statusID uint) ([]*entities.Request, error) {
	var requests []*entities.Request
	err := r.db.WithContext(ctx).
		Where("id_status_fk = ? AND deleted = ?", statusID, false).
		Preload("Municipalities").
		Preload("Status").
		Preload("CeoChief").
		Preload("LegalChief").
		Find(&requests).Error
	return requests, err
}

func (r *RequestRepositoryImpl) GetByMunicipality(ctx context.Context, municipalityID uint) ([]*entities.Request, error) {
	var requests []*entities.Request
	err := r.db.WithContext(ctx).
		Where("id_municipalities_fk = ? AND deleted = ?", municipalityID, false).
		Preload("Municipalities").
		Preload("Status").
		Preload("CeoChief").
		Preload("LegalChief").
		Find(&requests).Error
	return requests, err
}