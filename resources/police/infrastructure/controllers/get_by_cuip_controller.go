//api-seguridad/resources/police/infrastructure/controllers/get_by_cuip_controller.go
package controllers

import (
	"net/http"
	"api-seguridad/core/utils"
	"api-seguridad/resources/police/application"
	"github.com/gin-gonic/gin"
)

type GetPoliceByCUIPController struct {
	useCase *application.GetPoliceByCUIPUseCase
}

func NewGetPoliceByCUIPController(useCase *application.GetPoliceByCUIPUseCase) *GetPoliceByCUIPController {
	return &GetPoliceByCUIPController{useCase: useCase}
}

func (c *GetPoliceByCUIPController) Handle(ctx *gin.Context) {
	cuip := ctx.Query("cuip")
	if cuip == "" {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "CUIP parameter is required", nil)
		return
	}

	police, err := c.useCase.Execute(ctx.Request.Context(), cuip)
	if err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "police not found" {
			status = http.StatusNotFound
		}
		utils.ErrorResponse(ctx, status, err.Error(), nil)
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Police retrieved successfully", police)
}