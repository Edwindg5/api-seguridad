// api-seguridad/resources/request_status/infrastructure/adapters/request_status_repository_impl.go
package adapters

import (
	"api-seguridad/resources/request_status/domain/entities"
	"api-seguridad/resources/request_status/domain/repository"
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
)

type RequestStatusRepositoryImpl struct {
	db *gorm.DB
}

func NewRequestStatusRepository(db *gorm.DB) repository.RequestStatusRepository {
	return &RequestStatusRepositoryImpl{db: db}
}

func (r *RequestStatusRepositoryImpl) Create(ctx context.Context, status *entities.RequestStatus) error {
	// Ensure audit fields are set
	if status.CreatedAt.IsZero() {
		status.CreatedAt = time.Now()
	}
	if status.UpdatedAt.IsZero() {
		status.UpdatedAt = time.Now()
	}
	
	return r.db.WithContext(ctx).Create(status).Error
}

func (r *RequestStatusRepositoryImpl) GetByID(ctx context.Context, id uint) (*entities.RequestStatus, error) {
	var status entities.RequestStatus
	err := r.db.WithContext(ctx).
		Where("id_status = ? AND deleted = ?", id, false).
		Preload("CreatedByUser").
		Preload("UpdatedByUser").
		First(&status).Error
		
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &status, err
}

func (r *RequestStatusRepositoryImpl) GetAll(ctx context.Context) ([]*entities.RequestStatus, error) {
	var statusList []*entities.RequestStatus
	err := r.db.WithContext(ctx).
		Where("deleted = ?", false).
		Preload("CreatedByUser").
		Preload("UpdatedByUser").
		Find(&statusList).Error
	return statusList, err
}

func (r *RequestStatusRepositoryImpl) GetByName(ctx context.Context, name string) (*entities.RequestStatus, error) {
	var status entities.RequestStatus
	err := r.db.WithContext(ctx).
		Where("name = ? AND deleted = ?", name, false).
		First(&status).Error
		
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &status, err
}

func (r *RequestStatusRepositoryImpl) Update(ctx context.Context, status *entities.RequestStatus) error {
    return r.db.WithContext(ctx).Model(status).
        Updates(map[string]interface{}{
            "name":        status.Name,
            "description": status.Description,
            "updated_by":  status.UpdatedBy,
            "updated_at":  time.Now(),
        }).Error
}

func (r *RequestStatusRepositoryImpl) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).
		Model(&entities.RequestStatus{}).
		Where("id_status = ?", id).
		Updates(map[string]interface{}{
			"deleted":    true,
			"updated_at": time.Now(),
		}).Error
}