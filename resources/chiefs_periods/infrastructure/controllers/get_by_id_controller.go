package controllers

import (
	"net/http"
	"strconv"
	"api-seguridad/core/utils"
	"api-seguridad/resources/chiefs_periods/application"
	"github.com/gin-gonic/gin"
)

type GetChiefsPeriodByIDController struct {
	useCase *application.GetChiefsPeriodByIDUseCase
}

func NewGetChiefsPeriodByIDController(useCase *application.GetChiefsPeriodByIDUseCase) *GetChiefsPeriodByIDController {
	return &GetChiefsPeriodByIDController{useCase: useCase}
}

func (c *GetChiefsPeriodByIDController) Handle(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid period ID", err)
		return
	}

	period, err := c.useCase.Execute(ctx.Request.Context(), uint(id))
	if err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "period not found" {
			status = http.StatusNotFound
		}
		utils.ErrorResponse(ctx, status, err.Error(), nil)
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Chiefs period retrieved successfully", period)
}