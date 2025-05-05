package routes

import (
	"api-seguridad/core/utils"
	"api-seguridad/resources/users/application"
	"api-seguridad/resources/users/domain/entities"
	"api-seguridad/resources/users/infrastructure/controllers"
	"api-seguridad/resources/users/infrastructure/dependencies"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ConfigureRoutes(router *gin.RouterGroup) {
	userRepo := dependencies.GetUserRepository()
	
	createUC := application.NewCreateUserUseCase(userRepo)
	getByIDUC := application.NewGetUserByIDUseCase(userRepo)
	updateUC := application.NewUpdateUserUseCase(userRepo)
	deleteUC := application.NewDeleteUserUseCase(userRepo)
	listUC := application.NewListUsersUseCase(userRepo)
	getByUsernameUC := application.NewGetUserByUsernameUseCase(userRepo)
	getByEmailUC := application.NewGetUserByEmailUseCase(userRepo)
	loginUC := application.NewLoginUseCase(userRepo)

	createCtrl := controllers.NewUserCreateController(createUC)
	getByIDCtrl := controllers.NewUserGetByIDController(getByIDUC)
	updateCtrl := controllers.NewUserUpdateController(updateUC)
	deleteCtrl := controllers.NewUserDeleteController(deleteUC)
	listCtrl := controllers.NewUserListController(listUC)
	getByUsernameCtrl := controllers.NewUserGetByUsernameController(getByUsernameUC)
	getByEmailCtrl := controllers.NewUserGetByEmailController(getByEmailUC)
	loginCtrl := controllers.NewUserLoginController(loginUC)
	
	userRoutes := router.Group("/users")
	{
		userRoutes.POST("/login", loginCtrl.Handle)
		userRoutes.GET("", listCtrl.Handle)
		userRoutes.POST("", createCtrl.Handle)
		userRoutes.GET("/:id_user", getByIDCtrl.Handle)
		userRoutes.PUT("/:id_user", updateCtrl.Handle)
		userRoutes.PATCH("/:id_user", deleteCtrl.Handle)
		userRoutes.GET("/username/:username", getByUsernameCtrl.Handle)
		userRoutes.GET("/email/:email", getByEmailCtrl.Handle)
	}

	router.POST("/users/init-admin", func(c *gin.Context) {
		var user entities.User
		if err := c.ShouldBindJSON(&user); err != nil {
			utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request payload", err)
			return
		}

		if err := createUC.Execute(c.Request.Context(), &user); err != nil {
			utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create admin user", err)
			return
		}

		utils.SuccessResponse(c, http.StatusCreated, "Admin user created successfully", user)
	})
}