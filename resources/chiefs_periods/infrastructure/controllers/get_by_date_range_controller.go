package controllers

import (
	"net/http"
	"api-seguridad/core/utils"
	"api-seguridad/resources/chiefs_periods/application"
	"github.com/gin-gonic/gin"
	"time"
)

type GetChiefsPeriodsByDateRangeController struct {
	useCase *application.GetChiefsPeriodsByDateRangeUseCase
}

func NewGetChiefsPeriodsByDateRangeController(useCase *application.GetChiefsPeriodsByDateRangeUseCase) *GetChiefsPeriodsByDateRangeController {
	return &GetChiefsPeriodsByDateRangeController{useCase: useCase}
}

func (c *GetChiefsPeriodsByDateRangeController) Handle(ctx *gin.Context) {
	startStr := ctx.Query("start_date")
	endStr := ctx.Query("end_date")

	start, err := time.Parse("2006-01-02", startStr)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid start date format (YYYY-MM-DD required)", err)
		return
	}

	end, err := time.Parse("2006-01-02", endStr)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid end date format (YYYY-MM-DD required)", err)
		return
	}

	periods, err := c.useCase.Execute(ctx.Request.Context(), start, end)
	if err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "start date cannot be after end date" {
			status = http.StatusBadRequest
		}
		utils.ErrorResponse(ctx, status, err.Error(), nil)
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Chiefs periods retrieved successfully", periods)
}