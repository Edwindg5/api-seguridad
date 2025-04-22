//api-seguridad/resources/police/infrastructure/adapters/police_repository_impl.go
package adapters

import (
	"context"
	"errors"
	"api-seguridad/resources/police/domain/entities"
	"api-seguridad/resources/police/domain/repository"
	"gorm.io/gorm"
)

type PoliceRepositoryImpl struct {
	db *gorm.DB
}

func NewPoliceRepository(db *gorm.DB) repository.PoliceRepository {
	return &PoliceRepositoryImpl{db: db}
}

func (r *PoliceRepositoryImpl) Create(ctx context.Context, police *entities.Police) error {
	return r.db.WithContext(ctx).Create(police).Error
}

func (r *PoliceRepositoryImpl) GetByID(ctx context.Context, id uint) (*entities.Police, error) {
	var police entities.Police
	err := r.db.WithContext(ctx).
		Where("id_police = ? AND deleted = ?", id, false).
		Preload("TypePolice").
		Preload("CreatedByUser").
		Preload("UpdatedByUser").
		First(&police).Error
		
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &police, err
}

func (r *PoliceRepositoryImpl) GetAll(ctx context.Context) ([]*entities.Police, error) {
	var policeList []*entities.Police
	err := r.db.WithContext(ctx).
		Where("deleted = ?", false).
		Preload("TypePolice").
		Find(&policeList).Error
	return policeList, err
}

func (r *PoliceRepositoryImpl) GetByCUIP(ctx context.Context, cuip string) (*entities.Police, error) {
	var police entities.Police
	err := r.db.WithContext(ctx).
		Where("cuip = ? AND deleted = ?", cuip, false).
		Preload("TypePolice").
		First(&police).Error
		
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &police, err
}

func (r *PoliceRepositoryImpl) GetByRFC(ctx context.Context, rfc string) (*entities.Police, error) {
	var police entities.Police
	err := r.db.WithContext(ctx).
		Where("rfc = ? AND deleted = ?", rfc, false).
		Preload("TypePolice").
		First(&police).Error
		
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &police, err
}

func (r *PoliceRepositoryImpl) SearchByName(ctx context.Context, name string) ([]*entities.Police, error) {
	var policeList []*entities.Police
	err := r.db.WithContext(ctx).
		Where("(name_police LIKE ? OR paternal_lastname LIKE ? OR maternal_lastname LIKE ?) AND deleted = ?", 
			"%"+name+"%", "%"+name+"%", "%"+name+"%", false).
		Preload("TypePolice").
		Find(&policeList).Error
	return policeList, err
}

func (r *PoliceRepositoryImpl) SearchByFullName(ctx context.Context, name, paternalName, maternalName string) ([]*entities.Police, error) {
	var policeList []*entities.Police
	
	query := r.db.WithContext(ctx).
		Where("deleted = ?", false).
		Preload("TypePolice")

	if name != "" {
		query = query.Where("name_police LIKE ?", "%"+name+"%")
	}
	if paternalName != "" {
		query = query.Where("paternal_lastname LIKE ?", "%"+paternalName+"%")
	}
	if maternalName != "" {
		query = query.Where("maternal_lastname LIKE ?", "%"+maternalName+"%")
	}

	err := query.Find(&policeList).Error
	return policeList, err
}

func (r *PoliceRepositoryImpl) Update(ctx context.Context, police *entities.Police) error {
	return r.db.WithContext(ctx).Save(police).Error
}

func (r *PoliceRepositoryImpl) SoftDelete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).
		Model(&entities.Police{}).
		Where("id_police = ?", id).
		Update("deleted", true).Error
}