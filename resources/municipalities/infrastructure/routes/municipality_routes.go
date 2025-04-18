//api-seguridad/resources/municipalities/infrastructure/routes/municipality_routes.go
package routes

import (
	
	"api-seguridad/resources/municipalities/infrastructure/controllers"
	"api-seguridad/resources/municipalities/infrastructure/dependencies"

	"github.com/gin-gonic/gin"
)

func ConfigureRoutes(router *gin.RouterGroup) {
	municipalityController := controllers.NewMunicipalityController(dependencies.GetMunicipalityService())

	municipalityRoutes := router.Group("/municipalities")
	{
		municipalityRoutes.POST("", municipalityController.CreateMunicipality)
		municipalityRoutes.GET("/:id", municipalityController.GetMunicipality)

	}
}