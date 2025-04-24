// api-seguridad/resources/users/domain/entities/user.go
package entities

import (
	"time"
	rol "api-seguridad/resources/roles/domain/entities"
)

type User struct {
	ID        uint      `gorm:"primaryKey;column:id_user" json:"id"`
	FirstName string    `gorm:"size:50;not null" json:"first_name"`
	LastName  string    `gorm:"size:50;not null;column:lastname" json:"last_name"`
	Username  string    `gorm:"size:50;unique;not null" json:"username"`
	Email     string    `gorm:"size:50;unique;not null" json:"email"`
	Password  string    `gorm:"size:255;not null" json:"password"` 
	RoleID    uint      `gorm:"column:rol_id_fk;not null" json:"role_id"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	CreatedBy uint      `json:"created_by"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	UpdatedBy uint      `json:"updated_by"`
	Deleted   bool      `gorm:"default:false" json:"deleted"`

	Role    *rol.Role `gorm:"foreignKey:RoleID" json:"role,omitempty"`
	Creator *User     `gorm:"foreignKey:CreatedBy" json:"creator,omitempty"`
	Updater *User     `gorm:"foreignKey:UpdatedBy" json:"updater,omitempty"`
}

// Getters
func (u *User) GetID() uint {
	return u.ID
}

func (u *User) GetFirstName() string {
	return u.FirstName
}

func (u *User) GetLastName() string {
	return u.LastName
}

func (u *User) GetUsername() string {
	return u.Username
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetPassword() string {
	return u.Password
}

func (u *User) GetRoleID() uint {
	return u.RoleID
}

func (u *User) GetCreatedAt() time.Time {
	return u.CreatedAt
}

func (u *User) GetCreatedBy() uint {
	return u.CreatedBy
}

func (u *User) GetUpdatedAt() time.Time {
	return u.UpdatedAt
}

func (u *User) GetUpdatedBy() uint {
	return u.UpdatedBy
}

func (u *User) IsDeleted() bool {
	return u.Deleted
}

// Setters
func (u *User) SetID(id uint) {
	u.ID = id
}

func (u *User) SetFirstName(firstName string) {
	u.FirstName = firstName
}

func (u *User) SetLastName(lastName string) {
	u.LastName = lastName
}

func (u *User) SetUsername(username string) {
	u.Username = username
}

func (u *User) SetEmail(email string) {
	u.Email = email
}

func (u *User) SetPassword(password string) {
	u.Password = password
}

func (u *User) SetRoleID(roleID uint) {
	u.RoleID = roleID
}

func (u *User) SetCreatedBy(createdBy uint) {
	u.CreatedBy = createdBy
}

func (u *User) SetUpdatedBy(updatedBy uint) {
	u.UpdatedBy = updatedBy
}

func (u *User) SetDeleted(deleted bool) {
	u.Deleted = deleted
}


