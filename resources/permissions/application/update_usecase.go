package application

import (
	"context"
	"errors"
	"api-seguridad/resources/permissions/domain/entities"
	"api-seguridad/resources/permissions/domain/repository"
)

type UpdatePermissionUseCase struct {
	repo repository.PermissionRepository
}

func NewUpdatePermissionUseCase(repo repository.PermissionRepository) *UpdatePermissionUseCase {
	return &UpdatePermissionUseCase{repo: repo}
}

func (uc *UpdatePermissionUseCase) Execute(ctx context.Context, permission *entities.Permission) error {
	if permission.ID == 0 {
		return errors.New("invalid permission ID")
	}
	if permission.Name == "" {
		return errors.New("name is required")
	}

	existing, err := uc.repo.GetByID(ctx, permission.ID)
	if err != nil {
		return err
	}
	if existing == nil || existing.IsDeleted() {
		return errors.New("permission not found")
	}

	return uc.repo.Update(ctx, permission)
}