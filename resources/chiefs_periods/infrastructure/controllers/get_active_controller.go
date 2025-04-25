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
		status := http.StatusInternalServerError
		if err.Error() == "no active period found" {
			status = http.StatusNotFound
		}
		utils.ErrorResponse(ctx, status, err.Error(), nil)
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Active chiefs period retrieved successfully", period)
}