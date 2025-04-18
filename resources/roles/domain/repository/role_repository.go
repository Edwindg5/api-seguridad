//api-seguridad/resources/roles/domain/repository/role_repository.go
package repository

import (
	"context"
	"api-seguridad/resources/roles/domain/entities"
)

type RoleRepository interface {
	Create(ctx context.Context, role *entity.Role) error
	GetByID(ctx context.Context, id uint) (*entity.Role, error)
	GetByTitle(ctx context.Context, title string) (*entity.Role, error)
	Update(ctx context.Context, role *entity.Role) error
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context) ([]*entity.Role, error)
}