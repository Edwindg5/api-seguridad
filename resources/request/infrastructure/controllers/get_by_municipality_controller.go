// api-seguridad/resources/request/infrastructure/controllers/get_by_municipality_controller.go
package controllers

import (
	"net/http"
	"strconv"
	"api-seguridad/core/utils"
	"api-seguridad/resources/request/application"
	"github.com/gin-gonic/gin"
)

type GetRequestsByMunicipalityController struct {
	useCase *application.GetRequestsByMunicipalityUseCase
}

func NewGetRequestsByMunicipalityController(useCase *application.GetRequestsByMunicipalityUseCase) *GetRequestsByMunicipalityController {
	return &GetRequestsByMunicipalityController{useCase: useCase}
}

func (c *GetRequestsByMunicipalityController) Handle(ctx *gin.Context) {
	municipalityID, err := strconv.ParseUint(ctx.Query("municipality_id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Municipality ID parameter is required", nil)
		return
	}

	requests, err := c.useCase.Execute(ctx.Request.Context(), uint(municipalityID))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to get requests by municipality", err)
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Requests retrieved successfully", requests)
}