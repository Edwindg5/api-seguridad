//api-seguridad/resources/delegation/infrastructure/controllers/create_controller.go
package controllers

import (
	"net/http"
	"api-seguridad/core/utils"
	"api-seguridad/resources/delegation/application"
	"api-seguridad/resources/delegation/domain/entities"
	"github.com/gin-gonic/gin"
)

type CreateDelegationController struct {
	useCase *application.CreateDelegationUseCase
}

func NewCreateDelegationController(useCase *application.CreateDelegationUseCase) *CreateDelegationController {
	return &CreateDelegationController{useCase: useCase}
}

func (c *CreateDelegationController) Handle(ctx *gin.Context) {
	var delegation entities.Delegation
	if err := ctx.ShouldBindJSON(&delegation); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	if err := c.useCase.Execute(ctx.Request.Context(), &delegation); err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "delegation name is required" || 
		   err.Error() == "municipality ID is required" {
			status = http.StatusBadRequest
		}
		utils.ErrorResponse(ctx, status, err.Error(), nil)
		return
	}

	utils.SuccessResponse(ctx, http.StatusCreated, "Delegation created successfully", delegation)
}