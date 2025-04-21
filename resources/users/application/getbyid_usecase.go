// api-seguridad/resources/users/application/getbyid_usecase.go
package application

import (
	"api-seguridad/resources/users/domain/entities"
	"api-seguridad/resources/users/domain/repository"
	"context"
	"errors"
)

type GetUserByIDUseCase struct {
	userRepo repository.UserRepository
}

func NewGetUserByIDUseCase(userRepo repository.UserRepository) *GetUserByIDUseCase {
	return &GetUserByIDUseCase{userRepo: userRepo}
}

func (uc *GetUserByIDUseCase) Execute(ctx context.Context, id uint) (*entities.User, error) {
	user, err := uc.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if user == nil || user.Deleted {
		return nil, errors.New("user not found")
	}
	return user, nil
}
