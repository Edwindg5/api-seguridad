//api-seguridad/resources/type_police/domain/repository/type_police_repository.go
package repository

import (
	"context"
	"api-seguridad/resources/type_police/domain/entities"
)

type TypePoliceRepository interface {
	Create(ctx context.Context, typePolice *entities.TypePolice) error
	GetByID(ctx context.Context, id uint) (*entities.TypePolice, error)
	GetAll(ctx context.Context) ([]*entities.TypePolice, error)
	Update(ctx context.Context, typePolice *entities.TypePolice) error
	SoftDelete(ctx context.Context, id uint) error
}