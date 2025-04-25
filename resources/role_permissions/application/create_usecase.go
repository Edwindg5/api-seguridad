package application

import (
	"context"
	"errors"
	"api-seguridad/resources/role_permissions/domain/entities"
	"api-seguridad/resources/role_permissions/domain/repository"
)

type CreateRolePermissionUseCase struct {
	repo repository.RolePermissionRepository
}

func NewCreateRolePermissionUseCase(repo repository.RolePermissionRepository) *CreateRolePermissionUseCase {
	return &CreateRolePermissionUseCase{repo: repo}
}

func (uc *CreateRolePermissionUseCase) Execute(ctx context.Context, rolePermission *entities.RolePermission) error {
	// Validaciones básicas
	if rolePermission.RoleID == 0 {
		return errors.New("ID de rol es requerido")
	}
	if rolePermission.PermissionID == 0 {
		return errors.New("ID de permiso es requerido")
	}
	if rolePermission.CreatedBy == 0 {
		return errors.New("usuario creador es requerido")
	}

	// Verificar si ya existe la relación
	existing, err := uc.repo.GetByRoleAndPermission(ctx, rolePermission.RoleID, rolePermission.PermissionID)
	if err != nil {
		return err
	}
	if existing != nil && !existing.IsDeleted() {
		return errors.New("esta relación rol-permiso ya existe")
	}

	return uc.repo.Create(ctx, rolePermission)
}