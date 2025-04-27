//api-seguridad/resources/permissions/infrastructure/controllers/get_all_controller.go
package controllers

import (
	"net/http"
	"api-seguridad/core/utils"
	"api-seguridad/resources/permissions/application"
	"github.com/gin-gonic/gin"
)

type GetAllPermissionsController struct {
	useCase *application.GetAllPermissionsUseCase
}

func NewGetAllPermissionsController(useCase *application.GetAllPermissionsUseCase) *GetAllPermissionsController {
	return &GetAllPermissionsController{useCase: useCase}
}

func (c *GetAllPermissionsController) Handle(ctx *gin.Context) {
	permissions, err := c.useCase.Execute(ctx.Request.Context())
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Error al obtener permisos", err)
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Lista de permisos obtenida", permissions)
}