//api-seguridad/resources/users/domain/entities/user.go
package entity

import (
	"time"
	"api-seguridad/resources/roles/domain/entities" 
)

type User struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    FirstName string    `gorm:"size:50;not null" json:"first_name"`
    LastName  string    `gorm:"size:50;not null" json:"last_name"`
    Username  string    `gorm:"size:50;unique;not null" json:"username"`
    Email     string    `gorm:"size:50;unique;not null" json:"email"`
    Password  string    `gorm:"size:255;not null" json:"-"` // Cambiado a 255 para hashes
    RoleID    uint      `gorm:"column:rol_id_fk;not null" json:"role_id"`
    CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
    CreatedBy uint      `json:"created_by"`
    UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
    UpdatedBy uint      `json:"updated_by"`
    Deleted   bool      `gorm:"default:false" json:"deleted"`
    
    Role *entity.Role `gorm:"foreignKey:RoleID" json:"role,omitempty"`
}