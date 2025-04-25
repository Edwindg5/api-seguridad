// api-seguridad/resources/area_chiefs/infrastructure/routes/area_chief_routes.go)
package routes

import (
	"api-seguridad/core/middleware"
	"api-seguridad/resources/area_chiefs/infrastructure/controllers"
	"api-seguridad/resources/area_chiefs/infrastructure/dependencies"
	"github.com/gin-gonic/gin"
)

func ConfigureRoutes(router *gin.RouterGroup) {
	// Initialize controllers
	createCtrl := controllers.NewCreateAreaChiefController(dependencies.GetCreateAreaChiefUseCase())
	getByIDCtrl := controllers.NewGetAreaChiefByIDController(dependencies.GetAreaChiefByIDUseCase())
	getAllCtrl := controllers.NewGetAllAreaChiefsController(dependencies.GetAllAreaChiefsUseCase())
	updateCtrl := controllers.NewUpdateAreaChiefController(dependencies.GetUpdateAreaChiefUseCase())
	deleteCtrl := controllers.NewDeleteAreaChiefController(dependencies.GetDeleteAreaChiefUseCase())

	// Define routes with authentication middleware
	chiefRoutes := router.Group("/area-chiefs")
	chiefRoutes.Use(middleware.AuthMiddleware()) // Aplicar middleware a todas las rutas de area-chiefs
	{
		chiefRoutes.POST("", createCtrl.Handle)
		chiefRoutes.GET("", getAllCtrl.Handle)
		chiefRoutes.GET("/:id", getByIDCtrl.Handle)
		chiefRoutes.PUT("/:id", updateCtrl.Handle)
		chiefRoutes.DELETE("/:id", deleteCtrl.Handle)
	}
}