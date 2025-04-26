package entities

import (
	"time"
	entitypolice "api-seguridad/resources/police/domain/entities"
	entityusers "api-seguridad/resources/users/domain/entities"
)

type RequestDetail struct {
	ID                uint      `gorm:"primaryKey;column:id" json:"id"`
	RequestID         uint      `gorm:"column:request_id;not null" json:"request_id"`
	PoliceID          uint      `gorm:"column:id_police_fk;not null" json:"police_id"`
	Active            bool      `gorm:"default:false" json:"active"`
	Census            bool      `gorm:"default:false" json:"census"`
	Located           bool      `gorm:"default:false" json:"located"`
	Register          bool      `gorm:"default:false" json:"register"`
	Approved          bool      `gorm:"default:false" json:"approved"`
	Comments          string    `gorm:"type:TEXT" json:"comments"`
	MunicipalityActive bool     `gorm:"default:false" json:"municipality_active"`
	CreatedAt         time.Time `gorm:"autoCreateTime;column:created_at" json:"created_at"`
	CreatedBy         uint      `gorm:"column:created_by" json:"created_by"`
	UpdatedAt         time.Time `gorm:"autoUpdateTime;column:updated_at" json:"updated_at"`
	UpdatedBy         uint      `gorm:"column:updated_by" json:"updated_by"`
	Deleted           bool      `gorm:"default:false;column:deleted" json:"deleted"`

	// Relaciones
	Police    *entitypolice.Police `gorm:"foreignKey:PoliceID;references:ID" json:"police,omitempty"`
	Creator   *entityusers.User    `gorm:"foreignKey:CreatedBy;references:ID" json:"creator,omitempty"`
	Updater   *entityusers.User    `gorm:"foreignKey:UpdatedBy;references:ID" json:"updater,omitempty"`
}

// Getters
func (rd *RequestDetail) GetID() uint {
	return rd.ID
}

func (rd *RequestDetail) GetRequestID() uint {
	return rd.RequestID
}

func (rd *RequestDetail) GetPoliceID() uint {
	return rd.PoliceID
}

func (rd *RequestDetail) IsActive() bool {
	return rd.Active
}

func (rd *RequestDetail) IsCensus() bool {
	return rd.Census
}

func (rd *RequestDetail) IsLocated() bool {
	return rd.Located
}

func (rd *RequestDetail) IsRegister() bool {
	return rd.Register
}

func (rd *RequestDetail) IsApproved() bool {
	return rd.Approved
}

func (rd *RequestDetail) GetComments() string {
	return rd.Comments
}

func (rd *RequestDetail) IsMunicipalityActive() bool {
	return rd.MunicipalityActive
}

func (rd *RequestDetail) GetCreatedAt() time.Time {
	return rd.CreatedAt
}

func (rd *RequestDetail) GetCreatedBy() uint {
	return rd.CreatedBy
}

func (rd *RequestDetail) GetUpdatedAt() time.Time {
	return rd.UpdatedAt
}

func (rd *RequestDetail) GetUpdatedBy() uint {
	return rd.UpdatedBy
}

func (rd *RequestDetail) IsDeleted() bool {
	return rd.Deleted
}

// Setters
func (rd *RequestDetail) SetID(id uint) {
	rd.ID = id
}

func (rd *RequestDetail) SetRequestID(requestID uint) {
	rd.RequestID = requestID
}

func (rd *RequestDetail) SetPoliceID(policeID uint) {
	rd.PoliceID = policeID
}

func (rd *RequestDetail) SetActive(active bool) {
	rd.Active = active
}

func (rd *RequestDetail) SetCensus(census bool) {
	rd.Census = census
}

func (rd *RequestDetail) SetLocated(located bool) {
	rd.Located = located
}

func (rd *RequestDetail) SetRegister(register bool) {
	rd.Register = register
}

func (rd *RequestDetail) SetApproved(approved bool) {
	rd.Approved = approved
}

func (rd *RequestDetail) SetComments(comments string) {
	rd.Comments = comments
}

func (rd *RequestDetail) SetMunicipalityActive(active bool) {
	rd.MunicipalityActive = active
}

func (rd *RequestDetail) SetCreatedAt(createdAt time.Time) {
	rd.CreatedAt = createdAt
}

func (rd *RequestDetail) SetCreatedBy(createdBy uint) {
	rd.CreatedBy = createdBy
}

func (rd *RequestDetail) SetUpdatedAt(updatedAt time.Time) {
	rd.UpdatedAt = updatedAt
}

func (rd *RequestDetail) SetUpdatedBy(updatedBy uint) {
	rd.UpdatedBy = updatedBy
}

func (rd *RequestDetail) SetDeleted(deleted bool) {
	rd.Deleted = deleted
}