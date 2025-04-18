// api-seguridad/resources/request/application/request_service.go
package application

import (
	"context"
	"api-seguridad/resources/request/domain/entities"
	"api-seguridad/resources/request/domain/repository"
)

type RequestService struct {
	requestRepo repository.RequestRepository
}

func NewRequestService(requestRepo repository.RequestRepository) *RequestService {
	return &RequestService{requestRepo: requestRepo}
}

func (s *RequestService) CreateRequest(ctx context.Context, request *entity.Request) error {
	postUC := NewPostRequestUseCase(s.requestRepo)
	return postUC.Execute(ctx, request)
}

func (s *RequestService) GetRequestByID(ctx context.Context, id uint) (*entity.Request, error) {
	return s.requestRepo.GetByID(ctx, id)
}

func (s *RequestService) UpdateRequest(ctx context.Context, request *entity.Request) error {
	return s.requestRepo.Update(ctx, request)
}

func (s *RequestService) DeleteRequest(ctx context.Context, id uint) error {
	return s.requestRepo.Delete(ctx, id)
}

func (s *RequestService) ListRequests(ctx context.Context) ([]*entity.Request, error) {
	return s.requestRepo.List(ctx)
}