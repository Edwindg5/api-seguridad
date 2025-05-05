// api-seguridad/resources/users/domain/repository/user_repository.go
package repository

import (
	"api-seguridad/resources/users/domain/entities"
	"context"
)

type UserRepository interface {
	Create(ctx context.Context, user *entities.User) error
	Update(ctx context.Context, user *entities.User) error
	List(ctx context.Context) ([]*entities.User, error)
	GetByID(ctx context.Context, id uint) (*entities.User, error)
	GetByUsername(ctx context.Context, username string) (*entities.User, error)
	GetByEmail(ctx context.Context, email string) (*entities.User, error)
	CheckRoleExists(ctx context.Context, roleID uint) (bool, error) // Nuevo método añadido

	// Opcional: Métodos específicos para borrado lógico
	SoftDelete(ctx context.Context, id uint) error
	ListActiveUsers(ctx context.Context) ([]*entities.User, error)

	// Métodos para autenticación
	Login(ctx context.Context, username string, password string) (*entities.User, string, error)
	Exists(ctx context.Context, id uint) (bool, error)
}

