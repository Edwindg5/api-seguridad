package controllers

import (
	"net/http"
	"strconv"
	"api-seguridad/core/utils"
	"api-seguridad/resources/permissions/application"
	"github.com/gin-gonic/gin"
)

type GetPermissionByIDController struct {
	useCase *application.GetPermissionByIDUseCase
}

func NewGetPermissionByIDController(useCase *application.GetPermissionByIDUseCase) *GetPermissionByIDController {
	return &GetPermissionByIDController{useCase: useCase}
}

func (c *GetPermissionByIDController) Handle(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "ID de permiso inválido", err)
		return
	}

	permission, err := c.useCase.Execute(ctx.Request.Context(), uint(id))
	if err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "permiso no encontrado" {
			status = http.StatusNotFound
		}
		utils.ErrorResponse(ctx, status, err.Error(), nil)
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Permiso obtenido exitosamente", permission)
}