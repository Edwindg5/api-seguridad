//api-seguridad/resources/type_police/infrastructure/controllers/update_controller.go
package controllers

import (
	"net/http"
	"strconv"
	"api-seguridad/core/utils"
	"api-seguridad/resources/type_police/application"
	"api-seguridad/resources/type_police/domain/entities"
	"github.com/gin-gonic/gin"
)

type UpdateTypePoliceController struct {
	useCase *application.UpdateTypePoliceUseCase
}

func NewUpdateTypePoliceController(useCase *application.UpdateTypePoliceUseCase) *UpdateTypePoliceController {
	return &UpdateTypePoliceController{useCase: useCase}
}

func (c *UpdateTypePoliceController) Handle(ctx *gin.Context) {
	// Obtener el ID de la URL
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid type police ID", err)
		return
	}

	var typePolice entities.TypePolice
	if err := ctx.ShouldBindJSON(&typePolice); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	// Asignar el ID desde el parámetro de la URL
	typePolice.ID = uint(id)

	if err := c.useCase.Execute(ctx.Request.Context(), &typePolice); err != nil {
		status := http.StatusInternalServerError
		switch err.Error() {
		case "invalid type police ID", "type police title is required":
			status = http.StatusBadRequest
		case "type police not found":
			status = http.StatusNotFound
		}
		utils.ErrorResponse(ctx, status, err.Error(), nil)
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "TypePolice updated successfully", typePolice)
}