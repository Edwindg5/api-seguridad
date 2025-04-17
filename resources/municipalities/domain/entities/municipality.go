// api-seguridad/resources/municipalities/domain/entities/municipality.go
package entity

import (
	
	"time"
)

type Municipality struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"size:50;not null" json:"name"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	CreatedBy uint      `json:"created_by"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	UpdatedBy uint      `json:"updated_by"`
	Deleted   bool      `gorm:"default:false" json:"deleted"`

}