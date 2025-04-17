package application

import (
	"context"
	"errors"
	"api-seguridad/resources/request/domain/entities"
	"api-seguridad/resources/request/domain/repository"
	"time"
)

type PostRequestUseCase struct {
	requestRepo repository.RequestRepository
}

func NewPostRequestUseCase(requestRepo repository.RequestRepository) *PostRequestUseCase {
	return &PostRequestUseCase{requestRepo: requestRepo}
}

func (uc *PostRequestUseCase) Execute(ctx context.Context, request *entity.Request) error {
	if request.ReceiptDate.IsZero() {
		return errors.New("receipt date is required")
	}
	if request.MunicipalityID == 0 {
		return errors.New("municipality_id is required")
	}
	if request.StatusID == 0 {
		return errors.New("status_id is required")
	}
	if request.Date.IsZero() {
		request.Date = time.Now()
	}

	return uc.requestRepo.Create(ctx, request)
}