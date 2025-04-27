//api-seguridad/resources/request_details/infrastructure/controllers/get_by_request_id_controller.go
package controllers

import (
	"net/http"
	"strconv"
	"api-seguridad/core/utils"
	"api-seguridad/resources/request_details/application"
	"github.com/gin-gonic/gin"
)

type GetByRequestIDController struct {
	useCase *application.GetByRequestIDUseCase
}

func NewGetByRequestIDController(useCase *application.GetByRequestIDUseCase) *GetByRequestIDController {
	return &GetByRequestIDController{useCase: useCase}
}

func (c *GetByRequestIDController) Handle(ctx *gin.Context) {
	requestID, err := strconv.ParseUint(ctx.Param("request_id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "ID de solicitud inválido", err)
		return
	}

	details, err := c.useCase.Execute(ctx.Request.Context(), uint(requestID))
	if err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "ID de solicitud es requerido" {
			status = http.StatusBadRequest
		}
		utils.ErrorResponse(ctx, status, err.Error(), nil)
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Detalles de solicitud obtenidos", details)
}