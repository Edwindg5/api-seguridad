package controllers

import (
	"net/http"
	"api-seguridad/core/utils"
	"api-seguridad/resources/type_police/application"
	"github.com/gin-gonic/gin"
)

type GetAllTypePoliceController struct {
	useCase *application.GetAllTypePoliceUseCase
}

func NewGetAllTypePoliceController(useCase *application.GetAllTypePoliceUseCase) *GetAllTypePoliceController {
	return &GetAllTypePoliceController{useCase: useCase}
}

func (c *GetAllTypePoliceController) Handle(ctx *gin.Context) {
	typesPolice, err := c.useCase.Execute(ctx.Request.Context())
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to get types police", err)
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "TypesPolice retrieved successfully", typesPolice)
}