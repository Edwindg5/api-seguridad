package application

import (
	"context"
	"errors"
	"api-seguridad/resources/permissions/domain/entities"
	"api-seguridad/resources/permissions/domain/repository"
)

type GetPermissionByIDUseCase struct {
	repo repository.PermissionRepository
}

func NewGetPermissionByIDUseCase(repo repository.PermissionRepository) *GetPermissionByIDUseCase {
	return &GetPermissionByIDUseCase{repo: repo}
}

func (uc *GetPermissionByIDUseCase) Execute(ctx context.Context, id uint) (*entities.Permission, error) {
	// Validación de ID
	if id == 0 {
		return nil, errors.New("ID de permiso inválido")
	}

	// Obtener permiso del repositorio
	permission, err := uc.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Verificar si existe y no está eliminado
	if permission == nil || permission.IsDeleted() {
		return nil, errors.New("permiso no encontrado")
	}

	return permission, nil
}