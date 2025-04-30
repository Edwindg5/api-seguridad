// api-seguridad/resources/delegation/infrastructure/routes/delegation_routes.go
package routes

import (
	"api-seguridad/core/middleware"
	"api-seguridad/resources/delegation/infrastructure/controllers"
	"api-seguridad/resources/delegation/infrastructure/dependencies"
	"github.com/gin-gonic/gin"
)

func ConfigureRoutes(router *gin.RouterGroup) {
	createCtrl := controllers.NewCreateDelegationController(dependencies.GetCreateDelegationUseCase())
	getByIdCtrl := controllers.NewGetDelegationByIDController(dependencies.GetDelegationByIDUseCase())
	getAllCtrl := controllers.NewGetAllDelegationsController(dependencies.GetAllDelegationsUseCase())
	updateCtrl := controllers.NewUpdateDelegationController(dependencies.GetUpdateDelegationUseCase())
	softDeleteCtrl := controllers.NewSoftDeleteDelegationController(dependencies.GetSoftDeleteDelegationUseCase())

	delegationRoutes := router.Group("/delegations")
	delegationRoutes.Use(middleware.AuthMiddleware())
	{
		delegationRoutes.POST("", createCtrl.Handle)
		delegationRoutes.GET("", getAllCtrl.Handle)
		delegationRoutes.GET("/:id", getByIdCtrl.Handle)
		delegationRoutes.PUT("/:id", updateCtrl.Handle)
		delegationRoutes.DELETE("/:id", softDeleteCtrl.Handle)
	}
}
