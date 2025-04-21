// api-seguridad/resources/roles/domain/entities/role.go
package entities

import (
	"time"
)

type Role struct {
	ID          uint      `gorm:"primaryKey;column:id_rol" json:"id"`
	Title       string    `gorm:"size:50;not null;column:title_rol" json:"title"`
	Description string    `gorm:"size:50" json:"description"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	CreatedBy   uint      `json:"created_by"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	UpdatedBy   uint      `json:"updated_by"`
	Deleted     bool      `gorm:"default:false" json:"deleted"`
}

// Getters
func (r *Role) GetID() uint {
	return r.ID
}

func (r *Role) GetTitle() string {
	return r.Title
}

func (r *Role) GetDescription() string {
	return r.Description
}

func (r *Role) GetCreatedAt() time.Time {
	return r.CreatedAt
}

func (r *Role) GetCreatedBy() uint {
	return r.CreatedBy
}

func (r *Role) GetUpdatedAt() time.Time {
	return r.UpdatedAt
}

func (r *Role) GetUpdatedBy() uint {
	return r.UpdatedBy
}

func (r *Role) IsDeleted() bool {
	return r.Deleted
}

// Setters
func (r *Role) SetID(id uint) {
	r.ID = id
}

func (r *Role) SetTitle(title string) {
	r.Title = title
}

func (r *Role) SetDescription(description string) {
	r.Description = description
}

func (r *Role) SetCreatedAt(createdAt time.Time) {
	r.CreatedAt = createdAt
}

func (r *Role) SetCreatedBy(createdBy uint) {
	r.CreatedBy = createdBy
}

func (r *Role) SetUpdatedAt(updatedAt time.Time) {
	r.UpdatedAt = updatedAt
}

func (r *Role) SetUpdatedBy(updatedBy uint) {
	r.UpdatedBy = updatedBy
}

func (r *Role) SetDeleted(deleted bool) {
	r.Deleted = deleted
}