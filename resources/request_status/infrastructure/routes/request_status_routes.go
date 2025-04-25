package routes

import (
	"api-seguridad/core/middleware"
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

	// Configure API routes with authentication middleware
	statusRoutes := router.Group("/request-status")
	statusRoutes.Use(middleware.AuthMiddleware()) // Aplicar middleware a todas las rutas
	{
		// CRUD endpoints
		statusRoutes.POST("", createCtrl.Handle)          // Create new status
		statusRoutes.GET("", getAllCtrl.Handle)           // Get all statuses
		statusRoutes.GET("/:id", getByIdCtrl.Handle)      // Get status by ID
		statusRoutes.PUT("/:id", updateCtrl.Handle)       // Update status
		statusRoutes.DELETE("/:id", deleteCtrl.Handle)    // Delete status
	}
}