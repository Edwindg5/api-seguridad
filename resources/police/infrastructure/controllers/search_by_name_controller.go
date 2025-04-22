package controllers

import (
	"net/http"
	"api-seguridad/core/utils"
	"api-seguridad/resources/police/application"
	"github.com/gin-gonic/gin"
)

type SearchPoliceByNameController struct {
	useCase *application.SearchPoliceByNameUseCase
}

func NewSearchPoliceByNameController(useCase *application.SearchPoliceByNameUseCase) *SearchPoliceByNameController {
	return &SearchPoliceByNameController{useCase: useCase}
}

func (c *SearchPoliceByNameController) Handle(ctx *gin.Context) {
	name := ctx.Query("name")
	if name == "" {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Name parameter is required", nil)
		return
	}

	results, err := c.useCase.Execute(ctx.Request.Context(), name)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Search failed", err)
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Search completed successfully", results)
}