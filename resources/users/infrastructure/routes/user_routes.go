// api-seguridad/resources/users/infrastructure/routes/user_routes.go
package routes

import (
	"api-seguridad/core/middleware"
	"api-seguridad/core/utils"
	"api-seguridad/resources/users/application"
	"api-seguridad/resources/users/domain/entities"
	"api-seguridad/resources/users/infrastructure/controllers"
	"api-seguridad/resources/users/infrastructure/dependencies"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ConfigureRoutes(router *gin.RouterGroup) {
	// Obtener dependencias
	userRepo := dependencies.GetUserRepository()
	
	// Inicializar casos de uso
	createUC := application.NewCreateUserUseCase(userRepo)
	getByIDUC := application.NewGetUserByIDUseCase(userRepo)
	updateUC := application.NewUpdateUserUseCase(userRepo)
	deleteUC := application.NewDeleteUserUseCase(userRepo)
	listUC := application.NewListUsersUseCase(userRepo)
	getByUsernameUC := application.NewGetUserByUsernameUseCase(userRepo)
	getByEmailUC := application.NewGetUserByEmailUseCase(userRepo)

	// Inicializar controladores
	createCtrl := controllers.NewUserCreateController(createUC)
	getByIDCtrl := controllers.NewUserGetByIDController(getByIDUC)
	updateCtrl := controllers.NewUserUpdateController(updateUC)
	deleteCtrl := controllers.NewUserDeleteController(deleteUC)
	listCtrl := controllers.NewUserListController(listUC)
	getByUsernameCtrl := controllers.NewUserGetByUsernameController(getByUsernameUC)
	getByEmailCtrl := controllers.NewUserGetByEmailController(getByEmailUC)

	// Configurar rutas con middleware de autenticación
	userRoutes := router.Group("/users")
	userRoutes.Use(middleware.AuthMiddleware())
	{
		// CRUD básico
		userRoutes.POST("", createCtrl.Handle)
		userRoutes.GET("", listCtrl.Handle)
		userRoutes.GET("/:id_user", getByIDCtrl.Handle)
		userRoutes.PUT("/:id_user", updateCtrl.Handle)
		userRoutes.DELETE("/:id_user", deleteCtrl.Handle)

		// Rutas adicionales
		userRoutes.GET("/username/:username", getByUsernameCtrl.Handle)
		userRoutes.GET("/email/:email", getByEmailCtrl.Handle)
	}

	// Ruta especial sin autenticación para crear admin inicial
	router.POST("/users/init-admin", func(c *gin.Context) {
		var user entities.User
		if err := c.ShouldBindJSON(&user); err != nil {
			utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request payload", err)
			return
		}

		user.CreatedBy = 0
		user.UpdatedBy = 0

		if err := createUC.Execute(c.Request.Context(), &user); err != nil {
			utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create admin user", err)
			return
		}

		utils.SuccessResponse(c, http.StatusCreated, "Admin user created successfully", user)
	})
}