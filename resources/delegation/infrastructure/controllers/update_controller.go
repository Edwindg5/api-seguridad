package controllers

import (
	"net/http"
	"api-seguridad/core/utils"
	"api-seguridad/resources/delegation/application"
	"api-seguridad/resources/delegation/domain/entities"
	"github.com/gin-gonic/gin"
)

type UpdateDelegationController struct {
	useCase *application.UpdateDelegationUseCase
}

func NewUpdateDelegationController(useCase *application.UpdateDelegationUseCase) *UpdateDelegationController {
	return &UpdateDelegationController{useCase: useCase}
}

func (c *UpdateDelegationController) Handle(ctx *gin.Context) {
	var delegation entities.Delegation
	if err := ctx.ShouldBindJSON(&delegation); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	if err := c.useCase.Execute(ctx.Request.Context(), &delegation); err != nil {
		status := http.StatusInternalServerError
		switch err.Error() {
		case "invalid delegation ID", "delegation name is required":
			status = http.StatusBadRequest
		case "delegation not found":
			status = http.StatusNotFound
		}
		utils.ErrorResponse(ctx, status, err.Error(), nil)
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Delegation updated successfully", delegation)
}