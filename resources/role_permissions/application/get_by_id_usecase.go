//api-seguridad/resources/role_permissions/application/get_by_id_usecase.go
package application

import (
	"context"
	"errors"
	"api-seguridad/resources/role_permissions/domain/entities"
	"api-seguridad/resources/role_permissions/domain/repository"
)

type GetRolePermissionByIDUseCase struct {
	repo repository.RolePermissionRepository
}

func NewGetRolePermissionByIDUseCase(repo repository.RolePermissionRepository) *GetRolePermissionByIDUseCase {
	return &GetRolePermissionByIDUseCase{repo: repo}
}

func (uc *GetRolePermissionByIDUseCase) Execute(ctx context.Context, id uint) (*entities.RolePermission, error) {
	if id == 0 {
		return nil, errors.New("ID inválido")
	}

	rolePermission, err := uc.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if rolePermission == nil || rolePermission.IsDeleted() {
		return nil, errors.New("relación rol-permiso no encontrada")
	}

	return rolePermission, nil
}