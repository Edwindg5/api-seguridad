//api-seguridad/resources/municipalities/infrastructure/controllers/get_by_id_controller.go
package controllers

import (
	"net/http"
	"strconv"
	"api-seguridad/core/utils"
	"api-seguridad/resources/municipalities/application"
	"github.com/gin-gonic/gin"
)

type GetMunicipalityByIDController struct {
	useCase *application.GetMunicipalityByIDUseCase
}

func NewGetMunicipalityByIDController(useCase *application.GetMunicipalityByIDUseCase) *GetMunicipalityByIDController {
	return &GetMunicipalityByIDController{useCase: useCase}
}

func (c *GetMunicipalityByIDController) Handle(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid municipality ID", err)
		return
	}

	municipality, err := c.useCase.Execute(ctx.Request.Context(), uint(id))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to get municipality", err)
		return
	}

	if municipality == nil {
		utils.ErrorResponse(ctx, http.StatusNotFound, "Municipality not found", nil)
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Municipality retrieved successfully", municipality)
}