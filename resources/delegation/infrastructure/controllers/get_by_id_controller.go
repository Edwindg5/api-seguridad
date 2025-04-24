//api-seguridad/resource/delegation/infrastructure/controllers/get_by_id_controller.go
package controllers

import (
	"net/http"
	"strconv"
	"api-seguridad/core/utils"
	"api-seguridad/resources/delegation/application"
	"github.com/gin-gonic/gin"
)

type GetDelegationByIDController struct {
	useCase *application.GetDelegationByIDUseCase
}

func NewGetDelegationByIDController(useCase *application.GetDelegationByIDUseCase) *GetDelegationByIDController {
	return &GetDelegationByIDController{useCase: useCase}
}

func (c *GetDelegationByIDController) Handle(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid delegation ID", err)
		return
	}

	delegation, err := c.useCase.Execute(ctx.Request.Context(), uint(id))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to get delegation", err)
		return
	}

	if delegation == nil {
		utils.ErrorResponse(ctx, http.StatusNotFound, "Delegation not found", nil)
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Delegation retrieved successfully", delegation)
}