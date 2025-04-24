// api-seguridad/resources/request_status/infrastructure/controllers/get_by_id_controller.go
package controllers

import (
	"net/http"
	"strconv"
	"api-seguridad/core/utils"
	"api-seguridad/resources/request_status/application"
	"github.com/gin-gonic/gin"
)

type GetRequestStatusByIDController struct {
	useCase *application.GetRequestStatusByIDUseCase
}

func NewGetRequestStatusByIDController(useCase *application.GetRequestStatusByIDUseCase) *GetRequestStatusByIDController {
	return &GetRequestStatusByIDController{useCase: useCase}
}

func (c *GetRequestStatusByIDController) Handle(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid status ID", err)
		return
	}

	status, err := c.useCase.Execute(ctx.Request.Context(), uint(id))
	if err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "status not found" {
			status = http.StatusNotFound
		}
		utils.ErrorResponse(ctx, status, err.Error(), nil)
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Request status retrieved successfully", status)
}