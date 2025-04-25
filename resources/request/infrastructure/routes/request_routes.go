// api-seguridad/resources/request/infrastructure/routes/request_routes.go
package routes

import (
	"api-seguridad/core/middleware"
	"api-seguridad/resources/request/infrastructure/controllers"
	"api-seguridad/resources/request/infrastructure/dependencies"
	"github.com/gin-gonic/gin"
)

func ConfigureRoutes(router *gin.RouterGroup) {
	// Initialize controllers
	createCtrl := controllers.NewCreateRequestController(dependencies.GetCreateRequestUseCase())
	getByIdCtrl := controllers.NewGetRequestByIDController(dependencies.GetRequestByIDUseCase())
	updateCtrl := controllers.NewUpdateRequestController(dependencies.GetUpdateRequestUseCase())
	deleteCtrl := controllers.NewDeleteRequestController(dependencies.GetDeleteRequestUseCase())
	getByStatusCtrl := controllers.NewGetRequestsByStatusController(dependencies.GetRequestsByStatusUseCase())
	getByMunicipalityCtrl := controllers.NewGetRequestsByMunicipalityController(dependencies.GetRequestsByMunicipalityUseCase())

	// Configure API routes with authentication
	requestRoutes := router.Group("/requests")
	requestRoutes.Use(middleware.AuthMiddleware()) // Middleware aplicado a todas las rutas
	{
		// CRUD endpoints
		requestRoutes.POST("", createCtrl.Handle)
		requestRoutes.GET("/:id", getByIdCtrl.Handle)
		requestRoutes.PUT("/:id", updateCtrl.Handle)
		requestRoutes.DELETE("/:id", deleteCtrl.Handle)

		// Specialized endpoints
		requestRoutes.GET("/status", getByStatusCtrl.Handle)
		requestRoutes.GET("/municipality", getByMunicipalityCtrl.Handle)
	}
}