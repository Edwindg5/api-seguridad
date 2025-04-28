// api-seguridad/resources/delegation/infrastructure/adapters/delegation_repository_impl.go
package adapters

import (
	"context"
	"errors"

	"api-seguridad/resources/delegation/domain/entities"
	"api-seguridad/resources/delegation/domain/repository"

	"gorm.io/gorm"
)

type DelegationRepositoryImpl struct {
	db *gorm.DB
}

func NewDelegationRepository(db *gorm.DB) repository.DelegationRepository {
	return &DelegationRepositoryImpl{db: db}
}

func (r *DelegationRepositoryImpl) Create(ctx context.Context, delegation *entities.Delegation) error {
    // Asegurar que los campos requeridos estén establecidos
    if delegation.GetCreatedBy() == 0 {
        return errors.New("created_by is required")
    }
    
    return r.db.WithContext(ctx).Create(delegation).Error
}
func (r *DelegationRepositoryImpl) GetByID(ctx context.Context, id uint) (*entities.Delegation, error) {
	var delegation entities.Delegation
	err := r.db.WithContext(ctx).
		Where("id_delegation = ? AND deleted = ?", id, false).
		First(&delegation).Error
		
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &delegation, err
}

func (r *DelegationRepositoryImpl) GetAll(ctx context.Context) ([]*entities.Delegation, error) {
	var delegations []*entities.Delegation
	err := r.db.WithContext(ctx).
		Where("deleted = ?", false).
		Find(&delegations).Error
	return delegations, err
}

func (r *DelegationRepositoryImpl) Update(ctx context.Context, delegation *entities.Delegation) error {
	return r.db.WithContext(ctx).Save(delegation).Error
}



func (r *DelegationRepositoryImpl) SoftDelete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).
		Model(&entities.Delegation{}).
		Where("id_delegation = ?", id).
		Update("deleted", true).Error
}