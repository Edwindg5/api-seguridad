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
	
	// Inicializar casos de uso y controladores
	createUC := application.NewCreateUserUseCase(userRepo)
	getByIDUC := application.NewGetUserByIDUseCase(userRepo)
	updateUC := application.NewUpdateUserUseCase(userRepo)
	deleteUC := application.NewDeleteUserUseCase(userRepo)
	listUC := application.NewListUsersUseCase(userRepo)
	getByUsernameUC := application.NewGetUserByUsernameUseCase(userRepo)
	getByEmailUC := application.NewGetUserByEmailUseCase(userRepo)
	loginUC := application.NewLoginUseCase(userRepo) // Nuevo caso de uso

	// Inicializar controladores
	createCtrl := controllers.NewUserCreateController(createUC)
	getByIDCtrl := controllers.NewUserGetByIDController(getByIDUC)
	updateCtrl := controllers.NewUserUpdateController(updateUC)
	deleteCtrl := controllers.NewUserDeleteController(deleteUC)
	listCtrl := controllers.NewUserListController(listUC)
	getByUsernameCtrl := controllers.NewUserGetByUsernameController(getByUsernameUC)
	getByEmailCtrl := controllers.NewUserGetByEmailController(getByEmailUC)
	loginCtrl := controllers.NewUserLoginController(loginUC) // Nuevo controlador

	// Configurar rutas
	userRoutes := router.Group("/users")
	{
		// Rutas PÚBLICAS (sin autenticación)
		userRoutes.POST("/login", loginCtrl.Handle) // Ruta de login
		userRoutes.GET("", listCtrl.Handle)        // Listar usuarios (acceso público)

		// Rutas PROTEGIDAS (requieren autenticación)
		protectedRoutes := userRoutes.Group("")
		protectedRoutes.Use(middleware.AuthMiddleware())  // Middleware aplicado solo a estas rutas
		{
			protectedRoutes.POST("", createCtrl.Handle)
			protectedRoutes.GET("/:id_user", getByIDCtrl.Handle)
			protectedRoutes.PUT("/:id_user", updateCtrl.Handle)
			protectedRoutes.DELETE("/:id_user", deleteCtrl.Handle)
			protectedRoutes.GET("/username/:username", getByUsernameCtrl.Handle)
			protectedRoutes.GET("/email/:email", getByEmailCtrl.Handle)
		}
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