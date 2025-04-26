package routes

import (
	"api-seguridad/resources/request_details/infrastructure/controllers"
	"api-seguridad/resources/request_details/infrastructure/dependencies"
	"github.com/gin-gonic/gin"
)

func ConfigureRequestDetailRoutes(router *gin.RouterGroup, authMiddleware gin.HandlerFunc) {
	// Initialize controllers
	createCtrl := controllers.NewCreateRequestDetailController(dependencies.GetCreateRequestDetailUseCase())
	getByIdCtrl := controllers.NewGetRequestDetailByIDController(dependencies.GetRequestDetailByIDUseCase())
	getByRequestIdCtrl := controllers.NewGetByRequestIDController(dependencies.GetByRequestIDUseCase())
	getByPoliceIdCtrl := controllers.NewGetByPoliceIDController(dependencies.GetByPoliceIDUseCase())
	updateCtrl := controllers.NewUpdateRequestDetailController(dependencies.GetUpdateRequestDetailUseCase())
	deleteCtrl := controllers.NewSoftDeleteRequestDetailController(dependencies.GetSoftDeleteRequestDetailUseCase())

	// Configure routes
	detailRoutes := router.Group("/request-details")
	detailRoutes.Use(authMiddleware)
	{
		// CRUD endpoints
		detailRoutes.POST("", createCtrl.Handle)
		detailRoutes.GET("/:id", getByIdCtrl.Handle)
		detailRoutes.GET("/request/:request_id", getByRequestIdCtrl.Handle)
		detailRoutes.GET("/police/:police_id", getByPoliceIdCtrl.Handle)
		detailRoutes.PUT("/:id", updateCtrl.Handle)
		detailRoutes.DELETE("/:id", deleteCtrl.Handle)
	}
}