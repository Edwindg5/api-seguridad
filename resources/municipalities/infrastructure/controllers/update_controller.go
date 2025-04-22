package controllers

import (
	"net/http"
	"api-seguridad/core/utils"
	"api-seguridad/resources/municipalities/application"
	"api-seguridad/resources/municipalities/domain/entities"
	"github.com/gin-gonic/gin"
)

type UpdateMunicipalityController struct {
	useCase *application.UpdateMunicipalityUseCase
}

func NewUpdateMunicipalityController(useCase *application.UpdateMunicipalityUseCase) *UpdateMunicipalityController {
	return &UpdateMunicipalityController{useCase: useCase}
}

func (c *UpdateMunicipalityController) Handle(ctx *gin.Context) {
	var municipality entities.Municipality
	if err := ctx.ShouldBindJSON(&municipality); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	if err := c.useCase.Execute(ctx.Request.Context(), &municipality); err != nil {
		status := http.StatusInternalServerError
		switch err.Error() {
		case "invalid municipality ID", "municipality name is required":
			status = http.StatusBadRequest
		case "municipality not found":
			status = http.StatusNotFound
		case "municipality with this name already exists":
			status = http.StatusConflict
		}
		utils.ErrorResponse(ctx, status, err.Error(), nil)
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Municipality updated successfully", municipality)
}