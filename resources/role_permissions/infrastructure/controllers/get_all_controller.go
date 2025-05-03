//api-seguridad/resources/role_permissions/infrastructure/controllers/get_all_controller.go
package controllers

import (
	"net/http"
	"api-seguridad/core/utils"
	"api-seguridad/resources/role_permissions/application"
	"github.com/gin-gonic/gin"
)

type GetAllRolePermissionsController struct {
	useCase *application.GetAllUseCase
}

func NewGetAllRolePermissionsController(useCase *application.GetAllUseCase) *GetAllRolePermissionsController {
	return &GetAllRolePermissionsController{useCase: useCase}
}

func (c *GetAllRolePermissionsController) Handle(ctx *gin.Context) {
	// Ejecutar el caso de uso
	rolePermissions, err := c.useCase.Execute(ctx.Request.Context())
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Error al obtener relaciones rol-permiso", err)
		return
	}

	// Verificar si hay resultados
	if len(rolePermissions) == 0 {
		utils.SuccessResponse(ctx, http.StatusOK, "No se encontraron relaciones rol-permiso", []interface{}{})
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Relaciones rol-permiso obtenidas exitosamente", rolePermissions)
}