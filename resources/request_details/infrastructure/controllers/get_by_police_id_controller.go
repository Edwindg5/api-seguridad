package controllers

import (
	"net/http"
	"strconv"
	"api-seguridad/core/utils"
	"api-seguridad/resources/request_details/application"
	"github.com/gin-gonic/gin"
)

type GetByPoliceIDController struct {
	useCase *application.GetByPoliceIDUseCase
}

func NewGetByPoliceIDController(useCase *application.GetByPoliceIDUseCase) *GetByPoliceIDController {
	return &GetByPoliceIDController{useCase: useCase}
}

func (c *GetByPoliceIDController) Handle(ctx *gin.Context) {
	policeID, err := strconv.ParseUint(ctx.Param("police_id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "ID de policía inválido", err)
		return
	}

	details, err := c.useCase.Execute(ctx.Request.Context(), uint(policeID))
	if err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "ID de policía es requerido" {
			status = http.StatusBadRequest
		}
		utils.ErrorResponse(ctx, status, err.Error(), nil)
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Detalles por policía obtenidos", details)
}