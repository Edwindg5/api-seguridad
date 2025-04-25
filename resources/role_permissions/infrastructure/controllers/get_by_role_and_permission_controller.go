package controllers

import (
	"net/http"
	"strconv"
	"api-seguridad/core/utils"
	"api-seguridad/resources/role_permissions/application"
	"github.com/gin-gonic/gin"
)

type GetByRoleAndPermissionController struct {
	useCase *application.GetByRoleAndPermissionUseCase
}

func NewGetByRoleAndPermissionController(useCase *application.GetByRoleAndPermissionUseCase) *GetByRoleAndPermissionController {
	return &GetByRoleAndPermissionController{useCase: useCase}
}

func (c *GetByRoleAndPermissionController) Handle(ctx *gin.Context) {
	roleID, err := strconv.ParseUint(ctx.Query("role_id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "ID de rol inválido", err)
		return
	}

	permissionID, err := strconv.ParseUint(ctx.Query("permission_id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "ID de permiso inválido", err)
		return
	}

	rolePermission, err := c.useCase.Execute(ctx.Request.Context(), uint(roleID), uint(permissionID))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Error al obtener relación", err)
		return
	}

	if rolePermission == nil {
		utils.SuccessResponse(ctx, http.StatusOK, "No existe relación entre el rol y permiso", nil)
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Relación rol-permiso obtenida", rolePermission)
}