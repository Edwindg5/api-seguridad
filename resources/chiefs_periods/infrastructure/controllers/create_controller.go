//api-seguridad/resources/chiefs_periods/infrastructure/controllers/create_controller.go
package controllers

import (
	"net/http"
	"api-seguridad/core/utils"
	"api-seguridad/resources/chiefs_periods/application"
	"api-seguridad/resources/chiefs_periods/domain/entities"
	"github.com/gin-gonic/gin"
)

type CreateChiefsPeriodController struct {
	useCase *application.CreateChiefsPeriodUseCase
}

func NewCreateChiefsPeriodController(useCase *application.CreateChiefsPeriodUseCase) *CreateChiefsPeriodController {
	return &CreateChiefsPeriodController{useCase: useCase}
}

func (c *CreateChiefsPeriodController) Handle(ctx *gin.Context) {
	var period entities.ChiefsPeriod
	if err := ctx.ShouldBindJSON(&period); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	// Obtener ID del usuario que realiza la creación
	if creatorID, exists := ctx.Get("userID"); exists {
		if uid, ok := creatorID.(uint); ok {
			period.CreatedBy = uid
			period.UpdatedBy = uid
		}
	}

	if err := c.useCase.Execute(ctx.Request.Context(), &period); err != nil {
		status := http.StatusInternalServerError
		switch err.Error() {
		case "both CEO and Legal chief IDs are required", "start date is required", "start date cannot be after end date":
			status = http.StatusBadRequest
		case "there is already an active period":
			status = http.StatusConflict
		}
		utils.ErrorResponse(ctx, status, err.Error(), nil)
		return
	}

	utils.SuccessResponse(ctx, http.StatusCreated, "Chiefs period created successfully", period)
}