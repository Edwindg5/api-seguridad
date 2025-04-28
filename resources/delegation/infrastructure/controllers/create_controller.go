// api-seguridad/resources/delegation/infrastructure/controllers/create_controller.go
package controllers

import (
	"api-seguridad/core/utils"
	"api-seguridad/resources/delegation/application"
	"api-seguridad/resources/delegation/domain/entities"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type CreateDelegationController struct {
	useCase *application.CreateDelegationUseCase
}

func NewCreateDelegationController(useCase *application.CreateDelegationUseCase) *CreateDelegationController {
	return &CreateDelegationController{useCase: useCase}
}

func (c *CreateDelegationController) Handle(ctx *gin.Context) {
    var delegation entities.Delegation
    if err := ctx.ShouldBindJSON(&delegation); err != nil {
        utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request payload", err)
        return
    }

    // Get user ID from context
    userID, exists := ctx.Get("userID")
    if !exists {
        utils.ErrorResponse(ctx, http.StatusUnauthorized, "User not authenticated", nil)
        return
    }

    // Set user IDs and timestamps
    if uid, ok := userID.(uint); ok {
        delegation.SetCreatedBy(uid)
        delegation.SetUpdatedBy(uid)
    } else {
        utils.ErrorResponse(ctx, http.StatusInternalServerError, "Invalid user ID in context", nil)
        return
    }

    now := time.Now()
    delegation.SetCreatedAt(now)
    delegation.SetUpdatedAt(now)

    if err := c.useCase.Execute(ctx.Request.Context(), &delegation); err != nil {
        status := http.StatusInternalServerError
        if err.Error() == "delegation name is required" {
            status = http.StatusBadRequest
        }
        utils.ErrorResponse(ctx, status, err.Error(), nil)
        return
    }

    utils.SuccessResponse(ctx, http.StatusCreated, "Delegation created successfully", delegation)
}