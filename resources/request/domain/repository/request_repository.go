// api-seguridad/resources/request/domain/repository/request_repository.go
package repository

import (
	"context"
	"api-seguridad/resources/request/domain/entities"
)

type RequestRepository interface {
	// Basic CRUD operations
	Create(ctx context.Context, request *entities.Request) error
	GetByID(ctx context.Context, id uint) (*entities.Request, error)
	Update(ctx context.Context, request *entities.Request) error
	Delete(ctx context.Context, id uint) error
	GetAll(ctx context.Context) ([]*entities.Request, error)

	// Specialized search methods
	GetByStatus(ctx context.Context, statusID uint) ([]*entities.Request, error)
	GetByMunicipality(ctx context.Context, municipalityID uint) ([]*entities.Request, error)
}