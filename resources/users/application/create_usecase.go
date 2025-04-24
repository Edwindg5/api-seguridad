// api-seguridad/resources/users/application/create_usecase.go
package application

import (
	"api-seguridad/resources/users/domain/entities"
	"api-seguridad/resources/users/domain/repository"
	"context"
	"errors"
	"fmt"
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
        return errors.New("validation error: username is required")
    }
    if user.Email == "" {
        return errors.New("validation error: email is required")
    }
    if len(user.Password) < 8 {
        return fmt.Errorf("validation error: password must be at least 8 characters (received %d)", len(user.Password))
    }
    if user.RoleID == 0 {
        return errors.New("validation error: role is required")
    }

    // Validaciones de existencia
    if existingUser, err := uc.userRepo.GetByUsername(ctx, user.Username); err == nil && existingUser != nil {
        return errors.New("username already exists")
    }

    if existingUser, err := uc.userRepo.GetByEmail(ctx, user.Email); err == nil && existingUser != nil {
        return errors.New("email already exists")
    }

    // Establecer valores por defecto
    user.Deleted = false
    user.CreatedAt = time.Now()
    user.UpdatedAt = time.Now()

    // Si es el primer usuario, establecer created_by y updated_by como 0
    if user.CreatedBy == 0 && user.UpdatedBy == 0 {
        user.CreatedBy = 0 // O puedes dejarlo como 0 si prefieres
        user.UpdatedBy = 0
    }

    return uc.userRepo.Create(ctx, user)
}
