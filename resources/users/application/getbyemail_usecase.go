// api-seguridad/resources/users/application/getbyemail_usecase.go
package application

import (
	"api-seguridad/resources/users/domain/entities"
	"api-seguridad/resources/users/domain/repository"
	"context"
	"errors"
)

type GetUserByEmailUseCase struct {
	userRepo repository.UserRepository
}

func NewGetUserByEmailUseCase(userRepo repository.UserRepository) *GetUserByEmailUseCase {
	return &GetUserByEmailUseCase{userRepo: userRepo}
}

func (uc *GetUserByEmailUseCase) Execute(ctx context.Context, email string) (*entities.User, error) {
	user, err := uc.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if user == nil || user.Deleted {
		return nil, errors.New("user not found")
	}
	return user, nil
}
