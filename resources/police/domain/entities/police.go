// api-seguridad/resources/police/domain/entities/police.go
package entities

import (
	entity "api-seguridad/resources/type_police/domain/entities"
	"time"
)

type Police struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Name         string    `gorm:"column:name_police;size:50;not null" json:"name"`
	PaternalName string    `gorm:"column:paternal_lastname;size:50;not null" json:"paternal_lastname"`
	MaternalName string    `gorm:"column:maternal_lastname;size:50" json:"maternal_lastname"`
	TypePoliceID uint      `gorm:"column:id_type_police_fk;not null" json:"type_police_id"`
	Sex          string    `gorm:"size:1;not null" json:"sex"`
	CUIP         string    `gorm:"size:50;unique;not null" json:"cuip"`
	RFC          string    `gorm:"size:20" json:"rfc"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
	CreatedBy    uint      `json:"created_by"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	UpdatedBy    uint      `json:"updated_by"`
	Deleted      bool      `gorm:"default:false" json:"deleted"`

	TypePolice *entity.TypePolice `gorm:"foreignKey:TypePoliceID" json:"type_police,omitempty"`
}
