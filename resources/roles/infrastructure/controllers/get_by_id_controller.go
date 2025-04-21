package controllers

import (
	"net/http"
	"strconv"
	"api-seguridad/resources/roles/application"
	"github.com/gin-gonic/gin"
)

type GetRoleByIDController struct {
	useCase *application.GetRoleByIDUseCase
}

func NewGetRoleByIDController(useCase *application.GetRoleByIDUseCase) *GetRoleByIDController {
	return &GetRoleByIDController{useCase: useCase}
}

func (c *GetRoleByIDController) Handle(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid role ID"})
		return
	}

	role, err := c.useCase.Execute(ctx.Request.Context(), uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, role)
}