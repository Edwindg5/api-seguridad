//api-seguridad/resources/request_details/infrastructure/controllers/update_controller.go
package controllers

import (
	"net/http"
	"strconv"
	"api-seguridad/core/utils"
	"api-seguridad/resources/request_details/application"
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

	var updateData struct {
		Active             bool   `json:"active"`
		Census            bool   `json:"census"`
		Located           bool   `json:"located"`
		Register          bool   `json:"register"`
		Approved          bool   `json:"approved"`
		Comments          string `json:"comments"`
		MunicipalityActive bool  `json:"municipality_active"`
	}

	if err := ctx.ShouldBindJSON(&updateData); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Datos inválidos", err)
		return
	}

	// Obtener ID del usuario que actualiza
	var updaterID uint
	if uid, exists := ctx.Get("userID"); exists {
		if u, ok := uid.(uint); ok {
			updaterID = u
		}
	}

	// Crear DTO para la actualización
	updateDTO := application.UpdateRequestDetailDTO{
		ID:                 uint(id),
		Active:             updateData.Active,
		Census:             updateData.Census,
		Located:            updateData.Located,
		Register:           updateData.Register,
		Approved:           updateData.Approved,
		Comments:           updateData.Comments,
		MunicipalityActive: updateData.MunicipalityActive,
		UpdaterID:          updaterID,
	}

	updatedDetail, err := c.useCase.ExecuteWithDTO(ctx.Request.Context(), &updateDTO)
	if err != nil {
		status := http.StatusInternalServerError
		switch err.Error() {
		case "detalle de solicitud no encontrado":
			status = http.StatusNotFound
		case "no se pueden modificar los IDs de solicitud o policía":
			status = http.StatusBadRequest
		}
		utils.ErrorResponse(ctx, status, err.Error(), nil)
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Detalle de solicitud actualizado", updatedDetail)
}