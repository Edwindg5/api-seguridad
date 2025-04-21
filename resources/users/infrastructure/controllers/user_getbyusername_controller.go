// api-seguridad/resources/users/infrastructure/controllers/user_getbyusername_controller.go
package controllers

import (
	"net/http"
	"api-seguridad/resources/users/application"
	"api-seguridad/core/utils"

	"github.com/gin-gonic/gin"
)

type UserGetByUsernameController struct {
	getByUsernameUC *application.GetUserByUsernameUseCase
}

func NewUserGetByUsernameController(getByUsernameUC *application.GetUserByUsernameUseCase) *UserGetByUsernameController {
	return &UserGetByUsernameController{getByUsernameUC: getByUsernameUC}
}

func (c *UserGetByUsernameController) Handle(ctx *gin.Context) {
	username := ctx.Param("username")

	user, err := c.getByUsernameUC.Execute(ctx.Request.Context(), username)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "user not found" {
			statusCode = http.StatusNotFound
		}
		utils.ErrorResponse(ctx, statusCode, "Failed to get user", err)
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "User retrieved successfully", user)
}