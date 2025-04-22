//api-seguridad/resources/municipalities/domain/repository/municipality_repository.go
package repository

import (
	"context"
	"api-seguridad/resources/municipalities/domain/entities"
)

type MunicipalityRepository interface {
	Create(ctx context.Context, municipality *entities.Municipality) error
	GetByID(ctx context.Context, id uint) (*entities.Municipality, error)
	GetByName(ctx context.Context, name string) (*entities.Municipality, error)
	GetAll(ctx context.Context) ([]*entities.Municipality, error)
	Update(ctx context.Context, municipality *entities.Municipality) error
	SoftDelete(ctx context.Context, id uint) error
}