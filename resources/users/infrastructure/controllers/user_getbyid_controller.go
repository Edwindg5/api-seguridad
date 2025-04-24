// api-seguridad/resources/users/infrastructure/controllers/user_getbyid_controller.go
package controllers

import (
	"net/http"
	"strconv"
	"api-seguridad/resources/users/application"
	"api-seguridad/core/utils"

	"github.com/gin-gonic/gin"
)

type UserGetByIDController struct {
	getByIDUC *application.GetUserByIDUseCase
}

func NewUserGetByIDController(getByIDUC *application.GetUserByIDUseCase) *UserGetByIDController {
	return &UserGetByIDController{getByIDUC: getByIDUC}
}

func (c *UserGetByIDController) Handle(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id_user"), 10, 64)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid user ID", err)
		return
	}

	user, err := c.getByIDUC.Execute(ctx.Request.Context(), uint(id))
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