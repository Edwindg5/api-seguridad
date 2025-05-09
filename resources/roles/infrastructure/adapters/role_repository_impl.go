//api-seguridad/resources/roles/infrastructure/adapters/role_repository_impl.go
package adapters

import (
	"context"
	"errors"

	"api-seguridad/resources/roles/domain/entities"
	"api-seguridad/resources/roles/domain/repository"

	"gorm.io/gorm"
)

type RoleRepositoryImpl struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) repository.RoleRepository {
	return &RoleRepositoryImpl{db: db}
}

func (r *RoleRepositoryImpl) Create(ctx context.Context, role *entities.Role) error {
	return r.db.WithContext(ctx).Create(role).Error
}

func (r *RoleRepositoryImpl) GetByID(ctx context.Context, id uint) (*entities.Role, error) {
	var role entities.Role
	err := r.db.WithContext(ctx).First(&role, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &role, err
}

func (r *RoleRepositoryImpl) GetAll(ctx context.Context) ([]*entities.Role, error) {
	var roles []*entities.Role
	err := r.db.WithContext(ctx).Where("deleted = ?", false).Find(&roles).Error
	return roles, err
}
func (r *RoleRepositoryImpl) GetByTitle(ctx context.Context, title string) (*entities.Role, error) {
	var role entities.Role
	err := r.db.WithContext(ctx).
		Where("title_rol = ? AND deleted = ?", title, false).
		First(&role).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &role, err
}

func (r *RoleRepositoryImpl) Update(ctx context.Context, role *entities.Role) error {
	return r.db.WithContext(ctx).Model(role).
		Updates(map[string]interface{}{
			"title_rol":    role.Title,
			"description":  role.Description,
			"updated_by":   role.UpdatedBy,
			"updated_at":   role.UpdatedAt,
			"deleted":      role.Deleted,
		}).Error
}

func (r *RoleRepositoryImpl) SoftDelete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Model(&entities.Role{}).
		Where("id_rol = ?", id).
		Update("deleted", true).Error
}

