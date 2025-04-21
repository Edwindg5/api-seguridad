//api-seguridad/resources/roles/infrastructure/routes/role_routes.go
package routes

import (
	"api-seguridad/resources/roles/infrastructure/controllers"
	"api-seguridad/resources/roles/infrastructure/dependencies"
	"github.com/gin-gonic/gin"
)

func ConfigureRoutes(router *gin.RouterGroup) {
	// Inicializar controladores con sus respectivos casos de uso
	createCtrl := controllers.NewCreateRoleController(dependencies.GetCreateRoleUseCase())
	getByIdCtrl := controllers.NewGetRoleByIDController(dependencies.GetRoleByIDUseCase())
	getAllCtrl := controllers.NewGetAllRolesController(dependencies.GetAllRolesUseCase())
	updateCtrl := controllers.NewUpdateRoleController(dependencies.GetUpdateRoleUseCase())
	softDeleteCtrl := controllers.NewSoftDeleteRoleController(dependencies.GetSoftDeleteRoleUseCase())

	// Configurar rutas
	roleRoutes := router.Group("/roles")
	{
		roleRoutes.POST("", createCtrl.Handle)
		roleRoutes.GET("", getAllCtrl.Handle)
		roleRoutes.GET("/:id", getByIdCtrl.Handle)
		roleRoutes.PUT("/:id", updateCtrl.Handle)
		roleRoutes.DELETE("/:id", softDeleteCtrl.Handle)
	}
}