// api-seguridad/resources/police/domain/entities/police.go
package entities

import (
	entityusers "api-seguridad/resources/users/domain/entities"
	entitytypepolice "api-seguridad/resources/type_police/domain/entities"
	"time"
)

type Police struct {
	ID            uint      `gorm:"primaryKey;column:id_police" json:"id"`
	Name          string    `gorm:"column:name_police;size:50;not null" json:"name"`
	PaternalName  string    `gorm:"column:paternal_lastname;size:50;not null" json:"paternal_lastname"`
	MaternalName  string    `gorm:"column:maternal_lastname;size:50" json:"maternal_lastname"`
	TypePoliceID  uint      `gorm:"column:id_type_police_fk;not null" json:"type_police_id"`
	Sex           string    `gorm:"size:1;not null" json:"sex"`
    CUIP string `gorm:"column:c_ui_p;size:50;unique;not null" json:"cuip"`
	RFC           string    `gorm:"size:20" json:"rfc"`
	CreatedAt     time.Time `gorm:"autoCreateTime;column:created_at" json:"created_at"`
	CreatedBy     uint      `gorm:"column:created_by" json:"created_by"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime;column:updated_at" json:"updated_at"`
	UpdatedBy     uint      `gorm:"column:updated_by" json:"updated_by"`
	Deleted       bool      `gorm:"default:false;column:deleted" json:"deleted"`

	TypePolice    *entitytypepolice.TypePolice `gorm:"foreignKey:TypePoliceID" json:"type_police,omitempty"`
	CreatedByUser *entityusers.User            `gorm:"foreignKey:CreatedBy" json:"created_by_user,omitempty"`
	UpdatedByUser *entityusers.User            `gorm:"foreignKey:UpdatedBy" json:"updated_by_user,omitempty"`
}
// Getters
func (p *Police) GetID() uint {
	return p.ID
}

func (p *Police) GetName() string {
	return p.Name
}

func (p *Police) GetPaternalName() string {
	return p.PaternalName
}

func (p *Police) GetMaternalName() string {
	return p.MaternalName
}

func (p *Police) GetTypePoliceID() uint {
	return p.TypePoliceID
}

func (p *Police) GetSex() string {
	return p.Sex
}

func (p *Police) GetCUIP() string {
	return p.CUIP
}

func (p *Police) GetRFC() string {
	return p.RFC
}

func (p *Police) GetCreatedAt() time.Time {
	return p.CreatedAt
}

func (p *Police) GetCreatedBy() uint {
	return p.CreatedBy
}

func (p *Police) GetUpdatedAt() time.Time {
	return p.UpdatedAt
}

func (p *Police) GetUpdatedBy() uint {
	return p.UpdatedBy
}

func (p *Police) IsDeleted() bool {
	return p.Deleted
}

// Setters
func (p *Police) SetID(id uint) {
	p.ID = id
}

func (p *Police) SetName(name string) {
	p.Name = name
}

func (p *Police) SetPaternalName(name string) {
	p.PaternalName = name
}

func (p *Police) SetMaternalName(name string) {
	p.MaternalName = name
}

func (p *Police) SetTypePoliceID(id uint) {
	p.TypePoliceID = id
}

func (p *Police) SetSex(sex string) {
	p.Sex = sex
}

func (p *Police) SetCUIP(cuip string) {
	p.CUIP = cuip
}

func (p *Police) SetRFC(rfc string) {
	p.RFC = rfc
}

func (p *Police) SetCreatedAt(createdAt time.Time) {
	p.CreatedAt = createdAt
}

func (p *Police) SetCreatedBy(createdBy uint) {
	p.CreatedBy = createdBy
}

func (p *Police) SetUpdatedAt(updatedAt time.Time) {
	p.UpdatedAt = updatedAt
}

func (p *Police) SetUpdatedBy(updatedBy uint) {
	p.UpdatedBy = updatedBy
}

func (p *Police) SetDeleted(deleted bool) {
	p.Deleted = deleted
}