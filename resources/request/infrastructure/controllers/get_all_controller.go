// api-seguridad/resources/request/infrastructure/controllers/get_all_controller.go
package controllers

import (
	"net/http"
	"api-seguridad/core/utils"
	"api-seguridad/resources/request/application"
	"github.com/gin-gonic/gin"
)

type GetAllRequestsController struct {
	useCase *application.GetAllRequestsUseCase
}

func NewGetAllRequestsController(useCase *application.GetAllRequestsUseCase) *GetAllRequestsController {
	return &GetAllRequestsController{useCase: useCase}
}

func (c *GetAllRequestsController) Handle(ctx *gin.Context) {
	requests, err := c.useCase.Execute(ctx.Request.Context())
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Error retrieving requests", err)
		return
	}

	// Manejar caso cuando no hay requests
	if len(requests) == 0 {
		utils.SuccessResponse(ctx, http.StatusOK, "No requests found", []interface{}{})
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Requests retrieved successfully", requests)
}