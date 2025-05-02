// api-seguridad/resources/area_chiefs/infrastructure/controllers/create_controller.go
package controllers

import (
	"api-seguridad/core/utils"
	"api-seguridad/resources/area_chiefs/application"
	"api-seguridad/resources/area_chiefs/domain/entities"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)


type CreateAreaChiefController struct {
	useCase *application.CreateAreaChiefUseCase
	uploadPath string
}

func NewCreateAreaChiefController(useCase *application.CreateAreaChiefUseCase) *CreateAreaChiefController {
	// Define la ruta de uploads relativa al directorio del proyecto
	uploadPath := filepath.Join("core", "uploads", "signatures")
	return &CreateAreaChiefController{
		useCase: useCase,
		uploadPath: uploadPath,
	}
}

func (c *CreateAreaChiefController) Handle(ctx *gin.Context) {
	var chief entities.AreaChief
	if err := ctx.ShouldBindJSON(&chief); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	// Obtener userID del contexto de autenticación
	userID, exists := ctx.Get("userID")
	if !exists {
		utils.ErrorResponse(ctx, http.StatusUnauthorized, "Authentication required", nil)
		return
	}

	// Convertir y asignar el userID
	if uid, ok := userID.(uint); ok {
		chief.CreatedBy = uid
		chief.UpdatedBy = uid
	} else {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Invalid user ID in context", nil)
		return
	}

	if err := c.useCase.Execute(ctx.Request.Context(), &chief); err != nil {
		status := http.StatusInternalServerError
		switch err.Error() {
		case "chief name is required", "position is required":
			status = http.StatusBadRequest
		}
		utils.ErrorResponse(ctx, status, err.Error(), nil)
		return
	}

	utils.SuccessResponse(ctx, http.StatusCreated, "Area chief created successfully", chief)
}
