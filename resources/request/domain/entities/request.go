//api-seguridad/resources/request/domain/entities/request.go
package entity

import (
	"time"
	"api-seguridad/resources/police/domain/entities"
	municipalitiesentity "api-seguridad/resources/municipalities/domain/entities"

)

type Request struct {
	ID                     uint      `gorm:"primaryKey" json:"id"`
	ReceiptDate            time.Time `gorm:"not null" json:"receipt_date"`
	OfficeNumber           string    `gorm:"size:100" json:"office_number"`
	SignatureName          string    `gorm:"size:255" json:"signature_name"`
	MunicipalityID         uint      `gorm:"not null" json:"municipality_id"`
	NumberPost             int       `json:"number_post"`
	StatusID               uint      `gorm:"not null" json:"status_id"`
	Date                   time.Time `gorm:"not null" json:"date"`
	NumberOfLettersDelivered int     `json:"number_of_letters_delivered"`
	DeliveryName           string    `gorm:"size:255" json:"delivery_name"`
	ReceiveName            string    `gorm:"size:255" json:"receive_name"`
	DepartmentArea         string    `gorm:"size:100" json:"department_area"`
	Phone                  string    `gorm:"size:20" json:"phone"`
	CreatedAt              time.Time `gorm:"autoCreateTime" json:"created_at"`
	CreatedBy              uint      `json:"created_by"`
	UpdatedAt              time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	UpdatedBy              uint      `json:"updated_by"`
	Deleted                bool      `gorm:"default:false" json:"deleted"`

	Municipality *municipalitiesentity.Municipality `gorm:"foreignKey:MunicipalityID" json:"municipality,omitempty"`
	Status       *RequestStatus `gorm:"foreignKey:StatusID" json:"status,omitempty"`
	Police       []*entity.Police      `gorm:"many2many:request_details;" json:"police,omitempty"`
}