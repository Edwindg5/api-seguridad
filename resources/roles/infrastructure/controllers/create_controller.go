package controllers

import (
	"net/http"
	"api-seguridad/resources/roles/application"
	"api-seguridad/resources/roles/domain/entities"
	"github.com/gin-gonic/gin"
)

type CreateRoleController struct {
	useCase *application.CreateRoleUseCase
}

func NewCreateRoleController(useCase *application.CreateRoleUseCase) *CreateRoleController {
	return &CreateRoleController{useCase: useCase}
}

func (c *CreateRoleController) Handle(ctx *gin.Context) {
	var role entities.Role
	if err := ctx.ShouldBindJSON(&role); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.useCase.Execute(ctx.Request.Context(), &role); err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, role)
}