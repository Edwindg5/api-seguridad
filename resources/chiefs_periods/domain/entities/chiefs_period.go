package entities

import (
	"time"
	entityareachiefs "api-seguridad/resources/area_chiefs/domain/entities"
	entityusers "api-seguridad/resources/users/domain/entities"
)

type ChiefsPeriod struct {
	ID           uint      `gorm:"primaryKey;column:id" json:"id"`
	CeoChiefID   uint      `gorm:"column:ceo_chief_id;not null" json:"ceo_chief_id"`
	LegalChiefID uint      `gorm:"column:legal_chief_id;not null" json:"legal_chief_id"`
	StartDate    time.Time `gorm:"column:start_date;not null" json:"start_date"`
	EndDate      time.Time `gorm:"column:end_date" json:"end_date"`
	PeriodActive bool      `gorm:"column:period_active;default:false" json:"period_active"`
	CreatedAt    time.Time `gorm:"autoCreateTime;column:created_at" json:"created_at"`
	CreatedBy    uint      `gorm:"column:created_by" json:"created_by"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime;column:updated_at" json:"updated_at"`
	UpdatedBy    uint      `gorm:"column:updated_by" json:"updated_by"`
	Deleted      bool      `gorm:"column:deleted;default:false" json:"deleted"`

	// Relaciones
	CeoChief   *entityareachiefs.AreaChief `gorm:"foreignKey:CeoChiefID;references:ID" json:"ceo_chief,omitempty"`
	LegalChief *entityareachiefs.AreaChief `gorm:"foreignKey:LegalChiefID;references:ID" json:"legal_chief,omitempty"`
	Creator    *entityusers.User           `gorm:"foreignKey:CreatedBy;references:ID" json:"creator,omitempty"`
	Updater    *entityusers.User           `gorm:"foreignKey:UpdatedBy;references:ID" json:"updater,omitempty"`
}

// Getters
func (cp *ChiefsPeriod) GetID() uint {
	return cp.ID
}

func (cp *ChiefsPeriod) GetCeoChiefID() uint {
	return cp.CeoChiefID
}

func (cp *ChiefsPeriod) GetLegalChiefID() uint {
	return cp.LegalChiefID
}

func (cp *ChiefsPeriod) GetStartDate() time.Time {
	return cp.StartDate
}

func (cp *ChiefsPeriod) GetEndDate() time.Time {
	return cp.EndDate
}

func (cp *ChiefsPeriod) IsPeriodActive() bool {
	return cp.PeriodActive
}

func (cp *ChiefsPeriod) GetCreatedAt() time.Time {
	return cp.CreatedAt
}

func (cp *ChiefsPeriod) GetCreatedBy() uint {
	return cp.CreatedBy
}

func (cp *ChiefsPeriod) GetUpdatedAt() time.Time {
	return cp.UpdatedAt
}

func (cp *ChiefsPeriod) GetUpdatedBy() uint {
	return cp.UpdatedBy
}

func (cp *ChiefsPeriod) IsDeleted() bool {
	return cp.Deleted
}

// Setters
func (cp *ChiefsPeriod) SetID(id uint) {
	cp.ID = id
}

func (cp *ChiefsPeriod) SetCeoChiefID(id uint) {
	cp.CeoChiefID = id
}

func (cp *ChiefsPeriod) SetLegalChiefID(id uint) {
	cp.LegalChiefID = id
}

func (cp *ChiefsPeriod) SetStartDate(date time.Time) {
	cp.StartDate = date
}

func (cp *ChiefsPeriod) SetEndDate(date time.Time) {
	cp.EndDate = date
}

func (cp *ChiefsPeriod) SetPeriodActive(active bool) {
	cp.PeriodActive = active
}

func (cp *ChiefsPeriod) SetCreatedAt(createdAt time.Time) {
	cp.CreatedAt = createdAt
}

func (cp *ChiefsPeriod) SetCreatedBy(createdBy uint) {
	cp.CreatedBy = createdBy
}

func (cp *ChiefsPeriod) SetUpdatedAt(updatedAt time.Time) {
	cp.UpdatedAt = updatedAt
}

func (cp *ChiefsPeriod) SetUpdatedBy(updatedBy uint) {
	cp.UpdatedBy = updatedBy
}

func (cp *ChiefsPeriod) SetDeleted(deleted bool) {
	cp.Deleted = deleted
}