// api-seguridad/resources/users/application/delete_usecase.go
package application

import (
	"context"
	"errors"

	"api-seguridad/resources/users/domain/entities"
	"api-seguridad/resources/users/domain/repository"
)

type DeleteUserUseCase struct {
	userRepo repository.UserRepository
}

func NewDeleteUserUseCase(userRepo repository.UserRepository) *DeleteUserUseCase {
	return &DeleteUserUseCase{userRepo: userRepo}
}

// Cambiado para aceptar un *User en lugar de un uint
func (uc *DeleteUserUseCase) Execute(ctx context.Context, user *entities.User) error {
	// Obtener el usuario existente
	existingUser, err := uc.userRepo.GetByID(ctx, user.ID)
	if err != nil {
		return err
	}
	if existingUser == nil {
		return errors.New("user not found")
	}
	if existingUser.Deleted {
		return errors.New("user already deleted")
	}

	// Actualizar campos para borrado lógico
	existingUser.Deleted = true
	existingUser.UpdatedAt = user.UpdatedAt
	existingUser.UpdatedBy = user.UpdatedBy

	return uc.userRepo.Update(ctx, existingUser)
}