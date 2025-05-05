// api-seguridad/resources/users/application/update_usecase.go
package application

import (
	"api-seguridad/resources/users/domain/entities"
	"api-seguridad/resources/users/domain/repository"
	"context"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type UpdateUserUseCase struct {
	UserRepo repository.UserRepository
}

func NewUpdateUserUseCase(userRepo repository.UserRepository) *UpdateUserUseCase {
	return &UpdateUserUseCase{UserRepo: userRepo}
}

func (uc *UpdateUserUseCase) Execute(ctx context.Context, user *entities.User) error {
    // Validar que el usuario exista
    existingUser, err := uc.UserRepo.GetByID(ctx, user.ID)
    if err != nil {
        return fmt.Errorf("error checking user existence: %w", err)
    }
    if existingUser == nil || existingUser.Deleted {
        return errors.New("user not found")
    }

    // Validación estricta del RoleID
    if user.RoleID == 0 {
        return errors.New("role ID must be provided")
    }

    // Validar que el rol exista en la base de datos
    roleExists, err := uc.UserRepo.CheckRoleExists(ctx, user.RoleID)
    if err != nil {
        return fmt.Errorf("error checking role existence: %w", err)
    }
    if !roleExists {
        return errors.New("role not found")
    }

    // Validar campos únicos si han cambiado
    if user.Username != existingUser.Username {
        existing, err := uc.UserRepo.GetByUsername(ctx, user.Username)
        if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
            return fmt.Errorf("error checking username: %w", err)
        }
        if existing != nil {
            return errors.New("new username already exists")
        }
    }

    if user.Email != existingUser.Email {
        existing, err := uc.UserRepo.GetByEmail(ctx, user.Email)
        if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
            return fmt.Errorf("error checking email: %w", err)
        }
        if existing != nil {
            return errors.New("new email already exists")
        }
    }

    // Preparar datos para actualización
    user.UpdatedAt = time.Now()
    user.CreatedAt = existingUser.CreatedAt
    user.CreatedBy = existingUser.CreatedBy
    user.Deleted = existingUser.Deleted

    return uc.UserRepo.Update(ctx, user)
}