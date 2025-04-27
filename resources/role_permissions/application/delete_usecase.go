//api-seguridad/resources/role_permissions/application/delete_usecase.go
package application

import (
	"context"
	"errors"
	"api-seguridad/resources/role_permissions/domain/repository"
)

type DeleteRolePermissionUseCase struct {
	repo repository.RolePermissionRepository
}

func NewDeleteRolePermissionUseCase(repo repository.RolePermissionRepository) *DeleteRolePermissionUseCase {
	return &DeleteRolePermissionUseCase{repo: repo}
}

func (uc *DeleteRolePermissionUseCase) Execute(ctx context.Context, id uint) error {
	if id == 0 {
		return errors.New("ID inválido")
	}

	existing, err := uc.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if existing == nil || existing.IsDeleted() {
		return errors.New("relación rol-permiso no encontrada")
	}

	return uc.repo.Delete(ctx, id)
}