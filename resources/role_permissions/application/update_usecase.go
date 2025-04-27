//api-seguridad/resources/role_permissions/application/update_usecase.go
package application

import (
	"context"
	"errors"
	"api-seguridad/resources/role_permissions/domain/entities"
	"api-seguridad/resources/role_permissions/domain/repository"
)

type UpdateRolePermissionUseCase struct {
	repo repository.RolePermissionRepository
}

func NewUpdateRolePermissionUseCase(repo repository.RolePermissionRepository) *UpdateRolePermissionUseCase {
	return &UpdateRolePermissionUseCase{repo: repo}
}

func (uc *UpdateRolePermissionUseCase) Execute(ctx context.Context, rolePermission *entities.RolePermission) error {
	if rolePermission.ID == 0 {
		return errors.New("ID inválido")
	}

	existing, err := uc.repo.GetByID(ctx, rolePermission.ID)
	if err != nil {
		return err
	}
	if existing == nil || existing.IsDeleted() {
		return errors.New("relación rol-permiso no encontrada")
	}

	// Validar que no se modifiquen las claves foráneas
	if existing.RoleID != rolePermission.RoleID || existing.PermissionID != rolePermission.PermissionID {
		return errors.New("no se pueden modificar los IDs de rol o permiso")
	}

	return uc.repo.Update(ctx, rolePermission)
}