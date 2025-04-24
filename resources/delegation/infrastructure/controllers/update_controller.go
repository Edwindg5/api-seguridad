//api-seguridad/resource/delegation/infrastructure/controllers/update_controller.go
package controllers

import (
	"net/http"
	"strconv"
	"api-seguridad/core/utils"
	"api-seguridad/resources/delegation/application"
	"api-seguridad/resources/delegation/domain/entities"
	"github.com/gin-gonic/gin"
)

type UpdateDelegationController struct {
	useCase *application.UpdateDelegationUseCase
}

func NewUpdateDelegationController(useCase *application.UpdateDelegationUseCase) *UpdateDelegationController {
	return &UpdateDelegationController{useCase: useCase}
}

func (c *UpdateDelegationController) Handle(ctx *gin.Context) {
	// Obtener el ID de la URL
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid delegation ID", nil)
		return
	}

	var delegation entities.Delegation
	if err := ctx.ShouldBindJSON(&delegation); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	// Asignar el ID de la URL al objeto delegation
	delegation.ID = uint(id)

	// Obtener ID del usuario que realiza la actualización
	if updaterID, exists := ctx.Get("userID"); exists {
		if uid, ok := updaterID.(uint); ok {
			delegation.UpdatedBy = uid
		}
	}

	if err := c.useCase.Execute(ctx.Request.Context(), &delegation); err != nil {
		status := http.StatusInternalServerError
		switch err.Error() {
		case "invalid delegation ID", "delegation name is required":
			status = http.StatusBadRequest
		case "delegation not found":
			status = http.StatusNotFound
		}
		utils.ErrorResponse(ctx, status, err.Error(), nil)
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Delegation updated successfully", delegation)
}