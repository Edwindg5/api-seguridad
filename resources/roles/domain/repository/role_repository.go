//api-seguridad/resources/roles/domain/repository/role_repository.go
package repository

import (
	"context"
	"api-seguridad/resources/roles/domain/entities"
)

type RoleRepository interface {

	Create(ctx context.Context, role *entities.Role) error
	GetByID(ctx context.Context, id uint) (*entities.Role, error)
	GetAll(ctx context.Context) ([]*entities.Role, error)
	Update(ctx context.Context, role *entities.Role) error
	SoftDelete(ctx context.Context, id uint) error
	
}