//api-seguridad/resources/chiefs_periods/infrastructure/controllers/update_controller.go
package controllers

import (
	"net/http"
	"strconv"
	"api-seguridad/core/utils"
	"api-seguridad/resources/chiefs_periods/application"
	"api-seguridad/resources/chiefs_periods/domain/entities"
	"github.com/gin-gonic/gin"
)

type UpdateChiefsPeriodController struct {
	useCase *application.UpdateChiefsPeriodUseCase
}

func NewUpdateChiefsPeriodController(useCase *application.UpdateChiefsPeriodUseCase) *UpdateChiefsPeriodController {
	return &UpdateChiefsPeriodController{useCase: useCase}
}

func (c *UpdateChiefsPeriodController) Handle(ctx *gin.Context) {
	// Parsear el ID del periodo
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid period ID", err)
		return
	}

	// Bind del JSON de entrada
	var period entities.ChiefsPeriod
	if err := ctx.ShouldBindJSON(&period); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	// Asignar el ID
	period.ID = uint(id)

	// Obtener ID del usuario que realiza la actualización
	if updaterID, exists := ctx.Get("userID"); exists {
		if uid, ok := updaterID.(uint); ok {
			period.UpdatedBy = uid
		}
	}

	// Ejecutar el caso de uso para la actualización
	if err := c.useCase.Execute(ctx.Request.Context(), &period); err != nil {
		status := http.StatusInternalServerError
		switch err.Error() {
		case "invalid period ID", "both CEO and Legal chief IDs are required":
			status = http.StatusBadRequest
		case "period not found":
			status = http.StatusNotFound
		case "there is already an active period":
			status = http.StatusConflict
		case "start date cannot be after end date":
			status = http.StatusBadRequest
		}
		utils.ErrorResponse(ctx, status, err.Error(), nil)
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Chiefs period updated successfully", period)
}