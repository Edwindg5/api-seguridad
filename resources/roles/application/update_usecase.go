//api-seguridad/resources/roles/application/update_usecase.go
package application

import (
	"context"
	"errors"
	"time"
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

	if existingRole.Title != role.Title {
		if existing, err := uc.roleRepo.GetByTitle(ctx, role.Title); err == nil && existing != nil && !existing.IsDeleted() {
			return errors.New("role with this title already exists")
		}
	}

	role.CreatedAt = existingRole.CreatedAt
	role.CreatedBy = existingRole.CreatedBy
	role.UpdatedAt = time.Now()

	return uc.roleRepo.Update(ctx, role)
}

// Getter para acceder al repositorio
func (uc *UpdateRoleUseCase) GetRepository() repository.RoleRepository {
	return uc.roleRepo
}
