// api-seguridad/resources/request_status/infrastructure/controllers/get_all_controller.go
package controllers

import (
	"net/http"
	"api-seguridad/core/utils"
	"api-seguridad/resources/request_status/application"
	"github.com/gin-gonic/gin"
)

type GetAllRequestStatusController struct {
	useCase *application.GetAllRequestStatusUseCase
}

func NewGetAllRequestStatusController(useCase *application.GetAllRequestStatusUseCase) *GetAllRequestStatusController {
	return &GetAllRequestStatusController{useCase: useCase}
}

func (c *GetAllRequestStatusController) Handle(ctx *gin.Context) {
	statusList, err := c.useCase.Execute(ctx.Request.Context())
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to get status list", err)
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Request status list retrieved successfully", statusList)
}