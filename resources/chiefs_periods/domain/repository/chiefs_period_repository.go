//api-seguridad/resources/chiefs_periods/domain/repository/chiefs_period_repository.go
package repository

import (
	"context"
	"api-seguridad/resources/chiefs_periods/domain/entities"
	"time"
)

type ChiefsPeriodRepository interface {
	// Basic CRUD operations
	Create(ctx context.Context, period *entities.ChiefsPeriod) error
	GetByID(ctx context.Context, id uint) (*entities.ChiefsPeriod, error)
	GetAll(ctx context.Context) ([]*entities.ChiefsPeriod, error)
	Update(ctx context.Context, period *entities.ChiefsPeriod) error
	SoftDelete(ctx context.Context, id uint) error

	// Specialized methods
	GetActivePeriod(ctx context.Context) (*entities.ChiefsPeriod, error)
	GetPeriodsByDateRange(ctx context.Context, start, end time.Time) ([]*entities.ChiefsPeriod, error)
}