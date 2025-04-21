// api-seguridad/resources/users/infrastructure/controllers/user_list_controller.go
package controllers

import (
	"net/http"
	"api-seguridad/resources/users/application"
	"api-seguridad/core/utils"

	"github.com/gin-gonic/gin"
)

type UserListController struct {
	listUC *application.ListUsersUseCase
}

func NewUserListController(listUC *application.ListUsersUseCase) *UserListController {
	return &UserListController{listUC: listUC}
}

func (c *UserListController) Handle(ctx *gin.Context) {
	users, err := c.listUC.Execute(ctx.Request.Context())
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to list users", err)
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Users retrieved successfully", users)
}