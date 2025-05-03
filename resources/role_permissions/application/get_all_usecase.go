//api-seguridad/resources/role_permissions/application/get_all_usecase.go
package application

import (
	"context"
	"api-seguridad/resources/role_permissions/domain/entities"
	"api-seguridad/resources/role_permissions/domain/repository"
)

type GetAllUseCase struct {
	repo repository.RolePermissionRepository
}

func NewGetAllUseCase(repo repository.RolePermissionRepository) *GetAllUseCase {
	return &GetAllUseCase{repo: repo}
}

func (uc *GetAllUseCase) Execute(ctx context.Context) ([]*entities.RolePermission, error) {
	// Obtener todos los role permissions del repositorio
	rolePermissions, err := uc.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	// Filtrar registros eliminados (soft delete)
	var activeRolePermissions []*entities.RolePermission
	for _, rp := range rolePermissions {
		if !rp.IsDeleted() {
			activeRolePermissions = append(activeRolePermissions, rp)
		}
	}

	return activeRolePermissions, nil
}