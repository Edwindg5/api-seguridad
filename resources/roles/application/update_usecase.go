//api-seguridad/resources/roles/application/update_usecase.go
package application

import (
	"context"
	"errors"
	"api-seguridad/resources/roles/domain/entities"
	"api-seguridad/resources/roles/domain/repository"
)

type UpdateRoleUseCase struct {
	roleRepo repository.RoleRepository
}

func NewUpdateRoleUseCase(roleRepo repository.RoleRepository) *UpdateRoleUseCase {
	return &UpdateRoleUseCase{roleRepo: roleRepo}
}

func (uc *UpdateRoleUseCase) Execute(ctx context.Context, role *entities.Role) error {
	if role.ID == 0 {
		return errors.New("invalid role ID")
	}
	if role.Title == "" {
		return errors.New("role title is required")
	}

	existingRole, err := uc.roleRepo.GetByID(ctx, role.ID)
	if err != nil {
		return err
	}
	if existingRole == nil || existingRole.IsDeleted() {
		return errors.New("role not found")
	}

	// Verificar unicidad del título solo si está cambiando
	if existingRole.Title != role.Title {
		// Obtener todos los roles y verificar si el nuevo título ya existe
		roles, err := uc.roleRepo.GetAll(ctx)
		if err != nil {
			return err
		}
		for _, r := range roles {
			if r.Title == role.Title && !r.IsDeleted() {
				return errors.New("role with this title already exists")
			}
		}
	}

	return uc.roleRepo.Update(ctx, role)
}