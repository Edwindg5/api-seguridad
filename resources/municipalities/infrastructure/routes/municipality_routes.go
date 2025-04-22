//api-seguridad/resources/municipalities/infrastructure/routes/municipality_routes.go
package routes

import (
	"api-seguridad/resources/municipalities/infrastructure/controllers"
	"api-seguridad/resources/municipalities/infrastructure/dependencies"
	"github.com/gin-gonic/gin"
)

func ConfigureRoutes(router *gin.RouterGroup) {
	// Inicializar controladores con sus respectivos casos de uso
	createCtrl := controllers.NewCreateMunicipalityController(dependencies.GetCreateUseCase())
	getByIdCtrl := controllers.NewGetMunicipalityByIDController(dependencies.GetByIDUseCase())
	getAllCtrl := controllers.NewGetAllMunicipalitiesController(dependencies.GetAllUseCase())
	updateCtrl := controllers.NewUpdateMunicipalityController(dependencies.GetUpdateUseCase())
	softDeleteCtrl := controllers.NewSoftDeleteMunicipalityController(dependencies.GetSoftDeleteUseCase())

	// Configurar rutas
	municipalityRoutes := router.Group("/municipalities")
	{
		municipalityRoutes.POST("", createCtrl.Handle)
		municipalityRoutes.GET("", getAllCtrl.Handle)
		municipalityRoutes.GET("/:id", getByIdCtrl.Handle)
		municipalityRoutes.PUT("/:id", updateCtrl.Handle)
		municipalityRoutes.DELETE("/:id", softDeleteCtrl.Handle)
	}
}