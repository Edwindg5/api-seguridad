// api-seguridad/resources/request/domain/entities/request.go
package entities

import (
	municipalitiesentity "api-seguridad/resources/municipalities/domain/entities"
	policeentity "api-seguridad/resources/police/domain/entities"
	"time"
)

type Request struct {
	ID                       uint      `gorm:"primaryKey" json:"id"`
	ReceiptDate              time.Time `gorm:"not null" json:"receipt_date"`
	OfficeNumber             string    `gorm:"size:100" json:"office_number"`
	SignatureName            string    `gorm:"size:255" json:"signature_name"`
	MunicipalityID           uint      `gorm:"column:id_municipalities_fk;not null" json:"municipality_id"`
	NumberPost               int       `json:"number_post"`
	StatusID                 uint      `gorm:"column:id_status_fk;not null" json:"status_id"`
	Date                     time.Time `gorm:"not null" json:"date"`
	NumberOfLettersDelivered int       `json:"number_of_letters_delivered"`
	DeliveryName             string    `gorm:"size:255" json:"delivery_name"`
	ReceiveName              string    `gorm:"size:255" json:"receive_name"`
	DepartmentArea           string    `gorm:"size:100" json:"department_area"`
	Phone                    string    `gorm:"size:20" json:"phone"`
	CeoChiefID               uint      `gorm:"column:ceo_chief_id" json:"ceo_chief_id"`
	LegalChiefID             uint      `gorm:"column:legal_chief_id" json:"legal_chief_id"`
	CreatedAt                time.Time `gorm:"autoCreateTime" json:"created_at"`
	CreatedBy                uint      `json:"created_by"`
	UpdatedAt                time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	UpdatedBy                uint      `json:"updated_by"`
	Deleted                  bool      `gorm:"default:false" json:"deleted"`

	Municipality *municipalitiesentity.Municipality `gorm:"foreignKey:MunicipalityID" json:"municipality,omitempty"`
	Status       *RequestStatus                     `gorm:"foreignKey:StatusID" json:"status,omitempty"`
	Police       []*policeentity.Police             `gorm:"many2many:request_details;" json:"police,omitempty"`
	CeoChief     *AreaChief                         `gorm:"foreignKey:CeoChiefID" json:"ceo_chief,omitempty"` // crear la tabla areachief
	LegalChief   *AreaChief                         `gorm:"foreignKey:LegalChiefID" json:"legal_chief,omitempty"`
}

//crear el recurso request details o integralo en el recurso request:package entity

/*import (
    "time"
)

type RequestDetails struct {
    ID               uint      `gorm:"primaryKey" json:"id"`
    RequestID        uint      `gorm:"column:request_id" json:"request_id"`
    PoliceID         uint      `gorm:"column:id_police_fk" json:"police_id"`
    Active           bool      `gorm:"default:true" json:"active"`
    Census           bool      `gorm:"default:false" json:"census"`
    Located          bool      `gorm:"default:false" json:"located"`
    Register         bool      `gorm:"default:false" json:"register"`
    Approved         bool      `gorm:"default:false" json:"approved"`
    Comments         string    `gorm:"type:text" json:"comments"`
    MunicipalityActive bool    `gorm:"default:false" json:"municipality_active"`
    CreatedAt        time.Time `gorm:"autoCreateTime" json:"created_at"`
    CreatedBy        uint      `json:"created_by"`
    UpdatedAt        time.Time `gorm:"autoUpdateTime" json:"updated_at"`
    UpdatedBy        uint      `json:"updated_by"`
    Deleted          bool      `gorm:"default:false" json:"deleted"`

    Request *Request `gorm:"foreignKey:RequestID" json:"request,omitempty"`
    Police  *Police  `gorm:"foreignKey:PoliceID" json:"police,omitempty"`
}*/
