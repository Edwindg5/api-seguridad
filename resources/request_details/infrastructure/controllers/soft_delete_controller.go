package controllers

import (
	"net/http"
	"strconv"
	"api-seguridad/core/utils"
	"api-seguridad/resources/request_details/application"
	"github.com/gin-gonic/gin"
)

type SoftDeleteRequestDetailController struct {
	useCase *application.SoftDeleteRequestDetailUseCase
}

func NewSoftDeleteRequestDetailController(useCase *application.SoftDeleteRequestDetailUseCase) *SoftDeleteRequestDetailController {
	return &SoftDeleteRequestDetailController{useCase: useCase}
}

func (c *SoftDeleteRequestDetailController) Handle(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "ID inválido", err)
		return
	}

	// Obtener ID del usuario que elimina
	userID, exists := ctx.Get("userID")
	if !exists {
		utils.ErrorResponse(ctx, http.StatusUnauthorized, "Usuario no autenticado", nil)
		return
	}

	uid, ok := userID.(uint)
	if !ok {
		utils.ErrorResponse(ctx, http.StatusUnauthorized, "ID de usuario inválido", nil)
		return
	}

	if err := c.useCase.Execute(ctx.Request.Context(), uint(id), uid); err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "detalle de solicitud no encontrado" {
			status = http.StatusNotFound
		}
		utils.ErrorResponse(ctx, status, err.Error(), nil)
		return
	}

	utils.SuccessResponse(ctx, http.StatusNoContent, "Detalle de solicitud eliminado", nil)
}