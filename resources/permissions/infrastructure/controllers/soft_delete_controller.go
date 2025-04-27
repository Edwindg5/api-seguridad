//api-seguridad/resources/permissions/infrastructure/controllers/soft_delete_controller.go
package controllers

import (
	"net/http"
	"strconv"
	"api-seguridad/core/utils"
	"api-seguridad/resources/permissions/application"
	"github.com/gin-gonic/gin"
)

type SoftDeletePermissionController struct {
	useCase *application.SoftDeletePermissionUseCase
}

func NewSoftDeletePermissionController(useCase *application.SoftDeletePermissionUseCase) *SoftDeletePermissionController {
	return &SoftDeletePermissionController{useCase: useCase}
}

func (c *SoftDeletePermissionController) Handle(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "ID de permiso inválido", err)
		return
	}

	if err := c.useCase.Execute(ctx.Request.Context(), uint(id)); err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "permission not found" {
			status = http.StatusNotFound
		}
		utils.ErrorResponse(ctx, status, err.Error(), nil)
		return
	}

	utils.SuccessResponse(ctx, http.StatusNoContent, "Permiso eliminado exitosamente", nil)
}