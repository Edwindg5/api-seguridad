// api-seguridad/resources/request/infrastructure/controllers/update_controller.go
package controllers

import (
	"net/http"
	"strconv"
	"api-seguridad/core/utils"
	"api-seguridad/resources/request/application"
	
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

	// Bind solo campos actualizables
	var updateData struct {
		OfficeNumber           string `json:"office_number"`
		NumberOfLettersDelivered int  `json:"number_of_letters_delivered"`
		DepartmentArea         string `json:"department_area"`
		Phone                  string `json:"phone"`
	}
	if err := ctx.ShouldBindJSON(&updateData); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	// Actualizar campos permitidos
	existingRequest.OfficeNumber = updateData.OfficeNumber
	existingRequest.NumberOfLettersDelivered = updateData.NumberOfLettersDelivered
	existingRequest.DepartmentArea = updateData.DepartmentArea
	existingRequest.Phone = updateData.Phone

	// Establecer usuario actualizador
	if updaterID, exists := ctx.Get("userID"); exists {
		if uid, ok := updaterID.(uint); ok {
			existingRequest.UpdatedBy = uid
		}
	}

	if err := c.useCase.Execute(ctx.Request.Context(), existingRequest); err != nil {
		status := http.StatusInternalServerError
		switch err.Error() {
		case "invalid request ID", "office number is required", "updater user is required":
			status = http.StatusBadRequest
		case "request not found":
			status = http.StatusNotFound
		}
		utils.ErrorResponse(ctx, status, err.Error(), nil)
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Request updated successfully", existingRequest)
}