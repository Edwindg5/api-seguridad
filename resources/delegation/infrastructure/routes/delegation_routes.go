// api-seguridad/resources/delegation/infrastructure/routes/delegation_routes.go
package routes

import (
	"api-seguridad/resources/delegation/infrastructure/controllers"
	"api-seguridad/resources/delegation/infrastructure/dependencies"
	"github.com/gin-gonic/gin"
	"api-seguridad/core/middleware"
)

func ConfigureRoutes(router *gin.RouterGroup) {
	// Initialize controllers
	createCtrl := controllers.NewCreateDelegationController(dependencies.GetCreateDelegationUseCase())
	getByIdCtrl := controllers.NewGetDelegationByIDController(dependencies.GetDelegationByIDUseCase())
	getAllCtrl := controllers.NewGetAllDelegationsController(dependencies.GetAllDelegationsUseCase())
	updateCtrl := controllers.NewUpdateDelegationController(dependencies.GetUpdateDelegationUseCase())
	softDeleteCtrl := controllers.NewSoftDeleteDelegationController(dependencies.GetSoftDeleteDelegationUseCase())

	// Configure routes with auth middleware
	delegationRoutes := router.Group("/delegations")
	delegationRoutes.Use(middleware.AuthMiddleware()) // Add this line
	{
		delegationRoutes.POST("", createCtrl.Handle)
		delegationRoutes.GET("", getAllCtrl.Handle)
		delegationRoutes.GET("/:id", getByIdCtrl.Handle)
		delegationRoutes.PUT("/:id", updateCtrl.Handle)
		delegationRoutes.DELETE("/:id", softDeleteCtrl.Handle)
	}
}