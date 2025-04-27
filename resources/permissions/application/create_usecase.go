//api-seguridad/resources/permissions/application/create_usecase.go
package application

import (
	"context"
	"errors"
	"api-seguridad/resources/permissions/domain/entities"
	"time"
	"api-seguridad/resources/permissions/domain/repository"
)

type CreatePermissionUseCase struct {
	repo repository.PermissionRepository
}

func NewCreatePermissionUseCase(repo repository.PermissionRepository) *CreatePermissionUseCase {
	return &CreatePermissionUseCase{repo: repo}
}

func (uc *CreatePermissionUseCase) Execute(ctx context.Context, permission *entities.Permission) error {
	// Validaciones básicas
	if permission.Name == "" {
		return errors.New("el nombre del permiso es requerido")
	}
	if len(permission.Name) > 100 {
		return errors.New("el nombre no puede exceder los 100 caracteres")
	}
	if len(permission.Description) > 255 {
		return errors.New("la descripción no puede exceder los 255 caracteres")
	}

	// Validar usuario creador
	if permission.CreatedBy == 0 {
		return errors.New("se requiere un usuario creador válido")
	}

	// Verificar si el permiso ya existe
	existing, err := uc.repo.GetByID(ctx, permission.ID)
	if err != nil {
		return err
	}
	if existing != nil {
		return errors.New("ya existe un permiso con este ID")
	}

	// Establecer fechas de auditoría
	if permission.CreatedAt.IsZero() {
		permission.CreatedAt = time.Now()
	}
	permission.UpdatedAt = time.Now()

	return uc.repo.Create(ctx, permission)
}