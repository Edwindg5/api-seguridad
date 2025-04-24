// api-seguridad/resources/area_chiefs/infrastructure/adapters/area_chief_repository_impl.go
package adapters

import (
	"api-seguridad/resources/area_chiefs/domain/entities"
	"api-seguridad/resources/area_chiefs/domain/repository"
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
)

type AreaChiefRepositoryImpl struct {
	db *gorm.DB
}

func NewAreaChiefRepository(db *gorm.DB) repository.AreaChiefRepository {
	return &AreaChiefRepositoryImpl{db: db}
}

func (r *AreaChiefRepositoryImpl) Create(ctx context.Context, chief *entities.AreaChief) error {
	// Set default timestamps if not provided
	if chief.CreatedAt.IsZero() {
		chief.CreatedAt = time.Now()
	}
	if chief.UpdatedAt.IsZero() {
		chief.UpdatedAt = time.Now()
	}

	return r.db.WithContext(ctx).Create(chief).Error
}

func (r *AreaChiefRepositoryImpl) GetByID(ctx context.Context, id uint) (*entities.AreaChief, error) {
	var chief entities.AreaChief
	err := r.db.WithContext(ctx).
		Where("id_official = ? AND deleted = ?", id, false).
		Preload("CreatedByUser").
		Preload("UpdatedByUser").
		First(&chief).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &chief, err
}

func (r *AreaChiefRepositoryImpl) GetAll(ctx context.Context) ([]*entities.AreaChief, error) {
	var chiefs []*entities.AreaChief
	err := r.db.WithContext(ctx).
		Where("deleted = ?", false).
		Preload("CreatedByUser").
		Preload("UpdatedByUser").
		Find(&chiefs).Error
	return chiefs, err
}

func (r *AreaChiefRepositoryImpl) Update(ctx context.Context, chief *entities.AreaChief) error {
	chief.UpdatedAt = time.Now()
	return r.db.WithContext(ctx).Save(chief).Error
}

func (r *AreaChiefRepositoryImpl) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).
		Model(&entities.AreaChief{}).
		Where("id_official = ?", id).
		Updates(map[string]interface{}{
			"deleted":    true,
			"updated_at": time.Now(),
			"updated_by": ctx.Value("userID"), // Set updater from context
		}).Error
}