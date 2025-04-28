// api-seguridad/resources/municipalities/infrastructure/routes/municipality_routes.go
package routes

import (
	"api-seguridad/core/middleware"
	"api-seguridad/resources/municipalities/infrastructure/controllers"
	"api-seguridad/resources/municipalities/infrastructure/dependencies"
	"github.com/gin-gonic/gin"
)

func ConfigureRoutes(router *gin.RouterGroup) {
	// Initialize controllers
	createCtrl := controllers.NewCreateMunicipalityController(dependencies.GetCreateUseCase())
	getByIdCtrl := controllers.NewGetMunicipalityByIDController(dependencies.GetByIDUseCase())
	getAllCtrl := controllers.NewGetAllMunicipalitiesController(dependencies.GetAllUseCase())
	updateCtrl := controllers.NewUpdateMunicipalityController(dependencies.GetUpdateUseCase())
	softDeleteCtrl := controllers.NewSoftDeleteMunicipalityController(dependencies.GetSoftDeleteUseCase())

	// Configure routes with auth middleware
	municipalityRoutes := router.Group("/municipalities")
	municipalityRoutes.Use(middleware.AuthMiddleware())
	{
		municipalityRoutes.POST("", createCtrl.Handle)
		municipalityRoutes.GET("", getAllCtrl.Handle)
		municipalityRoutes.GET("/:id", getByIdCtrl.Handle)
		municipalityRoutes.PUT("/:id", updateCtrl.Handle)
		municipalityRoutes.DELETE("/:id", softDeleteCtrl.Handle)
	}
}