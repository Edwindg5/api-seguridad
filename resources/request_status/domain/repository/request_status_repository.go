// api-seguridad/resources/request_status/domain/repository/request_status_repository.go
package repository

import (
	"context"
	"api-seguridad/resources/request_status/domain/entities"
)

type RequestStatusRepository interface {
	// Basic CRUD operations
	Create(ctx context.Context, status *entities.RequestStatus) error
	GetByID(ctx context.Context, id uint) (*entities.RequestStatus, error)
	GetAll(ctx context.Context) ([]*entities.RequestStatus, error)
	Update(ctx context.Context, status *entities.RequestStatus) error
	Delete(ctx context.Context, id uint) error

	// Specialized methods
	GetByName(ctx context.Context, name string) (*entities.RequestStatus, error)
}