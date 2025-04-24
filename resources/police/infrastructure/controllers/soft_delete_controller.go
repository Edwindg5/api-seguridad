//api-seguridad/resources/police/infrastructure/controllers/soft_delete_controller.go
package controllers

import (
	"net/http"
	"strconv"
	"api-seguridad/core/utils"
	"api-seguridad/resources/police/application"
	"github.com/gin-gonic/gin"
)

type SoftDeletePoliceController struct {
	useCase *application.SoftDeletePoliceUseCase
}

func NewSoftDeletePoliceController(useCase *application.SoftDeletePoliceUseCase) *SoftDeletePoliceController {
	return &SoftDeletePoliceController{useCase: useCase}
}

func (c *SoftDeletePoliceController) Handle(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid police ID", err)
		return
	}

	if err := c.useCase.Execute(ctx.Request.Context(), uint(id)); err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "police not found" {
			status = http.StatusNotFound
		}
		utils.ErrorResponse(ctx, status, err.Error(), nil)
		return
	}

	utils.SuccessResponse(ctx, http.StatusNoContent, "Police deleted successfully", nil)
}