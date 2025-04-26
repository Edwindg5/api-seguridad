package controllers

import (
	"net/http"
	"strconv"
	"api-seguridad/core/utils"
	"api-seguridad/resources/request_details/application"
	"github.com/gin-gonic/gin"
)

type GetRequestDetailByIDController struct {
	useCase *application.GetRequestDetailByIDUseCase
}

func NewGetRequestDetailByIDController(useCase *application.GetRequestDetailByIDUseCase) *GetRequestDetailByIDController {
	return &GetRequestDetailByIDController{useCase: useCase}
}

func (c *GetRequestDetailByIDController) Handle(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "ID inválido", err)
		return
	}

	detail, err := c.useCase.Execute(ctx.Request.Context(), uint(id))
	if err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "detalle de solicitud no encontrado" {
			status = http.StatusNotFound
		}
		utils.ErrorResponse(ctx, status, err.Error(), nil)
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Detalle de solicitud obtenido", detail)
}