// api-seguridad/resources/users/application/delete_usecase.go
package application

import (
	"context"
	"errors"
	"time"
	"api-seguridad/resources/users/domain/repository"
)

type DeleteUserUseCase struct {
	userRepo repository.UserRepository
}

func NewDeleteUserUseCase(userRepo repository.UserRepository) *DeleteUserUseCase {
	return &DeleteUserUseCase{userRepo: userRepo}
}

func (uc *DeleteUserUseCase) Execute(ctx context.Context, id uint) error {
	// Obtener el usuario existente
	user, err := uc.userRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("user not found")
	}
	if user.Deleted {
		return errors.New("user already deleted")
	}

	// Realizar borrado lógico
	user.Deleted = true
	user.UpdatedAt = time.Now()

	return uc.userRepo.Update(ctx, user)
}