// api-seguridad/resources/users/infrastructure/adapters/user_repository_impl.go
package adapters

import (
	"context"
	"errors"
	"time"

	"api-seguridad/resources/users/domain/entities"
	"api-seguridad/resources/users/domain/repository"
	rolEntities "api-seguridad/resources/roles/domain/entities"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db         *gorm.DB
	jwtSecret  string
}

func NewUserRepository(db *gorm.DB, jwtSecret string) repository.UserRepository {
	return &UserRepositoryImpl{
		db:        db,
		jwtSecret: jwtSecret,
	}
}

func (r *UserRepositoryImpl) Create(ctx context.Context, user *entities.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	// Asignar valores por defecto si no están establecidos
	if user.CreatedAt.IsZero() {
		user.CreatedAt = time.Now()
	}
	if user.UpdatedAt.IsZero() {
		user.UpdatedAt = time.Now()
	}

	return r.db.WithContext(ctx).Create(user).Error
}
func (r *UserRepositoryImpl) GetByID(ctx context.Context, id uint) (*entities.User, error) {
	var user entities.User
	err := r.db.WithContext(ctx).
		Preload("Role").
		Where("id_user = ? AND deleted = ?", id, false).
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
    // Obtener usuario existente
    existingUser, err := r.GetByID(ctx, user.ID)
    if err != nil {
        return err
    }
    if existingUser == nil {
        return errors.New("user not found")
    }

    // Verificar que el rol exista en la base de datos
    var roleExists bool
    err = r.db.WithContext(ctx).Model(&rolEntities.Role{}).
        Select("count(*) > 0").
        Where("id_rol = ? AND deleted = ?", user.RoleID, false).
        Find(&roleExists).Error
    
    if err != nil {
        return err
    }
    if !roleExists {
        return errors.New("role not found")
    }

    // Actualizar solo campos permitidos
    updates := map[string]interface{}{
        "first_name": user.FirstName,
        "lastname":   user.LastName,
        "username":   user.Username,
        "email":      user.Email,
        "rol_id_fk":  user.RoleID,
        "updated_at": user.UpdatedAt,
        "updated_by": user.UpdatedBy,
        "deleted":    user.Deleted,
    }

    return r.db.WithContext(ctx).Model(&entities.User{}).
        Where("id_user = ?", user.ID).
        Updates(updates).Error
}
func (r *UserRepositoryImpl) SoftDelete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).
		Model(&entities.User{}).
		Where("id_user = ?", id).
		Updates(map[string]interface{}{
			"deleted":    true,
			"updated_at": time.Now(),
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
	return r.List(ctx)
}

func (r *UserRepositoryImpl) Login(ctx context.Context, username string, password string) (*entities.User, string, error) {
	// Buscar usuario por username (activo)
	var user entities.User
	err := r.db.WithContext(ctx).
		Preload("Role").
		Where("username = ? AND deleted = ?", username, false).
		First(&user).Error
	
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, "", errors.New("invalid credentials")
		}
		return nil, "", errors.New("authentication error")
	}

	// Verificar contraseña
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, "", errors.New("invalid credentials")
	}

	// Generar token JWT
	token, err := r.generateJWTToken(&user)
	if err != nil {
		return nil, "", errors.New("failed to generate token")
	}

	return &user, token, nil
}

func (r *UserRepositoryImpl) generateJWTToken(user *entities.User) (string, error) {
	claims := jwt.MapClaims{
		"sub":       user.ID,                   // Subject (user ID)
		"username":  user.Username,             // Username
		"role_id":   user.RoleID,               // Role ID
		"exp":       time.Now().Add(24 * time.Hour).Unix(), // Expira en 24 horas
		"iat":       time.Now().Unix(),         // Issued at
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(r.jwtSecret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}