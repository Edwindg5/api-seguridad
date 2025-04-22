//api-seguridad/resources/type_police/domain/entities/type_police.go
package entities

import (
	"time"
	entityusers "api-seguridad/resources/users/domain/entities"
)

type TypePolice struct {
	ID             uint      `gorm:"primaryKey;column:id_type_police" json:"id"`
	TitleKindPolice string    `gorm:"size:50;not null;column:title_kind_police" json:"title_kind_police"`
	CreatedAt      time.Time `gorm:"autoCreateTime;column:created_at" json:"created_at"`
	CreatedBy      uint      `gorm:"column:created_by" json:"created_by"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime;column:updated_at" json:"updated_at"`
	UpdatedBy      uint      `gorm:"column:updated_by" json:"updated_by"`
	Deleted        bool      `gorm:"default:false;column:deleted" json:"deleted"`

	CreatedByUser *entityusers.User `gorm:"foreignKey:CreatedBy" json:"created_by_user,omitempty"`
	UpdatedByUser *entityusers.User `gorm:"foreignKey:UpdatedBy" json:"updated_by_user,omitempty"`
}

// Getters
func (t *TypePolice) GetID() uint {
	return t.ID
}

func (t *TypePolice) GetTitleKindPolice() string {
	return t.TitleKindPolice
}

func (t *TypePolice) GetCreatedAt() time.Time {
	return t.CreatedAt
}

func (t *TypePolice) GetCreatedBy() uint {
	return t.CreatedBy
}

func (t *TypePolice) GetUpdatedAt() time.Time {
	return t.UpdatedAt
}

func (t *TypePolice) GetUpdatedBy() uint {
	return t.UpdatedBy
}

func (t *TypePolice) IsDeleted() bool {
	return t.Deleted
}

// Setters
func (t *TypePolice) SetID(id uint) {
	t.ID = id
}

func (t *TypePolice) SetTitleKindPolice(title string) {
	t.TitleKindPolice = title
}

func (t *TypePolice) SetCreatedAt(createdAt time.Time) {
	t.CreatedAt = createdAt
}

func (t *TypePolice) SetCreatedBy(createdBy uint) {
	t.CreatedBy = createdBy
}

func (t *TypePolice) SetUpdatedAt(updatedAt time.Time) {
	t.UpdatedAt = updatedAt
}

func (t *TypePolice) SetUpdatedBy(updatedBy uint) {
	t.UpdatedBy = updatedBy
}

func (t *TypePolice) SetDeleted(deleted bool) {
	t.Deleted = deleted
}