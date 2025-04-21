package controllers

import (
	"net/http"
	"strconv"
	"api-seguridad/resources/roles/application"
	"github.com/gin-gonic/gin"
)

type SoftDeleteRoleController struct {
	useCase *application.SoftDeleteRoleUseCase
}

func NewSoftDeleteRoleController(useCase *application.SoftDeleteRoleUseCase) *SoftDeleteRoleController {
	return &SoftDeleteRoleController{useCase: useCase}
}

func (c *SoftDeleteRoleController) Handle(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid role ID"})
		return
	}

	if err := c.useCase.Execute(ctx.Request.Context(), uint(id)); err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "role not found" {
			status = http.StatusNotFound
		}
		ctx.JSON(status, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}