//api-seguridad/resources/police/infrastructure/controllers/get_by_id_controller.go
package controllers

import (
	"net/http"
	"strconv"
	"api-seguridad/core/utils"
	"api-seguridad/resources/police/application"
	"github.com/gin-gonic/gin"
)

type GetPoliceByIDController struct {
	useCase *application.GetPoliceByIDUseCase
}

func NewGetPoliceByIDController(useCase *application.GetPoliceByIDUseCase) *GetPoliceByIDController {
	return &GetPoliceByIDController{useCase: useCase}
}

func (c *GetPoliceByIDController) Handle(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid police ID", err)
		return
	}

	police, err := c.useCase.Execute(ctx.Request.Context(), uint(id))
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