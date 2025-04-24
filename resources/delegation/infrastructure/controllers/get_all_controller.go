//api-seguridad/resource/delegation/infrastructure/controllers/get_all_controller.go
package controllers

import (
	"net/http"
	"api-seguridad/core/utils"
	"api-seguridad/resources/delegation/application"
	"github.com/gin-gonic/gin"
)

type GetAllDelegationsController struct {
	useCase *application.GetAllDelegationsUseCase
}

func NewGetAllDelegationsController(useCase *application.GetAllDelegationsUseCase) *GetAllDelegationsController {
	return &GetAllDelegationsController{useCase: useCase}
}

func (c *GetAllDelegationsController) Handle(ctx *gin.Context) {
	delegations, err := c.useCase.Execute(ctx.Request.Context())
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to get delegations", err)
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Delegations retrieved successfully", delegations)
}