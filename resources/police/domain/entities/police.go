//api-seguridad/resources/police/domain/entities/police.go
package entity

import (
	"time"
)

type Police struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	Name          string    `gorm:"size:50;not null" json:"name"`
	PaternalName  string    `gorm:"size:50;not null" json:"paternal_lastname"`
	MaternalName  string    `gorm:"size:50" json:"maternal_lastname"`
	TypePoliceID  uint      `gorm:"not null" json:"type_police_id"`
	Sex           string    `gorm:"size:1;not null" json:"sex"`
	CUIP          string    `gorm:"size:50;unique;not null" json:"cuip"`
	RFC           string    `gorm:"size:20" json:"rfc"`
	CreatedAt     time.Time `gorm:"autoCreateTime" json:"created_at"`
	CreatedBy     uint      `json:"created_by"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	UpdatedBy     uint      `json:"updated_by"`
	Deleted       bool      `gorm:"default:false" json:"deleted"`

	TypePolice *TypePolice `gorm:"foreignKey:TypePoliceID" json:"type_police,omitempty"`
}