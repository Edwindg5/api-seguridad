package entities

import (
	"time"
	entityusers "api-seguridad/resources/users/domain/entities"
)

type Permission struct {
	ID          uint      `gorm:"primaryKey;column:id_permission" json:"id"`
	Name        string    `gorm:"size:100;not null" json:"name"`
	Description string    `gorm:"size:255" json:"description"`
	CreatedAt   time.Time `gorm:"autoCreateTime;column:created_at" json:"created_at"`
	CreatedBy   uint      `gorm:"column:created_by" json:"created_by"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime;column:updated_at" json:"updated_at"`
	UpdatedBy   uint      `gorm:"column:updated_by" json:"updated_by"`
	Deleted     bool      `gorm:"default:false;column:deleted" json:"deleted"`

	CreatedByUser *entityusers.User `gorm:"foreignKey:CreatedBy;references:ID" json:"created_by_user,omitempty"`
	UpdatedByUser *entityusers.User `gorm:"foreignKey:UpdatedBy;references:ID" json:"updated_by_user,omitempty"`
}

// Getters
func (p *Permission) GetID() uint {
	return p.ID
}

func (p *Permission) GetName() string {
	return p.Name
}

func (p *Permission) GetDescription() string {
	return p.Description
}

func (p *Permission) GetCreatedAt() time.Time {
	return p.CreatedAt
}

func (p *Permission) GetCreatedBy() uint {
	return p.CreatedBy
}

func (p *Permission) GetUpdatedAt() time.Time {
	return p.UpdatedAt
}

func (p *Permission) GetUpdatedBy() uint {
	return p.UpdatedBy
}

func (p *Permission) IsDeleted() bool {
	return p.Deleted
}

// Setters
func (p *Permission) SetID(id uint) {
	p.ID = id
}

func (p *Permission) SetName(name string) {
	p.Name = name
}

func (p *Permission) SetDescription(description string) {
	p.Description = description
}

func (p *Permission) SetCreatedAt(createdAt time.Time) {
	p.CreatedAt = createdAt
}

func (p *Permission) SetCreatedBy(createdBy uint) {
	p.CreatedBy = createdBy
}

func (p *Permission) SetUpdatedAt(updatedAt time.Time) {
	p.UpdatedAt = updatedAt
}

func (p *Permission) SetUpdatedBy(updatedBy uint) {
	p.UpdatedBy = updatedBy
}

func (p *Permission) SetDeleted(deleted bool) {
	p.Deleted = deleted
}