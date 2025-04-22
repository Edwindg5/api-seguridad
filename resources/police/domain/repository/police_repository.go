//api-seguridad/resources/police/domain/repository/police_repository.go
package repository

import (
	"context"
	"api-seguridad/resources/police/domain/entities"
)

type PoliceRepository interface {
	// Basic CRUD operations
	Create(ctx context.Context, police *entities.Police) error
	GetByID(ctx context.Context, id uint) (*entities.Police, error)
	GetAll(ctx context.Context) ([]*entities.Police, error)
	Update(ctx context.Context, police *entities.Police) error
	SoftDelete(ctx context.Context, id uint) error

	// Specialized search methods
	GetByCUIP(ctx context.Context, cuip string) (*entities.Police, error)
	GetByRFC(ctx context.Context, rfc string) (*entities.Police, error)
	SearchByName(ctx context.Context, name string) ([]*entities.Police, error)
	SearchByFullName(ctx context.Context, name, paternalName, maternalName string) ([]*entities.Police, error)
}