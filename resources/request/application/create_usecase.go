// api-seguridad/resources/request/application/create_usecase.go
package application

import (
	"context"
	"errors"
	"api-seguridad/resources/request/domain/entities"
	"api-seguridad/resources/request/domain/repository"
	"time"
)

type CreateRequestUseCase struct {
	repo repository.RequestRepository
}

func NewCreateRequestUseCase(repo repository.RequestRepository) *CreateRequestUseCase {
	return &CreateRequestUseCase{repo: repo}
}

func (uc *CreateRequestUseCase) Execute(ctx context.Context, request *entities.Request) error {
	// Validations
	if request.OfficeNumber == "" {
		return errors.New("office number is required")
	}
	if request.MunicipalitiesID == 0 {
		return errors.New("municipality is required")
	}
	if request.StatusID == 0 {
		return errors.New("status is required")
	}
	if request.CreatedBy == 0 {
		return errors.New("creator user is required")
	}

	// Set timestamps
	if request.ReceiptDate.IsZero() {
		request.ReceiptDate = time.Now()
	}
	if request.Date.IsZero() {
		request.Date = time.Now()
	}
	if request.CreatedAt.IsZero() {
		request.CreatedAt = time.Now()
	}
	request.UpdatedAt = time.Now()

	return uc.repo.Create(ctx, request)
}