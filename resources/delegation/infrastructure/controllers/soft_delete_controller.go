package controllers

import (
	"net/http"
	"strconv"
	"api-seguridad/core/utils"
	"api-seguridad/resources/delegation/application"
	"github.com/gin-gonic/gin"
)

type SoftDeleteDelegationController struct {
	useCase *application.SoftDeleteDelegationUseCase
}

func NewSoftDeleteDelegationController(useCase *application.SoftDeleteDelegationUseCase) *SoftDeleteDelegationController {
	return &SoftDeleteDelegationController{useCase: useCase}
}

func (c *SoftDeleteDelegationController) Handle(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid delegation ID", err)
		return
	}

	if err := c.useCase.Execute(ctx.Request.Context(), uint(id)); err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "invalid delegation ID" {
			status = http.StatusBadRequest
		} else if err.Error() == "delegation not found" {
			status = http.StatusNotFound
		}
		utils.ErrorResponse(ctx, status, err.Error(), nil)
		return
	}

	utils.SuccessResponse(ctx, http.StatusNoContent, "Delegation deleted successfully", nil)
}