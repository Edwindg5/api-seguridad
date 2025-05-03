// api-seguridad/resources/users/application/update_usecase.go
package application

import (
	"api-seguridad/resources/users/domain/entities"
	"api-seguridad/resources/users/domain/repository"
	"context"
	"errors"
	"time"
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
		return err
	}
	if existingUser == nil || existingUser.Deleted {
		return errors.New("user not found")
	}

	// Validar campos únicos si han cambiado
	if user.Username != existingUser.Username {
		if existing, err := uc.UserRepo.GetByUsername(ctx, user.Username); err == nil && existing != nil {
			return errors.New("new username already exists")
		}
	}

	if user.Email != existingUser.Email {
		if existing, err := uc.UserRepo.GetByEmail(ctx, user.Email); err == nil && existing != nil {
			return errors.New("new email already exists")
		}
	}

	// Preparar datos para actualización
	user.UpdatedAt = time.Now()
	
	// Mantener datos originales de creación
	user.CreatedAt = existingUser.CreatedAt
	user.CreatedBy = existingUser.CreatedBy
	user.Deleted = existingUser.Deleted

	return uc.UserRepo.Update(ctx, user)
}