// api-seguridad/resources/request_status/infrastructure/controllers/delete_controller.go
package controllers

import (
	"net/http"
	"strconv"
	"api-seguridad/core/utils"
	"api-seguridad/resources/request_status/application"
	"github.com/gin-gonic/gin"
)

type DeleteRequestStatusController struct {
	useCase *application.DeleteRequestStatusUseCase
}

func NewDeleteRequestStatusController(useCase *application.DeleteRequestStatusUseCase) *DeleteRequestStatusController {
	return &DeleteRequestStatusController{useCase: useCase}
}

func (c *DeleteRequestStatusController) Handle(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid status ID", err)
		return
	}

	if err := c.useCase.Execute(ctx.Request.Context(), uint(id)); err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "status not found" {
			status = http.StatusNotFound
		}
		utils.ErrorResponse(ctx, status, err.Error(), nil)
		return
	}

	utils.SuccessResponse(ctx, http.StatusNoContent, "Request status deleted successfully", nil)
}