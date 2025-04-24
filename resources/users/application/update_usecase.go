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
	UserRepo repository.UserRepository // Cambiado a exportado (con mayúscula)
}

func NewUpdateUserUseCase(userRepo repository.UserRepository) *UpdateUserUseCase {
	return &UpdateUserUseCase{UserRepo: userRepo}
}

// GetUserRepo permite acceder al repositorio (opcional)
func (uc *UpdateUserUseCase) GetUserRepo() repository.UserRepository {
	return uc.UserRepo
}

func (uc *UpdateUserUseCase) Execute(ctx context.Context, user *entities.User, updaterID uint) error {
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
	user.UpdatedBy = updaterID
	
	// No actualizar created_at y created_by
	user.CreatedAt = existingUser.CreatedAt
	user.CreatedBy = existingUser.CreatedBy

	return uc.UserRepo.Update(ctx, user)
}