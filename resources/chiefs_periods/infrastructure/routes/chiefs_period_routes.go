//api-seguridad/resources/chiefs_periods/infrastructure/routes/chiefs_period_routes.go
package routes

import (
	"api-seguridad/resources/chiefs_periods/infrastructure/controllers"
	"api-seguridad/resources/chiefs_periods/infrastructure/dependencies"
	"github.com/gin-gonic/gin"
	"api-seguridad/core/middleware"
)

func ConfigureRoutes(router *gin.RouterGroup) {
	// Initialize controllers
	createCtrl := controllers.NewCreateChiefsPeriodController(dependencies.GetCreateChiefsPeriodUseCase())
	getByIdCtrl := controllers.NewGetChiefsPeriodByIDController(dependencies.GetChiefsPeriodByIDUseCase())
	getAllCtrl := controllers.NewGetAllChiefsPeriodsController(dependencies.GetAllChiefsPeriodsUseCase())
	updateCtrl := controllers.NewUpdateChiefsPeriodController(dependencies.GetUpdateChiefsPeriodUseCase())
	deleteCtrl := controllers.NewSoftDeleteChiefsPeriodController(dependencies.GetSoftDeleteChiefsPeriodUseCase())
	getActiveCtrl := controllers.NewGetActiveChiefsPeriodController(dependencies.GetActiveChiefsPeriodUseCase())
	getByDateRangeCtrl := controllers.NewGetChiefsPeriodsByDateRangeController(dependencies.GetChiefsPeriodsByDateRangeUseCase())

	// Configure routes
	chiefsRoutes := router.Group("/chiefs-periods")
	chiefsRoutes.Use(middleware.AuthMiddleware())
	{
		// CRUD endpoints
		chiefsRoutes.POST("", createCtrl.Handle)
		chiefsRoutes.GET("",  getAllCtrl.Handle)
		chiefsRoutes.GET("/:id",  getByIdCtrl.Handle)
		chiefsRoutes.PUT("/:id",  updateCtrl.Handle)			
		chiefsRoutes.DELETE("/:id",  deleteCtrl.Handle)

		// Special endpoints
		chiefsRoutes.GET("/active",  getActiveCtrl.Handle)
		chiefsRoutes.GET("/search",  getByDateRangeCtrl.Handle)
	}
}