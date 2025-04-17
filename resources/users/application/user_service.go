package application

import (
	"context"
	"api-seguridad/resources/users/domain/entities"
	"api-seguridad/resources/users/domain/repository"
)

type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) CreateUser(ctx context.Context, user *entity.User) error {
	postUC := NewPostUserUseCase(s.userRepo)
	return postUC.Execute(ctx, user)
}

func (s *UserService) GetUserByID(ctx context.Context, id uint) (*entity.User, error) {
	return s.userRepo.GetByID(ctx, id)
}

func (s *UserService) GetUserByUsername(ctx context.Context, username string) (*entity.User, error) {
	return s.userRepo.GetByUsername(ctx, username)
}

func (s *UserService) UpdateUser(ctx context.Context, user *entity.User) error {
	return s.userRepo.Update(ctx, user)
}

func (s *UserService) DeleteUser(ctx context.Context, id uint) error {
	return s.userRepo.Delete(ctx, id)
}

func (s *UserService) ListUsers(ctx context.Context) ([]*entity.User, error) {
	return s.userRepo.List(ctx)
}