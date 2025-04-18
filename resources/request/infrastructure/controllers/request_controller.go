//api-seguridad/resources/request/infrastructure/controllers/request_controller.go
package controllers

import (
	"net/http"
	"api-seguridad/core/utils"
	"api-seguridad/resources/request/application"
	"api-seguridad/resources/request/domain/entities"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RequestController struct {
	requestService *application.RequestService
}

func NewRequestController(requestService *application.RequestService) *RequestController {
	return &RequestController{requestService: requestService}
}

func (c *RequestController) CreateRequest(ctx *gin.Context) {
	var request entity.Request
	if err := ctx.ShouldBindJSON(&request); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	if err := c.requestService.CreateRequest(ctx.Request.Context(), &request); err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to create request", err)
		return
	}

	utils.SuccessResponse(ctx, http.StatusCreated, "Request created successfully", request)
}

func (c *RequestController) GetRequest(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request ID", err)
		return
	}

	request, err := c.requestService.GetRequestByID(ctx.Request.Context(), uint(id))
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to get request", err)
		return
	}

	if request == nil {
		utils.ErrorResponse(ctx, http.StatusNotFound, "Request not found", nil)
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Request retrieved successfully", request)
}