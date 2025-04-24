// api-seguridad/resources/request/infrastructure/controllers/delete_controller.go
package controllers

import (
	"net/http"
	"strconv"
	"api-seguridad/core/utils"
	"api-seguridad/resources/request/application"
	"github.com/gin-gonic/gin"
)

type DeleteRequestController struct {
	useCase *application.DeleteRequestUseCase
}

func NewDeleteRequestController(useCase *application.DeleteRequestUseCase) *DeleteRequestController {
	return &DeleteRequestController{useCase: useCase}
}

func (c *DeleteRequestController) Handle(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request ID", err)
		return
	}

	if err := c.useCase.Execute(ctx.Request.Context(), uint(id)); err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "request not found" {
			status = http.StatusNotFound
		}
		utils.ErrorResponse(ctx, status, err.Error(), nil)
		return
	}

	utils.SuccessResponse(ctx, http.StatusNoContent, "Request deleted successfully", nil)
}