// api-seguridad/resources/users/infrastructure/controllers/user_create_controller.go
package controllers

import (
	"net/http"
	"api-seguridad/resources/users/application"
	"api-seguridad/resources/users/domain/entities"
	"api-seguridad/core/utils"

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
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	if err := c.createUC.Execute(ctx.Request.Context(), &user); err != nil {
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