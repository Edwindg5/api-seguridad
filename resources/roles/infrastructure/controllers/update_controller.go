//api-seguridad/resources/roles/infrastructure/controllers/update_controller.go
package controllers

import (
	"net/http"
	"strconv"
	"api-seguridad/core/utils"
	"api-seguridad/resources/roles/application"
	"api-seguridad/resources/roles/domain/entities"
	"github.com/gin-gonic/gin"
)

type UpdateRoleController struct {
	useCase *application.UpdateRoleUseCase
}

func NewUpdateRoleController(useCase *application.UpdateRoleUseCase) *UpdateRoleController {
	return &UpdateRoleController{useCase: useCase}
}

func (c *UpdateRoleController) Handle(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid role ID", err)
		return
	}

	var role entities.Role
	if err := ctx.ShouldBindJSON(&role); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	role.ID = uint(id)

	updaterID, _ := ctx.Get("userID")
	if updaterID != nil {
		if uid, ok := updaterID.(uint); ok {
			role.UpdatedBy = uid
		}
	}

	if err := c.useCase.Execute(ctx.Request.Context(), &role); err != nil {
		status := http.StatusInternalServerError
		switch err.Error() {
		case "role not found":
			status = http.StatusNotFound
		case "role with this title already exists":
			status = http.StatusConflict
		case "role title is required":
			status = http.StatusBadRequest
		}
		utils.ErrorResponse(ctx, status, "Failed to update role", err)
		return
	}

	updatedRole, err := c.useCase.GetRepository().GetByID(ctx.Request.Context(), role.ID)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to get updated role", err)
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Role updated successfully", updatedRole)
}
