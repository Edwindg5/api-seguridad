package adapters

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
	"api-seguridad/resources/chiefs_periods/domain/entities"
	"api-seguridad/resources/chiefs_periods/domain/repository"
)

type ChiefsPeriodRepositoryImpl struct {
	db *gorm.DB
}

func NewChiefsPeriodRepository(db *gorm.DB) repository.ChiefsPeriodRepository {
	return &ChiefsPeriodRepositoryImpl{db: db}
}

// Create implementa la creación de un nuevo periodo
func (r *ChiefsPeriodRepositoryImpl) Create(ctx context.Context, period *entities.ChiefsPeriod) error {
	// Establecer campos de auditoría
	if period.CreatedAt.IsZero() {
		period.CreatedAt = time.Now()
	}
	period.UpdatedAt = time.Now()

	return r.db.WithContext(ctx).Create(period).Error
}

// GetByID obtiene un periodo por su ID
func (r *ChiefsPeriodRepositoryImpl) GetByID(ctx context.Context, id uint) (*entities.ChiefsPeriod, error) {
	var period entities.ChiefsPeriod
	err := r.db.WithContext(ctx).
		Where("id = ? AND deleted = ?", id, false).
		Preload("CeoChief").
		Preload("LegalChief").
		Preload("Creator").
		Preload("Updater").
		First(&period).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &period, err
}

// GetAll obtiene todos los periodos no eliminados
func (r *ChiefsPeriodRepositoryImpl) GetAll(ctx context.Context) ([]*entities.ChiefsPeriod, error) {
	var periods []*entities.ChiefsPeriod
	err := r.db.WithContext(ctx).
		Where("deleted = ?", false).
		Preload("CeoChief").
		Preload("LegalChief").
		Order("start_date DESC").
		Find(&periods).Error
	return periods, err
}

// Update actualiza un periodo existente
func (r *ChiefsPeriodRepositoryImpl) Update(ctx context.Context, period *entities.ChiefsPeriod) error {
	period.UpdatedAt = time.Now()
	return r.db.WithContext(ctx).Save(period).Error
}

// SoftDelete realiza un borrado lógico
func (r *ChiefsPeriodRepositoryImpl) SoftDelete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).
		Model(&entities.ChiefsPeriod{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"deleted":    true,
			"updated_at": time.Now(),
		}).Error
}

// GetActivePeriod obtiene el periodo activo actual
func (r *ChiefsPeriodRepositoryImpl) GetActivePeriod(ctx context.Context) (*entities.ChiefsPeriod, error) {
	var period entities.ChiefsPeriod
	err := r.db.WithContext(ctx).
		Where("period_active = ? AND deleted = ?", true, false).
		Preload("CeoChief").
		Preload("LegalChief").
		First(&period).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &period, err
}

// GetPeriodsByDateRange busca periodos en un rango de fechas
func (r *ChiefsPeriodRepositoryImpl) GetPeriodsByDateRange(ctx context.Context, start, end time.Time) ([]*entities.ChiefsPeriod, error) {
	var periods []*entities.ChiefsPeriod
	err := r.db.WithContext(ctx).
		Where("((start_date BETWEEN ? AND ?) OR (end_date BETWEEN ? AND ?) OR (start_date <= ? AND (end_date >= ? OR end_date IS NULL)) AND deleted = ?",
			start, end, start, end, start, start, false).
		Preload("CeoChief").
		Preload("LegalChief").
		Order("start_date DESC").
		Find(&periods).Error
	return periods, err
}

// GetCurrentPeriod obtiene el periodo actual (que incluye la fecha actual)
func (r *ChiefsPeriodRepositoryImpl) GetCurrentPeriod(ctx context.Context) (*entities.ChiefsPeriod, error) {
	var period entities.ChiefsPeriod
	now := time.Now()
	err := r.db.WithContext(ctx).
		Where("(start_date <= ? AND (end_date >= ? OR end_date IS NULL)) AND deleted = ?",
			now, now, false).
		Preload("CeoChief").
		Preload("LegalChief").
		First(&period).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &period, err
}