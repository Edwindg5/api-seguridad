//api-seguridad/resources/permissions/application/soft_delete_usecase.go
package application

import (
	"context"
	"errors"
	"api-seguridad/resources/permissions/domain/repository"
)

type SoftDeletePermissionUseCase struct {
	repo repository.PermissionRepository
}

func NewSoftDeletePermissionUseCase(repo repository.PermissionRepository) *SoftDeletePermissionUseCase {
	return &SoftDeletePermissionUseCase{repo: repo}
}

func (uc *SoftDeletePermissionUseCase) Execute(ctx context.Context, id uint) error {
	if id == 0 {
		return errors.New("invalid permission ID")
	}

	existing, err := uc.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if existing == nil || existing.IsDeleted() {
		return errors.New("permission not found")
	}

	return uc.repo.SoftDelete(ctx, id)
}