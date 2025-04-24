// api-seguridad/resources/request/infrastructure/routes/request_routes.go
package routes

import (
	"api-seguridad/resources/request/infrastructure/controllers"
	"api-seguridad/resources/request/infrastructure/dependencies"
	"github.com/gin-gonic/gin"
)

func ConfigureRoutes(router *gin.RouterGroup) {
	// Initialize controllers with injected use cases
	createCtrl := controllers.NewCreateRequestController(dependencies.GetCreateRequestUseCase())
	getByIdCtrl := controllers.NewGetRequestByIDController(dependencies.GetRequestByIDUseCase())
	updateCtrl := controllers.NewUpdateRequestController(dependencies.GetUpdateRequestUseCase())
	deleteCtrl := controllers.NewDeleteRequestController(dependencies.GetDeleteRequestUseCase())
	getByStatusCtrl := controllers.NewGetRequestsByStatusController(dependencies.GetRequestsByStatusUseCase())
	getByMunicipalityCtrl := controllers.NewGetRequestsByMunicipalityController(dependencies.GetRequestsByMunicipalityUseCase())

	// Configure API routes
	requestRoutes := router.Group("/requests")
	{
		// CRUD endpoints
		requestRoutes.POST("", createCtrl.Handle)              // Create new request
		requestRoutes.GET("/:id", getByIdCtrl.Handle)         // Get request by ID
		requestRoutes.PUT("/:id", updateCtrl.Handle)          // Update request
		requestRoutes.DELETE("/:id", deleteCtrl.Handle)       // Delete request

		// Specialized endpoints
		requestRoutes.GET("/status", getByStatusCtrl.Handle)       // Get requests by status
		requestRoutes.GET("/municipality", getByMunicipalityCtrl.Handle) // Get requests by municipality
	}
}