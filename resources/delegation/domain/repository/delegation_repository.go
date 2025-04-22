package repository

import (
	"context"
	"api-seguridad/resources/delegation/domain/entities"
)

type DelegationRepository interface {
	Create(ctx context.Context, delegation *entities.Delegation) error
	GetByID(ctx context.Context, id uint) (*entities.Delegation, error)
	GetAll(ctx context.Context) ([]*entities.Delegation, error)
	Update(ctx context.Context, delegation *entities.Delegation) error
	SoftDelete(ctx context.Context, id uint) error
}