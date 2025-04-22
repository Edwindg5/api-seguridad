package controllers

import (
	"net/http"
	"strconv"
	"api-seguridad/core/utils"
	"api-seguridad/resources/type_police/application"
	"github.com/gin-gonic/gin"
)

type GetTypePoliceByIDController struct {
	useCase *application.GetTypePoliceByIDUseCase
}

func NewGetTypePoliceByIDController(useCase *application.GetTypePoliceByIDUseCase) *GetTypePoliceByIDController {
	return &GetTypePoliceByIDController{useCase: useCase}
}

func (c *GetTypePoliceByIDController) Handle(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid type police ID", err)
		return
	}

	typePolice, err := c.useCase.Execute(ctx.Request.Context(), uint(id))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to get type police", err)
		return
	}

	if typePolice == nil {
		utils.ErrorResponse(ctx, http.StatusNotFound, "TypePolice not found", nil)
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "TypePolice retrieved successfully", typePolice)
}