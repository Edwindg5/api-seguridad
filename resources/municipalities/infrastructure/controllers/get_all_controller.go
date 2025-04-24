//api-seguridad/resources/municipalities/infrastructure/controllers/get_all_controller.go
package controllers

import (
	"net/http"
	"api-seguridad/core/utils"
	"api-seguridad/resources/municipalities/application"
	"github.com/gin-gonic/gin"
)

type GetAllMunicipalitiesController struct {
	useCase *application.GetAllMunicipalitiesUseCase
}

func NewGetAllMunicipalitiesController(useCase *application.GetAllMunicipalitiesUseCase) *GetAllMunicipalitiesController {
	return &GetAllMunicipalitiesController{useCase: useCase}
}

func (c *GetAllMunicipalitiesController) Handle(ctx *gin.Context) {
	municipalities, err := c.useCase.Execute(ctx.Request.Context())
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to get municipalities", err)
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Municipalities retrieved successfully", municipalities)
}