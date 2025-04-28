//api-seguridad/resources/chiefs_periods/infrastructure/controllers/get_active_controller.go
package controllers

import (
	"net/http"
	"api-seguridad/core/utils"
	"api-seguridad/resources/chiefs_periods/application"
	"github.com/gin-gonic/gin"
)

type GetActiveChiefsPeriodController struct {
	useCase *application.GetActiveChiefsPeriodUseCase
}

func NewGetActiveChiefsPeriodController(useCase *application.GetActiveChiefsPeriodUseCase) *GetActiveChiefsPeriodController {
	return &GetActiveChiefsPeriodController{useCase: useCase}
}

func (c *GetActiveChiefsPeriodController) Handle(ctx *gin.Context) {
	period, err := c.useCase.Execute(ctx.Request.Context())
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to retrieve active period", err)
		return
	}

	if period == nil {
		// Cambiado a StatusOK con un mensaje claro
		utils.SuccessResponse(ctx, http.StatusOK, "No active period found", nil)
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Active chiefs period retrieved successfully", period)
}