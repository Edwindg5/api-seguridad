// api-seguridad/resources/users/application/getbyusername_usecase.go
package application

import (
	"api-seguridad/resources/users/domain/entities"
	"api-seguridad/resources/users/domain/repository"
	"context"
	"errors"
)

type GetUserByUsernameUseCase struct {
	userRepo repository.UserRepository
}

func NewGetUserByUsernameUseCase(userRepo repository.UserRepository) *GetUserByUsernameUseCase {
	return &GetUserByUsernameUseCase{userRepo: userRepo}
}

func (uc *GetUserByUsernameUseCase) Execute(ctx context.Context, username string) (*entities.User, error) {
	user, err := uc.userRepo.GetByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	if user == nil || user.Deleted {
		return nil, errors.New("user not found")
	}
	return user, nil
}
