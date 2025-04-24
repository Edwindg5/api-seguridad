// api-seguridad/resources/request/infrastructure/controllers/create_controller.go
package controllers

import (
	"net/http"
	"api-seguridad/core/utils"
	"api-seguridad/resources/request/application"
	"api-seguridad/resources/request/domain/entities"
	"github.com/gin-gonic/gin"
)

type CreateRequestController struct {
	useCase *application.CreateRequestUseCase
}

func NewCreateRequestController(useCase *application.CreateRequestUseCase) *CreateRequestController {
	return &CreateRequestController{useCase: useCase}
}

func (c *CreateRequestController) Handle(ctx *gin.Context) {
	var request entities.Request
	if err := ctx.ShouldBindJSON(&request); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	// Set creator user
	if creatorID, exists := ctx.Get("userID"); exists {
		if uid, ok := creatorID.(uint); ok {
			request.CreatedBy = uid
			request.UpdatedBy = uid
		}
	}

	if err := c.useCase.Execute(ctx.Request.Context(), &request); err != nil {
		status := http.StatusInternalServerError
		switch err.Error() {
		case "office number is required", "municipality is required", 
		     "status is required", "creator user is required":
			status = http.StatusBadRequest
		}
		utils.ErrorResponse(ctx, status, err.Error(), nil)
		return
	}

	utils.SuccessResponse(ctx, http.StatusCreated, "Request created successfully", request)
}