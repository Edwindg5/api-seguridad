package adapters

import (
	"context"
	"errors"
	"time"
	"gorm.io/gorm"
	"api-seguridad/resources/request_details/domain/entities"
	"api-seguridad/resources/request_details/domain/repository"
)

type RequestDetailRepositoryImpl struct {
	db *gorm.DB
}

func NewRequestDetailRepository(db *gorm.DB) repository.RequestDetailRepository {
	return &RequestDetailRepositoryImpl{db: db}
}

// Create implementa la creación de un nuevo detalle de solicitud
func (r *RequestDetailRepositoryImpl) Create(ctx context.Context, detail *entities.RequestDetail) error {
	// Establecer campos de auditoría
	if detail.CreatedAt.IsZero() {
		detail.CreatedAt = time.Now()
	}
	detail.UpdatedAt = time.Now()
	detail.Deleted = false

	return r.db.WithContext(ctx).Create(detail).Error
}

// GetByID obtiene un detalle por su ID
func (r *RequestDetailRepositoryImpl) GetByID(ctx context.Context, id uint) (*entities.RequestDetail, error) {
	var detail entities.RequestDetail
	err := r.db.WithContext(ctx).
		Where("id = ? AND deleted = ?", id, false).
		Preload("Police").
		Preload("Creator").
		Preload("Updater").
		First(&detail).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &detail, err
}

// GetByRequestID obtiene todos los detalles de una solicitud
func (r *RequestDetailRepositoryImpl) GetByRequestID(ctx context.Context, requestID uint) ([]*entities.RequestDetail, error) {
	var details []*entities.RequestDetail
	err := r.db.WithContext(ctx).
		Where("request_id = ? AND deleted = ?", requestID, false).
		Preload("Police").
		Order("created_at DESC").
		Find(&details).Error
	return details, err
}

// GetByPoliceID obtiene todos los detalles asociados a un policía
func (r *RequestDetailRepositoryImpl) GetByPoliceID(ctx context.Context, policeID uint) ([]*entities.RequestDetail, error) {
	var details []*entities.RequestDetail
	err := r.db.WithContext(ctx).
		Where("id_police_fk = ? AND deleted = ?", policeID, false).
		Order("created_at DESC").
		Find(&details).Error
	return details, err
}

// Update actualiza un detalle existente
func (r *RequestDetailRepositoryImpl) Update(ctx context.Context, detail *entities.RequestDetail) error {
	detail.UpdatedAt = time.Now()
	return r.db.WithContext(ctx).Save(detail).Error
}

// Delete realiza un borrado lógico
func (r *RequestDetailRepositoryImpl) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).
		Model(&entities.RequestDetail{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"deleted":     true,
			"updated_at":  time.Now(),
			// Mantiene el valor existente de updated_by
		}).Error
}

// ApproveDetail marca un detalle como aprobado
func (r *RequestDetailRepositoryImpl) ApproveDetail(ctx context.Context, id uint, approvedBy uint) error {
	return r.db.WithContext(ctx).
		Model(&entities.RequestDetail{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"approved":    true,
			"updated_at":  time.Now(),
			"updated_by":  approvedBy,
		}).Error
}

// RejectDetail marca un detalle como rechazado
func (r *RequestDetailRepositoryImpl) RejectDetail(ctx context.Context, id uint, rejectedBy uint, comments string) error {
	return r.db.WithContext(ctx).
		Model(&entities.RequestDetail{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"approved":    false,
			"comments":    comments,
			"updated_at":  time.Now(),
			"updated_by":  rejectedBy,
		}).Error
}