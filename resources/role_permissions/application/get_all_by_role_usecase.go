package application

import (
	"context"
	"errors"
	"api-seguridad/resources/role_permissions/domain/entities"
	"api-seguridad/resources/role_permissions/domain/repository"
)

type GetAllByRoleUseCase struct {
	repo repository.RolePermissionRepository
}

func NewGetAllByRoleUseCase(repo repository.RolePermissionRepository) *GetAllByRoleUseCase {
	return &GetAllByRoleUseCase{repo: repo}
}

func (uc *GetAllByRoleUseCase) Execute(ctx context.Context, roleID uint) ([]*entities.RolePermission, error) {
	if roleID == 0 {
		return nil, errors.New("ID de rol es requerido")
	}

	rolePermissions, err := uc.repo.GetAllByRole(ctx, roleID)
	if err != nil {
		return nil, err
	}

	// Filtrar registros eliminados (doble verificación)
	var activeRolePermissions []*entities.RolePermission
	for _, rp := range rolePermissions {
		if !rp.IsDeleted() {
			activeRolePermissions = append(activeRolePermissions, rp)
		}
	}

	return activeRolePermissions, nil
}