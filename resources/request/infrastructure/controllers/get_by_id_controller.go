// api-seguridad/resources/request/infrastructure/controllers/get_by_id_controller.go
package controllers

import (
	"net/http"
	"strconv"
	"api-seguridad/core/utils"
	"api-seguridad/resources/request/application"
	"github.com/gin-gonic/gin"
)

type GetRequestByIDController struct {
	useCase *application.GetRequestByIDUseCase
}

func NewGetRequestByIDController(useCase *application.GetRequestByIDUseCase) *GetRequestByIDController {
	return &GetRequestByIDController{useCase: useCase}
}

func (c *GetRequestByIDController) Handle(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request ID", err)
		return
	}

	request, err := c.useCase.Execute(ctx.Request.Context(), uint(id))
	if err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "request not found" {
			status = http.StatusNotFound
		}
		utils.ErrorResponse(ctx, status, err.Error(), nil)
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Request retrieved successfully", request)
}