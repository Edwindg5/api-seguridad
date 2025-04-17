//api-seguridad/resources/roles/domain/entities/role.go
package entity

import (
	"time"
)

type Role struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Title       string    `gorm:"size:50;not null" json:"title"`
	Description string    `gorm:"size:255" json:"description"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	CreatedBy   uint      `json:"created_by"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	UpdatedBy   uint      `json:"updated_by"`
	Deleted     bool      `gorm:"default:false" json:"deleted"`
}