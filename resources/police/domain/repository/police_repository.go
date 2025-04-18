//api-seguridad/resources/police/domain/repository/police_repository.go
package repository

import (
	"context"
	"api-seguridad/resources/police/domain/entities"
)

type PoliceRepository interface {
	Create(ctx context.Context, police *entity.Police) error
	GetByID(ctx context.Context, id uint) (*entity.Police, error)
	GetByCUIP(ctx context.Context, cuip string) (*entity.Police, error)
	Update(ctx context.Context, police *entity.Police) error
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context) ([]*entity.Police, error)
}