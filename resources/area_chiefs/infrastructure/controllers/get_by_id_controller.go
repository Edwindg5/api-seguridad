// api-seguridad/resources/area_chiefs/infrastructure/controllers/get_by_id_controller.go
package controllers

import (
	"net/http"
	"strconv"
	"api-seguridad/core/utils"
	"api-seguridad/resources/area_chiefs/application"
	"github.com/gin-gonic/gin"
)

type GetAreaChiefByIDController struct {
	useCase *application.GetAreaChiefByIDUseCase
}

func NewGetAreaChiefByIDController(useCase *application.GetAreaChiefByIDUseCase) *GetAreaChiefByIDController {
	return &GetAreaChiefByIDController{useCase: useCase}
}

func (c *GetAreaChiefByIDController) Handle(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid chief ID", err)
		return
	}

	chief, err := c.useCase.Execute(ctx.Request.Context(), uint(id))
	if err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "area chief not found" {
			status = http.StatusNotFound
		}
		utils.ErrorResponse(ctx, status, err.Error(), nil)
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Area chief retrieved successfully", chief)
}