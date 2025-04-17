package entity

import (
	"time"
)

type TypePolice struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	TitleKindPolice string    `gorm:"size:50;not null" json:"title_kind_police"`
	CreatedAt      time.Time `gorm:"autoCreateTime" json:"created_at"`
	CreatedBy      uint      `json:"created_by"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	UpdatedBy      uint      `json:"updated_by"`
	Deleted        bool      `gorm:"default:false" json:"deleted"`
}