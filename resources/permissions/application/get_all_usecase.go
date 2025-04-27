//api-seguridad/resources/permissions/application/get_all_usecase.go
package application

import (
	"context"
	"api-seguridad/resources/permissions/domain/entities"
	"api-seguridad/resources/permissions/domain/repository"
)

type GetAllPermissionsUseCase struct {
	repo repository.PermissionRepository
}

func NewGetAllPermissionsUseCase(repo repository.PermissionRepository) *GetAllPermissionsUseCase {
	return &GetAllPermissionsUseCase{repo: repo}
}

func (uc *GetAllPermissionsUseCase) Execute(ctx context.Context) ([]*entities.Permission, error) {
	permissions, err := uc.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	// Filtrar registros eliminados (aunque el repositorio ya lo hace, es doble verificación)
	var activePermissions []*entities.Permission
	for _, p := range permissions {
		if !p.IsDeleted() {
			activePermissions = append(activePermissions, p)
		}
	}

	return activePermissions, nil
}