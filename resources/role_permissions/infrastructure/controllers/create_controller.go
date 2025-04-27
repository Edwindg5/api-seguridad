//api-seguridad/resources/role_permissions/infrastructure/controllers/create_controller.go
package controllers

import (
	"net/http"
	"api-seguridad/core/utils"
	"api-seguridad/resources/role_permissions/application"
	"api-seguridad/resources/role_permissions/domain/entities"
	"github.com/gin-gonic/gin"
)

type CreateRolePermissionController struct {
	useCase *application.CreateRolePermissionUseCase
}

func NewCreateRolePermissionController(useCase *application.CreateRolePermissionUseCase) *CreateRolePermissionController {
	return &CreateRolePermissionController{useCase: useCase}
}

func (c *CreateRolePermissionController) Handle(ctx *gin.Context) {
	var rolePermission entities.RolePermission
	if err := ctx.ShouldBindJSON(&rolePermission); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Datos inválidos", err)
		return
	}

	// Obtener ID del usuario creador
	if creatorID, exists := ctx.Get("userID"); exists {
		if uid, ok := creatorID.(uint); ok {
			rolePermission.CreatedBy = uid
			rolePermission.UpdatedBy = uid
		}
	}

	if err := c.useCase.Execute(ctx.Request.Context(), &rolePermission); err != nil {
		status := http.StatusInternalServerError
		switch err.Error() {
		case "ID de rol es requerido", 
		     "ID de permiso es requerido",
		     "usuario creador es requerido":
			status = http.StatusBadRequest
		case "esta relación rol-permiso ya existe":
			status = http.StatusConflict
		}
		utils.ErrorResponse(ctx, status, err.Error(), nil)
		return
	}

	utils.SuccessResponse(ctx, http.StatusCreated, "Relación rol-permiso creada", rolePermission)
}