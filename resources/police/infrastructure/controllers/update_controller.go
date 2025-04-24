// api-seguridad/resources/police/infrastructure/controllers/update_controller.go
package controllers

import (
	"api-seguridad/core/utils"
	"api-seguridad/resources/police/application"
	"api-seguridad/resources/police/domain/entities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdatePoliceController struct {
	useCase *application.UpdatePoliceUseCase
}

func NewUpdatePoliceController(useCase *application.UpdatePoliceUseCase) *UpdatePoliceController {
	return &UpdatePoliceController{useCase: useCase}
}

func (c *UpdatePoliceController) Handle(ctx *gin.Context) {
    // Obtener el ID de la URL
    idStr := ctx.Param("id")
    id, err := strconv.ParseUint(idStr, 10, 32)
    if err != nil {
        utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid police ID", nil)
        return
    }

    var police entities.Police
    if err := ctx.ShouldBindJSON(&police); err != nil {
        utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request payload", err)
        return
    }

    // Asignar el ID de la URL al objeto police
    police.ID = uint(id)

    // Obtener ID del usuario que realiza la actualización
    if updaterID, exists := ctx.Get("userID"); exists {
        if uid, ok := updaterID.(uint); ok {
            police.UpdatedBy = uid
        }
    }

    if err := c.useCase.Execute(ctx.Request.Context(), &police); err != nil {
        status := http.StatusInternalServerError
        switch err.Error() {
        case "invalid police ID", "name and paternal lastname are required":
            status = http.StatusBadRequest
        case "police not found":
            status = http.StatusNotFound
        }
        utils.ErrorResponse(ctx, status, err.Error(), nil)
        return
    }

    utils.SuccessResponse(ctx, http.StatusOK, "Police updated successfully", police)
}