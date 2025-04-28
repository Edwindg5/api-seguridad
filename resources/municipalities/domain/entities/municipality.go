// api-seguridad/resources/municipalities/domain/entities/municipality.go
package entities

import (
	"time"
	entityusers "api-seguridad/resources/users/domain/entities"
	entitydelegations "api-seguridad/resources/delegation/domain/entities"
)

type Municipality struct {
	ID           uint      `gorm:"primaryKey;column:id_municipalities" json:"id"`
	Name         string    `gorm:"size:50;not null;column:name_municipalities" json:"name"`
	DelegationID uint      `gorm:"column:id_delegation_fk" json:"delegation_id"`
	CreatedAt    time.Time `gorm:"autoCreateTime;column:created_at" json:"created_at"`
	CreatedBy    uint      `gorm:"column:created_by" json:"created_by"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime;column:updated_at" json:"updated_at"`
	UpdatedBy    uint      `gorm:"column:updated_by" json:"updated_by"`
	Deleted      bool      `gorm:"default:false;column:deleted" json:"deleted"`

	// Relaciones
	CreatedByUser *entityusers.User          `gorm:"foreignKey:CreatedBy" json:"created_by_user,omitempty"`
	UpdatedByUser *entityusers.User          `gorm:"foreignKey:UpdatedBy" json:"updated_by_user,omitempty"`
	Delegation    *entitydelegations.Delegation `gorm:"foreignKey:DelegationID" json:"delegation,omitempty"`
}

// Getters
func (m *Municipality) GetID() uint {
	return m.ID
}

func (m *Municipality) GetName() string {
	return m.Name
}

func (m *Municipality) GetDelegationID() uint {
	return m.DelegationID
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

func (m *Municipality) SetDelegationID(delegationID uint) {
	m.DelegationID = delegationID
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