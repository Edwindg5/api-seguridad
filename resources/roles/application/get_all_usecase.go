package application

import (
	"context"
	"api-seguridad/resources/roles/domain/entities"
	"api-seguridad/resources/roles/domain/repository"
)

type GetAllRolesUseCase struct {
	roleRepo repository.RoleRepository
}

func NewGetAllRolesUseCase(roleRepo repository.RoleRepository) *GetAllRolesUseCase {
	return &GetAllRolesUseCase{roleRepo: roleRepo}
}

func (uc *GetAllRolesUseCase) Execute(ctx context.Context) ([]*entities.Role, error) {
	roles, err := uc.roleRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	// Filter out deleted roles if not already filtered in repository
	var activeRoles []*entities.Role
	for _, role := range roles {
		if !role.IsDeleted() {
			activeRoles = append(activeRoles, role)
		}
	}

	return activeRoles, nil
}