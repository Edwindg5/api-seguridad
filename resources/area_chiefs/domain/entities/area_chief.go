// api-seguridad/resources/area_chiefs/domain/entities/area_chief.go
package entities

import (
	"time"
	entityusers "api-seguridad/resources/users/domain/entities"
)

type AreaChief struct {
	ID           uint      `gorm:"primaryKey;column:id_official" json:"id"`
	Name         string    `gorm:"column:name_official;size:50" json:"name"`
	Position     string    `gorm:"size:50" json:"position"`
	Type         string    `gorm:"size:50" json:"type"`
	SignaturePath string   `gorm:"column:signature_path;size:100" json:"signature_path"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"created_at"`
	CreatedBy    uint      `gorm:"column:created_by" json:"created_by"`
	UpdatedBy    uint      `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"updated_at"`
	Deleted      bool      `gorm:"default:false;column:deleted" json:"deleted"`

	CreatedByUser *entityusers.User `gorm:"foreignKey:CreatedBy" json:"created_by_user,omitempty"`
	UpdatedByUser *entityusers.User `gorm:"foreignKey:UpdatedBy" json:"updated_by_user,omitempty"`
}

// Getters
func (a *AreaChief) GetID() uint {
	return a.ID
}

func (a *AreaChief) GetName() string {
	return a.Name
}

func (a *AreaChief) GetPosition() string {
	return a.Position
}

func (a *AreaChief) GetType() string {
	return a.Type
}

func (a *AreaChief) GetSignaturePath() string {
	return a.SignaturePath
}

func (a *AreaChief) GetCreatedAt() time.Time {
	return a.CreatedAt
}

func (a *AreaChief) GetCreatedBy() uint {
	return a.CreatedBy
}

func (a *AreaChief) GetUpdatedBy() uint {
	return a.UpdatedBy
}

func (a *AreaChief) GetUpdatedAt() time.Time {
	return a.UpdatedAt
}

func (a *AreaChief) IsDeleted() bool {
	return a.Deleted
}

// Setters
func (a *AreaChief) SetID(id uint) {
	a.ID = id
}

func (a *AreaChief) SetName(name string) {
	a.Name = name
}

func (a *AreaChief) SetPosition(position string) {
	a.Position = position
}

func (a *AreaChief) SetType(typeStr string) {
	a.Type = typeStr
}

func (a *AreaChief) SetSignaturePath(path string) {
	a.SignaturePath = path
}

func (a *AreaChief) SetCreatedAt(createdAt time.Time) {
	a.CreatedAt = createdAt
}

func (a *AreaChief) SetCreatedBy(createdBy uint) {
	a.CreatedBy = createdBy
}

func (a *AreaChief) SetUpdatedBy(updatedBy uint) {
	a.UpdatedBy = updatedBy
}

func (a *AreaChief) SetUpdatedAt(updatedAt time.Time) {
	a.UpdatedAt = updatedAt
}

func (a *AreaChief) SetDeleted(deleted bool) {
	a.Deleted = deleted
}