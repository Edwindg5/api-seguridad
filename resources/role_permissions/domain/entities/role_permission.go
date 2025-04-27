//api-seguridad/resources/role_permissions/domain/entities/role_permission.go
package entities

import (
	"time"
	entityroles "api-seguridad/resources/roles/domain/entities"
	entitypermissions "api-seguridad/resources/permissions/domain/entities"
	entityusers "api-seguridad/resources/users/domain/entities"
)

type RolePermission struct {
	ID            uint                     `gorm:"primaryKey;column:id_role_permission" json:"id"`
	RoleID        uint                     `gorm:"column:rol_id_fk;not null" json:"role_id"`
	PermissionID  uint                     `gorm:"column:permission_id_fk;not null" json:"permission_id"`
	Granted       bool                     `gorm:"default:true" json:"granted"`
	CreatedAt     time.Time                `gorm:"autoCreateTime;column:created_at" json:"created_at"`
	CreatedBy     uint                     `gorm:"column:created_by" json:"created_by"`
	UpdatedAt     time.Time                `gorm:"autoUpdateTime;column:updated_at" json:"updated_at"`
	UpdatedBy     uint                     `gorm:"column:updated_by" json:"updated_by"`
	Deleted       bool                     `gorm:"default:false;column:deleted" json:"deleted"`

	Role       *entityroles.Role            `gorm:"foreignKey:RoleID;references:ID" json:"role,omitempty"`
	Permission *entitypermissions.Permission `gorm:"foreignKey:PermissionID;references:ID" json:"permission,omitempty"`
	Creator    *entityusers.User            `gorm:"foreignKey:CreatedBy;references:ID" json:"creator,omitempty"`
	Updater    *entityusers.User            `gorm:"foreignKey:UpdatedBy;references:ID" json:"updater,omitempty"`
}

// Getters
func (rp *RolePermission) GetID() uint {
	return rp.ID
}

func (rp *RolePermission) GetRoleID() uint {
	return rp.RoleID
}

func (rp *RolePermission) GetPermissionID() uint {
	return rp.PermissionID
}

func (rp *RolePermission) IsGranted() bool {
	return rp.Granted
}

func (rp *RolePermission) GetCreatedAt() time.Time {
	return rp.CreatedAt
}

func (rp *RolePermission) GetCreatedBy() uint {
	return rp.CreatedBy
}

func (rp *RolePermission) GetUpdatedAt() time.Time {
	return rp.UpdatedAt
}

func (rp *RolePermission) GetUpdatedBy() uint {
	return rp.UpdatedBy
}

func (rp *RolePermission) IsDeleted() bool {
	return rp.Deleted
}

// Setters
func (rp *RolePermission) SetID(id uint) {
	rp.ID = id
}

func (rp *RolePermission) SetRoleID(id uint) {
	rp.RoleID = id
}

func (rp *RolePermission) SetPermissionID(id uint) {
	rp.PermissionID = id
}

func (rp *RolePermission) SetGranted(granted bool) {
	rp.Granted = granted
}

func (rp *RolePermission) SetCreatedAt(createdAt time.Time) {
	rp.CreatedAt = createdAt
}

func (rp *RolePermission) SetCreatedBy(createdBy uint) {
	rp.CreatedBy = createdBy
}

func (rp *RolePermission) SetUpdatedAt(updatedAt time.Time) {
	rp.UpdatedAt = updatedAt
}

func (rp *RolePermission) SetUpdatedBy(updatedBy uint) {
	rp.UpdatedBy = updatedBy
}

func (rp *RolePermission) SetDeleted(deleted bool) {
	rp.Deleted = deleted
}