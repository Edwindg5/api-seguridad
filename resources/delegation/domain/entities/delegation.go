// api-seguridad/resources/delegation/domain/entities/delegation.go
package entities

import (
	entityusers "api-seguridad/resources/users/domain/entities"
	"time"
)

type Delegation struct {
	ID               uint      `gorm:"primaryKey;column:id_delegation" json:"id"`
	Name             string    `gorm:"size:255;not null;column:name_delegation" json:"name"`
	Active           bool      `gorm:"column:active" json:"active"`
	CreatedAt        time.Time `gorm:"autoCreateTime;column:created_at" json:"created_at"`
	CreatedBy        uint      `gorm:"column:created_by" json:"created_by"`
	UpdatedAt        time.Time `gorm:"autoUpdateTime;column:updated_at" json:"updated_at"`
	UpdatedBy        uint      `gorm:"column:updated_by" json:"updated_by"`
	Deleted          bool      `gorm:"default:false;column:deleted" json:"deleted"`

	CreatedByUser    *entityusers.User `gorm:"foreignKey:CreatedBy" json:"created_by_user,omitempty"`
	UpdatedByUser    *entityusers.User `gorm:"foreignKey:UpdatedBy" json:"updated_by_user,omitempty"`
}

// Getters
func (d *Delegation) GetID() uint {
	return d.ID
}

func (d *Delegation) GetName() string {
	return d.Name
}

func (d *Delegation) IsActive() bool {
	return d.Active
}

func (d *Delegation) GetCreatedAt() time.Time {
	return d.CreatedAt
}

func (d *Delegation) GetCreatedBy() uint {
	return d.CreatedBy
}

func (d *Delegation) GetUpdatedAt() time.Time {
	return d.UpdatedAt
}

func (d *Delegation) GetUpdatedBy() uint {
	return d.UpdatedBy
}

func (d *Delegation) IsDeleted() bool {
	return d.Deleted
}

// Setters
func (d *Delegation) SetID(id uint) {
	d.ID = id
}

func (d *Delegation) SetName(name string) {
	d.Name = name
}

func (d *Delegation) SetActive(active bool) {
	d.Active = active
}

func (d *Delegation) SetCreatedAt(createdAt time.Time) {
	d.CreatedAt = createdAt
}

func (d *Delegation) SetCreatedBy(createdBy uint) {
	d.CreatedBy = createdBy
}

func (d *Delegation) SetUpdatedAt(updatedAt time.Time) {
	d.UpdatedAt = updatedAt
}

func (d *Delegation) SetUpdatedBy(updatedBy uint) {
	d.UpdatedBy = updatedBy
}

func (d *Delegation) SetDeleted(deleted bool) {
	d.Deleted = deleted
}