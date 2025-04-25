package adapters

import (
	"context"
	"errors"
	"time"
	"gorm.io/gorm"
	"api-seguridad/resources/role_permissions/domain/entities"
	"api-seguridad/resources/role_permissions/domain/repository"
)

type RolePermissionRepositoryImpl struct {
	db *gorm.DB
}

func NewRolePermissionRepository(db *gorm.DB) repository.RolePermissionRepository {
	return &RolePermissionRepositoryImpl{db: db}
}

// Create implementa la creación de una nueva relación rol-permiso
func (r *RolePermissionRepositoryImpl) Create(ctx context.Context, rolePermission *entities.RolePermission) error {
	// Establecer campos de auditoría
	if rolePermission.CreatedAt.IsZero() {
		rolePermission.CreatedAt = time.Now()
	}
	rolePermission.UpdatedAt = time.Now()
	rolePermission.Deleted = false

	return r.db.WithContext(ctx).Create(rolePermission).Error
}

// GetByID obtiene una relación por su ID
func (r *RolePermissionRepositoryImpl) GetByID(ctx context.Context, id uint) (*entities.RolePermission, error) {
	var rolePermission entities.RolePermission
	err := r.db.WithContext(ctx).
		Where("id_role_permission = ? AND deleted = ?", id, false).
		Preload("Role").
		Preload("Permission").
		Preload("Creator").
		Preload("Updater").
		First(&rolePermission).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &rolePermission, err
}

// GetByRoleAndPermission obtiene una relación específica
func (r *RolePermissionRepositoryImpl) GetByRoleAndPermission(ctx context.Context, roleID, permissionID uint) (*entities.RolePermission, error) {
    var rolePermission entities.RolePermission
    err := r.db.WithContext(ctx).
        Where("rol_id_fk = ? AND permission_id_fk = ? AND deleted = ?", 
            roleID, permissionID, false).
        First(&rolePermission).Error

    if errors.Is(err, gorm.ErrRecordNotFound) { // Cambiado de ExprRecordNotFound a ErrRecordNotFound
        return nil, nil
    }
    return &rolePermission, err
}

// GetAllByRole obtiene todos los permisos de un rol
func (r *RolePermissionRepositoryImpl) GetAllByRole(ctx context.Context, roleID uint) ([]*entities.RolePermission, error) {
	var rolePermissions []*entities.RolePermission
	err := r.db.WithContext(ctx).
		Where("rol_id_fk = ? AND deleted = ?", roleID, false).
		Preload("Permission").
		Order("created_at DESC").
		Find(&rolePermissions).Error
	return rolePermissions, err
}

// Update actualiza una relación existente
func (r *RolePermissionRepositoryImpl) Update(ctx context.Context, rolePermission *entities.RolePermission) error {
	rolePermission.UpdatedAt = time.Now()
	return r.db.WithContext(ctx).Save(rolePermission).Error
}

// Delete realiza un borrado lógico
func (r *RolePermissionRepositoryImpl) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).
		Model(&entities.RolePermission{}).
		Where("id_role_permission = ?", id).
		Updates(map[string]interface{}{
			"deleted": true,
			"updated_at": time.Now(),
			"updated_by": gorm.Expr("updated_by"), // Mantiene el valor existente
		}).Error
}

// GrantPermission concede un permiso a un rol (crea o actualiza)
func (r *RolePermissionRepositoryImpl) GrantPermission(ctx context.Context, roleID, permissionID, userID uint) error {
	// Buscar relación existente (incluyendo eliminados)
	var existing entities.RolePermission
	err := r.db.WithContext(ctx).
		Where("rol_id_fk = ? AND permission_id_fk = ?", roleID, permissionID).
		First(&existing).Error

	// Si no existe, crear nueva
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return r.Create(ctx, &entities.RolePermission{
			RoleID:       roleID,
			PermissionID: permissionID,
			Granted:      true,
			CreatedBy:    userID,
			UpdatedBy:    userID,
		})
	}

	// Si existe pero está eliminado, restaurar
	if existing.Deleted {
		return r.db.WithContext(ctx).
			Model(&entities.RolePermission{}).
			Where("id_role_permission = ?", existing.ID).
			Updates(map[string]interface{}{
				"deleted":    false,
				"granted":    true,
				"updated_at": time.Now(),
				"updated_by": userID,
			}).Error
	}

	// Si existe y no está eliminado, actualizar
	return r.db.WithContext(ctx).
		Model(&entities.RolePermission{}).
		Where("id_role_permission = ?", existing.ID).
		Updates(map[string]interface{}{
			"granted":    true,
			"updated_at": time.Now(),
			"updated_by": userID,
		}).Error
}

// RevokePermission revoca un permiso de un rol (soft delete)
func (r *RolePermissionRepositoryImpl) RevokePermission(ctx context.Context, roleID, permissionID, userID uint) error {
	return r.db.WithContext(ctx).
		Model(&entities.RolePermission{}).
		Where("rol_id_fk = ? AND permission_id_fk = ?", roleID, permissionID).
		Updates(map[string]interface{}{
			"granted":    false,
			"deleted":    true,
			"updated_at": time.Now(),
			"updated_by": userID,
		}).Error
}