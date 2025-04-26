package controllers

import (
	"net/http"
	"strconv"
	"api-seguridad/core/utils"
	"api-seguridad/resources/request_details/application"
	"api-seguridad/resources/request_details/domain/entities"
	"github.com/gin-gonic/gin"
)

type UpdateRequestDetailController struct {
	useCase *application.UpdateRequestDetailUseCase
}

func NewUpdateRequestDetailController(useCase *application.UpdateRequestDetailUseCase) *UpdateRequestDetailController {
	return &UpdateRequestDetailController{useCase: useCase}
}

func (c *UpdateRequestDetailController) Handle(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "ID inválido", err)
		return
	}

	var detail entities.RequestDetail
	if err := ctx.ShouldBindJSON(&detail); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Datos inválidos", err)
		return
	}

	detail.ID = uint(id)

	// Obtener ID del usuario que actualiza
	if updaterID, exists := ctx.Get("userID"); exists {
		if uid, ok := updaterID.(uint); ok {
			detail.UpdatedBy = uid
		}
	}

	if err := c.useCase.Execute(ctx.Request.Context(), &detail); err != nil {
		status := http.StatusInternalServerError
		switch err.Error() {
		case "ID inválido", "no se pueden modificar los IDs de solicitud o policía":
			status = http.StatusBadRequest
		case "detalle de solicitud no encontrado":
			status = http.StatusNotFound
		}
		utils.ErrorResponse(ctx, status, err.Error(), nil)
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Detalle de solicitud actualizado", detail)
}