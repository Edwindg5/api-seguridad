// api-seguridad/resources/users/application/list_usecase.go
package application

import (
	"api-seguridad/resources/users/domain/entities"
	"api-seguridad/resources/users/domain/repository"
	"context"
)

type ListUsersUseCase struct {
	userRepo repository.UserRepository
}

func NewListUsersUseCase(userRepo repository.UserRepository) *ListUsersUseCase {
	return &ListUsersUseCase{userRepo: userRepo}
}

func (uc *ListUsersUseCase) Execute(ctx context.Context) ([]*entities.User, error) {
	users, err := uc.userRepo.List(ctx)
	if err != nil {
		return nil, err
	}

	// Filtrar usuarios eliminados
	var activeUsers []*entities.User
	for _, user := range users {
		if !user.Deleted {
			activeUsers = append(activeUsers, user)
		}
	}

	return activeUsers, nil
}
