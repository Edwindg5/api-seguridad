package routes

import (
	"api-seguridad/resources/type_police/infrastructure/controllers"
	"api-seguridad/resources/type_police/infrastructure/dependencies"
	"github.com/gin-gonic/gin"
)

func ConfigureRoutes(router *gin.RouterGroup) {
	// Initialize controllers with injected use cases
	createCtrl := controllers.NewCreateTypePoliceController(dependencies.GetCreateTypePoliceUseCase())
	getByIdCtrl := controllers.NewGetTypePoliceByIDController(dependencies.GetTypePoliceByIDUseCase())
	getAllCtrl := controllers.NewGetAllTypePoliceController(dependencies.GetAllTypePoliceUseCase())
	updateCtrl := controllers.NewUpdateTypePoliceController(dependencies.GetUpdateTypePoliceUseCase())
	softDeleteCtrl := controllers.NewSoftDeleteTypePoliceController(dependencies.GetSoftDeleteTypePoliceUseCase())

	// Configure API routes
	typePoliceRoutes := router.Group("/type-police")
	{
		// POST /type-police - Create new type
		typePoliceRoutes.POST("", createCtrl.Handle)
		
		// GET /type-police - List all active types
		typePoliceRoutes.GET("", getAllCtrl.Handle)
		
		// GET /type-police/:id - Get type by ID
		typePoliceRoutes.GET("/:id", getByIdCtrl.Handle)
		
		// PUT /type-police/:id - Update type
		typePoliceRoutes.PUT("/:id", updateCtrl.Handle)
		
		// DELETE /type-police/:id - Soft delete type
		typePoliceRoutes.DELETE("/:id", softDeleteCtrl.Handle)
	}
}