// api-seguridad/resources/users/infrastructure/controllers/user_delete_controller.go
package controllers

import (
	"api-seguridad/core/utils"
	"api-seguridad/resources/users/application"
	"api-seguridad/resources/users/domain/entities"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type UserDeleteController struct {
	deleteUC *application.DeleteUserUseCase
}

func NewUserDeleteController(deleteUC *application.DeleteUserUseCase) *UserDeleteController {
	return &UserDeleteController{deleteUC: deleteUC}
}

// En user_delete_controller.go
// user_delete_controller.go
func (c *UserDeleteController) Handle(ctx *gin.Context) {
    id, err := strconv.ParseUint(ctx.Param("id_user"), 10, 64)
    if err != nil {
        utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid user ID", err)
        return
    }

    // Crear un usuario con los datos necesarios para el borrado
    user := &entities.User{
        ID:        uint(id),
        UpdatedBy: 1, // Usamos un ID por defecto (1 para admin)
        UpdatedAt: time.Now(),
        Deleted:   true,
    }

    if err := c.deleteUC.Execute(ctx.Request.Context(), user); err != nil {
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