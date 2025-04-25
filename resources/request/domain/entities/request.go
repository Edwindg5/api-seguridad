// api-seguridad/resources/request/domain/entities/request.go
package entities

import (
	"time"
	entitymunicipalities "api-seguridad/resources/municipalities/domain/entities"
	entitystatus "api-seguridad/resources/request_status/domain/entities"
	entitychiefs "api-seguridad/resources/area_chiefs/domain/entities"
	entityusers "api-seguridad/resources/users/domain/entities"
)

type Request struct {
    ID                      uint      `gorm:"primaryKey;column:id_request" json:"id"`
    ReceiptDate             time.Time `gorm:"column:receipt_date" json:"receipt_date"`
    OfficeNumber           string    `gorm:"column:office_number;size:100" json:"office_number"`
    SignatureName          string    `gorm:"column:signature_name;size:255" json:"signature_name"`
    MunicipalitiesID       uint      `gorm:"column:id_municipalities_fk" json:"municipalities_id"`
    NumberPost             int       `gorm:"column:number_post" json:"number_post"`
    StatusID               uint      `gorm:"column:id_status_fk" json:"status_id"`
    Date                   time.Time `gorm:"column:date" json:"date"`
    NumberOfLettersDelivered int     `gorm:"column:number_of_letters_delivered" json:"number_of_letters_delivered"`
    DeliveryName           string    `gorm:"column:delivery_name;size:255" json:"delivery_name"`
    ReceiveName            string    `gorm:"column:receive_name;size:255" json:"receive_name"`
    DepartmentArea         string    `gorm:"column:department_area;size:100" json:"department_area"`
    Phone                  string    `gorm:"column:phone;size:20" json:"phone"`
    CeoChiefID             uint      `gorm:"column:ceo_chief_id" json:"ceo_chief_id"`
    LegalChiefID           uint      `gorm:"column:legal_chief_id" json:"legal_chief_id"`
    CreatedAt              time.Time `gorm:"column:created_at" json:"created_at"`
    CreatedBy              uint      `gorm:"column:created_by" json:"created_by"`
    UpdatedBy              uint      `gorm:"column:updated_by" json:"updated_by"`
    UpdatedAt              time.Time `gorm:"column:updated_at" json:"updated_at"`
    Deleted                bool      `gorm:"column:deleted" json:"deleted"`

    Municipalities *entitymunicipalities.Municipality `gorm:"foreignKey:MunicipalitiesID;references:id_municipalities" json:"municipalities,omitempty"`
    Status         *entitystatus.RequestStatus        `gorm:"foreignKey:StatusID;references:id_status" json:"status,omitempty"`
    CeoChief       *entitychiefs.AreaChief           `gorm:"foreignKey:CeoChiefID;references:id_official" json:"ceo_chief,omitempty"`
    LegalChief     *entitychiefs.AreaChief           `gorm:"foreignKey:LegalChiefID;references:id_official" json:"legal_chief,omitempty"`
    CreatedByUser  *entityusers.User                 `gorm:"foreignKey:CreatedBy;references:id_user" json:"created_by_user,omitempty"`
    UpdatedByUser  *entityusers.User                 `gorm:"foreignKey:UpdatedBy;references:id_user" json:"updated_by_user,omitempty"`
}
// Getters
func (r *Request) GetID() uint {
	return r.ID
}

func (r *Request) GetReceiptDate() time.Time {
	return r.ReceiptDate
}

func (r *Request) GetOfficeNumber() string {
	return r.OfficeNumber
}

func (r *Request) GetSignatureName() string {
	return r.SignatureName
}

func (r *Request) GetMunicipalitiesID() uint {
	return r.MunicipalitiesID
}

func (r *Request) GetNumberPost() int {
	return r.NumberPost
}

func (r *Request) GetStatusID() uint {
	return r.StatusID
}

func (r *Request) GetDate() time.Time {
	return r.Date
}

func (r *Request) GetNumberOfLettersDelivered() int {
	return r.NumberOfLettersDelivered
}

func (r *Request) GetDeliveryName() string {
	return r.DeliveryName
}

func (r *Request) GetReceiveName() string {
	return r.ReceiveName
}

func (r *Request) GetDepartmentArea() string {
	return r.DepartmentArea
}

func (r *Request) GetPhone() string {
	return r.Phone
}

func (r *Request) GetCeoChiefID() uint {
	return r.CeoChiefID
}

func (r *Request) GetLegalChiefID() uint {
	return r.LegalChiefID
}

func (r *Request) GetCreatedAt() time.Time {
	return r.CreatedAt
}

func (r *Request) GetCreatedBy() uint {
	return r.CreatedBy
}

func (r *Request) GetUpdatedBy() uint {
	return r.UpdatedBy
}

func (r *Request) GetUpdatedAt() time.Time {
	return r.UpdatedAt
}

func (r *Request) IsDeleted() bool {
	return r.Deleted
}

// Setters
func (r *Request) SetID(id uint) {
	r.ID = id
}

func (r *Request) SetReceiptDate(date time.Time) {
	r.ReceiptDate = date
}

func (r *Request) SetOfficeNumber(number string) {
	r.OfficeNumber = number
}

func (r *Request) SetSignatureName(name string) {
	r.SignatureName = name
}

func (r *Request) SetMunicipalitiesID(id uint) {
	r.MunicipalitiesID = id
}

func (r *Request) SetNumberPost(number int) {
	r.NumberPost = number
}

func (r *Request) SetStatusID(id uint) {
	r.StatusID = id
}

func (r *Request) SetDate(date time.Time) {
	r.Date = date
}

func (r *Request) SetNumberOfLettersDelivered(number int) {
	r.NumberOfLettersDelivered = number
}

func (r *Request) SetDeliveryName(name string) {
	r.DeliveryName = name
}

func (r *Request) SetReceiveName(name string) {
	r.ReceiveName = name
}

func (r *Request) SetDepartmentArea(area string) {
	r.DepartmentArea = area
}

func (r *Request) SetPhone(phone string) {
	r.Phone = phone
}

func (r *Request) SetCeoChiefID(id uint) {
	r.CeoChiefID = id
}

func (r *Request) SetLegalChiefID(id uint) {
	r.LegalChiefID = id
}

func (r *Request) SetCreatedAt(createdAt time.Time) {
	r.CreatedAt = createdAt
}

func (r *Request) SetCreatedBy(createdBy uint) {
	r.CreatedBy = createdBy
}

func (r *Request) SetUpdatedBy(updatedBy uint) {
	r.UpdatedBy = updatedBy
}

func (r *Request) SetUpdatedAt(updatedAt time.Time) {
	r.UpdatedAt = updatedAt
}

func (r *Request) SetDeleted(deleted bool) {
	r.Deleted = deleted
}