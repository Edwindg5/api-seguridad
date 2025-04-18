// api-seguridad/resources/roles/application/post_usecase.go
package application

import (
	"context"
	"errors"
	"api-seguridad/resources/roles/domain/entities"
	"api-seguridad/resources/roles/domain/repository"
)

type PostRoleUseCase struct {
	roleRepo repository.RoleRepository
}

func NewPostRoleUseCase(roleRepo repository.RoleRepository) *PostRoleUseCase {
	return &PostRoleUseCase{roleRepo: roleRepo}
}

func (uc *PostRoleUseCase) Execute(ctx context.Context, role *entity.Role) error {
	if role.Title == "" {
		return errors.New("role title is required")
	}

	existingRole, err := uc.roleRepo.GetByTitle(ctx, role.Title)
	if err != nil {
		return err
	}
	if existingRole != nil {
		return errors.New("role with this title already exists")
	}

	return uc.roleRepo.Create(ctx, role)
}