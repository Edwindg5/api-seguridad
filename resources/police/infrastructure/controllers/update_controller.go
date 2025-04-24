//api-seguridad/resources/police/infrastructure/controllers/update_controller.go
package controllers

import (
	"net/http"
	"api-seguridad/core/utils"
	"api-seguridad/resources/police/application"
	"api-seguridad/resources/police/domain/entities"
	"github.com/gin-gonic/gin"
)

type UpdatePoliceController struct {
	useCase *application.UpdatePoliceUseCase
}

func NewUpdatePoliceController(useCase *application.UpdatePoliceUseCase) *UpdatePoliceController {
	return &UpdatePoliceController{useCase: useCase}
}

func (c *UpdatePoliceController) Handle(ctx *gin.Context) {
	var police entities.Police
	if err := ctx.ShouldBindJSON(&police); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	if err := c.useCase.Execute(ctx.Request.Context(), &police); err != nil {
		status := http.StatusInternalServerError
		switch err.Error() {
		case "invalid police ID", "name and paternal lastname are required":
			status = http.StatusBadRequest
		case "police not found":
			status = http.StatusNotFound
		}
		utils.ErrorResponse(ctx, status, err.Error(), nil)
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Police updated successfully", police)
}