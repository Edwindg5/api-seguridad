//api-seguridad/resources/chiefs_periods/infrastructure/controllers/get_all_controller.go
package controllers

import (
	"net/http"
	"api-seguridad/core/utils"
	"api-seguridad/resources/chiefs_periods/application"
	"github.com/gin-gonic/gin"
)

type GetAllChiefsPeriodsController struct {
	useCase *application.GetAllChiefsPeriodsUseCase
}

func NewGetAllChiefsPeriodsController(useCase *application.GetAllChiefsPeriodsUseCase) *GetAllChiefsPeriodsController {
	return &GetAllChiefsPeriodsController{useCase: useCase}
}

func (c *GetAllChiefsPeriodsController) Handle(ctx *gin.Context) {
	periods, err := c.useCase.Execute(ctx.Request.Context())
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to get chiefs periods", err)
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Chiefs periods retrieved successfully", periods)
}