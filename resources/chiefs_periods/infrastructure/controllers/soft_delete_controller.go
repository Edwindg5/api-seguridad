package controllers

import (
	"net/http"
	"strconv"
	"api-seguridad/core/utils"
	"api-seguridad/resources/chiefs_periods/application"
	"github.com/gin-gonic/gin"
)

type SoftDeleteChiefsPeriodController struct {
	useCase *application.SoftDeleteChiefsPeriodUseCase
}

func NewSoftDeleteChiefsPeriodController(useCase *application.SoftDeleteChiefsPeriodUseCase) *SoftDeleteChiefsPeriodController {
	return &SoftDeleteChiefsPeriodController{useCase: useCase}
}

func (c *SoftDeleteChiefsPeriodController) Handle(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid period ID", err)
		return
	}

	if err := c.useCase.Execute(ctx.Request.Context(), uint(id)); err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "period not found" {
			status = http.StatusNotFound
		}
		utils.ErrorResponse(ctx, status, err.Error(), nil)
		return
	}

	utils.SuccessResponse(ctx, http.StatusNoContent, "Chiefs period deleted successfully", nil)
}