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
	userRepo repository.UserRepository
}

func NewUpdateUserUseCase(userRepo repository.UserRepository) *UpdateUserUseCase {
	return &UpdateUserUseCase{userRepo: userRepo}
}

func (uc *UpdateUserUseCase) Execute(ctx context.Context, user *entities.User, updaterID uint) error {
	// Validar que el usuario exista
	existingUser, err := uc.userRepo.GetByID(ctx, user.ID)
	if err != nil {
		return err
	}
	if existingUser == nil {
		return errors.New("user not found")
	}

	// Validar campos
	if user.Username != existingUser.Username {
		if _, err := uc.userRepo.GetByUsername(ctx, user.Username); err == nil {
			return errors.New("new username already exists")
		}
	}

	if user.Email != existingUser.Email {
		if _, err := uc.userRepo.GetByEmail(ctx, user.Email); err == nil {
			return errors.New("new email already exists")
		}
	}

	// Actualizar campos
	existingUser.FirstName = user.FirstName
	existingUser.LastName = user.LastName
	existingUser.Username = user.Username
	existingUser.Email = user.Email
	if user.Password != "" {
		existingUser.Password = user.Password
	}
	existingUser.RoleID = user.RoleID
	existingUser.UpdatedAt = time.Now()
	existingUser.UpdatedBy = updaterID

	return uc.userRepo.Update(ctx, existingUser)
}
