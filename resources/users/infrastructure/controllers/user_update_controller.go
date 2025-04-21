// api-seguridad/resources/users/infrastructure/controllers/user_update_controller.go
package controllers

import (
	"net/http"
	"api-seguridad/resources/users/application"
	"api-seguridad/resources/users/domain/entities"
	"api-seguridad/core/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserUpdateController struct {
	updateUC *application.UpdateUserUseCase
}

func NewUserUpdateController(updateUC *application.UpdateUserUseCase) *UserUpdateController {
	return &UserUpdateController{updateUC: updateUC}
}

func (c *UserUpdateController) Handle(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid user ID", err)
		return
	}

	var user entities.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	// Set the ID from URL param
	user.ID = uint(id)

	// Get updater ID from context (assuming JWT middleware sets it)
	updaterID, _ := ctx.Get("userID")

	if err := c.updateUC.Execute(ctx.Request.Context(), &user, updaterID.(uint)); err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "user not found" || 
		   err.Error() == "new username already exists" ||
		   err.Error() == "new email already exists" {
			statusCode = http.StatusBadRequest
		}
		utils.ErrorResponse(ctx, statusCode, "Failed to update user", err)
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "User updated successfully", user)
}