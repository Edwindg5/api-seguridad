//api-seguridad/resources/permissions/infrastructure/adapters/permission_repository_impl.go
package adapters

import (
	"context"
	"errors"
	"time"
	"gorm.io/gorm"
	"api-seguridad/resources/permissions/domain/entities"
	"api-seguridad/resources/permissions/domain/repository"
)

type PermissionRepositoryImpl struct {
	db *gorm.DB
}

func NewPermissionRepository(db *gorm.DB) repository.PermissionRepository {
	return &PermissionRepositoryImpl{db: db}
}

// Create implementa la creación de un nuevo permiso
func (r *PermissionRepositoryImpl) Create(ctx context.Context, permission *entities.Permission) error {
	// Establecer campos de auditoría
	if permission.CreatedAt.IsZero() {
		permission.CreatedAt = time.Now()
	}
	permission.UpdatedAt = time.Now()
	permission.Deleted = false

	return r.db.WithContext(ctx).Create(permission).Error
}

// GetByID obtiene un permiso por su ID
func (r *PermissionRepositoryImpl) GetByID(ctx context.Context, id uint) (*entities.Permission, error) {
	var permission entities.Permission
	err := r.db.WithContext(ctx).
		Where("id_permission = ? AND deleted = ?", id, false).
		Preload("CreatedByUser").
		Preload("UpdatedByUser").
		First(&permission).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &permission, err
}

// GetAll obtiene todos los permisos no eliminados
func (r *PermissionRepositoryImpl) GetAll(ctx context.Context) ([]*entities.Permission, error) {
	var permissions []*entities.Permission
	err := r.db.WithContext(ctx).
		Where("deleted = ?", false).
		Preload("CreatedByUser").
		Preload("UpdatedByUser").
		Order("created_at DESC").
		Find(&permissions).Error
	return permissions, err
}

// Update actualiza un permiso existente
func (r *PermissionRepositoryImpl) Update(ctx context.Context, permission *entities.Permission) error {
	permission.UpdatedAt = time.Now()
	return r.db.WithContext(ctx).Save(permission).Error
}

// SoftDelete realiza un borrado lógico
func (r *PermissionRepositoryImpl) SoftDelete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).
		Model(&entities.Permission{}).
		Where("id_permission = ?", id).
		Updates(map[string]interface{}{
			"deleted":    true,
			"updated_at": time.Now(),
		}).Error
}