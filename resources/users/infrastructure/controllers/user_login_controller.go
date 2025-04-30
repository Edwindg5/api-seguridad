// api-seguridad/resources/users/infrastructure/controllers/user_login_controller.go
package controllers

import (
	"api-seguridad/core/utils"
	"api-seguridad/resources/users/application"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserLoginController struct {
	loginUC *application.LoginUseCase
}

func NewUserLoginController(loginUC *application.LoginUseCase) *UserLoginController {
	return &UserLoginController{loginUC: loginUC}
}

func (c *UserLoginController) Handle(ctx *gin.Context) {
	// Definimos una estructura para el request de login
	var loginRequest struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	// Parseamos el JSON de entrada
	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		fmt.Printf("Error binding JSON: %v\n", err) // Log de error
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	fmt.Printf("Login attempt for username: %s\n", loginRequest.Username) // Log del intento de login

	// Ejecutamos el caso de uso
	user, token, err := c.loginUC.Execute(ctx.Request.Context(), loginRequest.Username, loginRequest.Password)
	if err != nil {
		fmt.Printf("Login error: %v\n", err) // Log de error
		statusCode := http.StatusUnauthorized
		errorMessage := "Invalid credentials"
		
		// Podemos diferenciar entre distintos tipos de errores si lo deseamos
		if err.Error() == "validation error: username is required" || 
		   err.Error() == "validation error: password is required" ||
		   err.Error() == "validation error: password must be at least 8 characters" {
			statusCode = http.StatusBadRequest
			errorMessage = err.Error()
		}

		utils.ErrorResponse(ctx, statusCode, errorMessage, err)
		return
	}

	// Si todo va bien, retornamos el usuario y el token
	response := gin.H{
		"user":  user,
		"token": token,
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Login successful", response)
}