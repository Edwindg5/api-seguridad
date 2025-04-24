// api-seguridad/resources/area_chiefs/infrastructure/controllers/get_all_controller.go
package controllers

import (
	"net/http"
	"api-seguridad/core/utils"
	"api-seguridad/resources/area_chiefs/application"
	"github.com/gin-gonic/gin"
)

type GetAllAreaChiefsController struct {
	useCase *application.GetAllAreaChiefsUseCase
}

func NewGetAllAreaChiefsController(useCase *application.GetAllAreaChiefsUseCase) *GetAllAreaChiefsController {
	return &GetAllAreaChiefsController{useCase: useCase}
}

func (c *GetAllAreaChiefsController) Handle(ctx *gin.Context) {
	chiefs, err := c.useCase.Execute(ctx.Request.Context())
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to get area chiefs", err)
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Area chiefs retrieved successfully", chiefs)
}