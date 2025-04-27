//api-seguridad/resources/permissions/infrastructure/controllers/update_controller.go
package controllers

import (
	"net/http"
	"strconv"
	"api-seguridad/core/utils"
	"api-seguridad/resources/permissions/application"
	"api-seguridad/resources/permissions/domain/entities"
	"github.com/gin-gonic/gin"
)

type UpdatePermissionController struct {
	useCase *application.UpdatePermissionUseCase
}

func NewUpdatePermissionController(useCase *application.UpdatePermissionUseCase) *UpdatePermissionController {
	return &UpdatePermissionController{useCase: useCase}
}

func (c *UpdatePermissionController) Handle(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "ID de permiso inválido", err)
		return
	}

	var permission entities.Permission
	if err := ctx.ShouldBindJSON(&permission); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Datos de permiso inválidos", err)
		return
	}

	permission.ID = uint(id)

	// Obtener ID del usuario que actualiza
	if updaterID, exists := ctx.Get("userID"); exists {
		if uid, ok := updaterID.(uint); ok {
			permission.UpdatedBy = uid
		}
	}

	if err := c.useCase.Execute(ctx.Request.Context(), &permission); err != nil {
		status := http.StatusInternalServerError
		switch err.Error() {
		case "invalid permission ID", "name is required":
			status = http.StatusBadRequest
		case "permission not found":
			status = http.StatusNotFound
		}
		utils.ErrorResponse(ctx, status, err.Error(), nil)
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Permiso actualizado exitosamente", permission)
}