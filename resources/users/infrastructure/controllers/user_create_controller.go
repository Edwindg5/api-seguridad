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
    var request struct {
        FirstName string `json:"first_name" binding:"required"`
        LastName  string `json:"last_name" binding:"required"`
        Username  string `json:"username" binding:"required"`
        Email     string `json:"email" binding:"required,email"`
        Password  string `json:"password" binding:"required,min=8"`
        RoleID    uint   `json:"role_id" binding:"required"`
    }

    if err := ctx.ShouldBindJSON(&request); err != nil {
        fmt.Printf("Error binding JSON: %v\n", err)
        utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request payload", err)
        return
    }

    user := &entities.User{
        FirstName: request.FirstName,
        LastName:  request.LastName,
        Username:  request.Username,
        Email:     request.Email,
        Password:  request.Password,
        RoleID:    request.RoleID,
        CreatedBy: 1, // ID del usuario admin por defecto
    }

    if err := c.createUC.Execute(ctx.Request.Context(), user); err != nil {
        fmt.Printf("Create user error: %v\n", err)
        statusCode := http.StatusInternalServerError
        if err.Error() == "username already exists" || 
           err.Error() == "email already exists" ||
           err.Error() == "role is required" {
            statusCode = http.StatusBadRequest
        }
        utils.ErrorResponse(ctx, statusCode, "Failed to create user", err)
        return
    }

    // No devolver la contraseña en la respuesta
    user.Password = ""
    utils.SuccessResponse(ctx, http.StatusCreated, "User created successfully", user)
}