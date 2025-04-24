//api-seguridad/resources/police/infrastructure/controllers/create_controller.go
package controllers

import (
	"net/http"
	"api-seguridad/core/utils"
	"api-seguridad/resources/police/application"
	"api-seguridad/resources/police/domain/entities"
	"github.com/gin-gonic/gin"
)

type CreatePoliceController struct {
	useCase *application.CreatePoliceUseCase
}

func NewCreatePoliceController(useCase *application.CreatePoliceUseCase) *CreatePoliceController {
	return &CreatePoliceController{useCase: useCase}
}

func (c *CreatePoliceController) Handle(ctx *gin.Context) {
	var police entities.Police
	if err := ctx.ShouldBindJSON(&police); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	if err := c.useCase.Execute(ctx.Request.Context(), &police); err != nil {
		status := http.StatusInternalServerError
		switch err.Error() {
		case "name and paternal lastname are required", "CUIP is required", "invalid sex, must be M or F":
			status = http.StatusBadRequest
		case "police with this CUIP already exists":
			status = http.StatusConflict
		}
		utils.ErrorResponse(ctx, status, err.Error(), nil)
		return
	}

	utils.SuccessResponse(ctx, http.StatusCreated, "Police created successfully", police)
}