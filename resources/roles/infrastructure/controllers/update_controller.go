package controllers

import (
	"net/http"
	"api-seguridad/resources/roles/application"
	"api-seguridad/resources/roles/domain/entities"
	"github.com/gin-gonic/gin"
)

type UpdateRoleController struct {
	useCase *application.UpdateRoleUseCase
}

func NewUpdateRoleController(useCase *application.UpdateRoleUseCase) *UpdateRoleController {
	return &UpdateRoleController{useCase: useCase}
}

func (c *UpdateRoleController) Handle(ctx *gin.Context) {
	var role entities.Role
	if err := ctx.ShouldBindJSON(&role); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.useCase.Execute(ctx.Request.Context(), &role); err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "role not found" {
			status = http.StatusNotFound
		} else if err.Error() == "role with this title already exists" {
			status = http.StatusConflict
		}
		ctx.JSON(status, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, role)
}