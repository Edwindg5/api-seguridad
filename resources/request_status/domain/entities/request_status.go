// api-seguridad/resources/request_status/domain/entities/request_status.go
package entities

import (
	"time"
	entityusers "api-seguridad/resources/users/domain/entities"
)

type RequestStatus struct {
    ID          uint      `gorm:"primaryKey;column:id_status" json:"id"`
    Name        string    `gorm:"size:100;unique;not null" json:"name"`
    Description string    `gorm:"type:text" json:"description"`
    CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
    CreatedBy   uint      `gorm:"column:created_by" json:"created_by"`
    UpdatedBy   uint      `gorm:"column:updated_by" json:"updated_by"`
    UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at"`
    Deleted     bool      `gorm:"default:false;column:deleted" json:"deleted"`

    CreatedByUser *entityusers.User `gorm:"foreignKey:CreatedBy;references:id_user" json:"created_by_user,omitempty"`
    UpdatedByUser *entityusers.User `gorm:"foreignKey:UpdatedBy;references:id_user" json:"updated_by_user,omitempty"`
}
// Getters
func (rs *RequestStatus) GetID() uint {
	return rs.ID
}

func (rs *RequestStatus) GetName() string {
	return rs.Name
}

func (rs *RequestStatus) GetDescription() string {
	return rs.Description
}

func (rs *RequestStatus) GetCreatedAt() time.Time {
	return rs.CreatedAt
}

func (rs *RequestStatus) GetCreatedBy() uint {
	return rs.CreatedBy
}

func (rs *RequestStatus) GetUpdatedBy() uint {
	return rs.UpdatedBy
}

func (rs *RequestStatus) GetUpdatedAt() time.Time {
	return rs.UpdatedAt
}

func (rs *RequestStatus) IsDeleted() bool {
	return rs.Deleted
}

// Setters
func (rs *RequestStatus) SetID(id uint) {
	rs.ID = id
}

func (rs *RequestStatus) SetName(name string) {
	rs.Name = name
}

func (rs *RequestStatus) SetDescription(description string) {
	rs.Description = description
}

func (rs *RequestStatus) SetCreatedAt(createdAt time.Time) {
	rs.CreatedAt = createdAt
}

func (rs *RequestStatus) SetCreatedBy(createdBy uint) {
	rs.CreatedBy = createdBy
}

func (rs *RequestStatus) SetUpdatedBy(updatedBy uint) {
	rs.UpdatedBy = updatedBy
}

func (rs *RequestStatus) SetUpdatedAt(updatedAt time.Time) {
	rs.UpdatedAt = updatedAt
}

func (rs *RequestStatus) SetDeleted(deleted bool) {
	rs.Deleted = deleted
}