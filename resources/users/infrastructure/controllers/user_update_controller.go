// api-seguridad/resources/users/infrastructure/controllers/user_update_controller.go
package controllers

import (
	"errors"
	"net/http"
	"strconv"
	"api-seguridad/resources/users/application"
	"api-seguridad/resources/users/domain/entities"
	"api-seguridad/core/utils"

	"github.com/gin-gonic/gin"
)

type UserUpdateController struct {
	updateUC *application.UpdateUserUseCase
}

func NewUserUpdateController(updateUC *application.UpdateUserUseCase) *UserUpdateController {
	return &UserUpdateController{updateUC: updateUC}
}

func (c *UserUpdateController) Handle(ctx *gin.Context) {
	// Obtener ID del usuario a actualizar
	id, err := strconv.ParseUint(ctx.Param("id_user"), 10, 64)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid user ID", 
			errors.New("el ID de usuario debe ser un número válido"))
		return
	}

	// Obtener el ID del usuario autenticado
	authUserID, exists := ctx.Get("userID")
	if !exists {
		utils.ErrorResponse(ctx, http.StatusUnauthorized, "Authentication required", nil)
		return
	}

	var user entities.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	// Asignar ID desde el parámetro de la URL
	user.ID = uint(id)
	user.UpdatedBy = authUserID.(uint)

	// Ejecutar actualización
	if err := c.updateUC.Execute(ctx.Request.Context(), &user); err != nil {
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