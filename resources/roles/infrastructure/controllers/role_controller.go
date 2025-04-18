//api-seguridad/resources/roles/infrastructure/controllers/role_controller.go
package controllers

import (
	"net/http"
	"api-seguridad/core/utils"
	"api-seguridad/resources/roles/application"
	"api-seguridad/resources/roles/domain/entities"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RoleController struct {
	roleService *application.RoleService
}

func NewRoleController(roleService *application.RoleService) *RoleController {
	return &RoleController{roleService: roleService}
}

func (c *RoleController) CreateRole(ctx *gin.Context) {
	var role entity.Role
	if err := ctx.ShouldBindJSON(&role); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	if err := c.roleService.CreateRole(ctx.Request.Context(), &role); err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to create role", err)
		return
	}

	utils.SuccessResponse(ctx, http.StatusCreated, "Role created successfully", role)
}

func (c *RoleController) GetRole(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid role ID", err)
		return
	}

	role, err := c.roleService.GetRoleByID(ctx.Request.Context(), uint(id))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to get role", err)
		return
	}

	if role == nil {
		utils.ErrorResponse(ctx, http.StatusNotFound, "Role not found", nil)
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Role retrieved successfully", role)
}