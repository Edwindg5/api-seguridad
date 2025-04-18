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

func (r *RoleRepositoryImpl) Create(ctx context.Context, role *entity.Role) error {
	return r.db.WithContext(ctx).Create(role).Error
}

func (r *RoleRepositoryImpl) GetByID(ctx context.Context, id uint) (*entity.Role, error) {
	var role entity.Role
	err := r.db.WithContext(ctx).First(&role, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &role, err
}

func (r *RoleRepositoryImpl) GetByTitle(ctx context.Context, title string) (*entity.Role, error) {
	var role entity.Role
	err := r.db.WithContext(ctx).Where("title = ?", title).First(&role).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &role, err
}

func (r *RoleRepositoryImpl) Update(ctx context.Context, role *entity.Role) error {
	return r.db.WithContext(ctx).Save(role).Error
}

func (r *RoleRepositoryImpl) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&entity.Role{}, id).Error
}

func (r *RoleRepositoryImpl) List(ctx context.Context) ([]*entity.Role, error) {
	var roles []*entity.Role
	err := r.db.WithContext(ctx).Find(&roles).Error
	return roles, err
}