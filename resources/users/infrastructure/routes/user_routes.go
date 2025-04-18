//api-seguridad/resources/users/infrastructure/routes/user_routes.go
package routes

import (

	"api-seguridad/resources/users/infrastructure/controllers"
	"api-seguridad/resources/users/infrastructure/dependencies"

	"github.com/gin-gonic/gin"
)

func ConfigureRoutes(router *gin.RouterGroup) {
	userController := controllers.NewUserController(dependencies.GetUserService())

	userRoutes := router.Group("/users")
	{
		userRoutes.POST("", userController.CreateUser)
		userRoutes.GET("/:id", userController.GetUser)

	}
}