// api-seguridad/resources/users/application/create_usecase.go
package application

import (
	"api-seguridad/resources/users/domain/entities"
	"api-seguridad/resources/users/domain/repository"
	"context"
	"errors"
	"time"
)

type CreateUserUseCase struct {
	userRepo repository.UserRepository
}

func NewCreateUserUseCase(userRepo repository.UserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{userRepo: userRepo}
}

func (uc *CreateUserUseCase) Execute(ctx context.Context, user *entities.User) error {
	// Validaciones básicas
	if user.Username == "" {
		return errors.New("username is required")
	}
	if user.Email == "" {
		return errors.New("email is required")
	}
	if len(user.Password) < 8 {
		return errors.New("password must be at least 8 characters")
	}
	if user.RoleID == 0 {
		return errors.New("role is required")
	}

	// Validaciones de existencia
	if _, err := uc.userRepo.GetByUsername(ctx, user.Username); err == nil {
		return errors.New("username already exists")
	}

	if _, err := uc.userRepo.GetByEmail(ctx, user.Email); err == nil {
		return errors.New("email already exists")
	}

	// Establecer valores por defecto
	user.Deleted = false
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	return uc.userRepo.Create(ctx, user)
}
