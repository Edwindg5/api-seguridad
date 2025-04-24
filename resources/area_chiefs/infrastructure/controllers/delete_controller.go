// api-seguridad/resources/area_chiefs/infrastructure/controllers/delete_controller.go
package controllers

import (
	"net/http"
	"strconv"
	"api-seguridad/core/utils"
	"api-seguridad/resources/area_chiefs/application"
	"github.com/gin-gonic/gin"
)

type DeleteAreaChiefController struct {
	useCase *application.DeleteAreaChiefUseCase
}

func NewDeleteAreaChiefController(useCase *application.DeleteAreaChiefUseCase) *DeleteAreaChiefController {
	return &DeleteAreaChiefController{useCase: useCase}
}

func (c *DeleteAreaChiefController) Handle(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid chief ID", err)
		return
	}

	if err := c.useCase.Execute(ctx.Request.Context(), uint(id)); err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "area chief not found" {
			status = http.StatusNotFound
		}
		utils.ErrorResponse(ctx, status, err.Error(), nil)
		return
	}

	utils.SuccessResponse(ctx, http.StatusNoContent, "Area chief deleted successfully", nil)
}