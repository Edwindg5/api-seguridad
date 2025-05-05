package routes

import (
	"api-seguridad/core/middleware"
	"api-seguridad/resources/request_status/infrastructure/controllers"
	reqDeps "api-seguridad/resources/request_status/infrastructure/dependencies"
	userDeps "api-seguridad/resources/users/infrastructure/dependencies"
	"github.com/gin-gonic/gin"
)

func ConfigureRoutes(router *gin.RouterGroup) {
	// Initialize controllers with injected use cases
	createCtrl := controllers.NewCreateRequestStatusController(
		reqDeps.GetCreateRequestStatusUseCase(),
	)
	
	getByIdCtrl := controllers.NewGetRequestStatusByIDController(
		reqDeps.GetRequestStatusByIDUseCase(),
	)
	
	getAllCtrl := controllers.NewGetAllRequestStatusController(
		reqDeps.GetAllRequestStatusUseCase(),
	)
	
	updateCtrl := controllers.NewUpdateRequestStatusController(
		reqDeps.GetUpdateRequestStatusUseCase(),
		userDeps.GetUserRepository(),
	)
	
	deleteCtrl := controllers.NewDeleteRequestStatusController(
		reqDeps.GetDeleteRequestStatusUseCase(),
	)

	// Configure API routes with authentication middleware
	statusRoutes := router.Group("/request-status")
	statusRoutes.Use(middleware.AuthMiddleware())
	{
		statusRoutes.POST("", createCtrl.Handle)
		statusRoutes.GET("", getAllCtrl.Handle)
		statusRoutes.GET("/:id", getByIdCtrl.Handle)
		statusRoutes.PUT("/:id", updateCtrl.Handle)
		statusRoutes.DELETE("/:id", deleteCtrl.Handle)
	}
}