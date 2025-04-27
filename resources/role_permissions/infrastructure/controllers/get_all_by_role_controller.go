//api-seguridad/resources/role_permissions/infrastructure/controllers/get_all_by_role_controller.go
package controllers

import (
	"net/http"
	"strconv"
	"api-seguridad/core/utils"
	"api-seguridad/resources/role_permissions/application"
	"github.com/gin-gonic/gin"
)

type GetAllByRoleController struct {
	useCase *application.GetAllByRoleUseCase
}

func NewGetAllByRoleController(useCase *application.GetAllByRoleUseCase) *GetAllByRoleController {
	return &GetAllByRoleController{useCase: useCase}
}

func (c *GetAllByRoleController) Handle(ctx *gin.Context) {
	roleID, err := strconv.ParseUint(ctx.Param("role_id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "ID de rol inválido", err)
		return
	}

	rolePermissions, err := c.useCase.Execute(ctx.Request.Context(), uint(roleID))
	if err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "ID de rol es requerido" {
			status = http.StatusBadRequest
		}
		utils.ErrorResponse(ctx, status, err.Error(), nil)
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Permisos por rol obtenidos", rolePermissions)
}