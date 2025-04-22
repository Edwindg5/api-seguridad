package routes

import (
	"api-seguridad/resources/delegation/infrastructure/controllers"
	"api-seguridad/resources/delegation/infrastructure/dependencies"
	"github.com/gin-gonic/gin"
)

func ConfigureRoutes(router *gin.RouterGroup) {
	// Initialize controllers with the correct dependency getters
	createCtrl := controllers.NewCreateDelegationController(dependencies.GetCreateDelegationUseCase())
	getByIdCtrl := controllers.NewGetDelegationByIDController(dependencies.GetDelegationByIDUseCase())
	getAllCtrl := controllers.NewGetAllDelegationsController(dependencies.GetAllDelegationsUseCase())
	updateCtrl := controllers.NewUpdateDelegationController(dependencies.GetUpdateDelegationUseCase())
	softDeleteCtrl := controllers.NewSoftDeleteDelegationController(dependencies.GetSoftDeleteDelegationUseCase())

	// Configure routes
	delegationRoutes := router.Group("/delegations")
	{
		delegationRoutes.POST("", createCtrl.Handle)          // Create delegation
		delegationRoutes.GET("", getAllCtrl.Handle)           // Get all delegations
		delegationRoutes.GET("/:id", getByIdCtrl.Handle)      // Get delegation by ID
		delegationRoutes.PUT("/:id", updateCtrl.Handle)       // Update delegation
		delegationRoutes.DELETE("/:id", softDeleteCtrl.Handle) // Soft delete delegation
	}
}