package entity

import (
	"time"
)

type RequestStatus struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"size:100;unique;not null" json:"name"`
	Description string    `gorm:"type:text" json:"description"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	CreatedBy   uint      `json:"created_by"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	UpdatedBy   uint      `json:"updated_by"`
	Deleted     bool      `gorm:"default:false" json:"deleted"`
}