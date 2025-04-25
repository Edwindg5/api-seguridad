// api-seguridad/resources/request_status/infrastructure/controllers/update_controller.go
package controllers

import (
	"net/http"
	"strconv"
	"api-seguridad/core/utils"
	"api-seguridad/resources/request_status/application"
	
	"github.com/gin-gonic/gin"
)

type UpdateRequestStatusController struct {
	useCase *application.UpdateRequestStatusUseCase
}

func NewUpdateRequestStatusController(useCase *application.UpdateRequestStatusUseCase) *UpdateRequestStatusController {
	return &UpdateRequestStatusController{useCase: useCase}
}

func (c *UpdateRequestStatusController) Handle(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid status ID", err)
		return
	}

	// Obtener el estado existente a través del use case
	existingStatus, err := c.useCase.GetByID(ctx.Request.Context(), uint(id))
	if err != nil || existingStatus == nil {
		utils.ErrorResponse(ctx, http.StatusNotFound, "Status not found", err)
		return
	}

	// Bind solo los campos actualizables
	var updateData struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}
	if err := ctx.ShouldBindJSON(&updateData); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	// Actualizar campos
	existingStatus.Name = updateData.Name
	existingStatus.Description = updateData.Description
	
	// Establecer usuario actualizador
	if updaterID, exists := ctx.Get("userID"); exists {
		if uid, ok := updaterID.(uint); ok {
			existingStatus.UpdatedBy = uid
		}
	}

	// Usar el use case para ejecutar la actualización
	if err := c.useCase.Execute(ctx.Request.Context(), existingStatus); err != nil {
		status := http.StatusInternalServerError
		switch err.Error() {
		case "invalid status ID", "status name is required", "updater user is required":
			status = http.StatusBadRequest
		case "status not found", "status with this name already exists":
			status = http.StatusConflict
		}
		utils.ErrorResponse(ctx, status, err.Error(), nil)
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Request status updated successfully", existingStatus)
}