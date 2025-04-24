// api-seguridad/resources/request_status/infrastructure/controllers/create_controller.go
package controllers

import (
	"net/http"
	"api-seguridad/core/utils"
	"api-seguridad/resources/request_status/application"
	"api-seguridad/resources/request_status/domain/entities"
	"github.com/gin-gonic/gin"
)

type CreateRequestStatusController struct {
	useCase *application.CreateRequestStatusUseCase
}

func NewCreateRequestStatusController(useCase *application.CreateRequestStatusUseCase) *CreateRequestStatusController {
	return &CreateRequestStatusController{useCase: useCase}
}

func (c *CreateRequestStatusController) Handle(ctx *gin.Context) {
	var status entities.RequestStatus
	if err := ctx.ShouldBindJSON(&status); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	// Set creator user
	if creatorID, exists := ctx.Get("userID"); exists {
		if uid, ok := creatorID.(uint); ok {
			status.CreatedBy = uid
			status.UpdatedBy = uid
		}
	}

	if err := c.useCase.Execute(ctx.Request.Context(), &status); err != nil {
		status := http.StatusInternalServerError
		switch err.Error() {
		case "status name is required", "creator user is required":
			status = http.StatusBadRequest
		case "status with this name already exists":
			status = http.StatusConflict
		}
		utils.ErrorResponse(ctx, status, err.Error(), nil)
		return
	}

	utils.SuccessResponse(ctx, http.StatusCreated, "Request status created successfully", status)
}