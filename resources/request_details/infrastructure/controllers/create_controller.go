package controllers

import (
	"net/http"
	"api-seguridad/core/utils"
	"api-seguridad/resources/request_details/application"
	"api-seguridad/resources/request_details/domain/entities"
	"github.com/gin-gonic/gin"
)

type CreateRequestDetailController struct {
	useCase *application.CreateRequestDetailUseCase
}

func NewCreateRequestDetailController(useCase *application.CreateRequestDetailUseCase) *CreateRequestDetailController {
	return &CreateRequestDetailController{useCase: useCase}
}

func (c *CreateRequestDetailController) Handle(ctx *gin.Context) {
	var detail entities.RequestDetail
	if err := ctx.ShouldBindJSON(&detail); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Datos inválidos", err)
		return
	}

	// Obtener ID del usuario creador
	if creatorID, exists := ctx.Get("userID"); exists {
		if uid, ok := creatorID.(uint); ok {
			detail.CreatedBy = uid
			detail.UpdatedBy = uid
		}
	}

	if err := c.useCase.Execute(ctx.Request.Context(), &detail); err != nil {
		status := http.StatusInternalServerError
		switch err.Error() {
		case "ID de solicitud es requerido", 
		     "ID de policía es requerido",
		     "usuario creador es requerido":
			status = http.StatusBadRequest
		}
		utils.ErrorResponse(ctx, status, err.Error(), nil)
		return
	}

	utils.SuccessResponse(ctx, http.StatusCreated, "Detalle de solicitud creado", detail)
}