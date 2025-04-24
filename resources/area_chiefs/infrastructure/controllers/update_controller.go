// api-seguridad/resources/area_chiefs/infrastructure/controllers/update_controller.go
package controllers

import (
	"net/http"
	"strconv"
	"api-seguridad/core/utils"
	"api-seguridad/resources/area_chiefs/application"
	"api-seguridad/resources/area_chiefs/domain/entities"
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

	var chief entities.AreaChief
	if err := ctx.ShouldBindJSON(&chief); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	// Set ID from URL and updater user
	chief.ID = uint(id)
	if updaterID, exists := ctx.Get("userID"); exists {
		if uid, ok := updaterID.(uint); ok {
			chief.UpdatedBy = uid
		}
	}

	if err := c.useCase.Execute(ctx.Request.Context(), &chief); err != nil {
		status := http.StatusInternalServerError
		switch err.Error() {
		case "invalid chief ID", "chief name is required", "updater user is required":
			status = http.StatusBadRequest
		case "area chief not found":
			status = http.StatusNotFound
		}
		utils.ErrorResponse(ctx, status, err.Error(), nil)
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Area chief updated successfully", chief)
}