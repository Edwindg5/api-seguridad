//api-seguridad/resources/municipalities/infrastructure/controllers/create_controller.go
package controllers

import (
	"net/http"
	"api-seguridad/core/utils"
	"api-seguridad/resources/municipalities/application"
	"api-seguridad/resources/municipalities/domain/entities"
	"github.com/gin-gonic/gin"
)

type CreateMunicipalityController struct {
	useCase *application.PostMunicipalityUseCase
}

func NewCreateMunicipalityController(useCase *application.PostMunicipalityUseCase) *CreateMunicipalityController {
	return &CreateMunicipalityController{useCase: useCase}
}

func (c *CreateMunicipalityController) Handle(ctx *gin.Context) {
	var municipality entities.Municipality
	if err := ctx.ShouldBindJSON(&municipality); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	// Set created_by from authenticated user
	if userID, exists := ctx.Get("userID"); exists {
		if uid, ok := userID.(uint); ok {
			municipality.CreatedBy = uid
			municipality.UpdatedBy = uid
		}
	}

	if err := c.useCase.Execute(ctx.Request.Context(), &municipality); err != nil {
		status := http.StatusInternalServerError
		switch err.Error() {
		case "municipality name is required":
			status = http.StatusBadRequest
		case "municipality with this name already exists":
			status = http.StatusConflict
		}
		utils.ErrorResponse(ctx, status, err.Error(), nil)
		return
	}

	// Get the created municipality with delegation data
	createdMunicipality, err := c.useCase.GetByID(ctx.Request.Context(), municipality.GetID())
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to retrieve created municipality", err)
		return
	}

	utils.SuccessResponse(ctx, http.StatusCreated, "Municipality created successfully", createdMunicipality)
}