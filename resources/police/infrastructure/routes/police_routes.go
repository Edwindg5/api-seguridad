//api-seguridad/resources/police/infrastructure/routes/police_routes.go
package routes

import (
	"api-seguridad/resources/police/infrastructure/controllers"
	"api-seguridad/resources/police/infrastructure/dependencies"
	"github.com/gin-gonic/gin"
)

func ConfigureRoutes(router *gin.RouterGroup) {
	// Initialize controllers with injected use cases
	createCtrl := controllers.NewCreatePoliceController(dependencies.GetCreatePoliceUseCase())
	getByIdCtrl := controllers.NewGetPoliceByIDController(dependencies.GetPoliceByIDUseCase())
	getAllCtrl := controllers.NewGetAllPoliceController(dependencies.GetAllPoliceUseCase())
	updateCtrl := controllers.NewUpdatePoliceController(dependencies.GetUpdatePoliceUseCase())
	softDeleteCtrl := controllers.NewSoftDeletePoliceController(dependencies.GetSoftDeletePoliceUseCase())
	getByCUIPCtrl := controllers.NewGetPoliceByCUIPController(dependencies.GetPoliceByCUIPUseCase())
	searchByNameCtrl := controllers.NewSearchPoliceByNameController(dependencies.GetSearchPoliceByNameUseCase())

	// Configure API routes
	policeRoutes := router.Group("/police")
	{
		// CRUD endpoints
		policeRoutes.POST("", createCtrl.Handle)          // Create new police record
		policeRoutes.GET("", getAllCtrl.Handle)           // Get all police records
		policeRoutes.GET("/:id", getByIdCtrl.Handle)      // Get police by ID
		policeRoutes.PUT("/:id", updateCtrl.Handle)       // Update police record
		policeRoutes.DELETE("/:id", softDeleteCtrl.Handle) // Soft delete police record

		// Search endpoints
		policeRoutes.GET("/search/cuip", getByCUIPCtrl.Handle) // Get police by CUIP
		policeRoutes.GET("/search/name", searchByNameCtrl.Handle) // Search police by name
	}
}