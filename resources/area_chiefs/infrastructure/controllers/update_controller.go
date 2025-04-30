// api-seguridad/resources/area_chiefs/infrastructure/controllers/update_controller.go
package controllers

import (
	"net/http"
	"strconv"

	"api-seguridad/core/utils"
	"api-seguridad/resources/area_chiefs/application"
	"github.com/gin-gonic/gin"
)

type UpdateAreaChiefController struct {
	useCase *application.UpdateAreaChiefUseCase
}

func NewUpdateAreaChiefController(useCase *application.UpdateAreaChiefUseCase) *UpdateAreaChiefController {
	return &UpdateAreaChiefController{useCase: useCase}
}

func (c *UpdateAreaChiefController) Handle(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid chief ID", err)
		return
	}

	var updateData struct {
		Name     string `json:"name"`
		Position string `json:"position"`
		Type     string `json:"type"`
	}
	if err := ctx.ShouldBindJSON(&updateData); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	updateChief := &application.AreaChiefUpdate{
		ID:       uint(id),
		Name:     updateData.Name,
		Position: updateData.Position,
		Type:     updateData.Type,
		// UpdatedBy is now omitted entirely
	}

	if err := c.useCase.Execute(ctx.Request.Context(), updateChief); err != nil {
		status := http.StatusInternalServerError
		switch err.Error() {
		case "invalid chief ID", "chief name is required":
			status = http.StatusBadRequest
		case "area chief not found":
			status = http.StatusNotFound
		}
		utils.ErrorResponse(ctx, status, err.Error(), nil)
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Area chief updated successfully", nil)
}
