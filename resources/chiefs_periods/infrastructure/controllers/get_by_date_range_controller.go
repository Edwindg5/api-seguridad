//api-seguridad/resources/chiefs_periods/infrastructure/controllers/get_by_date_range_controller.go
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
    // Validar parámetros requeridos
    startStr := ctx.Query("start_date")
    if startStr == "" {
        utils.ErrorResponse(ctx, http.StatusBadRequest, "start_date parameter is required", nil)
        return
    }
    
    endStr := ctx.Query("end_date")
    if endStr == "" {
        utils.ErrorResponse(ctx, http.StatusBadRequest, "end_date parameter is required", nil)
        return
    }

    // Parsear fechas con manejo de errores
    start, err := time.Parse("2006-01-02", startStr)
    if err != nil {
        utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid start date format (use YYYY-MM-DD)", err)
        return
    }

    end, err := time.Parse("2006-01-02", endStr)
    if err != nil {
        utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid end date format (use YYYY-MM-DD)", err)
        return
    }

    // Validar rango de fechas
    if start.After(end) {
        utils.ErrorResponse(ctx, http.StatusBadRequest, "start_date cannot be after end_date", nil)
        return
    }

    // Ajustar las fechas para incluir todo el día
    start = start.UTC().Truncate(24 * time.Hour)
    end = end.UTC().Add(23*time.Hour + 59*time.Minute + 59*time.Second)

    // Ejecutar consulta
    periods, err := c.useCase.Execute(ctx.Request.Context(), start, end)
    if err != nil {
        utils.ErrorResponse(ctx, http.StatusInternalServerError, "Error retrieving periods", err)
        return
    }

    // Manejar caso de no resultados
    if len(periods) == 0 {
        utils.SuccessResponse(ctx, http.StatusOK, "No periods found for the given date range", []interface{}{})
        return
    }

    utils.SuccessResponse(ctx, http.StatusOK, "Chiefs periods retrieved successfully", periods)
}