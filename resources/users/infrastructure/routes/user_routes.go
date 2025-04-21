// api-seguridad/resources/users/infrastructure/routes/user_routes.go
package routes

import (
	"api-seguridad/resources/users/application"
	"api-seguridad/resources/users/infrastructure/controllers"
	"api-seguridad/resources/users/infrastructure/dependencies"

	"github.com/gin-gonic/gin"
)

func ConfigureRoutes(router *gin.RouterGroup) {
	// Obtener el repositorio de dependencias
	userRepo := dependencies.GetUserRepository()

	// Inicializar todos los casos de uso
	createUC := application.NewCreateUserUseCase(userRepo)
	getByIDUC := application.NewGetUserByIDUseCase(userRepo)
	updateUC := application.NewUpdateUserUseCase(userRepo)
	deleteUC := application.NewDeleteUserUseCase(userRepo)
	listUC := application.NewListUsersUseCase(userRepo)
	getByUsernameUC := application.NewGetUserByUsernameUseCase(userRepo)
	getByEmailUC := application.NewGetUserByEmailUseCase(userRepo)

	// Inicializar todos los controladores
	createCtrl := controllers.NewUserCreateController(createUC)
	getByIDCtrl := controllers.NewUserGetByIDController(getByIDUC)
	updateCtrl := controllers.NewUserUpdateController(updateUC)
	deleteCtrl := controllers.NewUserDeleteController(deleteUC)
	listCtrl := controllers.NewUserListController(listUC)
	getByUsernameCtrl := controllers.NewUserGetByUsernameController(getByUsernameUC)
	getByEmailCtrl := controllers.NewUserGetByEmailController(getByEmailUC)

	// Configurar rutas
	userRoutes := router.Group("/users")
	{
		// CRUD básico
		userRoutes.POST("", createCtrl.Handle)
		userRoutes.GET("", listCtrl.Handle)
		userRoutes.GET("/:id", getByIDCtrl.Handle)
		userRoutes.PUT("/:id", updateCtrl.Handle)
		userRoutes.DELETE("/:id", deleteCtrl.Handle)

		// Rutas adicionales
		userRoutes.GET("/username/:username", getByUsernameCtrl.Handle)
		userRoutes.GET("/email/:email", getByEmailCtrl.Handle)
	}
}