// api-seguridad/resources/delegation/domain/entities/delegation.go
package entity

import (

	entityusers "api-seguridad/resources/users/domain/entities"
	"time"
)

type Delegation struct {
	ID              uint       `gorm:"primaryKey" json:"id"`
	Name            string     `gorm:"size:255;not null" json:"name"`
	Active          bool       `json:"active"`
	MunicipalityID  uint       `gorm:"not null" json:"municipality_id"`
	CreatedAt       time.Time  `gorm:"autoCreateTime" json:"created_at"`
	CreatedBy       uint       `json:"created_by"`
	UpdatedAt       time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	UpdatedBy       uint       `json:"updated_by"`
	Deleted         bool       `gorm:"default:false" json:"deleted"`

	CreatedByUser   *entityusers.User         `gorm:"foreignKey:CreatedBy" json:"created_by_user,omitempty"`
	UpdatedByUser   *entityusers.User         `gorm:"foreignKey:UpdatedBy" json:"updated_by_user,omitempty"`
}