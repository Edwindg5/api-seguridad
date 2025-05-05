// api-seguridad/resources/request/infrastructure/controllers/update_controller.go
package controllers

import (
	"net/http"
	"strconv"
	"api-seguridad/core/utils"
	"api-seguridad/resources/request/application"
	
	"time"
	"github.com/gin-gonic/gin"
)

type UpdateRequestController struct {
	useCase *application.UpdateRequestUseCase
}

func NewUpdateRequestController(useCase *application.UpdateRequestUseCase) *UpdateRequestController {
	return &UpdateRequestController{useCase: useCase}
}

func (c *UpdateRequestController) Handle(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request ID", err)
		return
	}

	// Obtener request existente
	existingRequest, err := c.useCase.GetByID(ctx.Request.Context(), uint(id))
	if err != nil || existingRequest == nil {
		utils.ErrorResponse(ctx, http.StatusNotFound, "Request not found", err)
		return
	}

	// Bind de todos los campos actualizables
	var updateData struct {
		OfficeNumber           string    `json:"office_number"`
		MunicipalitiesID       uint      `json:"municipalities_id"`
		StatusID               uint      `json:"status_id"`
		ReceiptDate           time.Time `json:"receipt_date"`
		Date                  time.Time `json:"date"`
		SignatureName         string    `json:"signature_name"`
		NumberPost            int       `json:"number_post"`
		NumberOfLettersDelivered int    `json:"number_of_letters_delivered"`
		DeliveryName          string    `json:"delivery_name"`
		ReceiveName           string    `json:"receive_name"`
		DepartmentArea        string    `json:"department_area"`
		Phone                 string    `json:"phone"`
		CeoChiefID            uint      `json:"ceo_chief_id"`
		LegalChiefID          uint      `json:"legal_chief_id"`
	}
	
	if err := ctx.ShouldBindJSON(&updateData); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	// Actualizar todos los campos permitidos
	existingRequest.OfficeNumber = updateData.OfficeNumber
	existingRequest.MunicipalitiesID = updateData.MunicipalitiesID
	existingRequest.StatusID = updateData.StatusID
	existingRequest.ReceiptDate = updateData.ReceiptDate
	existingRequest.Date = updateData.Date
	existingRequest.SignatureName = updateData.SignatureName
	existingRequest.NumberPost = updateData.NumberPost
	existingRequest.NumberOfLettersDelivered = updateData.NumberOfLettersDelivered
	existingRequest.DeliveryName = updateData.DeliveryName
	existingRequest.ReceiveName = updateData.ReceiveName
	existingRequest.DepartmentArea = updateData.DepartmentArea
	existingRequest.Phone = updateData.Phone
	existingRequest.CeoChiefID = updateData.CeoChiefID
	existingRequest.LegalChiefID = updateData.LegalChiefID

	// Establecer usuario actualizador si está disponible (opcional)
	if updaterID, exists := ctx.Get("userID"); exists {
		if uid, ok := updaterID.(uint); ok {
			existingRequest.UpdatedBy = uid
		}
	}

	if err := c.useCase.Execute(ctx.Request.Context(), existingRequest); err != nil {
		status := http.StatusInternalServerError
		switch err.Error() {
		case "invalid request ID", "office number is required":
			status = http.StatusBadRequest
		case "request not found":
			status = http.StatusNotFound
		}
		utils.ErrorResponse(ctx, status, err.Error(), nil)
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Request updated successfully", existingRequest)
}