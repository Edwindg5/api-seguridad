// api-seguridad/resources/municipalities/infrastructure/controllers/update_controller.go
package controllers

import (
	"api-seguridad/core/utils"
	"api-seguridad/resources/municipalities/application"
	"api-seguridad/resources/municipalities/domain/entities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateMunicipalityController struct {
	useCase *application.UpdateMunicipalityUseCase
}

func NewUpdateMunicipalityController(useCase *application.UpdateMunicipalityUseCase) *UpdateMunicipalityController {
	return &UpdateMunicipalityController{useCase: useCase}
}

func (c *UpdateMunicipalityController) Handle(ctx *gin.Context) {
    idStr := ctx.Param("id")
    id, err := strconv.ParseUint(idStr, 10, 32)
    if err != nil {
        utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid municipality ID", nil)
        return
    }

    var updateData struct {
        Name         string `json:"name"`
        DelegationID uint   `json:"delegation_id"`
        Active       bool   `json:"active"`
    }
    
    if err := ctx.ShouldBindJSON(&updateData); err != nil {
        utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request payload", err)
        return
    }

    municipality := &entities.Municipality{
        ID:           uint(id),
        Name:         updateData.Name,
        DelegationID: updateData.DelegationID,
    }

    // Set updated_by from authenticated user
    if updaterID, exists := ctx.Get("userID"); exists {
        if uid, ok := updaterID.(uint); ok {
            municipality.UpdatedBy = uid
        }
    }

    if err := c.useCase.Execute(ctx.Request.Context(), municipality); err != nil {
        status := http.StatusInternalServerError
        switch err.Error() {
        case "invalid municipality ID", "municipality name is required":
            status = http.StatusBadRequest
        case "municipality not found":
            status = http.StatusNotFound
        case "municipality with this name already exists":
            status = http.StatusConflict
        }
        utils.ErrorResponse(ctx, status, err.Error(), nil)
        return
    }

    utils.SuccessResponse(ctx, http.StatusOK, "Municipality updated successfully", municipality)
}