//api-seguridad/resources/municipalities/infrastructure/controllers/municipality_controller.go
package controllers

import (
	"net/http"
	"api-seguridad/core/utils"
	"api-seguridad/resources/municipalities/application"
	"api-seguridad/resources/municipalities/domain/entities"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MunicipalityController struct {
	municipalityService *application.MunicipalityService
}

func NewMunicipalityController(municipalityService *application.MunicipalityService) *MunicipalityController {
	return &MunicipalityController{municipalityService: municipalityService}
}

func (c *MunicipalityController) CreateMunicipality(ctx *gin.Context) {
	var municipality entity.Municipality
	if err := ctx.ShouldBindJSON(&municipality); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	if err := c.municipalityService.CreateMunicipality(ctx.Request.Context(), &municipality); err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to create municipality", err)
		return
	}

	utils.SuccessResponse(ctx, http.StatusCreated, "Municipality created successfully", municipality)
}

func (c *MunicipalityController) GetMunicipality(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid municipality ID", err)
		return
	}

	municipality, err := c.municipalityService.GetMunicipalityByID(ctx.Request.Context(), uint(id))
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