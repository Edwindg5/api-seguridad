package controllers

import (
	"net/http"
	"strconv"
	"api-seguridad/core/utils"
	"api-seguridad/resources/role_permissions/application"
	"github.com/gin-gonic/gin"
)

type GetRolePermissionByIDController struct {
	useCase *application.GetRolePermissionByIDUseCase
}

func NewGetRolePermissionByIDController(useCase *application.GetRolePermissionByIDUseCase) *GetRolePermissionByIDController {
	return &GetRolePermissionByIDController{useCase: useCase}
}

func (c *GetRolePermissionByIDController) Handle(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "ID inválido", err)
		return
	}

	rolePermission, err := c.useCase.Execute(ctx.Request.Context(), uint(id))
	if err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "relación rol-permiso no encontrada" {
			status = http.StatusNotFound
		}
		utils.ErrorResponse(ctx, status, err.Error(), nil)
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Relación rol-permiso obtenida", rolePermission)
}