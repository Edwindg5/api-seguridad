//api-seguridad/resources/permissions/infrastructure/controllers/create_controller.go
package controllers

import (
	"net/http"
	"api-seguridad/core/utils"
	"api-seguridad/resources/permissions/application"
	"api-seguridad/resources/permissions/domain/entities"
	"github.com/gin-gonic/gin"
)

type CreatePermissionController struct {
	useCase *application.CreatePermissionUseCase
}

func NewCreatePermissionController(useCase *application.CreatePermissionUseCase) *CreatePermissionController {
	return &CreatePermissionController{useCase: useCase}
}

func (c *CreatePermissionController) Handle(ctx *gin.Context) {
	var permission entities.Permission
	if err := ctx.ShouldBindJSON(&permission); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Datos de permiso inválidos", err)
		return
	}

	// Obtener ID del usuario creador
	if creatorID, exists := ctx.Get("userID"); exists {
		if uid, ok := creatorID.(uint); ok {
			permission.CreatedBy = uid
			permission.UpdatedBy = uid
		}
	}

	if err := c.useCase.Execute(ctx.Request.Context(), &permission); err != nil {
		status := http.StatusInternalServerError
		switch err.Error() {
		case "el nombre del permiso es requerido", 
		     "el nombre no puede exceder los 100 caracteres",
		     "la descripción no puede exceder los 255 caracteres",
		     "se requiere un usuario creador válido":
			status = http.StatusBadRequest
		case "ya existe un permiso con este ID":
			status = http.StatusConflict
		}
		utils.ErrorResponse(ctx, status, err.Error(), nil)
		return
	}

	utils.SuccessResponse(ctx, http.StatusCreated, "Permiso creado exitosamente", permission)
}