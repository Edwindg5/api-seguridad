package routes

import (

	"api-seguridad/resources/roles/infrastructure/controllers"
	"api-seguridad/resources/roles/infrastructure/dependencies"

	"github.com/gin-gonic/gin"
)

func ConfigureRoutes(router *gin.RouterGroup) {
	roleController := controllers.NewRoleController(dependencies.GetRoleService())

	roleRoutes := router.Group("/roles")
	{
		roleRoutes.POST("", roleController.CreateRole)
		roleRoutes.GET("/:id", roleController.GetRole)
	
	}
}