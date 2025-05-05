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

    var request struct {
        FirstName string `json:"first_name"`
        LastName  string `json:"last_name"`  // Mapea a 'lastname' en la entidad
        Username  string `json:"username"`
        Email     string `json:"email"`
        RoleID    uint   `json:"rol_id_fk"`  // Cambiado a rol_id_fk para coincidir con el frontend
    }

    if err := ctx.ShouldBindJSON(&request); err != nil {
        utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request payload", err)
        return
    }

    // Validación del RoleID
    if request.RoleID == 0 {
        utils.ErrorResponse(ctx, http.StatusBadRequest, "Role ID is required", 
            errors.New("el ID de rol es requerido"))
        return
    }

    user := &entities.User{
        ID:        uint(id),
        FirstName: request.FirstName,
        LastName:  request.LastName, // Se mapeará a 'lastname' en la base de datos
        Username:  request.Username,
        Email:     request.Email,
        RoleID:    request.RoleID, // Se mapeará a 'rol_id_fk' en la base de datos
        UpdatedAt: time.Now(),
        UpdatedBy: 1, // Usuario admin por defecto
    }

    if err := c.updateUC.Execute(ctx.Request.Context(), user); err != nil {
        statusCode := http.StatusInternalServerError
        errorMessage := "Failed to update user"
        
        switch err.Error() {
        case "user not found":
            statusCode = http.StatusNotFound
            errorMessage = "User not found"
        case "new username already exists":
            statusCode = http.StatusConflict
            errorMessage = "Username already in use"
        case "new email already exists":
            statusCode = http.StatusConflict
            errorMessage = "Email already in use"
        case "role not found":
            statusCode = http.StatusBadRequest
            errorMessage = "Invalid role ID"
        case "role ID must be provided":
            statusCode = http.StatusBadRequest
            errorMessage = "Role ID is required"
        }

        utils.ErrorResponse(ctx, statusCode, errorMessage, err)
        return
    }

    utils.SuccessResponse(ctx, http.StatusOK, "User updated successfully", user)
}