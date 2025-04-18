//api-seguridad/resources/police/infrastructure/controllers/police_controller.go
package controllers

import (
	"net/http"
	"api-seguridad/core/utils"
	"api-seguridad/resources/police/application"
	"api-seguridad/resources/police/domain/entities"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PoliceController struct {
	policeService *application.PoliceService
}

func NewPoliceController(policeService *application.PoliceService) *PoliceController {
	return &PoliceController{policeService: policeService}
}

func (c *PoliceController) CreatePolice(ctx *gin.Context) {
	var police entity.Police
	if err := ctx.ShouldBindJSON(&police); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	if err := c.policeService.CreatePolice(ctx.Request.Context(), &police); err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to create police", err)
		return
	}

	utils.SuccessResponse(ctx, http.StatusCreated, "Police created successfully", police)
}

func (c *PoliceController) GetPolice(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid police ID", err)
		return
	}

	police, err := c.policeService.GetPoliceByID(ctx.Request.Context(), uint(id))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to get police", err)
		return
	}

	if police == nil {
		utils.ErrorResponse(ctx, http.StatusNotFound, "Police not found", nil)
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Police retrieved successfully", police)
}