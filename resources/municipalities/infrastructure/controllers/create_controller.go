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

	if err := c.useCase.Execute(ctx.Request.Context(), &municipality); err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "municipality name is required" {
			status = http.StatusBadRequest
		} else if err.Error() == "municipality with this name already exists" {
			status = http.StatusConflict
		}
		utils.ErrorResponse(ctx, status, err.Error(), nil)
		return
	}

	utils.SuccessResponse(ctx, http.StatusCreated, "Municipality created successfully", municipality)
}