package application

import (
	"context"
	"errors"
	"api-seguridad/resources/users/domain/entities"
	"api-seguridad/resources/users/domain/repository"
)

type PostUserUseCase struct {
	userRepo repository.UserRepository
}

func NewPostUserUseCase(userRepo repository.UserRepository) *PostUserUseCase {
	return &PostUserUseCase{userRepo: userRepo}
}

func (uc *PostUserUseCase) Execute(ctx context.Context, user *entity.User) error {
	// Validaciones de negocio
	if user.Username == "" {
		return errors.New("username is required")
	}
	if user.Email == "" {
		return errors.New("email is required")
	}
	if user.Password == "" {
		return errors.New("password is required")
	}
	if user.RoleID == 0 {
		return errors.New("role_id is required")
	}

	// Verificar si el usuario ya existe
	existingUser, err := uc.userRepo.GetByUsername(ctx, user.Username)
	if err != nil {
		return err
	}
	if existingUser != nil {
		return errors.New("username already exists")
	}

	// Verificar si el email ya existe
	existingEmail, err := uc.userRepo.GetByEmail(ctx, user.Email)
	if err != nil {
		return err
	}
	if existingEmail != nil {
		return errors.New("email already exists")
	}

	return uc.userRepo.Create(ctx, user)
}