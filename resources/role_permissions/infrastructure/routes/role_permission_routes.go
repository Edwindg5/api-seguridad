//api-seguridad/resources/role_permissions/infrastructure/routes/role_permission_routes.go
package routes

import (
	"api-seguridad/resources/role_permissions/infrastructure/controllers"
	"api-seguridad/resources/role_permissions/infrastructure/dependencies"
	"github.com/gin-gonic/gin"
	//"api-seguridad/core/middleware"
)

func ConfigureRoutes(router *gin.RouterGroup) {
	// Initialize controllers
	createCtrl := controllers.NewCreateRolePermissionController(dependencies.GetCreateRolePermissionUseCase())
	getByIdCtrl := controllers.NewGetRolePermissionByIDController(dependencies.GetRolePermissionByIDUseCase())
	getByRolePermCtrl := controllers.NewGetByRoleAndPermissionController(dependencies.GetByRoleAndPermissionUseCase())
	getAllByRoleCtrl := controllers.NewGetAllByRoleController(dependencies.GetAllByRoleUseCase())
	getAllCtrl := controllers.NewGetAllRolePermissionsController(dependencies.GetAllUseCase()) // Nuevo controller
	updateCtrl := controllers.NewUpdateRolePermissionController(dependencies.GetUpdateRolePermissionUseCase())
	deleteCtrl := controllers.NewDeleteRolePermissionController(dependencies.GetDeleteRolePermissionUseCase())	

	// Configure routes
	rpRoutes := router.Group("/role-permissions")
	//rpRoutes.Use(middleware.AuthMiddleware())
	{
		// CRUD endpoints
		rpRoutes.POST("", createCtrl.Handle)
		rpRoutes.GET("", getAllCtrl.Handle) // Nuevo endpoint para obtener todos xddd 
		rpRoutes.GET("/:id", getByIdCtrl.Handle)
		rpRoutes.GET("/check", getByRolePermCtrl.Handle) // ?role_id=X&permission_id=Y
		rpRoutes.GET("/role/:role_id", getAllByRoleCtrl.Handle)
		rpRoutes.PUT("/:id", updateCtrl.Handle)
		rpRoutes.DELETE("/:id", deleteCtrl.Handle)
	}
}