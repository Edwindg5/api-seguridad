// api-seguridad/resources/users/infrastructure/controllers/user_update_controller.go
package controllers

import (
	"api-seguridad/core/utils"
	"api-seguridad/resources/users/application"
	"api-seguridad/resources/users/domain/entities"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
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

    // Estructura que coincide exactamente con lo que envía el frontend
    var request struct {
        FirstName string `json:"first_name"`
        LastName  string `json:"lastname"`  // Cambiado a lastname para coincidir
        Username  string `json:"username"`
        Email     string `json:"email"`
        RoleID    uint   `json:"rol_id_fk"`  // Nombre exacto del campo
    }

    if err := ctx.ShouldBindJSON(&request); err != nil {
        utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request payload", err)
        return
    }

    // Validación más detallada del RoleID
    if request.RoleID == 0 {
        utils.ErrorResponse(ctx, http.StatusBadRequest, "Role ID is required and must be greater than 0", 
            errors.New("el ID de rol es requerido y debe ser mayor que 0"))
        return
    }

    user := &entities.User{
        ID:        uint(id),
        FirstName: request.FirstName,
        LastName:  request.LastName,
        Username:  request.Username,
        Email:     request.Email,
        RoleID:    request.RoleID,
        UpdatedAt: time.Now(),
        UpdatedBy: 1,
    }

    // Debug: Mostrar datos recibidos
    fmt.Printf("Datos recibidos para actualización: %+v\n", user)

    if err := c.updateUC.Execute(ctx.Request.Context(), user); err != nil {
        // Manejo de errores mejorado
        statusCode := http.StatusInternalServerError
        errorMessage := err.Error()
        
        switch {
        case strings.Contains(err.Error(), "user not found"):
            statusCode = http.StatusNotFound
        case strings.Contains(err.Error(), "already exists"):
            statusCode = http.StatusConflict
        case strings.Contains(err.Error(), "role"):
            statusCode = http.StatusBadRequest
        }

        utils.ErrorResponse(ctx, statusCode, errorMessage, err)
        return
    }

    utils.SuccessResponse(ctx, http.StatusOK, "User updated successfully", user)
}