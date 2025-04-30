// api-seguridad/resource/delegation/infrastructure/controllers/update_controller.go
package controllers

import (
	"api-seguridad/core/utils"
	"api-seguridad/resources/delegation/application"

	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type UpdateDelegationController struct {
	useCase *application.UpdateDelegationUseCase
}

func NewUpdateDelegationController(useCase *application.UpdateDelegationUseCase) *UpdateDelegationController {
	return &UpdateDelegationController{useCase: useCase}
}

func (c *UpdateDelegationController) Handle(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid delegation ID", nil)
		return
	}

	existingDelegation, err := c.useCase.GetExistingDelegation(ctx.Request.Context(), uint(id))
	if err != nil || existingDelegation == nil {
		utils.ErrorResponse(ctx, http.StatusNotFound, "Delegation not found", nil)
		return
	}

	var updateData struct {
		Name   string `json:"name"`
		Active bool   `json:"active"`
	}
	if err := ctx.ShouldBindJSON(&updateData); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	existingDelegation.SetName(updateData.Name)
	existingDelegation.SetActive(updateData.Active)

	// Set fixed UpdatedBy and timestamp
	existingDelegation.SetUpdatedBy(1)
	existingDelegation.SetUpdatedAt(time.Now())

	if err := c.useCase.Execute(ctx.Request.Context(), existingDelegation); err != nil {
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

	utils.SuccessResponse(ctx, http.StatusOK, "Delegation updated successfully", existingDelegation)
}
