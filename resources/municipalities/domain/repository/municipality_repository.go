package repository

import (
	"context"
	"api-seguridad/resources/municipalities/domain/entities"
)

type MunicipalityRepository interface {
	Create(ctx context.Context, municipality *entity.Municipality) error
	GetByID(ctx context.Context, id uint) (*entity.Municipality, error)
	GetByName(ctx context.Context, name string) (*entity.Municipality, error)
	Update(ctx context.Context, municipality *entity.Municipality) error
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context) ([]*entity.Municipality, error)
}