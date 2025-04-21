package controllers

import (
	"net/http"
	"api-seguridad/resources/roles/application"
	"github.com/gin-gonic/gin"
)

type GetAllRolesController struct {
	useCase *application.GetAllRolesUseCase
}

func NewGetAllRolesController(useCase *application.GetAllRolesUseCase) *GetAllRolesController {
	return &GetAllRolesController{useCase: useCase}
}

func (c *GetAllRolesController) Handle(ctx *gin.Context) {
	roles, err := c.useCase.Execute(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, roles)
}