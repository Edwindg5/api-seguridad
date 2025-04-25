package controllers

import (
	"net/http"
	"strconv"
	"api-seguridad/core/utils"
	"api-seguridad/resources/role_permissions/application"
	"github.com/gin-gonic/gin"
)

type DeleteRolePermissionController struct {
	useCase *application.DeleteRolePermissionUseCase
}

func NewDeleteRolePermissionController(useCase *application.DeleteRolePermissionUseCase) *DeleteRolePermissionController {
	return &DeleteRolePermissionController{useCase: useCase}
}

func (c *DeleteRolePermissionController) Handle(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "ID inválido", err)
		return
	}

	if err := c.useCase.Execute(ctx.Request.Context(), uint(id)); err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "relación rol-permiso no encontrada" {
			status = http.StatusNotFound
		}
		utils.ErrorResponse(ctx, status, err.Error(), nil)
		return
	}

	utils.SuccessResponse(ctx, http.StatusNoContent, "Relación rol-permiso eliminada", nil)
}