package controllers

import (
	"net/http"
	"api-seguridad/core/utils"
	"api-seguridad/resources/police/application"
	"github.com/gin-gonic/gin"
)

type GetAllPoliceController struct {
	useCase *application.GetAllPoliceUseCase
}

func NewGetAllPoliceController(useCase *application.GetAllPoliceUseCase) *GetAllPoliceController {
	return &GetAllPoliceController{useCase: useCase}
}

func (c *GetAllPoliceController) Handle(ctx *gin.Context) {
	policeList, err := c.useCase.Execute(ctx.Request.Context())
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to get police list", err)
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Police list retrieved successfully", policeList)
}