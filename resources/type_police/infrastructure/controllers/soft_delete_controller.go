//api-seguridad/resources/type_police/infrastructure/controllers/soft_delete_controller.go
package controllers

import (
	"net/http"
	"strconv"
	"api-seguridad/core/utils"
	"api-seguridad/resources/type_police/application"
	"github.com/gin-gonic/gin"
)

type SoftDeleteTypePoliceController struct {
	useCase *application.SoftDeleteTypePoliceUseCase
}

func NewSoftDeleteTypePoliceController(useCase *application.SoftDeleteTypePoliceUseCase) *SoftDeleteTypePoliceController {
	return &SoftDeleteTypePoliceController{useCase: useCase}
}

func (c *SoftDeleteTypePoliceController) Handle(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid type police ID", err)
		return
	}

	if err := c.useCase.Execute(ctx.Request.Context(), uint(id)); err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "invalid type police ID" {
			status = http.StatusBadRequest
		} else if err.Error() == "type police not found" {
			status = http.StatusNotFound
		}
		utils.ErrorResponse(ctx, status, err.Error(), nil)
		return
	}

	utils.SuccessResponse(ctx, http.StatusNoContent, "TypePolice deleted successfully", nil)
}