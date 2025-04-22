package controllers

import (
	"net/http"
	"api-seguridad/core/utils"
	"api-seguridad/resources/type_police/application"
	"api-seguridad/resources/type_police/domain/entities"
	"github.com/gin-gonic/gin"
)

type CreateTypePoliceController struct {
	useCase *application.CreateTypePoliceUseCase
}

func NewCreateTypePoliceController(useCase *application.CreateTypePoliceUseCase) *CreateTypePoliceController {
	return &CreateTypePoliceController{useCase: useCase}
}

func (c *CreateTypePoliceController) Handle(ctx *gin.Context) {
	var typePolice entities.TypePolice
	if err := ctx.ShouldBindJSON(&typePolice); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	if err := c.useCase.Execute(ctx.Request.Context(), &typePolice); err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "type police title is required" {
			status = http.StatusBadRequest
		}
		utils.ErrorResponse(ctx, status, err.Error(), nil)
		return
	}

	utils.SuccessResponse(ctx, http.StatusCreated, "TypePolice created successfully", typePolice)
}