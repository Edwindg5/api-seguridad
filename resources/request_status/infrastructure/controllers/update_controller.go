// api-seguridad/resources/request_status/infrastructure/controllers/update_controller.go
package controllers

import (
	"api-seguridad/core/utils"
	"api-seguridad/resources/request_status/application"
	"api-seguridad/resources/users/domain/repository"
	userRepo "api-seguridad/resources/users/domain/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateRequestStatusController struct {
	useCase     *application.UpdateRequestStatusUseCase
	userRepo    userRepo.UserRepository
}

func NewUpdateRequestStatusController(
	useCase *application.UpdateRequestStatusUseCase,
	userRepo repository.UserRepository,
) *UpdateRequestStatusController {
	return &UpdateRequestStatusController{
		useCase:     useCase,
		userRepo:    userRepo,
	}
}

func (c *UpdateRequestStatusController) Handle(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid status ID", err)
		return
	}

	existingStatus, err := c.useCase.GetByID(ctx.Request.Context(), uint(id))
	if err != nil || existingStatus == nil {
		utils.ErrorResponse(ctx, http.StatusNotFound, "Status not found", err)
		return
	}

	var updateData struct {
		Name        string `json:"name" binding:"required,min=3,max=100"`
		Description string `json:"description" binding:"max=255"`
	}
	if err := ctx.ShouldBindJSON(&updateData); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	// Get updater user from context
	updaterID, exists := ctx.Get("userID")
	if !exists {
		utils.ErrorResponse(ctx, http.StatusUnauthorized, "User not authenticated", nil)
		return
	}

	uid, ok := updaterID.(uint)
	if !ok {
		utils.ErrorResponse(ctx, http.StatusUnauthorized, "Invalid user ID", nil)
		return
	}

	// Verify user exists
	userExists, err := c.userRepo.Exists(ctx.Request.Context(), uid)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Error verifying user", err)
		return
	}
	if !userExists {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Updater user not found", nil)
		return
	}

	existingStatus.Name = updateData.Name
	existingStatus.Description = updateData.Description
	existingStatus.UpdatedBy = uid

	if err := c.useCase.Execute(ctx.Request.Context(), existingStatus); err != nil {
		status := http.StatusInternalServerError
		switch err.Error() {
		case "invalid status ID", "status name is required", "updater user is required":
			status = http.StatusBadRequest
		case "status not found", "status with this name already exists":
			status = http.StatusConflict
		}
		utils.ErrorResponse(ctx, status, err.Error(), nil)
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Request status updated successfully", existingStatus)
}