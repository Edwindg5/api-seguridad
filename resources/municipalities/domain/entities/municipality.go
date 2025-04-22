// api-seguridad/resources/municipalities/domain/entities/municipality.go
package entities

import (
	"time"
)

type Municipality struct {
	ID        uint      `gorm:"primaryKey;column:id_municipalities" json:"id"`
	Name      string    `gorm:"size:50;not null;column:name_municipalities" json:"name"`
	CreatedAt time.Time `gorm:"autoCreateTime;column:created_at" json:"created_at"`
	CreatedBy uint      `gorm:"column:created_by" json:"created_by"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;column:updated_at" json:"updated_at"`
	UpdatedBy uint      `gorm:"column:updated_by" json:"updated_by"`
	Deleted   bool      `gorm:"default:false;column:deleted" json:"deleted"`
}

// Getters
func (m *Municipality) GetID() uint {
	return m.ID
}

func (m *Municipality) GetName() string {
	return m.Name
}

func (m *Municipality) GetCreatedAt() time.Time {
	return m.CreatedAt
}

func (m *Municipality) GetCreatedBy() uint {
	return m.CreatedBy
}

func (m *Municipality) GetUpdatedAt() time.Time {
	return m.UpdatedAt
}

func (m *Municipality) GetUpdatedBy() uint {
	return m.UpdatedBy
}

func (m *Municipality) IsDeleted() bool {
	return m.Deleted
}

// Setters
func (m *Municipality) SetID(id uint) {
	m.ID = id
}

func (m *Municipality) SetName(name string) {
	m.Name = name
}

func (m *Municipality) SetCreatedAt(createdAt time.Time) {
	m.CreatedAt = createdAt
}

func (m *Municipality) SetCreatedBy(createdBy uint) {
	m.CreatedBy = createdBy
}

func (m *Municipality) SetUpdatedAt(updatedAt time.Time) {
	m.UpdatedAt = updatedAt
}

func (m *Municipality) SetUpdatedBy(updatedBy uint) {
	m.UpdatedBy = updatedBy
}

func (m *Municipality) SetDeleted(deleted bool) {
	m.Deleted = deleted
}