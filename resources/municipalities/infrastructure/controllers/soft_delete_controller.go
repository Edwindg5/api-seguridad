package controllers

import (
	"net/http"
	"strconv"
	"api-seguridad/core/utils"
	"api-seguridad/resources/municipalities/application"
	"github.com/gin-gonic/gin"
)

type SoftDeleteMunicipalityController struct {
	useCase *application.SoftDeleteMunicipalityUseCase
}

func NewSoftDeleteMunicipalityController(useCase *application.SoftDeleteMunicipalityUseCase) *SoftDeleteMunicipalityController {
	return &SoftDeleteMunicipalityController{useCase: useCase}
}

func (c *SoftDeleteMunicipalityController) Handle(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid municipality ID", err)
		return
	}

	if err := c.useCase.Execute(ctx.Request.Context(), uint(id)); err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "invalid municipality ID" {
			status = http.StatusBadRequest
		} else if err.Error() == "municipality not found" {
			status = http.StatusNotFound
		}
		utils.ErrorResponse(ctx, status, err.Error(), nil)
		return
	}

	utils.SuccessResponse(ctx, http.StatusNoContent, "Municipality deleted successfully", nil)
}