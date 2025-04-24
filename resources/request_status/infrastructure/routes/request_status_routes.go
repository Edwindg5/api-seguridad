// api-seguridad/resources/request_status/infrastructure/routes/request_status_routes.go
package routes

import (
	"api-seguridad/resources/request_status/infrastructure/controllers"
	"api-seguridad/resources/request_status/infrastructure/dependencies"
	"github.com/gin-gonic/gin"
)

func ConfigureRoutes(router *gin.RouterGroup) {
	// Initialize controllers with injected use cases
	createCtrl := controllers.NewCreateRequestStatusController(dependencies.GetCreateRequestStatusUseCase())
	getByIdCtrl := controllers.NewGetRequestStatusByIDController(dependencies.GetRequestStatusByIDUseCase())
	getAllCtrl := controllers.NewGetAllRequestStatusController(dependencies.GetAllRequestStatusUseCase())
	updateCtrl := controllers.NewUpdateRequestStatusController(dependencies.GetUpdateRequestStatusUseCase())
	deleteCtrl := controllers.NewDeleteRequestStatusController(dependencies.GetDeleteRequestStatusUseCase())

	// Configure API routes
	statusRoutes := router.Group("/request-status")
	{
		// CRUD endpoints
		statusRoutes.POST("", createCtrl.Handle)          // Create new status
		statusRoutes.GET("", getAllCtrl.Handle)           // Get all statuses
		statusRoutes.GET("/:id", getByIdCtrl.Handle)      // Get status by ID
		statusRoutes.PUT("/:id", updateCtrl.Handle)       // Update status
		statusRoutes.DELETE("/:id", deleteCtrl.Handle)    // Delete status
	}
}