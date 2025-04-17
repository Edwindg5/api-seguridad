package application

import (
	"context"
	"api-seguridad/resources/roles/domain/entities"
	"api-seguridad/resources/roles/domain/repository"
)

type RoleService struct {
	roleRepo repository.RoleRepository
}

func NewRoleService(roleRepo repository.RoleRepository) *RoleService {
	return &RoleService{roleRepo: roleRepo}
}

func (s *RoleService) CreateRole(ctx context.Context, role *entity.Role) error {
	postUC := NewPostRoleUseCase(s.roleRepo)
	return postUC.Execute(ctx, role)
}

func (s *RoleService) GetRoleByID(ctx context.Context, id uint) (*entity.Role, error) {
	return s.roleRepo.GetByID(ctx, id)
}

func (s *RoleService) UpdateRole(ctx context.Context, role *entity.Role) error {
	return s.roleRepo.Update(ctx, role)
}

func (s *RoleService) DeleteRole(ctx context.Context, id uint) error {
	return s.roleRepo.Delete(ctx, id)
}

func (s *RoleService) ListRoles(ctx context.Context) ([]*entity.Role, error) {
	return s.roleRepo.List(ctx)
}