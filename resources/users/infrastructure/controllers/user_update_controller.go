// api-seguridad/resources/users/infrastructure/controllers/user_update_controller.go
package controllers

import (
	"errors"
	"net/http"
	"strconv"
	"api-seguridad/resources/users/application"
	"api-seguridad/resources/users/domain/entities"
	"api-seguridad/core/utils"
	"time"

	"github.com/gin-gonic/gin"
)

type UserUpdateController struct {
	updateUC *application.UpdateUserUseCase
}

func NewUserUpdateController(updateUC *application.UpdateUserUseCase) *UserUpdateController {
	return &UserUpdateController{updateUC: updateUC}
}

func (c *UserUpdateController) Handle(ctx *gin.Context) {
    id, err := strconv.ParseUint(ctx.Param("id_user"), 10, 64)
    if err != nil {
        utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid user ID", 
            errors.New("el ID de usuario debe ser un número válido"))
        return
    }

    var user entities.User
    if err := ctx.ShouldBindJSON(&user); err != nil {
        utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request payload", err)
        return
    }

    // Validar que el rol_id_fk sea válido (debe ser > 0)
    if user.RoleID <= 0 {
        utils.ErrorResponse(ctx, http.StatusBadRequest, "Role ID is required and must be valid", 
            errors.New("el ID de rol debe ser un número válido mayor a 0"))
        return
    }

    // Asignar ID y valores por defecto
    user.ID = uint(id)
    user.UpdatedAt = time.Now()
    user.UpdatedBy = 1 // Usuario admin por defecto

    if err := c.updateUC.Execute(ctx.Request.Context(), &user); err != nil {
        statusCode := http.StatusInternalServerError
        if err.Error() == "user not found" || 
           err.Error() == "new username already exists" ||
           err.Error() == "new email already exists" ||
           err.Error() == "role not found" {
            statusCode = http.StatusBadRequest
        }
        utils.ErrorResponse(ctx, statusCode, "Failed to update user", err)
        return
    }

    utils.SuccessResponse(ctx, http.StatusOK, "User updated successfully", user)
}