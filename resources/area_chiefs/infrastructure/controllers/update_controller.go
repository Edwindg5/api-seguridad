// api-seguridad/resources/area_chiefs/infrastructure/controllers/update_controller.go
package controllers

import (
	"net/http"

	"path/filepath"
	"strconv"
	
	"api-seguridad/core/utils"
	"api-seguridad/resources/area_chiefs/application"
	"github.com/gin-gonic/gin"
)

type UpdateAreaChiefController struct {
	useCase    *application.UpdateAreaChiefUseCase
	uploadPath string
}

func NewUpdateAreaChiefController(useCase *application.UpdateAreaChiefUseCase) *UpdateAreaChiefController {
	uploadPath := filepath.Join("core", "uploads", "signatures")
	return &UpdateAreaChiefController{
		useCase:    useCase,
		uploadPath: uploadPath,
	}
}

func (c *UpdateAreaChiefController) Handle(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid chief ID", err)
		return
	}

    // Bind only the fields that should be updated
    var updateData struct {
        Name     string `json:"name"`
        Position string `json:"position"`
        Type     string `json:"type"`
    }
    if err := ctx.ShouldBindJSON(&updateData); err != nil {
        utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request payload", err)
        return
    }

    // Create a partial update object
    updateChief := &application.AreaChiefUpdate{
        ID:       uint(id),
        Name:     updateData.Name,
        Position: updateData.Position,
        Type:     updateData.Type,
    }
    
    // Set updater user
    if updaterID, exists := ctx.Get("userID"); exists {
        if uid, ok := updaterID.(uint); ok {
            updateChief.UpdatedBy = uid
        }
    }

    if err := c.useCase.Execute(ctx.Request.Context(), updateChief); err != nil {
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

	utils.SuccessResponse(ctx, http.StatusOK, "Area chief updated successfully", nil)
}
