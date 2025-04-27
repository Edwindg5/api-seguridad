//api-seguridad/resources/role_permissions/infrastructure/controllers/update_controller.go
package controllers

import (
	"net/http"
	"strconv"
	"api-seguridad/core/utils"
	"api-seguridad/resources/role_permissions/application"
	"api-seguridad/resources/role_permissions/domain/entities"
	"github.com/gin-gonic/gin"
)

type UpdateRolePermissionController struct {
	useCase *application.UpdateRolePermissionUseCase
}

func NewUpdateRolePermissionController(useCase *application.UpdateRolePermissionUseCase) *UpdateRolePermissionController {
	return &UpdateRolePermissionController{useCase: useCase}
}

func (c *UpdateRolePermissionController) Handle(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "ID inválido", err)
		return
	}

	var rolePermission entities.RolePermission
	if err := ctx.ShouldBindJSON(&rolePermission); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Datos inválidos", err)
		return
	}

	rolePermission.ID = uint(id)

	// Obtener ID del usuario que actualiza
	if updaterID, exists := ctx.Get("userID"); exists {
		if uid, ok := updaterID.(uint); ok {
			rolePermission.UpdatedBy = uid
		}
	}

	if err := c.useCase.Execute(ctx.Request.Context(), &rolePermission); err != nil {
		status := http.StatusInternalServerError
		switch err.Error() {
		case "ID inválido", "no se pueden modificar los IDs de rol o permiso":
			status = http.StatusBadRequest
		case "relación rol-permiso no encontrada":
			status = http.StatusNotFound
		}
		utils.ErrorResponse(ctx, status, err.Error(), nil)
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Relación rol-permiso actualizada", rolePermission)
}