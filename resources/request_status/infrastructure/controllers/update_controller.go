// api-seguridad/resources/request_status/infrastructure/controllers/update_controller.go
package controllers

import (
	"net/http"
	"strconv"
	"api-seguridad/core/utils"
	"api-seguridad/resources/request_status/application"
	"api-seguridad/resources/request_status/domain/entities"
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

	var status entities.RequestStatus
	if err := ctx.ShouldBindJSON(&status); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	// Set ID from URL and updater user
	status.ID = uint(id)
	if updaterID, exists := ctx.Get("userID"); exists {
		if uid, ok := updaterID.(uint); ok {
			status.UpdatedBy = uid
		}
	}

	if err := c.useCase.Execute(ctx.Request.Context(), &status); err != nil {
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

	utils.SuccessResponse(ctx, http.StatusOK, "Request status updated successfully", status)
}