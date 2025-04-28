//api-seguridad/resources/permissions/infrastructure/routes/permission_routes.go
package routes

import (
	"api-seguridad/resources/permissions/infrastructure/controllers"
	"api-seguridad/resources/permissions/infrastructure/dependencies"
	"github.com/gin-gonic/gin"
	"api-seguridad/core/middleware"
)

func ConfigureRoutes(router *gin.RouterGroup) {
	// Initialize controllers
	createCtrl := controllers.NewCreatePermissionController(dependencies.GetCreatePermissionUseCase())
	getByIdCtrl := controllers.NewGetPermissionByIDController(dependencies.GetPermissionByIDUseCase())
	getAllCtrl := controllers.NewGetAllPermissionsController(dependencies.GetAllPermissionsUseCase())
	updateCtrl := controllers.NewUpdatePermissionController(dependencies.GetUpdatePermissionUseCase())
	deleteCtrl := controllers.NewSoftDeletePermissionController(dependencies.GetSoftDeletePermissionUseCase())

	// Configure routes
	permissionRoutes := router.Group("/permissions")
	permissionRoutes.Use(middleware.AuthMiddleware()) // Middleware de autenticación para todas las rutas
	{
		permissionRoutes.POST("", createCtrl.Handle)          // Crear permiso
		permissionRoutes.GET("", getAllCtrl.Handle)           // Listar todos los permisos
		permissionRoutes.GET("/:id", getByIdCtrl.Handle)      // Obtener permiso por ID
		permissionRoutes.PUT("/:id", updateCtrl.Handle)       // Actualizar permiso
		permissionRoutes.DELETE("/:id", deleteCtrl.Handle)    // Eliminar permiso (soft delete)
	}
}