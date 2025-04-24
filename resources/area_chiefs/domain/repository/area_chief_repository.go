// api-seguridad/resources/area_chiefs/domain/repository/area_chief_repository.go
package repository

import (
	"context"
	"api-seguridad/resources/area_chiefs/domain/entities"
)

type AreaChiefRepository interface {
	// Basic CRUD operations
	Create(ctx context.Context, chief *entities.AreaChief) error
	GetByID(ctx context.Context, id uint) (*entities.AreaChief, error)
	GetAll(ctx context.Context) ([]*entities.AreaChief, error)
	Update(ctx context.Context, chief *entities.AreaChief) error
	Delete(ctx context.Context, id uint) error
}