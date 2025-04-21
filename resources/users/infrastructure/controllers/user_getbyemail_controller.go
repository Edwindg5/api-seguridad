// api-seguridad/resources/users/infrastructure/controllers/user_getbyemail_controller.go
package controllers

import (
	"net/http"
	"api-seguridad/resources/users/application"
	"api-seguridad/core/utils"

	"github.com/gin-gonic/gin"
)

type UserGetByEmailController struct {
	getByEmailUC *application.GetUserByEmailUseCase
}

func NewUserGetByEmailController(getByEmailUC *application.GetUserByEmailUseCase) *UserGetByEmailController {
	return &UserGetByEmailController{getByEmailUC: getByEmailUC}
}

func (c *UserGetByEmailController) Handle(ctx *gin.Context) {
	email := ctx.Param("email")

	user, err := c.getByEmailUC.Execute(ctx.Request.Context(), email)
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