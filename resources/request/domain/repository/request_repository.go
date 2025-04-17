package repository

import (
	"context"
	"api-seguridad/resources/request/domain/entities"
)

type RequestRepository interface {
	Create(ctx context.Context, request *entity.Request) error
	GetByID(ctx context.Context, id uint) (*entity.Request, error)
	Update(ctx context.Context, request *entity.Request) error
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context) ([]*entity.Request, error)
}