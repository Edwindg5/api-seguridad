//api-seguridad/resources/role_permissions/domain/repository/role_permission_repository.go
package repository

import (
	"context"
	"api-seguridad/resources/role_permissions/domain/entities"
)

type RolePermissionRepository interface {
	Create(ctx context.Context, rolePermission *entities.RolePermission) error
	GetByID(ctx context.Context, id uint) (*entities.RolePermission, error)
	GetByRoleAndPermission(ctx context.Context, roleID, permissionID uint) (*entities.RolePermission, error)
	GetAllByRole(ctx context.Context, roleID uint) ([]*entities.RolePermission, error)
	GetAll(ctx context.Context) ([]*entities.RolePermission, error) // New method added
	Update(ctx context.Context, rolePermission *entities.RolePermission) error
	Delete(ctx context.Context, id uint) error
}