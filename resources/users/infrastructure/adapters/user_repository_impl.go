// api-seguridad/resources/users/infrastructure/adapters/user_repository_impl.go
package adapters

import (
	"context"
	"errors"

	"api-seguridad/resources/users/domain/entities"
	"api-seguridad/resources/users/domain/repository"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) Create(ctx context.Context, user *entities.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *UserRepositoryImpl) GetByID(ctx context.Context, id uint) (*entities.User, error) {
	var user entities.User
	err := r.db.WithContext(ctx).
		Preload("Role").
		Preload("Creator").
		Preload("Updater").
		Where("id = ? AND deleted = ?", id, false).
		First(&user).Error
	
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, err
}

func (r *UserRepositoryImpl) GetByUsername(ctx context.Context, username string) (*entities.User, error) {
	var user entities.User
	err := r.db.WithContext(ctx).
		Where("username = ? AND deleted = ?", username, false).
		First(&user).Error
	
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, err
}

func (r *UserRepositoryImpl) GetByEmail(ctx context.Context, email string) (*entities.User, error) {
	var user entities.User
	err := r.db.WithContext(ctx).
		Where("email = ? AND deleted = ?", email, false).
		First(&user).Error
	
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, err
}

func (r *UserRepositoryImpl) Update(ctx context.Context, user *entities.User) error {
	return r.db.WithContext(ctx).Save(user).Error
}

func (r *UserRepositoryImpl) SoftDelete(ctx context.Context, id uint, deleterID uint) error {
	return r.db.WithContext(ctx).
		Model(&entities.User{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"deleted":     true,
			"updated_by":  deleterID,
			"updated_at":  gorm.Expr("CURRENT_TIMESTAMP"),
		}).Error
}

func (r *UserRepositoryImpl) List(ctx context.Context) ([]*entities.User, error) {
	var users []*entities.User
	err := r.db.WithContext(ctx).
		Preload("Role").
		Preload("Creator").
		Preload("Updater").
		Where("deleted = ?", false).
		Find(&users).Error
	return users, err
}

func (r *UserRepositoryImpl) ListActiveUsers(ctx context.Context) ([]*entities.User, error) {
	// Como List ya filtra por deleted = false, podemos simplemente llamar a List
	return r.List(ctx)
}
