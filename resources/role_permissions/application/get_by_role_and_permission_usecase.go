package application

import (
	"context"
	"errors"
	"api-seguridad/resources/role_permissions/domain/entities"
	"api-seguridad/resources/role_permissions/domain/repository"
)

type GetByRoleAndPermissionUseCase struct {
	repo repository.RolePermissionRepository
}

func NewGetByRoleAndPermissionUseCase(repo repository.RolePermissionRepository) *GetByRoleAndPermissionUseCase {
	return &GetByRoleAndPermissionUseCase{repo: repo}
}

func (uc *GetByRoleAndPermissionUseCase) Execute(ctx context.Context, roleID, permissionID uint) (*entities.RolePermission, error) {
	if roleID == 0 || permissionID == 0 {
		return nil, errors.New("IDs de rol y permiso son requeridos")
	}

	rolePermission, err := uc.repo.GetByRoleAndPermission(ctx, roleID, permissionID)
	if err != nil {
		return nil, err
	}

	if rolePermission == nil || rolePermission.IsDeleted() {
		return nil, nil // No es error, simplemente no existe
	}

	return rolePermission, nil
}