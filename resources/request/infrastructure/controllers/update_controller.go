// api-seguridad/resources/request/infrastructure/controllers/update_controller.go
package controllers

import (
	"net/http"
	"strconv"
	"api-seguridad/core/utils"
	"api-seguridad/resources/request/application"
	"api-seguridad/resources/request/domain/entities"
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

	var request entities.Request
	if err := ctx.ShouldBindJSON(&request); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	// Set ID from URL and updater user
	request.ID = uint(id)
	if updaterID, exists := ctx.Get("userID"); exists {
		if uid, ok := updaterID.(uint); ok {
			request.UpdatedBy = uid
		}
	}

	if err := c.useCase.Execute(ctx.Request.Context(), &request); err != nil {
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

	utils.SuccessResponse(ctx, http.StatusOK, "Request updated successfully", request)
}