package application

import (
	"context"
	"errors"
	"api-seguridad/resources/roles/domain/entities"
	"api-seguridad/resources/roles/domain/repository"
)

type GetRoleByIDUseCase struct {
	roleRepo repository.RoleRepository
}

func NewGetRoleByIDUseCase(roleRepo repository.RoleRepository) *GetRoleByIDUseCase {
	return &GetRoleByIDUseCase{roleRepo: roleRepo}
}

func (uc *GetRoleByIDUseCase) Execute(ctx context.Context, id uint) (*entities.Role, error) {
	if id == 0 {
		return nil, errors.New("invalid role ID")
	}

	role, err := uc.roleRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if role == nil || role.IsDeleted() {
		return nil, errors.New("role not found")
	}

	return role, nil
}