// api-seguridad/resources/area_chiefs/infrastructure/controllers/create_controller.go
package controllers

import (
	"net/http"

	"api-seguridad/core/utils"
	"api-seguridad/resources/area_chiefs/application"
	"api-seguridad/resources/area_chiefs/domain/entities"
	"github.com/gin-gonic/gin"
)

type CreateAreaChiefController struct {
	useCase *application.CreateAreaChiefUseCase
}

func NewCreateAreaChiefController(useCase *application.CreateAreaChiefUseCase) *CreateAreaChiefController {
	return &CreateAreaChiefController{useCase: useCase}
}

func (c *CreateAreaChiefController) Handle(ctx *gin.Context) {
	var chief entities.AreaChief
	if err := ctx.ShouldBindJSON(&chief); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	// Se eliminó la asignación del usuario creador

	if err := c.useCase.Execute(ctx.Request.Context(), &chief); err != nil {
		status := http.StatusInternalServerError
		switch err.Error() {
		case "chief name is required", "position is required":
			status = http.StatusBadRequest
		}
		utils.ErrorResponse(ctx, status, err.Error(), nil)
		return
	}

	utils.SuccessResponse(ctx, http.StatusCreated, "Area chief created successfully", chief)
}
