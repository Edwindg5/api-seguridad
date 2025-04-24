//api-seguridad/resources/roles/application/soft_delete_usecase.go
package application

import (
	"context"
	"errors"
	"api-seguridad/resources/roles/domain/repository"
)

type SoftDeleteRoleUseCase struct {
	roleRepo repository.RoleRepository
}

func NewSoftDeleteRoleUseCase(roleRepo repository.RoleRepository) *SoftDeleteRoleUseCase {
	return &SoftDeleteRoleUseCase{roleRepo: roleRepo}
}

func (uc *SoftDeleteRoleUseCase) Execute(ctx context.Context, id uint) error {
	if id == 0 {
		return errors.New("invalid role ID")
	}

	// Verify role exists before deleting
	role, err := uc.roleRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if role == nil || role.IsDeleted() {
		return errors.New("role not found")
	}

	return uc.roleRepo.SoftDelete(ctx, id)
}