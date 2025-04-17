package controllers

import (
	"net/http"
	"api-seguridad/resources/users/application"
	"api-seguridad/resources/users/domain/entities"
	"api-seguridad/core/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *application.UserService
}

func NewUserController(userService *application.UserService) *UserController {
	return &UserController{userService: userService}
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var user entity.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	if err := c.userService.CreateUser(ctx.Request.Context(), &user); err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to create user", err)
		return
	}

	utils.SuccessResponse(ctx, http.StatusCreated, "User created successfully", user)
}

func (c *UserController) GetUser(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid user ID", err)
		return
	}

	user, err := c.userService.GetUserByID(ctx.Request.Context(), uint(id))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to get user", err)
		return
	}

	if user == nil {
		utils.ErrorResponse(ctx, http.StatusNotFound, "User not found", nil)
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "User retrieved successfully", user)
}

// Implementar otros métodos del controlador...