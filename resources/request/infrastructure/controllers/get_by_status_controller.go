// api-seguridad/resources/request/infrastructure/controllers/get_by_status_controller.go
package controllers

import (
	"net/http"
	"strconv"
	"api-seguridad/core/utils"
	"api-seguridad/resources/request/application"
	"github.com/gin-gonic/gin"
)

type GetRequestsByStatusController struct {
	useCase *application.GetRequestsByStatusUseCase
}

func NewGetRequestsByStatusController(useCase *application.GetRequestsByStatusUseCase) *GetRequestsByStatusController {
	return &GetRequestsByStatusController{useCase: useCase}
}

func (c *GetRequestsByStatusController) Handle(ctx *gin.Context) {
	statusID, err := strconv.ParseUint(ctx.Query("status_id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Status ID parameter is required", nil)
		return
	}

	requests, err := c.useCase.Execute(ctx.Request.Context(), uint(statusID))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to get requests by status", err)
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Requests retrieved successfully", requests)
}