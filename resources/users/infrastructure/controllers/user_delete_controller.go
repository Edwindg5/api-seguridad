// api-seguridad/resources/users/infrastructure/controllers/user_delete_controller.go
package controllers

import (
	"net/http"
	"strconv"
	"api-seguridad/resources/users/application"
	"api-seguridad/core/utils"

	"github.com/gin-gonic/gin"
)

type UserDeleteController struct {
	deleteUC *application.DeleteUserUseCase
}

func NewUserDeleteController(deleteUC *application.DeleteUserUseCase) *UserDeleteController {
	return &UserDeleteController{deleteUC: deleteUC}
}

func (c *UserDeleteController) Handle(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid user ID", err)
		return
	}

	// Get deleter ID from context (assuming JWT middleware sets it)
	deleterID, _ := ctx.Get("userID")

	if err := c.deleteUC.Execute(ctx.Request.Context(), uint(id), deleterID.(uint)); err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "user not found" || 
		   err.Error() == "user already deleted" {
			statusCode = http.StatusBadRequest
		}
		utils.ErrorResponse(ctx, statusCode, "Failed to delete user", err)
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "User deleted successfully", nil)
}