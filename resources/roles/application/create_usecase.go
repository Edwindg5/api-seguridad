//api-seguridad/resources/roles/application/create_usecase.go
package application

import (
	"context"
	"errors"
	"api-seguridad/resources/roles/domain/entities"
	"api-seguridad/resources/roles/domain/repository"
)

type CreateRoleUseCase struct {
	roleRepo repository.RoleRepository
}

func NewCreateRoleUseCase(roleRepo repository.RoleRepository) *CreateRoleUseCase {
	return &CreateRoleUseCase{roleRepo: roleRepo}
}

func (uc *CreateRoleUseCase) Execute(ctx context.Context, role *entities.Role) error {
	if role.Title == "" {
		return errors.New("role title is required")
	}

	// Verificar si ya existe un rol con el mismo título
	roles, err := uc.roleRepo.GetAll(ctx)
	if err != nil {
		return err
	}
	
	for _, r := range roles {
		if r.Title == role.Title && !r.IsDeleted() {
			return errors.New("role with this title already exists")
		}
	}

	// Establecer valores por defecto
	role.Deleted = false
	
	return uc.roleRepo.Create(ctx, role)
}