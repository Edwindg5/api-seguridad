// api-seguridad/resources/users/infrastructure/controllers/user_create_controller.go
package controllers

import (
	"api-seguridad/core/utils"
	"api-seguridad/resources/users/application"
	"api-seguridad/resources/users/domain/entities"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserCreateController struct {
	createUC *application.CreateUserUseCase
}

func NewUserCreateController(createUC *application.CreateUserUseCase) *UserCreateController {
	return &UserCreateController{createUC: createUC}
}

func (c *UserCreateController) Handle(ctx *gin.Context) {
    var user entities.User
    if err := ctx.ShouldBindJSON(&user); err != nil {
        fmt.Printf("Error binding JSON: %v\n", err) // Log de error
 
        utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request payload", err)
        return
    }

    fmt.Printf("User object before validation: %+v\n", user) // Log del objeto

    if err := c.createUC.Execute(ctx.Request.Context(), &user); err != nil {
        fmt.Printf("Create user error: %v\n", err) // Log de error
        statusCode := http.StatusInternalServerError
        if err.Error() == "username already exists" || 
           err.Error() == "email already exists" ||
           err.Error() == "role_id is required" {
            statusCode = http.StatusBadRequest
        }
        utils.ErrorResponse(ctx, statusCode, "Failed to create user", err)
        return
    }

    utils.SuccessResponse(ctx, http.StatusCreated, "User created successfully", user)
}