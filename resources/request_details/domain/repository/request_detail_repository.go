//api-seguridad/resources/request_details/domain/repository/request_detail_repository.go
package repository

import (
	"context"
	"api-seguridad/resources/request_details/domain/entities"
)

type RequestDetailRepository interface {
	Create(ctx context.Context, detail *entities.RequestDetail) error
	GetByID(ctx context.Context, id uint) (*entities.RequestDetail, error)
	GetByRequestID(ctx context.Context, requestID uint) ([]*entities.RequestDetail, error)
	Update(ctx context.Context, detail *entities.RequestDetail) error
	Delete(ctx context.Context, id uint) error
	GetByPoliceID(ctx context.Context, policeID uint) ([]*entities.RequestDetail, error)
}