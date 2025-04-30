// api-seguridad/resources/users/application/login_usecase.go
package application

import (
	"api-seguridad/resources/users/domain/entities"
	"api-seguridad/resources/users/domain/repository"
	"context"
	"errors"
)

type LoginUseCase struct {
	userRepo repository.UserRepository
}

func NewLoginUseCase(userRepo repository.UserRepository) *LoginUseCase {
	return &LoginUseCase{userRepo: userRepo}
}

func (uc *LoginUseCase) Execute(ctx context.Context, username string, password string) (*entities.User, string, error) {
	// Validaciones básicas
	if username == "" {
		return nil, "", errors.New("validation error: username is required")
	}
	if password == "" {
		return nil, "", errors.New("validation error: password is required")
	}
	if len(password) < 8 {
		return nil, "", errors.New("validation error: password must be at least 8 characters")
	}

	// Intentar autenticar al usuario
	user, token, err := uc.userRepo.Login(ctx, username, password)
	if err != nil {
		return nil, "", errors.New("authentication failed: invalid credentials")
	}

	// Verificar que el usuario tenga un rol asignado
	if user.RoleID == 0 {
		return nil, "", errors.New("authentication failed: user has no role assigned")
	}

	return user, token, nil
}