//api-seguridad/resources/permissions/domain/repository/permission_repository.go
package repository

import (
	"context"
	"api-seguridad/resources/permissions/domain/entities"
)

type PermissionRepository interface {
	Create(ctx context.Context, permission *entities.Permission) error
	GetByID(ctx context.Context, id uint) (*entities.Permission, error)
	GetAll(ctx context.Context) ([]*entities.Permission, error)
	Update(ctx context.Context, permission *entities.Permission) error
	SoftDelete(ctx context.Context, id uint) error
}